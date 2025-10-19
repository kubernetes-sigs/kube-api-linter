/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package runner

import (
	"encoding/json"
	"fmt"
	"go/parser"
	"go/token"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/scanner"

	"github.com/golangci/golangci-lint/v2/pkg/result"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const keyword = "want"

type jsonResult struct {
	Issues []*result.Issue
}

type expectation struct {
	kind string // either "fact" or "diagnostic"
	name string // name of object to which fact belongs, or "package" ("fact" only)
	rx   *regexp.Regexp
}

type key struct {
	file string
	line int
}

// Analyze analyzes the test expectations ('want').
// Inspired by:
// https://github.com/golang/tools/blob/1261a24ceb1867ea7439eda244e53e7ace4ad777/go/analysis/analysistest/analysistest.go#L655-L672
func analyze(sourcePath string, rawData []byte) {
	GinkgoHelper()

	want, err := getWantFromSourcePath(sourcePath)
	Expect(err).ToNot(HaveOccurred())

	var reportData jsonResult

	Expect(json.Unmarshal(rawData, &reportData)).To(Succeed())

	var failures []string
	for _, issue := range reportData.Issues {
		failures = append(failures, checkMessage(want, issue.Pos, "diagnostic", issue.FromLinter, issue.Text)...)
	}

	var surplus []string

	for key, expects := range want {
		for _, exp := range expects {
			err := fmt.Sprintf("%s:%d: no %s was reported matching %#q", key.file, key.line, exp.kind, exp.rx)
			surplus = append(surplus, err)
		}
	}

	sort.Strings(surplus)
	failures = append(failures, surplus...)

	Expect(failures).To(BeEmpty())
}

// Inspired by:
// https://github.com/golang/tools/blob/1261a24ceb1867ea7439eda244e53e7ace4ad777/go/analysis/analysistest/analysistest.go#L524-L553
func parseComments(sourcePath string, fileData []byte) (map[key][]expectation, error) {
	fset := token.NewFileSet()

	// the error is ignored to let 'typecheck' handle compilation error
	f, _ := parser.ParseFile(fset, sourcePath, fileData, parser.ParseComments)

	want := make(map[key][]expectation)

	for _, comment := range f.Comments {
		for _, c := range comment.List {
			text := strings.TrimPrefix(c.Text, "//")
			if text == c.Text { // not a //-comment.
				text = strings.TrimPrefix(text, "/*")
				text = strings.TrimSuffix(text, "*/")
			}

			if i := strings.Index(text, "// "+keyword); i >= 0 {
				text = text[i+len("// "):]
			}

			posn := fset.Position(c.Pos())

			text = strings.TrimSpace(text)

			if rest := strings.TrimPrefix(text, keyword); rest != text {
				delta, expects, err := parseExpectations(rest)
				if err != nil {
					return nil, err
				}

				want[key{sourcePath, posn.Line + delta}] = expects
			}
		}
	}

	return want, nil
}

// Inspired by:
// https://github.com/golang/tools/blob/1261a24ceb1867ea7439eda244e53e7ace4ad777/go/analysis/analysistest/analysistest.go#L685-L745
//
//nolint:cyclop
func parseExpectations(text string) (lineDelta int, expects []expectation, err error) {
	var scanErr string

	sc := new(scanner.Scanner).Init(strings.NewReader(text))
	sc.Error = func(_ *scanner.Scanner, msg string) {
		scanErr = msg // e.g. bad string escape
	}
	sc.Mode = scanner.ScanIdents | scanner.ScanStrings | scanner.ScanRawStrings | scanner.ScanInts

	scanRegexp := func(tok rune) (*regexp.Regexp, error) {
		if tok != scanner.String && tok != scanner.RawString {
			return nil, fmt.Errorf("got %s, want regular expression",
				scanner.TokenString(tok))
		}

		pattern, _ := strconv.Unquote(sc.TokenText()) // can't fail

		return regexp.Compile(pattern)
	}

	for {
		tok := sc.Scan()
		switch tok {
		case '+':
			tok = sc.Scan()
			if tok != scanner.Int {
				return 0, nil, fmt.Errorf("got +%s, want +Int", scanner.TokenString(tok))
			}

			lineDelta, _ = strconv.Atoi(sc.TokenText())
		case scanner.String, scanner.RawString:
			rx, err := scanRegexp(tok)
			if err != nil {
				return 0, nil, err
			}

			expects = append(expects, expectation{"diagnostic", "", rx})

		case scanner.Ident:
			name := sc.TokenText()

			tok = sc.Scan()
			if tok != ':' {
				return 0, nil, fmt.Errorf("got %s after %s, want ':'",
					scanner.TokenString(tok), name)
			}

			tok = sc.Scan()

			rx, err := scanRegexp(tok)
			if err != nil {
				return 0, nil, err
			}

			expects = append(expects, expectation{"diagnostic", name, rx})

		case scanner.EOF:
			if scanErr != "" {
				return 0, nil, fmt.Errorf("%s", scanErr)
			}

			return lineDelta, expects, nil

		default:
			return 0, nil, fmt.Errorf("unexpected %s", scanner.TokenString(tok))
		}
	}
}

// Inspired by:
// https://github.com/golang/tools/blob/1261a24ceb1867ea7439eda244e53e7ace4ad777/go/analysis/analysistest/analysistest.go#L594-L617
func checkMessage(want map[key][]expectation, posn token.Position, kind, name, message string) []string {
	GinkgoHelper()

	if !filepath.IsAbs(posn.Filename) {
		// Output from golangci-lint prefixes the filename with several layers of "../"
		// and then has the absolute path to the file.
		// Strip the "../" so that we use the absolute path.
		posn.Filename = "/" + strings.ReplaceAll(posn.Filename, "../", "")
	}

	k := key{posn.Filename, posn.Line}
	expects := want[k]

	var unmatched []string

	for i, exp := range expects {
		if exp.kind == kind && (exp.name == "" || exp.name == name) {
			if exp.rx.MatchString(message) {
				// matched: remove the expectation.
				expects[i] = expects[len(expects)-1]
				expects = expects[:len(expects)-1]
				want[k] = expects

				return []string{}
			}

			unmatched = append(unmatched, fmt.Sprintf("%#q", exp.rx))
		}
	}

	if unmatched == nil {
		return []string{fmt.Sprintf("%v: unexpected %s: %v", posn, kind, message)}
	} else {
		return []string{fmt.Sprintf("%v: %s %q does not match pattern %s", posn, kind, message, strings.Join(unmatched, " or "))}
	}
}

// getWantFromSourcePath combines all source files from the directory to get a complete list of `want` directives.
func getWantFromSourcePath(sourcePath string) (map[key][]expectation, error) {
	var err error

	// Make everything absolute so that we always compare absolute paths with the wants.
	sourcePath, err = filepath.Abs(sourcePath)
	if err != nil {
		return nil, fmt.Errorf("error getting absolute path: %w", err)
	}

	var sourcePaths []string

	sourcePathInfo, err := os.Stat(sourcePath)
	if err != nil {
		return nil, fmt.Errorf("error getting file info: %w", err)
	}

	if sourcePathInfo.IsDir() {
		sourcePaths, err = filepath.Glob(sourcePath + "/*.go")
		if err != nil {
			return nil, fmt.Errorf("error searching for go files in dir %s: %w", sourcePathInfo.Name(), err)
		}
	} else {
		sourcePaths = []string{sourcePath}
	}

	want := make(map[key][]expectation)

	for _, path := range sourcePaths {
		fileData, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("error reading file %s: %w", path, err)
		}

		fileWant, err := parseComments(path, fileData)
		if err != nil {
			return nil, fmt.Errorf("could not parse file %s comments: %w", path, err)
		}

		maps.Copy(want, fileWant)
	}

	return want, nil
}

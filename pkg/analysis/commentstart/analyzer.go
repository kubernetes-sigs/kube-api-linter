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
package commentstart

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
	"k8s.io/apimachinery/pkg/util/sets"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
)

const name = "commentstart"

type analyzer struct {
	excludePrefixes []string
}

// newAnalyzer creates a new analysis.Analyzer with the given config.
func newAnalyzer(cfg *Config) *analysis.Analyzer {
	if cfg == nil {
		cfg = &Config{}
	}

	defaultConfig(cfg)

	a := &analyzer{
		excludePrefixes: cfg.ExcludePrefixes,
	}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      "Check that all struct fields in an API have a godoc, and that the godoc starts with the serialised field name",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspector.Analyzer},
	}
}

// defaultConfig merges the built-in default exclude prefixes with
// any user-configured prefixes, deduplicating user entries against
// the defaults and against each other.
func defaultConfig(cfg *Config) {
	merged := make([]string, 0, len(defaultExcludePrefixes)+len(cfg.ExcludePrefixes))
	merged = append(merged, defaultExcludePrefixes...)

	seen := sets.New[string](defaultExcludePrefixes...)

	for _, p := range cfg.ExcludePrefixes {
		if !seen.Has(p) {
			merged = append(merged, p)
			seen.Insert(p)
		}
	}

	cfg.ExcludePrefixes = merged
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	for f := range inspect.Fields() {
		a.checkField(pass, f.Field, f.JSONTagInfo, f.QualifiedFieldName)
	}

	return nil, nil //nolint:nilnil
}

func (a *analyzer) checkField(pass *analysis.Pass, field *ast.Field, tagInfo extractjsontags.FieldTagInfo, qualifiedFieldName string) {
	if tagInfo.Name == "" {
		return
	}

	if field.Doc == nil {
		pass.Reportf(field.Pos(), "field %s is missing godoc comment", qualifiedFieldName)
		return
	}

	firstLine := field.Doc.List[0]

	// Check if the comment starts with an excluded prefix (e.g., "// Deprecated: ...").
	if a.hasExcludedPrefix(firstLine.Text) {
		return
	}

	if !strings.HasPrefix(firstLine.Text, "// "+tagInfo.Name+" ") {
		if strings.HasPrefix(strings.ToLower(firstLine.Text), strings.ToLower("// "+tagInfo.Name+" ")) {
			// The comment start is correct, apart from the casing, we can fix that.
			pass.Report(analysis.Diagnostic{
				Pos:     firstLine.Pos(),
				Message: fmt.Sprintf("godoc for field %s should start with '%s ...'", qualifiedFieldName, tagInfo.Name),
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: fmt.Sprintf("should replace first word with `%s`", tagInfo.Name),
						TextEdits: []analysis.TextEdit{
							{
								Pos:     firstLine.Pos(),
								End:     firstLine.Pos() + token.Pos(len(tagInfo.Name)) + token.Pos(4),
								NewText: []byte("// " + tagInfo.Name + " "),
							},
						},
					},
				},
			})
		} else {
			pass.Reportf(field.Doc.List[0].Pos(), "godoc for field %s should start with '%s ...'", qualifiedFieldName, tagInfo.Name)
		}
	}
}

// hasExcludedPrefix checks if the comment text starts with "// <prefix>"
// for any configured excluded prefix.
func (a *analyzer) hasExcludedPrefix(commentText string) bool {
	for _, prefix := range a.excludePrefixes {
		if strings.HasPrefix(commentText, "// "+prefix) {
			return true
		}
	}

	return false
}

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
package fieldnaming

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
)

const name = "fieldnaming"

// Analyzer is the analyzer for the nophase package.
// It checks that no struct fields named 'phase', or that contain phase as a
// substring are present.
var Analyzer = &analysis.Analyzer{
	Name:     name,
	Doc:      "phase fields are deprecated and conditions should be preferred, avoid phase like enum fields",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer, extractjsontags.Analyzer},
}

var defaults = []NamingConvention{
	{
		Matcher:   *regexp.MustCompile("(?i)phase"),
		Operation: "Drop",
		Message:   "phase fields are deprecated and discouraged. conditions should be used instead.",
	},
	{
		Matcher:   *regexp.MustCompile("(?i)timestamp"),
		Operation: "Replace",
		Message:   "prefer use of the term 'time' over 'timestamp'",
		// replacement of `Time` follows CamelCase principles for field names and JSON tags
		// TODO: Handle case where it is _not_ the second word in CamelCase for json tag
		Replacement: "Time",
	},
	{
		Matcher:   *regexp.MustCompile("(?i)reference"),
		Operation: "Replace",
		Message:   "prefer use of the term 'ref' over 'reference'",
		// replacement of `Ref` follows CamelCase principles for field names and JSON tags
		// TODO: Handle case where it is _not_ the second word in CamelCase for json tag
		Replacement: "Ref",
	},
}

type NamingConvention struct {
	// Matcher is a regular expression
	// used to identify field names
	// where this convention applies
	Matcher regexp.Regexp

	// Replacement is an optional
	// string value used to replace the matched content
	// in a suggested fix.
	// Only used when Operation is Replace.
	Replacement string

	// Operation is the type of operation that should take place for this
	// naming convention.
	// One of Drop, Replace.
	Operation string

	// Message is the message that should be included in the
	// linter report when this naming convention is applied.
	Message string
}

func run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	jsonTags, ok := pass.ResultOf[extractjsontags.Analyzer].(extractjsontags.StructFieldTags)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetJSONTags
	}

	// Filter to fields so that we can iterate over fields in a struct.
	nodeFilter := []ast.Node{
		(*ast.Field)(nil),
	}

	// Preorder visits all the nodes of the AST in depth-first order. It calls
	// f(n) for each node n before it visits n's children.
	//
	// We use the filter defined above, ensuring we only look at struct fields.
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		field, ok := n.(*ast.Field)
		if !ok {
			return
		}

		if field == nil || len(field.Names) == 0 {
			return
		}

		fieldName := utils.FieldName(field)
		for _, convention := range defaults {
			if convention.Operation == "Drop" {
				if convention.Matcher.MatchString(fieldName) {
					pass.Reportf(field.Pos(), "field %s: %s", fieldName, convention.Message)
				}
			}

			if convention.Operation == "Replace" {
				replacement := convention.Matcher.ReplaceAllString(fieldName, convention.Replacement)
				if replacement != fieldName {
					pass.Report(analysis.Diagnostic{
						Pos:     field.Pos(),
						Message: fmt.Sprintf("field %s: %s", fieldName, convention.Message),
						SuggestedFixes: []analysis.SuggestedFix{
							{
								Message: fmt.Sprintf("replace with %s", convention.Replacement),
								TextEdits: []analysis.TextEdit{
									{
										Pos:     field.Pos(),
										NewText: []byte(replacement),
										End:     field.Pos()+token.Pos(len(fieldName)),
									},
								},
							},
						},
					})
				}
			}
		}

		// Then check if the json serialization of the field contains 'phase'
		_ = jsonTags.FieldTags(field)
	})

	return nil, nil //nolint:nilnil
}

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

package notimestamp

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
)

const name = "notimestamp"

// Analyzer is the analyzer for the notimestamp package.
// It checks that no struct fields named 'timestamp', or that contain timestamp as a
// substring are present.
var Analyzer = &analysis.Analyzer{
	Name:     name,
	Doc:      "check that non of the struct field named timestamp or contain timestamp as a substring",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer, extractjsontags.Analyzer},
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

		// First check if the struct field name contains 'timestamp'
		if strings.Contains(strings.ToLower(fieldName), "timestamp") {
			pass.Reportf(field.Pos(),
				"field %s: fields with timestamp substring should be avoided",
				fieldName,
			)

			return
		}

		// Then check if the json serialization of the field contains 'timestamp'
		tagInfo := jsonTags.FieldTags(field)

		if strings.Contains(strings.ToLower(tagInfo.Name), "timestamp") {
			pass.Reportf(field.Pos(),
				"field %s: fields with timestamp substring should be avoided",
				fieldName,
			)
		}
	})

	return nil, nil //nolint:nilnil
}

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
package references

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
)

const name = "references"

type analyzer struct {
	allowRefAndRefs bool
}

// newAnalyzer creates a new analysis.Analyzer for the references linter.
func newAnalyzer(cfg *Config) *analysis.Analyzer {
	if cfg == nil {
		cfg = &Config{}
	}

	a := &analyzer{
		allowRefAndRefs: cfg.AllowRefAndRefs,
	}

	analyzer := &analysis.Analyzer{
		Name:     name,
		Doc:      "Enforces that fields use Ref/Refs and not Reference/References",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspector.Analyzer, extractjsontags.Analyzer},
	}

	return analyzer
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, stack []ast.Node, jsonTagInfo extractjsontags.FieldTagInfo, markersAccess markers.Markers) {
		a.checkField(pass, field, jsonTagInfo)
	})

	return nil, nil //nolint:nilnil
}

func (a *analyzer) checkField(pass *analysis.Pass, field *ast.Field, jsonTagInfo extractjsontags.FieldTagInfo) {
	if field == nil || len(field.Names) == 0 {
		return
	}

	fieldName := utils.FieldName(field)

	if strings.HasSuffix(fieldName, "Reference") {
		suggestedName := strings.TrimSuffix(fieldName, "Reference") + "Ref"
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("field %s should use 'Ref' instead of 'Reference'", fieldName),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "replace 'Reference' with 'Ref'",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     field.Names[0].Pos(),
							NewText: []byte(suggestedName),
							End:     field.Names[0].End(),
						},
					},
				},
			},
		})
		return
	}

	if strings.HasSuffix(fieldName, "References") {
		suggestedName := strings.TrimSuffix(fieldName, "References") + "Refs"
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("field %s should use 'Refs' instead of 'References'", fieldName),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "replace 'References' with 'Refs'",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     field.Names[0].Pos(),
							NewText: []byte(suggestedName),
							End:     field.Names[0].End(),
						},
					},
				},
			},
		})
		return
	}

	// If allowRefAndRefs is false, report errors for standalone Ref/Refs fields
	// If allowRefAndRefs is true (OpenShift), don't report errors for Ref/Refs fields
	if !a.allowRefAndRefs {
		// Check for fields ending with Ref or Refs (excluding those already handled above)
		if fieldName != "Ref" && strings.HasSuffix(fieldName, "Ref") && !strings.HasSuffix(fieldName, "eRef") {
			pass.Reportf(field.Pos(), "field %s should not use 'Ref' suffix", fieldName)
		}

		if fieldName != "Refs" && strings.HasSuffix(fieldName, "Refs") && !strings.HasSuffix(fieldName, "eRefs") {
			pass.Reportf(field.Pos(), "field %s should not use 'Refs' suffix", fieldName)
		}
	}
}

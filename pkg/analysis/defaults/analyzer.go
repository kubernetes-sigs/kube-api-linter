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
package defaults

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

const (
	name = "defaults"
)

func init() {
	markershelper.DefaultRegistry().Register(
		markers.DefaultMarker,
		markers.KubebuilderDefaultMarker,
		markers.K8sDefaultMarker,
		markers.OptionalMarker,
		markers.KubebuilderOptionalMarker,
		markers.K8sOptionalMarker,
	)
}

// Analyzer is the analyzer for the defaults package.
var Analyzer = &analysis.Analyzer{
	Name: name,
	Doc: `Checks that fields with default markers are configured correctly.
Fields with default markers (+default, +kubebuilder:default, or +k8s:default) should also be marked as optional.
Additionally, fields with default markers should have "omitempty" or "omitzero" in their json tags to ensure that the default values are applied correctly during serialization and deserialization.
`,
	Run:      run,
	Requires: []*analysis.Analyzer{inspector.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, jsonTagInfo extractjsontags.FieldTagInfo, markersAccess markershelper.Markers, qualifiedFieldName string) {
		checkField(pass, field, jsonTagInfo, markersAccess, qualifiedFieldName)
	})

	return nil, nil //nolint:nilnil
}

func checkField(pass *analysis.Pass, field *ast.Field, jsonTagInfo extractjsontags.FieldTagInfo, markersAccess markershelper.Markers, qualifiedFieldName string) {
	if field == nil || len(field.Names) == 0 {
		return
	}

	fieldMarkers := markersAccess.FieldMarkers(field)

	// Check for any default marker (+default, +kubebuilder:default, or +k8s:default)
	hasDefault := fieldMarkers.Has(markers.DefaultMarker)
	hasKubebuilderDefault := fieldMarkers.Has(markers.KubebuilderDefaultMarker)
	hasK8sDefault := fieldMarkers.Has(markers.K8sDefaultMarker)

	if !hasDefault && !hasKubebuilderDefault && !hasK8sDefault {
		return
	}

	if hasKubebuilderDefault {
		checkKubebuilderDefault(pass, field, fieldMarkers, qualifiedFieldName)
	}

	checkDefaultOptional(pass, field, markersAccess, qualifiedFieldName)

	checkDefaultOmitEmptyOrOmitZero(pass, field, jsonTagInfo, qualifiedFieldName)
}

func checkKubebuilderDefault(pass *analysis.Pass, field *ast.Field, fieldMarkers markershelper.MarkerSet, qualifiedFieldName string) {
	kubebuilderDefaultMarkers := fieldMarkers.Get(markers.KubebuilderDefaultMarker)
	for _, marker := range kubebuilderDefaultMarkers {
		payloadValue := marker.Payload.Value
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("field %s should use +default or +k8s:default marker instead of +kubebuilder:default", qualifiedFieldName),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: fmt.Sprintf("replace +kubebuilder:default with +default=%s", payloadValue),
					TextEdits: []analysis.TextEdit{
						{
							Pos:     marker.Pos,
							End:     marker.End,
							NewText: fmt.Appendf(nil, "// +default=%s", payloadValue),
						},
					},
				},
			},
		})
	}
}

func checkDefaultOptional(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers, qualifiedFieldName string) {
	if !utils.IsFieldOptional(field, markersAccess) {
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("field %s has a default value but is not marked as optional", qualifiedFieldName),
		})
	}
}

func checkDefaultOmitEmptyOrOmitZero(pass *analysis.Pass, field *ast.Field, jsonTagInfo extractjsontags.FieldTagInfo, qualifiedFieldName string) {
	if !jsonTagInfo.OmitEmpty && !jsonTagInfo.OmitZero && !jsonTagInfo.Inline && !jsonTagInfo.Ignored {
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("field %s has a default value but does not have omitempty or omitzero in its json tag", qualifiedFieldName),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: fmt.Sprintf("add omitempty to the json tag of field %s", qualifiedFieldName),
					TextEdits: []analysis.TextEdit{
						{
							Pos:     jsonTagInfo.Pos,
							End:     jsonTagInfo.End,
							NewText: fmt.Appendf([]byte{}, "%s,omitempty,omitzero", jsonTagInfo.RawValue),
						},
					},
				},
			},
		})
	}
}

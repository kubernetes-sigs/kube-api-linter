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
package requiredfields

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils/serialization"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

const (
	name = "requiredfields"
)

func init() {
	markershelper.DefaultRegistry().Register(
		// Required markers
		markers.RequiredMarker,
		markers.KubebuilderRequiredMarker,
		markers.K8sRequiredMarker,

		// MinItems markers
		markers.KubebuilderMinItemsMarker,
		markers.K8sMinItemsMarker,

		// MinLength markers
		markers.KubebuilderMinLengthMarker,
		markers.K8sMinLengthMarker,

		// MinProperties markers
		markers.KubebuilderMinPropertiesMarker,

		// Minimum/Maximum markers
		markers.KubebuilderMinimumMarker,
		markers.K8sMinimumMarker,
		markers.KubebuilderMaximumMarker,
		markers.K8sMaximumMarker,

		// Enum markers
		markers.KubebuilderEnumMarker,
		markers.K8sEnumMarker,
	)
}

type analyzer struct {
	serializationCheck      serialization.SerializationCheck
	preferredRequiredMarker string
}

// newAnalyzer creates a new analyzer.
func newAnalyzer(cfg *RequiredFieldsConfig) *analysis.Analyzer {
	if cfg == nil {
		cfg = &RequiredFieldsConfig{}
	}

	defaultConfig(cfg)

	serializationCheck := serialization.New(&serialization.Config{
		Pointers: serialization.PointersConfig{
			Policy: serialization.PointersPolicy(cfg.Pointers.Policy),
			// We only allow the WhenRequired preference for required fields.
			// This works for both built-in types and custom resources, and
			// avoids pointers unless absolutely necessary.
			Preference: serialization.PointersPreferenceWhenRequired,
		},
		OmitEmpty: serialization.OmitEmptyConfig{
			Policy: serialization.OmitEmptyPolicy(cfg.OmitEmpty.Policy),
		},
		OmitZero: serialization.OmitZeroConfig{
			Policy: serialization.OmitZeroPolicy(cfg.OmitZero.Policy),
		},
	})

	a := &analyzer{
		serializationCheck:      serializationCheck,
		preferredRequiredMarker: cfg.PreferredRequiredMarker,
	}

	return &analysis.Analyzer{
		Name: name,
		Doc: `Checks that all required fields are serialized correctly.
		Where the zero value is not valid, this means the field should not be a pointer, and should have the omitempty tag.
		Where the zero value is valid, this means the field should be a pointer and should not have the omitempty tag.
		`,
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspector.Analyzer, extractjsontags.Analyzer},
	}
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, jsonTagInfo extractjsontags.FieldTagInfo, markersAccess markershelper.Markers, qualifiedFieldName string) {
		a.checkField(pass, field, markersAccess, jsonTagInfo, qualifiedFieldName)
	})

	return nil, nil //nolint:nilnil
}

func (a *analyzer) checkField(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers, jsonTags extractjsontags.FieldTagInfo, qualifiedFieldName string) {
	if field == nil || len(field.Names) == 0 {
		return
	}

	if !utils.IsFieldRequired(field, markersAccess) {
		// The field is not marked required, so we don't need to check it.
		return
	}

	fieldMarkers := markersAccess.FieldMarkers(field)

	// Check +k8s:required marker (for declarative validation)
	// If +k8s:required is present but the field doesn't have the preferred required marker
	// or any other non-k8s required marker, suggest adding the preferred one.
	// When preferredRequiredMarker is "k8s:required", having +k8s:required alone is sufficient.
	hasPreferredRequired := fieldMarkers.Has(a.preferredRequiredMarker)
	hasAnyNonK8sRequired := fieldMarkers.Has(markers.RequiredMarker) || fieldMarkers.Has(markers.KubebuilderRequiredMarker)
	a.checkK8sRequired(pass, field, fieldMarkers, qualifiedFieldName, hasPreferredRequired || hasAnyNonK8sRequired)

	if field.Type == nil {
		// The field has no type? We can't check if it's a pointer.
		return
	}

	a.serializationCheck.Check(pass, field, markersAccess, jsonTags, qualifiedFieldName)
}

// checkK8sRequired checks for +k8s:required marker usage.
// +k8s:required is for declarative validation and is separate from +required or +kubebuilder:validation:Required markers.
// If the field has +k8s:required but doesn't have the preferred required marker or any other non-k8s required marker,
// we suggest adding the preferred one. When preferredRequiredMarker is "k8s:required", having +k8s:required alone is sufficient.
func (a *analyzer) checkK8sRequired(pass *analysis.Pass, field *ast.Field, fieldMarkers markershelper.MarkerSet, qualifiedFieldName string, hasOtherRequired bool) {
	if !fieldMarkers.Has(markers.K8sRequiredMarker) {
		return
	}

	// If the field already has +required or +kubebuilder:validation:Required, +k8s:required is acceptable alongside them
	// (e.g., in K/K where both are needed during transition period)
	if hasOtherRequired {
		return
	}

	// If only +k8s:required is present, suggest adding the preferred required marker
	k8sRequiredMarkers := fieldMarkers.Get(markers.K8sRequiredMarker)
	for _, marker := range k8sRequiredMarkers {
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("field %s has +%s but should also have +%s marker", qualifiedFieldName, markers.K8sRequiredMarker, a.preferredRequiredMarker),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: fmt.Sprintf("add +%s", a.preferredRequiredMarker),
					TextEdits: []analysis.TextEdit{
						{
							Pos:     marker.Pos,
							End:     marker.Pos,
							NewText: fmt.Appendf(nil, "// +%s\n\t", a.preferredRequiredMarker),
						},
					},
				},
			},
		})
	}
}

func defaultConfig(cfg *RequiredFieldsConfig) {
	if cfg.Pointers.Policy == "" {
		cfg.Pointers.Policy = RequiredFieldsPointerPolicySuggestFix
	}

	if cfg.OmitEmpty.Policy == "" {
		cfg.OmitEmpty.Policy = RequiredFieldsOmitEmptyPolicySuggestFix
	}

	if cfg.OmitZero.Policy == "" {
		cfg.OmitZero.Policy = RequiredFieldsOmitZeroPolicySuggestFix
	}

	if cfg.PreferredRequiredMarker == "" {
		cfg.PreferredRequiredMarker = markers.RequiredMarker
	}
}

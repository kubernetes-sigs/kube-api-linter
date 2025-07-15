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
package conflictingmarkers

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"k8s.io/apimachinery/pkg/util/sets"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	markersconsts "sigs.k8s.io/kube-api-linter/pkg/markers"
)

const name = "conflictingmarkers"

func init() {
	// Register all the markers we care about with the marker registry
	markers.DefaultRegistry().Register(markersconsts.OptionalMarker)
	markers.DefaultRegistry().Register(markersconsts.RequiredMarker)
	markers.DefaultRegistry().Register(markersconsts.DefaultMarker)
	markers.DefaultRegistry().Register(markersconsts.KubebuilderOptionalMarker)
	markers.DefaultRegistry().Register(markersconsts.KubebuilderRequiredMarker)
	markers.DefaultRegistry().Register(markersconsts.KubebuilderDefaultMarker)
	markers.DefaultRegistry().Register(markersconsts.K8sOptionalMarker)
	markers.DefaultRegistry().Register(markersconsts.K8sRequiredMarker)
}

type analyzer struct {
	conflictSets []ConflictSet
}

func newAnalyzer(cfg *ConflictingMarkersConfig) *analysis.Analyzer {
	if cfg == nil {
		cfg = &ConflictingMarkersConfig{}
	}

	// Register custom markers from configuration
	for _, conflictSet := range cfg.CustomConflicts {
		for _, markerID := range conflictSet.SetA {
			markers.DefaultRegistry().Register(markerID)
		}

		for _, markerID := range conflictSet.SetB {
			markers.DefaultRegistry().Register(markerID)
		}
	}

	a := &analyzer{
		conflictSets: append(defaultConflictSets(), cfg.CustomConflicts...),
	}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      "Check that fields do not have conflicting markers from mutually exclusive sets",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspector.Analyzer},
	}
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, stack []ast.Node, _ extractjsontags.FieldTagInfo, markersAccess markers.Markers) {
		checkField(pass, field, markersAccess, a.conflictSets)
	})

	return nil, nil //nolint:nilnil
}

func checkField(pass *analysis.Pass, field *ast.Field, markersAccess markers.Markers, conflictSets []ConflictSet) {
	if field == nil || len(field.Names) == 0 {
		return
	}

	markers := utils.TypeAwareMarkerCollectionForField(pass, markersAccess, field)

	for _, conflictSet := range conflictSets {
		checkConflict(pass, field, markers, conflictSet)
	}
}

func checkConflict(pass *analysis.Pass, field *ast.Field, markers markers.MarkerSet, conflictSet ConflictSet) {
	setAMarkers := sets.New[string]()
	setBMarkers := sets.New[string]()

	// Collect markers from set A
	for _, markerID := range conflictSet.SetA {
		if markers.Has(markerID) {
			setAMarkers.Insert(markerID)
		}
	}

	// Collect markers from set B
	for _, markerID := range conflictSet.SetB {
		if markers.Has(markerID) {
			setBMarkers.Insert(markerID)
		}
	}

	// If both sets have markers, report the conflict
	if setAMarkers.Len() > 0 && setBMarkers.Len() > 0 {
		reportConflict(pass, field, conflictSet, setAMarkers, setBMarkers)
	}
}

func reportConflict(pass *analysis.Pass, field *ast.Field, conflictSet ConflictSet, setAMarkers, setBMarkers sets.Set[string]) {
	pass.Report(analysis.Diagnostic{
		Pos: field.Pos(),
		Message: fmt.Sprintf("field %s has conflicting markers: %s: %v and %v. %s",
			field.Names[0].Name,
			conflictSet.Name,
			sets.List(setAMarkers), sets.List(setBMarkers),
			conflictSet.Description),
	})
}

func defaultConflictSets() []ConflictSet {
	return []ConflictSet{
		{
			Name:        "optional_vs_required",
			SetA:        []string{markersconsts.OptionalMarker, markersconsts.KubebuilderOptionalMarker, markersconsts.K8sOptionalMarker},
			SetB:        []string{markersconsts.RequiredMarker, markersconsts.KubebuilderRequiredMarker, markersconsts.K8sRequiredMarker},
			Description: "A field cannot be both optional and required",
		},
		{
			Name:        "default_vs_required",
			SetA:        []string{markersconsts.DefaultMarker, markersconsts.KubebuilderDefaultMarker},
			SetB:        []string{markersconsts.RequiredMarker, markersconsts.KubebuilderRequiredMarker, markersconsts.K8sRequiredMarker},
			Description: "A field with a default value cannot be required",
		},
	}
}

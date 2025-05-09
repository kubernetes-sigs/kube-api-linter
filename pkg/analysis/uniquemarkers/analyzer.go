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
package uniquemarkers

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	"sigs.k8s.io/kube-api-linter/pkg/config"
	markersutil "sigs.k8s.io/kube-api-linter/pkg/markers"
)

const name = "uniquemarkers"

func init() {
	markers.DefaultRegistry().Register(defaultUniqueMarkers()...)
}

type analyzer struct {
	uniqueMarkers []string
}

func newAnalyzer(cfg config.UniqueMarkersConfig) *analysis.Analyzer {
	a := &analyzer{
		uniqueMarkers: append(defaultUniqueMarkers(), cfg.CustomMarkers...),
	}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      "Check that all markers that should be unique on a field/type are only present once",
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
		checkField(pass, field, markersAccess, a.uniqueMarkers)
	})

	inspect.InspectTypeSpec(func(typeSpec *ast.TypeSpec, markersAccess markers.Markers) {
		checkType(pass, typeSpec, markersAccess, a.uniqueMarkers)
	})

	return nil, nil //nolint:nilnil
}

func checkField(pass *analysis.Pass, field *ast.Field, markersAccess markers.Markers, uniqueMarkers []string) {
	if field == nil || len(field.Names) == 0 {
		return
	}

	markers := utils.TypeAwareMarkerCollectionForField(pass, markersAccess, field)
	check(markers, uniqueMarkers, reportField(pass, field))
}

func checkType(pass *analysis.Pass, typeSpec *ast.TypeSpec, markersAccess markers.Markers, uniqueMarkers []string) {
	if typeSpec == nil {
		return
	}

	markers := markersAccess.TypeMarkers(typeSpec)
	check(markers, uniqueMarkers, reportType(pass, typeSpec))
}

func check(markers markers.MarkerSet, uniqueMarkers []string, reportFunc func(id string)) {
	for _, identifier := range uniqueMarkers {
		marks := markers.MarkersForIdentifier(identifier)
		if len(marks) > 1 {
			reportFunc(identifier)
		}
	}
}

func reportField(pass *analysis.Pass, field *ast.Field) func(id string) {
	return func(id string) {
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("field %s has multiple definitions of marker %s when only a single definition should exist", field.Names[0].Name, id),
		})
	}
}

func reportType(pass *analysis.Pass, typeSpec *ast.TypeSpec) func(id string) {
	return func(id string) {
		pass.Report(analysis.Diagnostic{
			Pos:     typeSpec.Pos(),
			Message: fmt.Sprintf("type %s has multiple definitions of marker %s when only a single definition should exist", typeSpec.Name, id),
		})
	}
}

func defaultUniqueMarkers() []string {
	return []string{
		// Basic unique markers
		// ------
		markersutil.DefaultMarker,
		// ------

		// Kubebuilder-specific unique markers
		// ------
		markersutil.KubebuilderDefaultMarker,
		markersutil.KubebuilderExampleMarker,
		markersutil.KubebuilderEnumMarker,
		markersutil.KubebuilderExclusiveMaximumMarker,
		markersutil.KubebuilderExclusiveMinimumMarker,
		markersutil.KubebuilderFormatMarker,
		markersutil.KubebuilderMaxItemsMarker,
		markersutil.KubebuilderMaxLengthMarker,
		markersutil.KubebuilderMaxPropertiesMarker,
		markersutil.KubebuilderMaximumMarker,
		markersutil.KubebuilderMinItemsMarker,
		markersutil.KubebuilderMinLengthMarker,
		markersutil.KubebuilderMinPropertiesMarker,
		markersutil.KubebuilderMinimumMarker,
		markersutil.KubebuilderMultipleOfMarker,
		markersutil.KubebuilderPatternMarker,
		markersutil.KubebuilderTypeMarker,
		markersutil.KubebuilderUniqueItemsMarker,

		markersutil.KubebuilderItemsEnumMarker,
		markersutil.KubebuilderItemsFormatMarker,
		markersutil.KubebuilderItemsMaxLengthMarker,
		markersutil.KubebuilderItemsMaxItemsMarker,
		markersutil.KubebuilderItemsMaxPropertiesMarker,
		markersutil.KubebuilderItemsMaximumMarker,
		markersutil.KubebuilderItemsMinLengthMarker,
		markersutil.KubebuilderItemsMinItemsMarker,
		markersutil.KubebuilderItemsMinPropertiesMarker,
		markersutil.KubebuilderItemsMinimumMarker,
		markersutil.KubebuilderItemsExclusiveMaximumMarker,
		markersutil.KubebuilderItemsExclusiveMinimumMarker,
		markersutil.KubebuilderItemsMultipleOfMarker,
		markersutil.KubebuilderItemsPatternMarker,
		markersutil.KubebuilderItemsTypeMarker,
		markersutil.KubebuilderItemsUniqueItemsMarker,
		// ------
	}
}

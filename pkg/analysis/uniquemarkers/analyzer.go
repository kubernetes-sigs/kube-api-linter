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
	"strings"

	"golang.org/x/tools/go/analysis"
	"k8s.io/apimachinery/pkg/util/sets"
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
	for _, uniqueMarker := range defaultUniqueMarkers() {
		markers.DefaultRegistry().Register(uniqueMarker.Identifier)
	}
}

type analyzer struct {
	uniqueMarkers []config.UniqueMarker
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

func checkField(pass *analysis.Pass, field *ast.Field, markersAccess markers.Markers, uniqueMarkers []config.UniqueMarker) {
	if field == nil || len(field.Names) == 0 {
		return
	}

	markers := utils.TypeAwareMarkerCollectionForField(pass, markersAccess, field)
	check(markers, uniqueMarkers, reportField(pass, field))
}

func checkType(pass *analysis.Pass, typeSpec *ast.TypeSpec, markersAccess markers.Markers, uniqueMarkers []config.UniqueMarker) {
	if typeSpec == nil {
		return
	}

	markers := markersAccess.TypeMarkers(typeSpec)
	check(markers, uniqueMarkers, reportType(pass, typeSpec))
}

func check(markerSet markers.MarkerSet, uniqueMarkers []config.UniqueMarker, reportFunc func(id string)) {
	for _, marker := range uniqueMarkers {
		marks := markerSet.Get(marker.Identifier)
		markSet := sets.New[string]()
		for _, mark := range marks {
			id := mark.Identifier

			if len(marker.Attributes) > 0 {
				id += ":"
			}

			for _, attr := range marker.Attributes {
				id += fmt.Sprintf("%s=%s,", attr, mark.Expressions[attr])
			}

			id = strings.TrimSuffix(id, ",")

			if markSet.Has(id) {
				reportFunc(id)
				continue
			}

			markSet.Insert(id)
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

func defaultUniqueMarkers() []config.UniqueMarker {
	return []config.UniqueMarker{
		// Basic unique markers
		// ------
		{
			Identifier: markersutil.DefaultMarker,
		},
		// ------

		// Kubebuilder-specific unique markers
		// ------
		{
			Identifier: markersutil.KubebuilderDefaultMarker,
		},
		{
			Identifier: markersutil.KubebuilderExampleMarker,
		},
		{
			Identifier: markersutil.KubebuilderEnumMarker,
		},
		{
			Identifier: markersutil.KubebuilderExclusiveMaximumMarker,
		},
		{
			Identifier: markersutil.KubebuilderExclusiveMinimumMarker,
		},
		{
			Identifier: markersutil.KubebuilderFormatMarker,
		},
		{
			Identifier: markersutil.KubebuilderMaxItemsMarker,
		},
		{
			Identifier: markersutil.KubebuilderMaxLengthMarker,
		},
		{
			Identifier: markersutil.KubebuilderMaxPropertiesMarker,
		},
		{
			Identifier: markersutil.KubebuilderMaximumMarker,
		},
		{
			Identifier: markersutil.KubebuilderMinItemsMarker,
		},
		{
			Identifier: markersutil.KubebuilderMinLengthMarker,
		},
		{
			Identifier: markersutil.KubebuilderMinPropertiesMarker,
		},
		{
			Identifier: markersutil.KubebuilderMinimumMarker,
		},
		{
			Identifier: markersutil.KubebuilderMultipleOfMarker,
		},
		{
			Identifier: markersutil.KubebuilderPatternMarker,
		},
		{
			Identifier: markersutil.KubebuilderTypeMarker,
		},
		{
			Identifier: markersutil.KubebuilderUniqueItemsMarker,
		},

		{
			Identifier: markersutil.KubebuilderItemsEnumMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsFormatMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsMaxLengthMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsMaxItemsMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsMaxPropertiesMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsMaximumMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsMinLengthMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsMinItemsMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsMinPropertiesMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsMinimumMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsExclusiveMaximumMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsExclusiveMinimumMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsMultipleOfMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsPatternMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsTypeMarker,
		},
		{
			Identifier: markersutil.KubebuilderItemsUniqueItemsMarker,
		},
		// ------
	}
}

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
package ssatags

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
	"sigs.k8s.io/kube-api-linter/pkg/config"
	kubebuildermarkers "sigs.k8s.io/kube-api-linter/pkg/markers"
)

const name = "ssatags"

const (
	listTypeAtomic = "atomic"
	listTypeSet    = "set"
	listTypeMap    = "map"
)

type analyzer struct {
	cfg config.SSATagsListTypeSetUsage
}

func newAnalyzer(cfg config.SSATagsConfig) *analysis.Analyzer {
	defaultConfig(&cfg)

	a := &analyzer{
		cfg: cfg.ListTypeSetUsage,
	}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      "Check that all array types in the API have the SSA tags and the usage of the tags is correct",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspector.Analyzer},
	}
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, stack []ast.Node, jsonTagInfo extractjsontags.FieldTagInfo, markersAccess markers.Markers) {
		a.checkField(pass, field, markersAccess)
	})

	return nil, nil //nolint:nilnil
}

func (a *analyzer) checkField(pass *analysis.Pass, field *ast.Field, markersAccess markers.Markers) {
	if _, ok := field.Type.(*ast.ArrayType); !ok {
		return
	}

	fieldMarkers := markersAccess.FieldMarkers(field)
	if fieldMarkers == nil {
		return
	}

	fieldName := utils.FieldName(field)
	listTypeMarkers := fieldMarkers.Get(kubebuildermarkers.KubebuilderListTypeMarker)

	if len(listTypeMarkers) == 0 {
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("%s should have a listType marker (atomic, set, or map)", fieldName),
		})

		return
	}

	for _, marker := range listTypeMarkers {
		listType := strings.TrimSpace(marker.Expressions[""])
		a.checkListTypeMarker(pass, listType, field)

		if listType == listTypeMap {
			a.checkListTypeMap(pass, fieldMarkers, field)
		}

		if listType == listTypeSet {
			a.checkListTypeSet(pass, marker, field)
		}
	}
}

func (a *analyzer) checkListTypeMarker(pass *analysis.Pass, listType string, field *ast.Field) {
	fieldName := utils.FieldName(field)

	if !validListType(listType) {
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("%s has invalid listType %q, must be one of: atomic, set, map", fieldName, listType),
		})

		return
	}
}

func (a *analyzer) checkListTypeMap(pass *analysis.Pass, fieldMarkers markers.MarkerSet, field *ast.Field) {
	listMapKeyMarkers := fieldMarkers.Get(kubebuildermarkers.KubebuilderListMapKeyMarker)
	fieldName := utils.FieldName(field)

	if len(listMapKeyMarkers) == 0 {
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("%s with listType=map must have at least one listMapKey marker", fieldName),
		})
	}
}

func (a *analyzer) checkListTypeSet(pass *analysis.Pass, listTypeMarker markers.Marker, field *ast.Field) {
	if a.cfg == config.SSATagsListTypeSetUsageIgnore {
		return
	}

	fieldName := utils.FieldName(field)

	diagnostic := analysis.Diagnostic{
		Pos:     field.Pos(),
		Message: fmt.Sprintf("listType=set is forbidden, use listType=%s or listType=%s instead", listTypeAtomic, listTypeMap),
	}

	if a.cfg == config.SSATagsListTypeSetUsageWarn {
		pass.Report(diagnostic)
	}

	if a.cfg == config.SSATagsListTypeSetUsageSuggestFix {
		diagnostic.SuggestedFixes = []analysis.SuggestedFix{
			{
				Message: fmt.Sprintf("Remove listType=set from %s", fieldName),
				TextEdits: []analysis.TextEdit{
					{
						Pos:     listTypeMarker.Pos,
						End:     listTypeMarker.End,
						NewText: []byte{},
					},
				},
			},
		}
		pass.Report(diagnostic)
	}
}

func validListType(listType string) bool {
	switch listType {
	case listTypeAtomic, listTypeSet, listTypeMap:
		return true
	default:
		return false
	}
}

func defaultConfig(cfg *config.SSATagsConfig) {
	if cfg.ListTypeSetUsage == "" {
		cfg.ListTypeSetUsage = config.SSATagsListTypeSetUsageWarn
	}
}

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
package maxlength

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
	name = "maxlength"
)

func init() {
	markershelper.DefaultRegistry().Register(
		markers.K8sMaxLengthMarker,
		markers.K8sMaxBytesMarker,
		markers.K8sMaxItemsMarker,
		markers.K8sMaxPropertiesMarker,
		markers.K8sEnumMarker,
		markers.K8sFormatMarker,
	)
}

// Analyzer is the analyzer for the maxlength package with the default (kubebuilder-preferred) configuration.
// It checks that strings and arrays have maximum lengths and maximum items respectively.
var Analyzer = newAnalyzer(nil)

type analyzer struct {
	preferredMaxLengthMarker     string
	preferredMaxItemsMarker      string
	preferredMaxPropertiesMarker string
}

// newAnalyzer creates a new analysis.Analyzer for the given MaxLengthConfig.
// If cfg is nil, the default configuration is used (kubebuilder markers preferred).
func newAnalyzer(cfg *MaxLengthConfig) *analysis.Analyzer {
	if cfg == nil {
		cfg = &MaxLengthConfig{}
	}

	defaultConfig(cfg)

	a := &analyzer{
		preferredMaxLengthMarker:     cfg.PreferredMaxLengthMarker,
		preferredMaxItemsMarker:      cfg.PreferredMaxItemsMarker,
		preferredMaxPropertiesMarker: cfg.PreferredMaxPropertiesMarker,
	}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      "Checks that all string fields are marked with a maximum length, arrays are marked with max items, and maps are marked with max properties.",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspector.Analyzer},
	}
}

func defaultConfig(cfg *MaxLengthConfig) {
	if cfg.PreferredMaxLengthMarker == "" {
		cfg.PreferredMaxLengthMarker = markers.KubebuilderMaxLengthMarker
	}

	if cfg.PreferredMaxItemsMarker == "" {
		cfg.PreferredMaxItemsMarker = markers.KubebuilderMaxItemsMarker
	}

	if cfg.PreferredMaxPropertiesMarker == "" {
		cfg.PreferredMaxPropertiesMarker = markers.KubebuilderMaxPropertiesMarker
	}
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, _ extractjsontags.FieldTagInfo, markersAccess markershelper.Markers, qualifiedFieldName string) {
		a.checkField(pass, field, markersAccess, qualifiedFieldName)
	})

	return nil, nil //nolint:nilnil
}

func (a *analyzer) checkField(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers, qualifiedFieldName string) {
	prefix := fmt.Sprintf("field %s", qualifiedFieldName)

	a.checkTypeExpr(pass, field.Type, field, nil, markersAccess, prefix, a.preferredMaxLengthMarker, a.needsStringMaxLength)
}

func (a *analyzer) checkIdent(pass *analysis.Pass, ident *ast.Ident, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix, marker string, needsMaxLength func(markershelper.MarkerSet) bool) {
	if utils.IsBasicType(pass, ident) { // Built-in type
		checkString(pass, ident, node, aliases, markersAccess, prefix, marker, needsMaxLength)

		return
	}

	tSpec, ok := utils.LookupTypeSpec(pass, ident)
	if !ok {
		return
	}

	a.checkTypeSpec(pass, tSpec, node, append(aliases, tSpec), markersAccess, fmt.Sprintf("%s type", prefix), marker, needsMaxLength)
}

func checkString(pass *analysis.Pass, ident *ast.Ident, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix, marker string, needsMaxLength func(markershelper.MarkerSet) bool) {
	if ident.Name != "string" {
		return
	}

	markerSet := getCombinedMarkers(markersAccess, node, aliases)

	if needsMaxLength(markerSet) {
		pass.Reportf(node.Pos(), "%s must have a maximum length, add %s marker", prefix, marker)
	}
}

func (a *analyzer) checkTypeSpec(pass *analysis.Pass, tSpec *ast.TypeSpec, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix, marker string, needsMaxLength func(markershelper.MarkerSet) bool) {
	if tSpec.Name == nil {
		return
	}

	typeName := tSpec.Name.Name
	prefix = fmt.Sprintf("%s %s", prefix, typeName)

	a.checkTypeExpr(pass, tSpec.Type, node, aliases, markersAccess, prefix, marker, needsMaxLength)
}

func (a *analyzer) checkTypeExpr(pass *analysis.Pass, typeExpr ast.Expr, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix, marker string, needsMaxLength func(markershelper.MarkerSet) bool) {
	switch typ := typeExpr.(type) {
	case *ast.Ident:
		a.checkIdent(pass, typ, node, aliases, markersAccess, prefix, marker, needsMaxLength)
	case *ast.StarExpr:
		a.checkTypeExpr(pass, typ.X, node, aliases, markersAccess, prefix, marker, needsMaxLength)
	case *ast.ArrayType:
		a.checkArrayType(pass, typ, node, aliases, markersAccess, prefix)
	case *ast.MapType:
		a.checkMapType(pass, typ, node, aliases, markersAccess, prefix)
	}
}

func (a *analyzer) checkArrayType(pass *analysis.Pass, arrayType *ast.ArrayType, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix string) {
	if arrayType.Elt != nil {
		if ident, ok := arrayType.Elt.(*ast.Ident); ok {
			if ident.Name == "byte" {
				// byte slices are a special case as they are treated as strings.
				// Pretend the ident is a string so that checkString can process it as expected.
				// In DV, +k8s:maxBytes (not +k8s:maxLength) is the correct tag for []byte fields,
				// as it constrains byte count rather than Unicode character count.
				i := &ast.Ident{
					NamePos: ident.NamePos,
					Name:    "string",
				}
				checkString(pass, i, node, aliases, markersAccess, prefix, a.preferredMaxLengthMarker, a.needsByteSliceMaxLength)

				return
			}

			a.checkArrayElementIdent(pass, ident, node, aliases, markersAccess, fmt.Sprintf("%s array element", prefix))
		}
	}

	markerSet := getCombinedMarkers(markersAccess, node, aliases)

	if !markerSet.Has(markers.KubebuilderMaxItemsMarker) && !markerSet.Has(markers.K8sMaxItemsMarker) {
		pass.Reportf(node.Pos(), "%s must have a maximum items, add %s marker", prefix, a.preferredMaxItemsMarker)
	}
}

// checkMapType checks that map[string]V fields have a maximum properties marker.
// Only string-keyed maps are checked, mirroring the DV constraint that +k8s:maxProperties
// applies to map[string]V types only.
func (a *analyzer) checkMapType(pass *analysis.Pass, mapType *ast.MapType, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix string) {
	// DV +k8s:maxProperties only supports string-keyed maps.
	// Skip non-string-keyed maps (e.g. map[int]string) to avoid false positives.
	if keyIdent, ok := mapType.Key.(*ast.Ident); !ok || keyIdent.Name != "string" {
		return
	}

	markerSet := getCombinedMarkers(markersAccess, node, aliases)

	if !markerSet.Has(markers.KubebuilderMaxPropertiesMarker) && !markerSet.Has(markers.K8sMaxPropertiesMarker) {
		pass.Reportf(node.Pos(), "%s must have a maximum properties, add %s marker", prefix, a.preferredMaxPropertiesMarker)
	}
}

func (a *analyzer) checkArrayElementIdent(pass *analysis.Pass, ident *ast.Ident, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix string) {
	if ident.Obj == nil { // Built-in type
		checkString(pass, ident, node, aliases, markersAccess, prefix, markers.KubebuilderItemsMaxLengthMarker, a.needsItemsMaxLength)

		return
	}

	tSpec, ok := ident.Obj.Decl.(*ast.TypeSpec)
	if !ok {
		return
	}

	// If the array element wasn't directly a string, allow a string alias to be used
	// with either the items style markers or the on alias style markers.
	a.checkTypeSpec(pass, tSpec, node, append(aliases, tSpec), markersAccess, fmt.Sprintf("%s type", prefix), markers.KubebuilderMaxLengthMarker, func(ms markershelper.MarkerSet) bool {
		return a.needsStringMaxLength(ms) && a.needsItemsMaxLength(ms)
	})
}

func getCombinedMarkers(markersAccess markershelper.Markers, node ast.Node, aliases []*ast.TypeSpec) markershelper.MarkerSet {
	base := markershelper.NewMarkerSet(getMarkers(markersAccess, node).UnsortedList()...)

	for _, alias := range aliases {
		base.Insert(getMarkers(markersAccess, alias).UnsortedList()...)
	}

	return base
}

func getMarkers(markersAccess markershelper.Markers, node ast.Node) markershelper.MarkerSet {
	switch t := node.(type) {
	case *ast.Field:
		return markersAccess.FieldMarkers(t)
	case *ast.TypeSpec:
		return markersAccess.TypeMarkers(t)
	}

	return nil
}

// needsStringMaxLength returns true if the field needs a maximum length.
// Returns false if either a kubebuilder or DV max-length marker is already present,
// or if the field is an enum, or is formatted as a date, date-time or duration.
func (a *analyzer) needsStringMaxLength(markerSet markershelper.MarkerSet) bool {
	switch {
	case markerSet.Has(markers.KubebuilderMaxLengthMarker),
		markerSet.Has(markers.K8sMaxLengthMarker),
		markerSet.Has(markers.KubebuilderEnumMarker),
		markerSet.Has(markers.K8sEnumMarker),
		markerSet.HasWithValue(kubebuilderFormatWithValue("date")),
		markerSet.HasWithValue(kubebuilderFormatWithValue("date-time")),
		markerSet.HasWithValue(kubebuilderFormatWithValue("duration")),
		markerSet.HasWithValue(k8sFormatWithValue("date")),
		markerSet.HasWithValue(k8sFormatWithValue("date-time")),
		markerSet.HasWithValue(k8sFormatWithValue("duration")):
		return false
	}

	return true
}

// needsByteSliceMaxLength is like needsStringMaxLength but also accepts +k8s:maxBytes,
// which is the DV-correct marker for byte-length constraints on []byte fields.
// (+k8s:maxLength counts Unicode characters; +k8s:maxBytes counts raw bytes.)
func (a *analyzer) needsByteSliceMaxLength(markerSet markershelper.MarkerSet) bool {
	if markerSet.Has(markers.K8sMaxBytesMarker) {
		return false
	}

	return a.needsStringMaxLength(markerSet)
}

func (a *analyzer) needsItemsMaxLength(markerSet markershelper.MarkerSet) bool {
	switch {
	case markerSet.Has(markers.KubebuilderItemsMaxLengthMarker),
		markerSet.Has(markers.KubebuilderItemsEnumMarker),
		markerSet.HasWithValue(kubebuilderItemsFormatWithValue("date")),
		markerSet.HasWithValue(kubebuilderItemsFormatWithValue("date-time")),
		markerSet.HasWithValue(kubebuilderItemsFormatWithValue("duration")):
		return false
	}

	return true
}

func kubebuilderFormatWithValue(value string) string {
	return fmt.Sprintf("%s:=%s", markers.KubebuilderFormatMarker, value)
}

func kubebuilderItemsFormatWithValue(value string) string {
	return fmt.Sprintf("%s:=%s", markers.KubebuilderItemsFormatMarker, value)
}

func k8sFormatWithValue(value string) string {
	return fmt.Sprintf("%s:=%s", markers.K8sFormatMarker, value)
}

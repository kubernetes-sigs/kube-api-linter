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
package minlength

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
	name = "minlength"
)

// NewAnalyzer creates a new analysis.Analyzer for the minlength
// linter based on the provided MinLengthConfig
func newAnalyzer(cfg *Config) *analysis.Analyzer {
	if cfg == nil {
		cfg = &Config{}
	}

	defaultConfig(cfg)

	a := &analyzer{
		preferredSuggestionMarkerType: cfg.PreferredSuggestionMarkerType,
	}

	analyzer := &analysis.Analyzer{
		Name:     name,
		Doc:      "Checks that all strings formatted fields are marked with a minimum length, and that arrays are marked with min items, maps are marked with min properties, and structs that do not have required fields are marked with min properties",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspector.Analyzer},
	}

	return analyzer
}

func defaultConfig(cfg *Config) {
	if cfg.PreferredSuggestionMarkerType == "" {
		cfg.PreferredSuggestionMarkerType = PreferredSuggestionMarkerTypeKubebuilder
	}
}

type analyzer struct {
	preferredSuggestionMarkerType PreferredSuggestionMarkerType
}

type markerPreferences struct {
	MinLengthMarker      string
	MinPropertiesMarker  string
	MinItemsMarker       string
	ItemsMinLengthMarker string
}

func markerPreferencesForConfigType(markerType PreferredSuggestionMarkerType) markerPreferences {
	switch markerType {
	case PreferredSuggestionMarkerTypeKubebuilder:
		return markerPreferences{
			MinLengthMarker:      markers.KubebuilderMinLengthMarker,
			MinPropertiesMarker:  markers.KubebuilderMinPropertiesMarker,
			MinItemsMarker:       markers.KubebuilderMinItemsMarker,
			ItemsMinLengthMarker: markers.KubebuilderItemsMinLengthMarker,
		}
	case PreferredSuggestionMarkerTypeDeclarativeValidation:
		return markerPreferences{
			MinLengthMarker: markers.K8sMinLengthMarker,
			// TODO: This marker needs to be added to DV marker set
			MinPropertiesMarker: markers.K8sMinPropertiesMarker,
			// TODO: This marker needs to be added to DV marker set
			MinItemsMarker:       markers.K8sMinItemsMarker,
			ItemsMinLengthMarker: fmt.Sprintf("%s=+%s", markers.K8sEachValMarker, markers.K8sMinLengthMarker),
		}
	default:
		// default to Kubebuilder markers for backwards compatibility
		return markerPreferences{
			MinLengthMarker:      markers.KubebuilderMinLengthMarker,
			MinPropertiesMarker:  markers.KubebuilderMinPropertiesMarker,
			MinItemsMarker:       markers.KubebuilderMinItemsMarker,
			ItemsMinLengthMarker: markers.KubebuilderItemsMinLengthMarker,
		}
	}
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, _ extractjsontags.FieldTagInfo, markersAccess markershelper.Markers, qualifiedFieldName string) {
		checkField(pass, field, markersAccess, qualifiedFieldName, markerPreferencesForConfigType(a.preferredSuggestionMarkerType))
	})

	return nil, nil //nolint:nilnil
}

func checkField(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers, qualifiedFieldName string, markerPreference markerPreferences) {
	prefix := fmt.Sprintf("field %s", qualifiedFieldName)

	checkTypeExpr(pass, field.Type, field, nil, markersAccess, prefix, markerPreference, needsStringMinLength)
}

func checkIdent(pass *analysis.Pass, ident *ast.Ident, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix string, markerPreference markerPreferences, needsMinLength func(markershelper.MarkerSet) bool) {
	if utils.IsBasicType(pass, ident) { // Built-in type
		checkString(pass, ident, node, aliases, markersAccess, prefix, markerPreference.MinLengthMarker, needsMinLength)

		return
	}

	tSpec, ok := utils.LookupTypeSpec(pass, ident)
	if !ok {
		return
	}

	checkTypeSpec(pass, tSpec, node, append(aliases, tSpec), markersAccess, fmt.Sprintf("%s type", prefix), markerPreference, needsMinLength)
}

func checkString(pass *analysis.Pass, ident *ast.Ident, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix, marker string, needsMinLength func(markershelper.MarkerSet) bool) {
	if ident.Name != "string" {
		return
	}

	markers := getCombinedMarkers(markersAccess, node, aliases)

	if needsMinLength(markers) {
		pass.Reportf(node.Pos(), "%s must have a minimum length, add %s marker", prefix, marker)
	}
}

func checkTypeSpec(pass *analysis.Pass, tSpec *ast.TypeSpec, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix string, markerPreference markerPreferences, needsMinLength func(markershelper.MarkerSet) bool) {
	if tSpec.Name == nil {
		return
	}

	typeName := tSpec.Name.Name
	prefix = fmt.Sprintf("%s %s", prefix, typeName)

	checkTypeExpr(pass, tSpec.Type, node, aliases, markersAccess, prefix, markerPreference, needsMinLength)
}

func checkTypeExpr(pass *analysis.Pass, typeExpr ast.Expr, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix string, markerPreference markerPreferences, needsMinLength func(markershelper.MarkerSet) bool) {
	switch typ := typeExpr.(type) {
	case *ast.Ident:
		checkIdent(pass, typ, node, aliases, markersAccess, prefix, markerPreference, needsMinLength)
	case *ast.StarExpr:
		checkTypeExpr(pass, typ.X, node, aliases, markersAccess, prefix, markerPreference, needsMinLength)
	case *ast.ArrayType:
		checkArrayType(pass, typ, node, aliases, markersAccess, prefix, markerPreference)
	case *ast.MapType:
		checkMapType(pass, node, aliases, markersAccess, prefix, markerPreference)
	case *ast.StructType:
		checkStructType(pass, typ, node, aliases, markersAccess, prefix, markerPreference)
	}
}

func checkArrayType(pass *analysis.Pass, arrayType *ast.ArrayType, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix string, markerPreference markerPreferences) {
	if arrayType.Elt != nil {
		if ident, ok := arrayType.Elt.(*ast.Ident); ok {
			if ident.Name == "byte" {
				// byte slices are a special case as they are treated as strings.
				// Pretend the ident is a string so that checkString can process it as expected.
				// TODO: is this true of DV validations?
				i := &ast.Ident{
					NamePos: ident.NamePos,
					Name:    "string",
				}
				checkString(pass, i, node, aliases, markersAccess, prefix, markerPreference.MinLengthMarker, needsStringMinLength)

				return
			}

			checkArrayElementIdent(pass, ident, node, aliases, markersAccess, fmt.Sprintf("%s array element", prefix), markerPreference)
		}
	}

	markerSet := getCombinedMarkers(markersAccess, node, aliases)

	if !markerSet.Has(markers.KubebuilderMinItemsMarker) && !markerSet.Has(markers.K8sMinItemsMarker) {
		pass.Reportf(node.Pos(), "%s must have a minimum items, add %s marker", prefix, markerPreference.MinItemsMarker)
	}
}

func checkArrayElementIdent(pass *analysis.Pass, ident *ast.Ident, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix string, markerPreference markerPreferences) {
	if ident.Obj == nil { // Built-in type
		checkString(pass, ident, node, aliases, markersAccess, prefix, markerPreference.ItemsMinLengthMarker, needsItemsMinLength)

		return
	}

	tSpec, ok := ident.Obj.Decl.(*ast.TypeSpec)
	if !ok {
		return
	}

	// If the array element wasn't directly a string, allow a string alias to be used
	// with either the items style markers or the on alias style markers.
	checkTypeSpec(pass, tSpec, node, append(aliases, tSpec), markersAccess, fmt.Sprintf("%s type", prefix), markerPreference, func(ms markershelper.MarkerSet) bool {
		return needsStringMinLength(ms) && needsItemsMinLength(ms)
	})
}

func checkMapType(pass *analysis.Pass, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix string, markerPreference markerPreferences) {
	markerSet := getCombinedMarkers(markersAccess, node, aliases)

	if !markerSet.Has(markers.KubebuilderMinPropertiesMarker) && !markerSet.Has(markers.K8sMinPropertiesMarker) {
		pass.Reportf(node.Pos(), "%s must have a minimum properties, add %s marker", prefix, markerPreference.MinPropertiesMarker)
	}
}

func checkStructType(pass *analysis.Pass, structType *ast.StructType, node ast.Node, aliases []*ast.TypeSpec, markersAccess markershelper.Markers, prefix string, markerPreference markerPreferences) {
	markerSet := getCombinedMarkers(markersAccess, node, aliases)

	minProperties, err := utils.GetMinProperties(markerSet)
	if err != nil {
		pass.Reportf(node.Pos(), "could not get min properties for struct: %v", err)
		return
	}

	if minProperties != nil {
		// There's already a min properties specified.
		return
	}

	// Check if the struct has union markers that satisfy the required constraint
	if markerSet.Has(markers.KubebuilderExactlyOneOf) || markerSet.Has(markers.KubebuilderAtLeastOneOfMarker) {
		// ExactlyOneOf / AtLeastOneOf markers enforce that at least one field is required,
		// this means that `{}` is not valid.
		return
	}

	for _, field := range structType.Fields.List {
		if utils.IsFieldRequired(field, markersAccess) {
			// The struct has at least one required field,
			// this means that `{}` is not valid.
			return
		}
	}

	// The field does not have a min properties, and does not have any required fields.
	pass.Reportf(node.Pos(), "%s must have a minimum properties, add %s marker", prefix, markerPreference.MinPropertiesMarker)
}

func getCombinedMarkers(markersAccess markershelper.Markers, node ast.Node, aliases []*ast.TypeSpec) markershelper.MarkerSet {
	base := markershelper.NewMarkerSet(getMarkers(markersAccess, node).UnsortedList()...)

	for _, a := range aliases {
		base.Insert(getMarkers(markersAccess, a).UnsortedList()...)
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

// needsMinLength returns true if the field needs a minimum length.
// Fields do not need a minimum length if they are already marked with a minimum length,
// or if they are an enum, or if they are a format that already enforces a minimum length.
func needsStringMinLength(markerSet markershelper.MarkerSet) bool {
	return !hasKubebuilderMinLengthOrEquivalentMarker(markerSet) &&
		!hasDeclarativeValidationMinLengthOrEquivalentMarker(markerSet)
}

func needsItemsMinLength(markerSet markershelper.MarkerSet) bool {
	return !hasKubebuilderItemsMinLengthOrEquivalentMarker(markerSet) &&
		!hasDeclarativeEachValMinLengthOrEquivalentMarker(markerSet)
}

func hasKubebuilderMinLengthOrEquivalentMarker(markerSet markershelper.MarkerSet) bool {
	if markerSet.Has(markers.KubebuilderMinLengthMarker) || markerSet.Has(markers.KubebuilderEnumMarker) {
		return true
	}

	for _, value := range knownKubebuilderMinimumLengthConstrainedFormats {
		if markerSet.HasWithValue(kubebuilderFormatWithValue(value)) {
			return true
		}
	}

	return false
}

func hasKubebuilderItemsMinLengthOrEquivalentMarker(markerSet markershelper.MarkerSet) bool {
	if markerSet.Has(markers.KubebuilderItemsMinLengthMarker) || markerSet.Has(markers.KubebuilderItemsEnumMarker) {
		return true
	}

	for _, value := range knownKubebuilderMinimumLengthConstrainedFormats {
		if markerSet.HasWithValue(kubebuilderItemsFormatWithValue(value)) {
			return true
		}
	}

	return false
}

func hasDeclarativeValidationMinLengthOrEquivalentMarker(markerSet markershelper.MarkerSet) bool {
	if markerSet.Has(markers.K8sMinLengthMarker) || markerSet.Has(markers.K8sEnumMarker) {
		return true
	}

	for _, value := range knownDeclarativeValidationMinimumLengthConstrainedFormats {
		if markerSet.HasWithValue(fmt.Sprintf("%s=%s", markers.K8sFormatMarker, value)) {
			return true
		}
	}

	return false
}

func hasDeclarativeEachValMinLengthOrEquivalentMarker(markerSet markershelper.MarkerSet) bool {
	if markerSet.HasWithPayloadFunc(markers.K8sEachValMarker, markershelper.WithMarkerIdentifier(markers.K8sMinLengthMarker)) ||
		markerSet.HasWithPayloadFunc(markers.K8sEachValMarker, markershelper.WithMarkerIdentifier(markers.K8sEnumMarker)) {
		return true
	}

	for _, value := range knownKubebuilderMinimumLengthConstrainedFormats {
		formatMarker := fmt.Sprintf("%s=%s", markers.K8sFormatMarker, value)
		if markerSet.HasWithValue(fmt.Sprintf("%s=+%s", markers.K8sEachValMarker, formatMarker)) {
		
			return true
		}
	}

	return false
}

func kubebuilderFormatWithValue(value string) string {
	return fmt.Sprintf("%s:=%s", markers.KubebuilderFormatMarker, value)
}

func kubebuilderItemsFormatWithValue(value string) string {
	return fmt.Sprintf("%s:=%s", markers.KubebuilderItemsFormatMarker, value)
}

// knownKubebuilderMinimumLengthConstrainedFormats is a slice of the known
// formats that have minimum length constraints associated
// with them, specifically the formats supported by
// https://github.com/kubernetes/kube-openapi/tree/master/pkg/validation/strfmt
var knownKubebuilderMinimumLengthConstrainedFormats = []string{
	// date is an RFC3339 date format and is processed by
	// https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/date.go#L30
	// An inherent minimum length is enforced by this format.
	"date",

	// date-time is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/time.go#L67
	// which ensures an input adheres to one of RFC3339 micro, RFC3339 millis, RFC3339, RFC3339 nano, or ISO8601 local time formats.
	// An inherent minimum length is enforced by these formats.
	"date-time",

	// duration is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/duration.go#L59
	// which ensures that an input adheres to the format '{digits}{unit}', inherently
	// enforcing a minimum length of 2 (i.e '1s').
	"duration",

	// bsonobjectid is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/bson.go#L28
	// which ensures that an input is a valid hex string that is exactly 12 characters in length.
	"bsonobjectid",

	// uri is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L149
	// which uses Golang's net/url.ParseRequestURI function to validate the input.
	// net/url.ParseRequestURI does not accept an empty string as a valid input, inherently enforcing a minimum length.
	"uri",

	// email is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L152
	// which uses Golang's net/mail.ParseAddress function to validate the input.
	// net/mail.ParseAddress does not accept an emptry string as a valid input, inherently enforcing a minimum length.
	"email",

	// hostname is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L155
	// which enforces that the input is a valid representation of an internet hostname as defined by RFC1034 section 3.1.
	// This format inherently enforces a minimum length.
	"hostname",

	// ipv4 is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L158
	// which inherently enforces a minimum length.
	"ipv4",

	// ipv6 is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L161
	// which inherently enforces a minimum length.
	"ipv6",

	// cidr is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L164
	// which inherently enforces a minimum length by validating the input against CIDR notation formats
	// as defined by RFC 4632 and RFC 4291.
	"cidr",

	// mac is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L167
	// which inherently enforces a minimum length via Golang's net.ParseMAC function.
	"mac",

	// uuid is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L170
	// which ensures that the input is a valid UUID, inherently enforcing a minimum length.
	"uuid",

	// uuid3 is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L173
	// which ensures that the input is a valid UUID3, inherently enforcing a minimum length.
	"uuid3",

	// uuid4 is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L176
	// which ensures that the input is a valid UUID4, inherently enforcing a minimum length.
	"uuid4",

	// uuid5 is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L179
	// which ensures that the input is a valid UUID5, inherently enforcing a minimum length.
	"uuid5",

	// isbn is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L182
	// which ensures that the input is either a valid ISBN10 or ISBN13 value, inherently enforcing a minimum length.
	"isbn",

	// isbn10 is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L185
	// which ensures that the input is a valid ISBN10, inherently enforcing a minimum length.
	"isbn10",

	// isbn13 is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L188
	// which ensures that the input is a valid ISBN13, inherently enforcing a minimum length.
	"isbn13",

	// creditcard is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L191
	// which ensures that the input is a valid credit card number, inherently enforcing a minimum length.
	"creditcard",

	// ssn is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L194
	// which ensures that the input is a valid social security number, inherently enforcing a minimum length.
	"ssn",

	// hexcolor is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L197
	// which ensures that the input is a valid hexcolor, inherently enforcing a minimum length.
	"hexcolor",

	// rgbcolor is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L200
	// which ensures that the input is a valid rgbcolor, inherently enforcing a minimum length.
	"rgbcolor",

	// byte is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L203
	// which ensures that the input is a valid base64 string, inherently enforcing a minimum length.
	"byte",

	// k8s-short-name is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/kubernetes-extensions.go#L83
	// which ensures that the input is a valid Kubernetes short name. This is a variant of RFC1123 DNS Label names with the exception that only
	// lowercase letters are allowed.
	// This inherently enforces a minimum length.
	"k8s-short-name",

	// k8s-long-name is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/kubernetes-extensions.go#L140
	// which ensures that the input is a valid Kubernetes long name. This is a variant of RFC1123 DNS Subdomain names with the exception that only
	// lowercase letters are allowed and there is no 63 character limit for each of the dot-separated DNS labels.
	// This inherently enforces a minimum length.
	"k8s-long-name",
}

// knownDeclarativeValidationMinimumLengthConstrainedFormats is a slice of the known
// formats that have minimum length constraints associated
// with them, specifically the formats supported by
// https://github.com/kubernetes/kube-openapi/tree/master/pkg/validation/strfmt
// For Kubernetes native types, they may or may not
// have use cases for validating known OpenAPI spec formats.
// For future usage as a replacement of the `kubebuilder:validation:Format` marker,
// and supporting native types, this is a superset of knownKubebuilderMinimumLengthConstrainedFormats.
var knownDeclarativeValidationMinimumLengthConstrainedFormats = []string{
	// OpenAPI-available formats
	// ---

	// date is an RFC3339 date format and is processed by
	// https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/date.go#L30
	// An inherent minimum length is enforced by this format.
	"date",

	// date-time is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/time.go#L67
	// which ensures an input adheres to one of RFC3339 micro, RFC3339 millis, RFC3339, RFC3339 nano, or ISO8601 local time formats.
	// An inherent minimum length is enforced by these formats.
	"date-time",

	// duration is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/duration.go#L59
	// which ensures that an input adheres to the format '{digits}{unit}', inherently
	// enforcing a minimum length of 2 (i.e '1s').
	"duration",

	// bsonobjectid is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/bson.go#L28
	// which ensures that an input is a valid hex string that is exactly 12 characters in length.
	"bsonobjectid",

	// uri is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L149
	// which uses Golang's net/url.ParseRequestURI function to validate the input.
	// net/url.ParseRequestURI does not accept an empty string as a valid input, inherently enforcing a minimum length.
	"uri",

	// email is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L152
	// which uses Golang's net/mail.ParseAddress function to validate the input.
	// net/mail.ParseAddress does not accept an emptry string as a valid input, inherently enforcing a minimum length.
	"email",

	// hostname is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L155
	// which enforces that the input is a valid representation of an internet hostname as defined by RFC1034 section 3.1.
	// This format inherently enforces a minimum length.
	"hostname",

	// ipv4 is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L158
	// which inherently enforces a minimum length.
	"ipv4",

	// ipv6 is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L161
	// which inherently enforces a minimum length.
	"ipv6",

	// cidr is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L164
	// which inherently enforces a minimum length by validating the input against CIDR notation formats
	// as defined by RFC 4632 and RFC 4291.
	"cidr",

	// mac is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L167
	// which inherently enforces a minimum length via Golang's net.ParseMAC function.
	"mac",

	// uuid is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L170
	// which ensures that the input is a valid UUID, inherently enforcing a minimum length.
	"uuid",

	// uuid3 is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L173
	// which ensures that the input is a valid UUID3, inherently enforcing a minimum length.
	"uuid3",

	// uuid4 is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L176
	// which ensures that the input is a valid UUID4, inherently enforcing a minimum length.
	"uuid4",

	// uuid5 is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L179
	// which ensures that the input is a valid UUID5, inherently enforcing a minimum length.
	"uuid5",

	// isbn is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L182
	// which ensures that the input is either a valid ISBN10 or ISBN13 value, inherently enforcing a minimum length.
	"isbn",

	// isbn10 is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L185
	// which ensures that the input is a valid ISBN10, inherently enforcing a minimum length.
	"isbn10",

	// isbn13 is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L188
	// which ensures that the input is a valid ISBN13, inherently enforcing a minimum length.
	"isbn13",

	// creditcard is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L191
	// which ensures that the input is a valid credit card number, inherently enforcing a minimum length.
	"creditcard",

	// ssn is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L194
	// which ensures that the input is a valid social security number, inherently enforcing a minimum length.
	"ssn",

	// hexcolor is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L197
	// which ensures that the input is a valid hexcolor, inherently enforcing a minimum length.
	"hexcolor",

	// rgbcolor is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L200
	// which ensures that the input is a valid rgbcolor, inherently enforcing a minimum length.
	"rgbcolor",

	// byte is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/default.go#L203
	// which ensures that the input is a valid base64 string, inherently enforcing a minimum length.
	"byte",

	// k8s-short-name is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/kubernetes-extensions.go#L83
	// which ensures that the input is a valid Kubernetes short name. This is a variant of RFC1123 DNS Label names with the exception that only
	// lowercase letters are allowed.
	// This inherently enforces a minimum length.
	"k8s-short-name",

	// k8s-long-name is processed by https://github.com/kubernetes/kube-openapi/blob/a19766b6e2d458320f2b3d4d14a6533a870fcaba/pkg/validation/strfmt/kubernetes-extensions.go#L140
	// which ensures that the input is a valid Kubernetes long name. This is a variant of RFC1123 DNS Subdomain names with the exception that only
	// lowercase letters are allowed and there is no 63 character limit for each of the dot-separated DNS labels.
	// This inherently enforces a minimum length.
	"k8s-long-name",

	// Declarative Validation only formats
	// ---

	// k8s-extended-resource-name is processed by https://github.com/kubernetes/kubernetes/blob/23ea1ec286387f45f52e1189089bacc0702a00aa/staging/src/k8s.io/apimachinery/pkg/api/validate/strfmt.go#L216
	// which ensures that the input is a valid Kubernetes extended resource name, inherently enforcing a minimum length.
	"k8s-extended-resource-name",

	// k8s-label-key is processed by https://github.com/kubernetes/kubernetes/blob/23ea1ec286387f45f52e1189089bacc0702a00aa/staging/src/k8s.io/apimachinery/pkg/api/validate/strfmt.go#L89
	// which ensures that the input is a valid Kubernetes label key, inherently enforcing a minimum length.
	"k8s-label-key",

	// k8s-long-name-caseless is processed by https://github.com/kubernetes/kubernetes/blob/23ea1ec286387f45f52e1189089bacc0702a00aa/staging/src/k8s.io/apimachinery/pkg/api/validate/strfmt.go#L109
	// which ensures that the input is a valid Kubernetes long name with case insensitivity, inherently enforcing a minimum length.
	// Even though this format is deprecated, there may be existing native types that use this and we
	// shouldn't require an explicit minimum length here if it is already enforced.
	"k8s-long-name-caseless",

	// k8s-resource-fully-qualified-name is processed by https://github.com/kubernetes/kubernetes/blob/23ea1ec286387f45f52e1189089bacc0702a00aa/staging/src/k8s.io/apimachinery/pkg/api/validate/strfmt.go#L285
	// which ensures that the input is a fully qualified Kubernetes resource name, inherently enforcing a minimum length because it
	// does not accept an empty value.
	"k8s-resource-fully-qualified-name",

	// k8s-resource-pool-name is processed by https://github.com/kubernetes/kubernetes/blob/23ea1ec286387f45f52e1189089bacc0702a00aa/staging/src/k8s.io/apimachinery/pkg/api/validate/strfmt.go#L184
	// which ensures that the input is a fully qualified Kubernetes resource name, inherently enforcing a minimum length.
	"k8s-resource-pool-name",

	// k8s-uuid is processed by https://github.com/kubernetes/kubernetes/blob/23ea1ec286387f45f52e1189089bacc0702a00aa/staging/src/k8s.io/apimachinery/pkg/api/validate/strfmt.go#L157
	// which ensures that the input is a valid Kubernetes UUID - which follows the 8-4-4-4-12 format, inherently enfrocing a minimum length of 36.
	"k8s-uuid",
}

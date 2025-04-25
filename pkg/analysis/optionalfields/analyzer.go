package optionalfields

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"strconv"

	"golang.org/x/tools/go/analysis"
	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

const (
	name = "optionalfields"

	optionalMarker            = "optional"
	requiredMarker            = "required"
	kubebuilderOptinalMarker  = "kubebuilder:validation:Optional"
	kubebuilderRequiredMarker = "kubebuilder:validation:Required"

	minItemsMarker      = "kubebuilder:validation:MinItems"
	minLengthMarker     = "kubebuilder:validation:MinLength"
	minPropertiesMarker = "kubebuilder:validation:MinProperties"

	minimumMarker = "kubebuilder:validation:Minimum"
	maximumMarker = "kubebuilder:validation:Maximum"
)

func init() {
	markers.DefaultRegistry().Register(optionalMarker, kubebuilderOptinalMarker)
}

var (
	errMarkerMissingValue = errors.New("marker does not have a value")
)

type analyzer struct {
	pointerPolicy     config.OptionalFieldsPointerPolicy
	pointerPreference config.OptionalFieldsPointerPreference
	omitEmptyPolicy   config.OptionalFieldsOmitEmptyPolicy
}

// newAnalyzer creates a new analyzer.
func newAnalyzer(cfg config.OptionalFieldsConfig) *analysis.Analyzer {
	defaultConfig(&cfg)

	a := &analyzer{
		pointerPolicy:     cfg.Pointers.Policy,
		pointerPreference: cfg.Pointers.Preference,
		omitEmptyPolicy:   cfg.OmitEmpty.Policy,
	}

	return &analysis.Analyzer{
		Name: name,
		Doc: `Checks all optional fields comply with the configured policy.
		Depending on the configuration, this may include checking for the presence of the omitempty tag or
		whether the field is a pointer.
		For structs, this includes checking that if the field is marked as optional, it should be a pointer when it has omitempty.
		Where structs include required fields, they must be a pointer when they themselves are optional.
		`,
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
		a.checkField(pass, field, markersAccess, jsonTagInfo)
	})

	return nil, nil //nolint:nilnil
}

func (a *analyzer) checkField(pass *analysis.Pass, field *ast.Field, markersAccess markers.Markers, jsonTags extractjsontags.FieldTagInfo) {
	if field == nil || len(field.Names) == 0 {
		return
	}

	fieldMarkers := markersAccess.FieldMarkers(field)

	fieldName := field.Names[0].Name

	if !fieldMarkers.Has(optionalMarker) && !fieldMarkers.Has(kubebuilderOptinalMarker) {
		// The field is not marked optional, so we don't need to check it.
		return
	}

	a.checkFieldOmitEmpty(pass, field, fieldName, jsonTags)

	if field.Type == nil {
		// The field has no type? We can't check if it's a pointer.
		return
	}

	a.checkFieldPointers(pass, field, fieldName, markersAccess, jsonTags)
}

func (a *analyzer) checkFieldOmitEmpty(pass *analysis.Pass, field *ast.Field, fieldName string, jsonTags extractjsontags.FieldTagInfo) {
	if a.omitEmptyPolicy == config.OptionalFieldsOmitEmptyPolicyIgnore {
		return
	}

	if !jsonTags.OmitEmpty {
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf("field %s is optional and should be omitempty", fieldName),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "add the omitempty tag",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     jsonTags.Pos + token.Pos(len(jsonTags.Name)),
							NewText: []byte(",omitempty"),
						},
					},
				},
			},
		})
	}
}

func (a *analyzer) checkFieldPointers(pass *analysis.Pass, field *ast.Field, fieldName string, markersAccess markers.Markers, jsonTags extractjsontags.FieldTagInfo) {
	isStarExpr, underlyingType := isStarExpr(field.Type)

	if isPointerType(underlyingType) {
		a.checkFieldPointersPointerTypes(pass, field, fieldName, isStarExpr, markersAccess, jsonTags)

		return
	}

	switch a.pointerPreference {
	case config.OptionalFieldsPointerPreferenceAlways:
		a.checkFieldPointersPreferenceAlways(pass, field, fieldName, isStarExpr)
	case config.OptionalFieldsPointerPreferenceWhenRequired:
		a.checkFieldPointersPreferenceWhenRequired(pass, field, fieldName, isStarExpr, underlyingType, markersAccess, jsonTags)
	}
}

func (a *analyzer) checkFieldPointersPointerTypes(pass *analysis.Pass, field *ast.Field, fieldName string, isStarExpr bool, markersAccess markers.Markers, jsonTags extractjsontags.FieldTagInfo) {
	if a.omitEmptyPolicy == config.OptionalFieldsOmitEmptyPolicyIgnore && !jsonTags.OmitEmpty {
		a.checkFieldPointersPointerTypesWithoutOmitEmpty(pass, field, fieldName, markersAccess)
	}

	// Pointer types should not be pointered again.
	if !isStarExpr {
		return

	}

	reportShouldRemovePointer(pass, field, a.pointerPolicy, fieldName, "field %s is a pointer type and should not be a pointer")
}

func (a *analyzer) checkFieldPointersPointerTypesWithoutOmitEmpty(pass *analysis.Pass, field *ast.Field, fieldName string, markersAccess markers.Markers) {
	fieldMarkers := markersAccess.FieldMarkers(field)

	switch field.Type.(type) {
	case *ast.MapType:
		if fieldMarkers.Has(minPropertiesMarker) {
			for _, marker := range fieldMarkers.Get(minPropertiesMarker) {
				markerValue, err := getMarkerIntegerValue(marker)
				if err != nil {
					pass.Reportf(marker.Pos, fmt.Sprintf("invalid value for %s marker: %v", minPropertiesMarker, err))
					return
				}

				if markerValue > 0 {
					reportShouldRemoveMarker(pass, field, marker, fieldName, "field %s has a greater than zero minimum number of properties without omitempty. The minimum number of properties should be removed.")
				}
			}
		}
	case *ast.ArrayType:
		if fieldMarkers.Has(minItemsMarker) {
			for _, marker := range fieldMarkers.Get(minItemsMarker) {
				markerValue, err := getMarkerIntegerValue(marker)
				if err != nil {
					pass.Reportf(marker.Pos, fmt.Sprintf("invalid value for %s marker: %v", minItemsMarker, err))
					return
				}

				if markerValue > 0 {
					reportShouldRemoveMarker(pass, field, marker, fieldName, "field %s has a greater than zero minimum number of items without omitempty. The minimum number of items should be removed.")
				}
			}
		}
	}

}

func (a *analyzer) checkFieldPointersPreferenceAlways(pass *analysis.Pass, field *ast.Field, fieldName string, isStarExpr bool) {
	if isStarExpr {
		return // The field is already a pointer, so we don't need to do anything.
	}

	reportShouldAddPointer(pass, field, a.pointerPolicy, fieldName, "field %s is optional and should be a pointer")
}

func (a *analyzer) checkFieldPointersPreferenceWhenRequired(pass *analysis.Pass, field *ast.Field, fieldName string, isStarExpr bool, underlyingType ast.Expr, markersAccess markers.Markers, jsonTags extractjsontags.FieldTagInfo) {
	ident, ok := underlyingType.(*ast.Ident)
	if !ok {
		// All fields should be idents, not sure when this would happen?
		return
	}

	if ident.Obj != nil {
		// The field is not a simple type, check the object.
		a.checkFieldPointersPreferenceWhenRequiredIdentObj(pass, field, fieldName, isStarExpr, ident.Obj, markersAccess, jsonTags)
		return
	}

	switch ident.Name {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		a.checkFieldPointersPreferenceWhenRequiredInteger(pass, field, fieldName, isStarExpr, markersAccess, jsonTags)
	case "string":
		a.checkFieldPointersPreferenceWhenRequiredString(pass, field, fieldName, isStarExpr, markersAccess, jsonTags)
	case "bool":
		a.checkFieldPointersPreferenceWhenRequiredBool(pass, field, fieldName, isStarExpr, markersAccess, jsonTags)
	case "float32", "float64":
		a.checkFieldPointersPreferenceWhenRequiredFloat(pass, field, fieldName, isStarExpr, markersAccess, jsonTags)
	default:
		panic(fmt.Sprintf("unknown type: %s", ident.Name))
	}
}

func (a *analyzer) checkFieldPointersPreferenceWhenRequiredIdentObj(pass *analysis.Pass, field *ast.Field, fieldName string, isStarExpr bool, obj *ast.Object, markersAccess markers.Markers, jsonTags extractjsontags.FieldTagInfo) {
	decl, ok := obj.Decl.(*ast.TypeSpec)
	if !ok {
		return
	}

	switch decl.Type.(type) {
	case *ast.StructType:
		a.checkFieldPointersPreferenceWhenRequiredStructType(pass, field, fieldName, isStarExpr, decl.Type.(*ast.StructType), markersAccess, jsonTags)
	default:
		panic(fmt.Sprintf("unknown type: %T", decl.Type))
	}
}

func (a *analyzer) checkFieldPointersPreferenceWhenRequiredStructType(pass *analysis.Pass, field *ast.Field, fieldName string, isStarExpr bool, typeExpr *ast.StructType, markersAccess markers.Markers, jsonTags extractjsontags.FieldTagInfo) {
	if a.omitEmptyPolicy == config.OptionalFieldsOmitEmptyPolicyIgnore && !jsonTags.OmitEmpty {
		a.checkFieldPointersPreferenceWhenRequiredStructTypeWithoutOmitEmpty(pass, field, fieldName, isStarExpr, typeExpr, markersAccess)
		return
	}

	hasRequiredFields := structContainsRequiredFields(typeExpr, markersAccess)

	hasMinimumProperties, err := structHasGreaterThanZeroMinProperties(typeExpr, markersAccess.StructMarkers(typeExpr))
	if err != nil {
		pass.Reportf(field.Pos(), "error checking struct for min properties: %v", err)
		return
	}

	fieldHasMinimumProperties, err := structHasGreaterThanZeroMinProperties(typeExpr, markersAccess.FieldMarkers(field))
	if err != nil {
		pass.Reportf(field.Pos(), "error checking field for min properties: %v", err)
		return
	}

	switch {
	case hasRequiredFields && !isStarExpr:
		reportShouldAddPointer(pass, field, a.pointerPolicy, fieldName, "field %s is optional, but contains required field(s) and should be a pointer")
	case hasMinimumProperties && !isStarExpr, fieldHasMinimumProperties && !isStarExpr:
		reportShouldAddPointer(pass, field, a.pointerPolicy, fieldName, "field %s has a greater than zero minimum number of properties and should be a pointer")
	case isStarExpr && !hasRequiredFields && !hasMinimumProperties && !fieldHasMinimumProperties:
		reportShouldRemovePointer(pass, field, a.pointerPolicy, fieldName, "field %s is optional, and contains no required field(s) and does not need to be a pointer")
	}
}

func (a *analyzer) checkFieldPointersPreferenceWhenRequiredStructTypeWithoutOmitEmpty(pass *analysis.Pass, field *ast.Field, fieldName string, isStarExpr bool, typeExpr *ast.StructType, markersAccess markers.Markers) {
	hasMinimumProperties, err := structHasGreaterThanZeroMinProperties(typeExpr, markersAccess.StructMarkers(typeExpr))
	if err != nil {
		pass.Reportf(field.Pos(), "error checking struct for min properties: %v", err)
		return
	}

	fieldHasMinimumProperties, err := structHasGreaterThanZeroMinProperties(typeExpr, markersAccess.FieldMarkers(field))
	if err != nil {
		pass.Reportf(field.Pos(), "error checking field for min properties: %v", err)
		return
	}

	switch {
	case hasMinimumProperties && isStarExpr, fieldHasMinimumProperties && isStarExpr:
		// The field is already a pointer and should be a pointer, so we don't need to do anything.
	case hasMinimumProperties && !isStarExpr:
		reportShouldAddPointer(pass, field, a.pointerPolicy, fieldName, "field %s has a greater than zero minimum number of properties and should be a pointer")
	case fieldHasMinimumProperties && !isStarExpr:
		reportShouldRemoveMarker(pass, field, markersAccess.FieldMarkers(field).Get(minPropertiesMarker)[0], fieldName, "field %s has a greater than zero minimum number of properties without omitempty. The minimum number of properties should be removed.")
	case isStarExpr:
		// The field is a pointer and should not be a pointer, so we need to remove the pointer.
		reportShouldRemovePointer(pass, field, a.pointerPolicy, fieldName, "field %s is an optional struct without omitempty. It should not be a pointer")
	}
}

func (a *analyzer) checkFieldPointersPreferenceWhenRequiredString(pass *analysis.Pass, field *ast.Field, fieldName string, isStarExpr bool, markersAccess markers.Markers, jsonTags extractjsontags.FieldTagInfo) {
	fieldMarkers := markersAccess.FieldMarkers(field)

	if a.omitEmptyPolicy == config.OptionalFieldsOmitEmptyPolicyIgnore && !jsonTags.OmitEmpty {
		a.checkFieldPointersPreferenceWhenRequiredStringWithoutOmitEmpty(pass, field, fieldName, isStarExpr, fieldMarkers)
		return
	}

	if !fieldMarkers.Has(minLengthMarker) {
		if isStarExpr {
			pass.Reportf(field.Pos(), "field %s is an optional string and does not have a minimum length. Where the difference between omitted and the empty string is significant, set the minmum length to 0", fieldName)
		} else {
			pass.Reportf(field.Pos(), "field %s is an optional string and does not have a minimum length. Either set a minimum length or make %s a pointer where the difference between omitted and the empty string is significant", fieldName, fieldName)
		}

		return
	}

	if fieldMarkers.HasWithValue(minLengthMarker + "=0") {
		if isStarExpr {
			// With a minimum length of 0, the field should be a pointer.
		} else {
			reportShouldAddPointer(pass, field, a.pointerPolicy, fieldName, "field %s has a minimum length of 0. The empty string is a valid value and therefore the field should be a pointer")
		}

		return
	}

	// The field has a non-zero (assumed to be greater than zero) minimum length, so it doesn't need to be a pointer.
	if isStarExpr {
		reportShouldRemovePointer(pass, field, a.pointerPolicy, fieldName, "field %s has a greater than 0 length and does not need to be a pointer")
	}
}

func (a *analyzer) checkFieldPointersPreferenceWhenRequiredStringWithoutOmitEmpty(pass *analysis.Pass, field *ast.Field, fieldName string, isStarExpr bool, fieldMarkers markers.MarkerSet) {
	if fieldMarkers.Has(minLengthMarker) {
		for _, marker := range fieldMarkers.Get(minLengthMarker) {
			markerValue, err := getMarkerIntegerValue(marker)
			if err != nil {
				pass.Reportf(marker.Pos, fmt.Sprintf("invalid value for %s marker: %v", minLengthMarker, err))
				return
			}

			if markerValue > 0 {
				reportShouldRemoveMarker(pass, field, marker, fieldName, "field %s has a greater than zero minimum length without omitempty. The minimum length should be removed.")
			}
		}
	}

	// When non-omitempty, the string field should not be a pointer.
	// The empty string should be a valid/acceptable value.
	if isStarExpr {
		reportShouldRemovePointer(pass, field, a.pointerPolicy, fieldName, "field %s is an optional string without omitempty. It should not be a pointer.")
	}
}

func (a *analyzer) checkFieldPointersPreferenceWhenRequiredBool(pass *analysis.Pass, field *ast.Field, fieldName string, isStarExpr bool, markersAccess markers.Markers, jsonTags extractjsontags.FieldTagInfo) {
	if a.omitEmptyPolicy == config.OptionalFieldsOmitEmptyPolicyIgnore && !jsonTags.OmitEmpty {
		if isStarExpr {
			reportShouldRemovePointer(pass, field, a.pointerPolicy, fieldName, "field %s is an optional boolean without omitempty. It should not be a pointer")
		}
		return
	}

	// Optional bools should always be pointers.
	// When the bool is not a pointer, setting the value to false won't round trip.
	// This could be confusing for users.
	if !isStarExpr {
		reportShouldAddPointer(pass, field, a.pointerPolicy, fieldName, "field %s is an optional boolean and should be a pointer")
	}
}

func (a *analyzer) checkFieldPointersPreferenceWhenRequiredInteger(pass *analysis.Pass, field *ast.Field, fieldName string, isStarExpr bool, markersAccess markers.Markers, jsonTags extractjsontags.FieldTagInfo) {
	fieldMarkers := markersAccess.FieldMarkers(field)

	// These are pointers so that we can determine if they are set, or not.
	var minValue, maxValue *int

	if minimum := fieldMarkers.Get(minimumMarker); len(minimum) > 0 {
		// Where there are multiple minimum markers, we only care about the first one.
		// Other rules should deduplicate them.
		min, err := strconv.Atoi(minimum[0].Expressions[""])
		if err != nil {
			pass.Reportf(field.Pos(), "field %s has a minimum value of %s, but it is not an integer", fieldName, minimum[0].Expressions[""])
			return
		}
		minValue = &min
	}

	if maximum := fieldMarkers.Get(maximumMarker); len(maximum) > 0 {
		// Where there are multiple maximum markers, we only care about the first one.
		// Other rules should deduplicate them.
		max, err := strconv.Atoi(maximum[0].Expressions[""])
		if err != nil {
			pass.Reportf(field.Pos(), "field %s has a maximum value of %s, but it is not an integer", fieldName, maximum[0].Expressions[""])
			return
		}
		maxValue = &max
	}

	if a.omitEmptyPolicy == config.OptionalFieldsOmitEmptyPolicyIgnore && !jsonTags.OmitEmpty {
		a.checkFieldPointersPreferenceWhenRequiredIntegerWithoutOmitEmpty(pass, field, fieldName, minValue, maxValue, isStarExpr, fieldMarkers)
		return
	}

	switch {
	case minValue != nil && *minValue > 0:
		if isStarExpr {
			reportShouldRemovePointer(pass, field, a.pointerPolicy, fieldName, "field %s has a greater than 0 minimum value and does not need to be a pointer")
		}
	case maxValue != nil && *maxValue < 0:
		if isStarExpr {
			reportShouldRemovePointer(pass, field, a.pointerPolicy, fieldName, "field %s has a negative maximum value and does not need to be a pointer")
		}
	case minValue != nil && *minValue == 0,
		maxValue != nil && *maxValue == 0,
		minValue != nil && maxValue != nil && *minValue < 0 && *maxValue > 0:
		if !isStarExpr {
			reportShouldAddPointer(pass, field, a.pointerPolicy, fieldName, "field %s has a range of values including 0. The difference between omitted and 0 is significant and therefore the field should be a pointer")
		}
	case minValue != nil && *minValue < 0 && maxValue == nil:
		pass.Reportf(field.Pos(), "field %s has a negative minimum value and does not have a maximum value. A maximum value should be set", fieldName)
	case maxValue != nil && *maxValue > 0 && minValue == nil:
		pass.Reportf(field.Pos(), "field %s has a positive maximum value and does not have a minimum value. A minimum value should be set", fieldName)
	case minValue == nil || maxValue == nil:
		if isStarExpr {
			pass.Reportf(field.Pos(), "field %s is an optional integer and does not have a minimum/maximum value. Where the difference between omitted and 0 is significant, set the minimum/maximum value to a range including 0", fieldName)
		} else {
			pass.Reportf(field.Pos(), "field %s is an optional integer and does not have a minimum/maximum value. Either set a minimum/maximum value or make %s a pointer where the difference between omitted and 0 is significant", fieldName, fieldName)
		}
	}
}

func (a *analyzer) checkFieldPointersPreferenceWhenRequiredIntegerWithoutOmitEmpty(pass *analysis.Pass, field *ast.Field, fieldName string, minValue, maxValue *int, isStarExpr bool, fieldMarkers markers.MarkerSet) {
	switch {
	case minValue != nil && *minValue > 0:
		reportShouldRemoveMarker(pass, field, fieldMarkers.Get(minimumMarker)[0], fieldName, "field %s has a greater than zero minimum value without omitempty. The minimum value should be removed.")
	case maxValue != nil && *maxValue < 0:
		reportShouldRemoveMarker(pass, field, fieldMarkers.Get(maximumMarker)[0], fieldName, "field %s has a less than zero maximum value without omitempty. The maximum value should be removed.")
	}

	if isStarExpr {
		reportShouldRemovePointer(pass, field, a.pointerPolicy, fieldName, "field %s is an optional integer without omitempty. It should not be a pointer.")
	}
}

func (a *analyzer) checkFieldPointersPreferenceWhenRequiredFloat(pass *analysis.Pass, field *ast.Field, fieldName string, isStarExpr bool, markersAccess markers.Markers, jsonTags extractjsontags.FieldTagInfo) {
	fieldMarkers := markersAccess.FieldMarkers(field)

	// These are pointers so that we can determine if they are set, or not.
	var minValue, maxValue *float64

	if minimum := fieldMarkers.Get(minimumMarker); len(minimum) > 0 {
		// Where there are multiple minimum markers, we only care about the first one.
		// Other rules should deduplicate them.
		min, err := strconv.ParseFloat(minimum[0].Expressions[""], 64)
		if err != nil {
			pass.Reportf(field.Pos(), "field %s has a minimum value of %s, but it is not a float", fieldName, minimum[0].Expressions[""])
			return
		}
		minValue = &min
	}

	if maximum := fieldMarkers.Get(maximumMarker); len(maximum) > 0 {
		// Where there are multiple maximum markers, we only care about the first one.
		// Other rules should deduplicate them.
		max, err := strconv.ParseFloat(maximum[0].Expressions[""], 64)
		if err != nil {
			pass.Reportf(field.Pos(), "field %s has a maximum value of %s, but it is not a float", fieldName, maximum[0].Expressions[""])
			return
		}
		maxValue = &max
	}

	if a.omitEmptyPolicy == config.OptionalFieldsOmitEmptyPolicyIgnore && !jsonTags.OmitEmpty {
		a.checkFieldPointersPreferenceWhenRequiredFloatWithoutOmitEmpty(pass, field, fieldName, minValue, maxValue, isStarExpr, fieldMarkers)
		return
	}

	switch {
	case minValue != nil && *minValue > 0:
		if isStarExpr {
			reportShouldRemovePointer(pass, field, a.pointerPolicy, fieldName, "field %s has a greater than 0 minimum value and does not need to be a pointer")
		}
	case maxValue != nil && *maxValue < 0:
		if isStarExpr {
			reportShouldRemovePointer(pass, field, a.pointerPolicy, fieldName, "field %s has a negative maximum value and does not need to be a pointer")
		}
	case minValue != nil && *minValue == 0,
		maxValue != nil && *maxValue == 0,
		minValue != nil && maxValue != nil && *minValue < 0 && *maxValue > 0:
		if !isStarExpr {
			reportShouldAddPointer(pass, field, a.pointerPolicy, fieldName, "field %s has a range of values including 0. The difference between omitted and 0 is significant and therefore the field should be a pointer")
		}
	case minValue != nil && *minValue < 0 && maxValue == nil:
		pass.Reportf(field.Pos(), "field %s has a negative minimum value and does not have a maximum value. A maximum value should be set", fieldName)
	case maxValue != nil && *maxValue > 0 && minValue == nil:
		pass.Reportf(field.Pos(), "field %s has a positive maximum value and does not have a minimum value. A minimum value should be set", fieldName)
	case minValue == nil || maxValue == nil:
		if isStarExpr {
			pass.Reportf(field.Pos(), "field %s is an optional float and does not have a minimum/maximum value. Where the difference between omitted and 0 is significant, set the minimum/maximum value to a range including 0", fieldName)
		} else {
			pass.Reportf(field.Pos(), "field %s is an optional float and does not have a minimum/maximum value. Either set a minimum/maximum value or make %s a pointer where the difference between omitted and 0 is significant", fieldName, fieldName)
		}
	}
}

func (a *analyzer) checkFieldPointersPreferenceWhenRequiredFloatWithoutOmitEmpty(pass *analysis.Pass, field *ast.Field, fieldName string, minValue, maxValue *float64, isStarExpr bool, fieldMarkers markers.MarkerSet) {
	switch {
	case minValue != nil && *minValue > 0:
		reportShouldRemoveMarker(pass, field, fieldMarkers.Get(minimumMarker)[0], fieldName, "field %s has a greater than zero minimum value without omitempty. The minimum value should be removed.")
	case maxValue != nil && *maxValue < 0:
		reportShouldRemoveMarker(pass, field, fieldMarkers.Get(maximumMarker)[0], fieldName, "field %s has a less than zero maximum value without omitempty. The maximum value should be removed.")
	}

	if isStarExpr {
		reportShouldRemovePointer(pass, field, a.pointerPolicy, fieldName, "field %s is an optional float without omitempty. It should not be a pointer.")
	}
}

func defaultConfig(cfg *config.OptionalFieldsConfig) {
	if cfg.Pointers.Policy == "" {
		cfg.Pointers.Policy = config.OptionalFieldsPointerPolicySuggestFix
	}

	if cfg.Pointers.Preference == "" {
		cfg.Pointers.Preference = config.OptionalFieldsPointerPreferenceAlways
	}

	if cfg.OmitEmpty.Policy == "" {
		cfg.OmitEmpty.Policy = config.OptionalFieldsOmitEmptyPolicySuggestFix
	}
}

// isStarExpr checks if the expression is a pointer type.
func isStarExpr(expr ast.Expr) (bool, ast.Expr) {
	if ptrType, ok := expr.(*ast.StarExpr); ok {
		return true, ptrType.X
	}

	return false, expr
}

// isPointerType checks if the expression is a pointer type.
// This is for types that are always implemented as pointers and therefore should
// not be the underlying type of a star expr.
func isPointerType(expr ast.Expr) bool {
	switch expr.(type) {
	case *ast.StarExpr, *ast.MapType, *ast.ArrayType:
		return true
	default:
		return false
	}
}

func structContainsRequiredFields(structType *ast.StructType, markersAccess markers.Markers) bool {
	if structType == nil {
		return false
	}

	for _, field := range structType.Fields.List {
		fieldMarkers := markersAccess.FieldMarkers(field)

		if isFieldRequired(fieldMarkers) {
			return true
		}
	}

	return false
}

func isFieldRequired(fieldMarkers markers.MarkerSet) bool {
	return fieldMarkers.Has("required") || fieldMarkers.Has("kubebuilder:validation:Required")
}

func reportShouldAddPointer(pass *analysis.Pass, field *ast.Field, pointerPolicy config.OptionalFieldsPointerPolicy, fieldName, messageFmt string) {
	switch pointerPolicy {
	case config.OptionalFieldsPointerPolicySuggestFix:
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf(messageFmt, fieldName),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "should make the field a pointer",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     field.Pos() + token.Pos(len(fieldName)+1),
							NewText: []byte("*"),
						},
					},
				},
			},
		})
	case config.OptionalFieldsPointerPolicyWarn:
		pass.Reportf(field.Pos(), messageFmt, fieldName)
	default:
		panic(fmt.Sprintf("unknown pointer policy: %s", pointerPolicy))
	}

}

func reportShouldRemovePointer(pass *analysis.Pass, field *ast.Field, pointerPolicy config.OptionalFieldsPointerPolicy, fieldName, messageFmt string) {
	switch pointerPolicy {
	case config.OptionalFieldsPointerPolicySuggestFix:
		pass.Report(analysis.Diagnostic{
			Pos:     field.Pos(),
			Message: fmt.Sprintf(messageFmt, fieldName),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "should remove the pointer",
					TextEdits: []analysis.TextEdit{
						{
							Pos: field.Pos() + token.Pos(len(fieldName)+1),
							End: field.Pos() + token.Pos(len(fieldName)+2),
						},
					},
				},
			},
		})
	case config.OptionalFieldsPointerPolicyWarn:
		pass.Reportf(field.Pos(), messageFmt, fieldName)
	default:
		panic(fmt.Sprintf("unknown pointer policy: %s", pointerPolicy))
	}
}

func reportShouldRemoveMarker(pass *analysis.Pass, field *ast.Field, marker markers.Marker, fieldName, messageFmt string) {
	pass.Report(analysis.Diagnostic{
		Pos:     field.Pos(),
		Message: fmt.Sprintf(messageFmt, fieldName),
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: fmt.Sprintf("should remove the marker: %s", marker.RawComment),
				TextEdits: []analysis.TextEdit{
					{
						Pos: marker.Pos,
						End: marker.End + 1,
					},
				},
			},
		},
	})
}

// getMarkerIntegerValue extracts a numeric value from the default
// value of a marker.
// Works for markers like MaxLength, MinLength, etc.
func getMarkerIntegerValue(marker markers.Marker) (int, error) {
	rawValue, ok := marker.Expressions[""]
	if !ok {
		return 0, errMarkerMissingValue
	}

	value, err := strconv.Atoi(rawValue)
	if err != nil {
		return 0, fmt.Errorf("error converting value to integer: %w", err)
	}

	return value, nil
}

func structHasGreaterThanZeroMinProperties(structType *ast.StructType, structMarkers markers.MarkerSet) (bool, error) {
	if structType == nil {
		return false, nil
	}

	if structMarkers.Has(minPropertiesMarker) {
		for _, marker := range structMarkers.Get(minPropertiesMarker) {
			markerValue, err := getMarkerIntegerValue(marker)
			if err != nil {
				return false, fmt.Errorf("error getting marker value: %w", err)
			}

			if markerValue > 0 {
				return true, nil
			}
		}
	}

	return false, nil
}

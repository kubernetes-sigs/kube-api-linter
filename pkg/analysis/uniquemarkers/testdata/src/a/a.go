package a

// +kubebuilder:validation:MinProperties:=3
// +kubebuilder:validation:MinProperties:=5
type A struct { // want "type A has multiple definitions of marker kubebuilder:validation:MinProperties when only a single definition should exist"

	// +kubebuilder:validation:XValidation:rule='self.matches("[0-9A-Za-z]")',message='should match regex'
	// +kubebuilder:validation:XValidation:rule='isURL(self)',message='should be a URL'
	AllowedNonUniqueMarkers string

	// +kubebuilder:validation:XValidation:rule='self.matches("[0-9]")',message='should actually be a number'
	// +kubebuilder:validation:XValidation:rule='self.matches("[0-9]")',message='different message'
	NonUniqueXValidations string // want "field A.NonUniqueXValidations has multiple definitions of marker kubebuilder:validation:XValidation:rule='self\\.matches\\(\\\"\\[0-9]\\\"\\)' when only a single definition should exist"

	// +kubebuilder:validation:items:XValidation:rule='self.matches("[0-9]")',message='should actually be a number'
	// +kubebuilder:validation:items:XValidation:rule='self.matches("[0-9]")',message='different message'
	NonUniqueItemsXValidations []string // want "field A.NonUniqueItemsXValidations has multiple definitions of marker kubebuilder:validation:items:XValidation:rule='self\\.matches\\(\\\"\\[0-9]\\\"\\)' when only a single definition should exist"

	// +kubebuilder:validation:XValidation:rule='self.matches("[a-z]")',message='same'
	// +kubebuilder:validation:XValidation:rule='self.matches("[0-9]")',message='same'
	UniqueXValidationsFromMessage string

	// +custom:SomeCustomMarker:=value
	// +custom:SomeCustomMarker:=diffvalue
	CustomMarkers string

	// +default:value="homer"
	UniqueBasicDefault string

	// +default:value="homer"
	// +default:value="bart"
	NonUniqueBasicDefault string // want "field A.NonUniqueBasicDefault has multiple definitions of marker default when only a single definition should exist"

	// +default:value="homer"
	NonUniqueBasicDefaultFromAliasWithBasicDefault UniqueBasicDefaultAlias // want "field A.NonUniqueBasicDefaultFromAliasWithBasicDefault has multiple definitions of marker default when only a single definition should exist"

	NonUniqueBasicDefaultOnlyFromAliasWithBasicDefault NonUniqueBasicDefaultAlias // want "field A.NonUniqueBasicDefaultOnlyFromAliasWithBasicDefault has multiple definitions of marker default when only a single definition should exist"

	// +kubebuilder:default:="homer"
	UniqueKubebuilderDefault string

	// +kubebuilder:default:="homer"
	// +kubebuilder:default:="bart"
	NonUniqueKubebuilderDefault string // want "field A.NonUniqueKubebuilderDefault has multiple definitions of marker kubebuilder:default when only a single definition should exist"

	// +kubebuilder:default:="homer"
	NonUniqueKubebuilderDefaultFromAliasWithKubebuilderDefault UniqueKubebuilderDefaultAlias // want "field A.NonUniqueKubebuilderDefaultFromAliasWithKubebuilderDefault has multiple definitions of marker kubebuilder:default when only a single definition should exist"

	NonUniqueKubebuilderDefaultOnlyFromAliasWithKubebuilderDefault NonUniqueKubebuilderDefaultAlias // want "field A.NonUniqueKubebuilderDefaultOnlyFromAliasWithKubebuilderDefault has multiple definitions of marker kubebuilder:default when only a single definition should exist"

	// +kubebuilder:example:="homer"
	UniqueKubebuilderExample string

	// +kubebuilder:example:="homer"
	// +kubebuilder:example:="bart"
	NonUniqueKubebuilderExample string // want "field A.NonUniqueKubebuilderExample has multiple definitions of marker kubebuilder:example when only a single definition should exist"

	// +kubebuilder:example:="homer"
	NonUniqueKubebuilderExampleFromAliasWithKubebuilderExample UniqueKubebuilderExampleAlias // want "field A.NonUniqueKubebuilderExampleFromAliasWithKubebuilderExample has multiple definitions of marker kubebuilder:example when only a single definition should exist"

	NonUniqueKubebuilderExampleOnlyFromAliasWithKubebuilderExample NonUniqueKubebuilderExampleAlias // want "field A.NonUniqueKubebuilderExampleOnlyFromAliasWithKubebuilderExample has multiple definitions of marker kubebuilder:example when only a single definition should exist"

	// +kubebuilder:validation:MinLength:=1
	UniqueMinLength string

	// +kubebuilder:validation:MinLength:=1
	// +kubebuilder:validation:MinLength:=5
	NonUniqueMinLength string // want "field A.NonUniqueMinLength has multiple definitions of marker kubebuilder:validation:MinLength when only a single definition should exist"

	// +kubebuilder:validation:MinLength:=1
	NonUniqueMinLengthFromAliasWithMinLength UniqueMinLengthStringAlias // want "field A.NonUniqueMinLengthFromAliasWithMinLength has multiple definitions of marker kubebuilder:validation:MinLength when only a single definition should exist"

	NonUniqueMinLengthOnlyFromAliasWithMinLength NonUniqueMinLengthStringAlias // want "field A.NonUniqueMinLengthOnlyFromAliasWithMinLength has multiple definitions of marker kubebuilder:validation:MinLength when only a single definition should exist"

	// +kubebuilder:validation:Enum:=Foo;Bar
	UniqueEnum string

	// +kubebuilder:validation:Enum:=Foo;Bar
	// +kubebuilder:validation:Enum:=Baz;Qux
	NonUniqueEnum string // want "field A.NonUniqueEnum has multiple definitions of marker kubebuilder:validation:Enum when only a single definition should exist"

	// +kubebuilder:validation:Enum:=Foo;Bar
	NonUniqueEnumFromAliasWithEnum UniqueEnumAlias // want "field A.NonUniqueEnumFromAliasWithEnum has multiple definitions of marker kubebuilder:validation:Enum when only a single definition should exist"

	NonUniqueEnumOnlyFromAliasWithEnum NonUniqueEnumAlias // want "field A.NonUniqueEnumOnlyFromAliasWithEnum has multiple definitions of marker kubebuilder:validation:Enum when only a single definition should exist"

	// +kubebuilder:validation:Maximum:=10
	// +kubebuilder:validation:ExclusiveMaximum:=true
	UniqueExclusiveMaximum int

	// +kubebuilder:validation:Maximum:=10
	// +kubebuilder:validation:ExclusiveMaximum:=true
	// +kubebuilder:validation:ExclusiveMaximum:=false
	NonUniqueExclusiveMaximum int // want "field A.NonUniqueExclusiveMaximum has multiple definitions of marker kubebuilder:validation:ExclusiveMaximum when only a single definition should exist"

	// +kubebuilder:validation:ExclusiveMaximum:=true
	NonUniqueExclusiveMaximumFromAliasWithExclusiveMaximum UniqueExclusiveMaximumAlias // want "field A.NonUniqueExclusiveMaximumFromAliasWithExclusiveMaximum has multiple definitions of marker kubebuilder:validation:ExclusiveMaximum when only a single definition should exist"

	NonUniqueExclusiveMaximumOnlyFromAliasWithExclusiveMaximum NonUniqueExclusiveMaximumAlias // want "field A.NonUniqueExclusiveMaximumOnlyFromAliasWithExclusiveMaximum has multiple definitions of marker kubebuilder:validation:ExclusiveMaximum when only a single definition should exist"

	// +kubebuilder:validation:Minimum:=2
	// +kubebuilder:validation:ExclusiveMinimum:=true
	UniqueExclusiveMinimum int

	// +kubebuilder:validation:Minimum:=2
	// +kubebuilder:validation:ExclusiveMinimum:=true
	// +kubebuilder:validation:ExclusiveMinimum:=false
	NonUniqueExclusiveMinimum int // want "field A.NonUniqueExclusiveMinimum has multiple definitions of marker kubebuilder:validation:ExclusiveMinimum when only a single definition should exist"

	// +kubebuilder:validation:ExclusiveMinimum:=true
	NonUniqueExclusiveMinimumFromAliasWithExclusiveMinimum UniqueExclusiveMinimumAlias // want "field A.NonUniqueExclusiveMinimumFromAliasWithExclusiveMinimum has multiple definitions of marker kubebuilder:validation:ExclusiveMinimum when only a single definition should exist"

	NonUniqueExclusiveMinimumOnlyFromAliasWithExclusiveMinimum NonUniqueExclusiveMinimumAlias // want "field A.NonUniqueExclusiveMinimumOnlyFromAliasWithExclusiveMinimum has multiple definitions of marker kubebuilder:validation:ExclusiveMinimum when only a single definition should exist"

	// +kubebuilder:validation:Format:="date-time"
	UniqueFormat string

	// +kubebuilder:validation:Format:="date-time"
	// +kubebuilder:validation:Format:="password"
	NonUniqueFormat string // want "field A.NonUniqueFormat has multiple definitions of marker kubebuilder:validation:Format when only a single definition should exist"

	// +kubebuilder:validation:Format:="password"
	NonUniqueFormatFromAliasWithFormat UniqueFormatAlias // want "field A.NonUniqueFormatFromAliasWithFormat has multiple definitions of marker kubebuilder:validation:Format when only a single definition should exist"

	NonUniqueFormatOnlyFromAliasWithFormat NonUniqueFormatAlias // want "field A.NonUniqueFormatOnlyFromAliasWithFormat has multiple definitions of marker kubebuilder:validation:Format when only a single definition should exist"

	// +kubebuilder:validation:MaxItems:=10
	UniqueMaxItems []string

	// +kubebuilder:validation:MaxItems:=10
	// +kubebuilder:validation:MaxItems:=5
	NonUniqueMaxItems []string // want "field A.NonUniqueMaxItems has multiple definitions of marker kubebuilder:validation:MaxItems when only a single definition should exist"

	// +kubebuilder:validation:MaxItems:=8
	NonUniqueMaxItemsFromAliasWithMaxItems UniqueMaxItemsAlias // want "field A.NonUniqueMaxItemsFromAliasWithMaxItems has multiple definitions of marker kubebuilder:validation:MaxItems when only a single definition should exist"

	NonUniqueMaxItemsOnlyFromAliasWithMaxItems NonUniqueMaxItemsAlias // want "field A.NonUniqueMaxItemsOnlyFromAliasWithMaxItems has multiple definitions of marker kubebuilder:validation:MaxItems when only a single definition should exist"

	// +kubebuilder:validation:MaxProperties:=3
	UniqueMaxProperties map[string]string

	// +kubebuilder:validation:MaxProperties:=3
	// +kubebuilder:validation:MaxProperties:=5
	NonUniqueMaxProperties map[string]string // want "field A.NonUniqueMaxProperties has multiple definitions of marker kubebuilder:validation:MaxProperties when only a single definition should exist"

	// +kubebuilder:validation:MaxProperties:=4
	NonUniqueMaxPropertiesFromAliasWithMaxProperties UniqueMaxPropertiesAlias // want "field A.NonUniqueMaxPropertiesFromAliasWithMaxProperties has multiple definitions of marker kubebuilder:validation:MaxProperties when only a single definition should exist"

	NonUniqueMaxPropertiesOnlyFromAliasWithMaxProperties NonUniqueMaxPropertiesAlias // want "field A.NonUniqueMaxPropertiesOnlyFromAliasWithMaxProperties has multiple definitions of marker kubebuilder:validation:MaxProperties when only a single definition should exist"

	// +kubebuilder:validation:Maximum:=100
	UniqueMaximum int

	// +kubebuilder:validation:Maximum:=100
	// +kubebuilder:validation:Maximum:=200
	NonUniqueMaximum int // want "field A.NonUniqueMaximum has multiple definitions of marker kubebuilder:validation:Maximum when only a single definition should exist"

	// +kubebuilder:validation:Maximum:=150
	NonUniqueMaximumFromAliasWithMaximum UniqueMaximumAlias // want "field A.NonUniqueMaximumFromAliasWithMaximum has multiple definitions of marker kubebuilder:validation:Maximum when only a single definition should exist"

	NonUniqueMaximumOnlyFromAliasWithMaximum NonUniqueMaximumAlias // want "field A.NonUniqueMaximumOnlyFromAliasWithMaximum has multiple definitions of marker kubebuilder:validation:Maximum when only a single definition should exist"

	// +kubebuilder:validation:MinItems:=1
	UniqueMinItems []string

	// +kubebuilder:validation:MinItems:=1
	// +kubebuilder:validation:MinItems:=2
	NonUniqueMinItems []string // want "field A.NonUniqueMinItems has multiple definitions of marker kubebuilder:validation:MinItems when only a single definition should exist"

	// +kubebuilder:validation:MinItems:=3
	NonUniqueMinItemsFromAliasWithMinItems UniqueMinItemsAlias // want "field A.NonUniqueMinItemsFromAliasWithMinItems has multiple definitions of marker kubebuilder:validation:MinItems when only a single definition should exist"

	NonUniqueMinItemsOnlyFromAliasWithMinItems NonUniqueMinItemsAlias // want "field A.NonUniqueMinItemsOnlyFromAliasWithMinItems has multiple definitions of marker kubebuilder:validation:MinItems when only a single definition should exist"

	// +kubebuilder:validation:MinLength:=30
	UniqueMinLengthCustom string

	// +kubebuilder:validation:MinLength:=30
	// +kubebuilder:validation:MinLength:=35
	NonUniqueMinLengthCustom string // want "field A.NonUniqueMinLengthCustom has multiple definitions of marker kubebuilder:validation:MinLength when only a single definition should exist"

	// +kubebuilder:validation:MinLength:=30
	NonUniqueMinLengthCustomFromAliasWithMinLength UniqueMinLengthCustomAlias // want "field A.NonUniqueMinLengthCustomFromAliasWithMinLength has multiple definitions of marker kubebuilder:validation:MinLength when only a single definition should exist"

	NonUniqueMinLengthCustomOnlyFromAliasWithMinLength NonUniqueMinLengthCustomAlias // want "field A.NonUniqueMinLengthCustomOnlyFromAliasWithMinLength has multiple definitions of marker kubebuilder:validation:MinLength when only a single definition should exist"

	// +kubebuilder:validation:MinProperties:=1
	UniqueMinPropertiesField map[string]string

	// +kubebuilder:validation:MinProperties:=1
	// +kubebuilder:validation:MinProperties:=2
	NonUniqueMinPropertiesField map[string]string // want "field A.NonUniqueMinPropertiesField has multiple definitions of marker kubebuilder:validation:MinProperties when only a single definition should exist"

	// +kubebuilder:validation:MinProperties:=1
	NonUniqueMinPropertiesFieldFromAliasWithMinProperties UniqueMinPropertiesFieldAlias // want "field A.NonUniqueMinPropertiesFieldFromAliasWithMinProperties has multiple definitions of marker kubebuilder:validation:MinProperties when only a single definition should exist"

	NonUniqueMinPropertiesFieldOnlyFromAliasWithMinProperties NonUniqueMinPropertiesFieldAlias // want "field A.NonUniqueMinPropertiesFieldOnlyFromAliasWithMinProperties has multiple definitions of marker kubebuilder:validation:MinProperties when only a single definition should exist"

	// +kubebuilder:validation:Minimum:=10
	UniqueMinimum int

	// +kubebuilder:validation:Minimum:=10
	// +kubebuilder:validation:Minimum:=20
	NonUniqueMinimum int // want "field A.NonUniqueMinimum has multiple definitions of marker kubebuilder:validation:Minimum when only a single definition should exist"

	// +kubebuilder:validation:Minimum:=15
	NonUniqueMinimumFromAliasWithMinimum UniqueMinimumAlias // want "field A.NonUniqueMinimumFromAliasWithMinimum has multiple definitions of marker kubebuilder:validation:Minimum when only a single definition should exist"

	NonUniqueMinimumOnlyFromAliasWithMinimum NonUniqueMinimumAlias // want "field A.NonUniqueMinimumOnlyFromAliasWithMinimum has multiple definitions of marker kubebuilder:validation:Minimum when only a single definition should exist"

	// +kubebuilder:validation:MultipleOf:=3
	UniqueMultipleOf int

	// +kubebuilder:validation:MultipleOf:=3
	// +kubebuilder:validation:MultipleOf:=5
	NonUniqueMultipleOf int // want "field A.NonUniqueMultipleOf has multiple definitions of marker kubebuilder:validation:MultipleOf when only a single definition should exist"

	// +kubebuilder:validation:MultipleOf:=2
	NonUniqueMultipleOfFromAliasWithMultipleOf UniqueMultipleOfAlias // want "field A.NonUniqueMultipleOfFromAliasWithMultipleOf has multiple definitions of marker kubebuilder:validation:MultipleOf when only a single definition should exist"

	NonUniqueMultipleOfOnlyFromAliasWithMultipleOf NonUniqueMultipleOfAlias // want "field A.NonUniqueMultipleOfOnlyFromAliasWithMultipleOf has multiple definitions of marker kubebuilder:validation:MultipleOf when only a single definition should exist"

	// +kubebuilder:validation:Pattern:="^[a-z]+$"
	UniquePattern string

	// +kubebuilder:validation:Pattern:="^[a-z]+$"
	// +kubebuilder:validation:Pattern:="^[0-9]+$"
	NonUniquePattern string // want "field A.NonUniquePattern has multiple definitions of marker kubebuilder:validation:Pattern when only a single definition should exist"

	// +kubebuilder:validation:Pattern:="^[A-Z]+$"
	NonUniquePatternFromAliasWithPattern UniquePatternAlias // want "field A.NonUniquePatternFromAliasWithPattern has multiple definitions of marker kubebuilder:validation:Pattern when only a single definition should exist"

	NonUniquePatternOnlyFromAliasWithPattern NonUniquePatternAlias // want "field A.NonUniquePatternOnlyFromAliasWithPattern has multiple definitions of marker kubebuilder:validation:Pattern when only a single definition should exist"

	// +kubebuilder:validation:Type:="string"
	UniqueTypeValidationField string

	// +kubebuilder:validation:Type:="string"
	// +kubebuilder:validation:Type:="date"
	NonUniqueTypeValidationField string // want "field A.NonUniqueTypeValidationField has multiple definitions of marker kubebuilder:validation:Type when only a single definition should exist"

	// +kubebuilder:validation:Type:="date"
	NonUniqueTypeValidationFieldFromAliasWithType UniqueTypeValidationAlias // want "field A.NonUniqueTypeValidationFieldFromAliasWithType has multiple definitions of marker kubebuilder:validation:Type when only a single definition should exist"

	NonUniqueTypeValidationFieldOnlyFromAliasWithType NonUniqueTypeValidationAlias // want "field A.NonUniqueTypeValidationFieldOnlyFromAliasWithType has multiple definitions of marker kubebuilder:validation:Type when only a single definition should exist"

	// +kubebuilder:validation:UniqueItems:=true
	UniqueUniqueItems []string

	// +kubebuilder:validation:UniqueItems:=true
	// +kubebuilder:validation:UniqueItems:=false
	NonUniqueUniqueItems []string // want "field A.NonUniqueUniqueItems has multiple definitions of marker kubebuilder:validation:UniqueItems when only a single definition should exist"

	// +kubebuilder:validation:UniqueItems:=true
	NonUniqueUniqueItemsFromAliasWithUniqueItems UniqueUniqueItemsAlias // want "field A.NonUniqueUniqueItemsFromAliasWithUniqueItems has multiple definitions of marker kubebuilder:validation:UniqueItems when only a single definition should exist"

	NonUniqueUniqueItemsOnlyFromAliasWithUniqueItems NonUniqueUniqueItemsAlias // want "field A.NonUniqueUniqueItemsOnlyFromAliasWithUniqueItems has multiple definitions of marker kubebuilder:validation:UniqueItems when only a single definition should exist"

	// +kubebuilder:validation:items:Enum=Apple;Orange
	UniqueItemsEnum []string

	// +kubebuilder:validation:items:Enum=Apple;Orange
	// +kubebuilder:validation:items:Enum=Banana;Grape
	NonUniqueItemsEnum []string // want "field A.NonUniqueItemsEnum has multiple definitions of marker kubebuilder:validation:items:Enum when only a single definition should exist"

	// +kubebuilder:validation:items:Enum=Apple;Orange
	NonUniqueItemsEnumFromAliasWithItemsEnum UniqueItemsEnumAlias // want "field A.NonUniqueItemsEnumFromAliasWithItemsEnum has multiple definitions of marker kubebuilder:validation:items:Enum when only a single definition should exist"

	NonUniqueItemsEnumOnlyFromAliasWithItemsEnum NonUniqueItemsEnumAlias // want "field A.NonUniqueItemsEnumOnlyFromAliasWithItemsEnum has multiple definitions of marker kubebuilder:validation:items:Enum when only a single definition should exist"

	// +kubebuilder:validation:items:Maximum:=10
	// +kubebuilder:validation:items:ExclusiveMaximum:=true
	UniqueItemsExclusiveMaximum []int

	// +kubebuilder:validation:items:Maximum:=10
	// +kubebuilder:validation:items:ExclusiveMaximum:=true
	// +kubebuilder:validation:items:ExclusiveMaximum:=false
	NonUniqueItemsExclusiveMaximum []int // want "field A.NonUniqueItemsExclusiveMaximum has multiple definitions of marker kubebuilder:validation:items:ExclusiveMaximum when only a single definition should exist"

	// +kubebuilder:validation:items:ExclusiveMaximum:=true
	NonUniqueItemsExclusiveMaximumFromAliasWithItemsExclusiveMaximum UniqueItemsExclusiveMaximumAlias // want "field A.NonUniqueItemsExclusiveMaximumFromAliasWithItemsExclusiveMaximum has multiple definitions of marker kubebuilder:validation:items:ExclusiveMaximum when only a single definition should exist"

	NonUniqueItemsExclusiveMaximumOnlyFromAliasWithItemsExclusiveMaximum NonUniqueItemsExclusiveMaximumAlias // want "field A.NonUniqueItemsExclusiveMaximumOnlyFromAliasWithItemsExclusiveMaximum has multiple definitions of marker kubebuilder:validation:items:ExclusiveMaximum when only a single definition should exist"

	// +kubebuilder:validation:items:Minimum:=1
	// +kubebuilder:validation:items:ExclusiveMinimum:=true
	UniqueItemsExclusiveMinimum []int

	// +kubebuilder:validation:items:Minimum:=1
	// +kubebuilder:validation:items:ExclusiveMinimum:=true
	// +kubebuilder:validation:items:ExclusiveMinimum:=false
	NonUniqueItemsExclusiveMinimum []int // want "field A.NonUniqueItemsExclusiveMinimum has multiple definitions of marker kubebuilder:validation:items:ExclusiveMinimum when only a single definition should exist"

	// +kubebuilder:validation:items:ExclusiveMinimum:=true
	NonUniqueItemsExclusiveMinimumFromAliasWithItemsExclusiveMinimum UniqueItemsExclusiveMinimumAlias // want "field A.NonUniqueItemsExclusiveMinimumFromAliasWithItemsExclusiveMinimum has multiple definitions of marker kubebuilder:validation:items:ExclusiveMinimum when only a single definition should exist"

	NonUniqueItemsExclusiveMinimumOnlyFromAliasWithItemsExclusiveMinimum NonUniqueItemsExclusiveMinimumAlias // want "field A.NonUniqueItemsExclusiveMinimumOnlyFromAliasWithItemsExclusiveMinimum has multiple definitions of marker kubebuilder:validation:items:ExclusiveMinimum when only a single definition should exist"

	// +kubebuilder:validation:items:Format:="date-time"
	UniqueItemsFormat []string

	// +kubebuilder:validation:items:Format:="date-time"
	// +kubebuilder:validation:items:Format:="password"
	NonUniqueItemsFormat []string // want "field A.NonUniqueItemsFormat has multiple definitions of marker kubebuilder:validation:items:Format when only a single definition should exist"

	// +kubebuilder:validation:items:Format:="date-time"
	NonUniqueItemsFormatFromAliasWithItemsFormat UniqueItemsFormatAlias // want "field A.NonUniqueItemsFormatFromAliasWithItemsFormat has multiple definitions of marker kubebuilder:validation:items:Format when only a single definition should exist"

	NonUniqueItemsFormatOnlyFromAliasWithItemsFormat NonUniqueItemsFormatAlias // want "field A.NonUniqueItemsFormatOnlyFromAliasWithItemsFormat has multiple definitions of marker kubebuilder:validation:items:Format when only a single definition should exist"

	// +kubebuilder:validation:items:MaxItems:=5
	UniqueItemsMaxItems [][]string

	// +kubebuilder:validation:items:MaxItems:=5
	// +kubebuilder:validation:items:MaxItems:=10
	NonUniqueItemsMaxItems [][]string // want "field A.NonUniqueItemsMaxItems has multiple definitions of marker kubebuilder:validation:items:MaxItems when only a single definition should exist"

	// +kubebuilder:validation:items:MaxItems:=3
	NonUniqueItemsMaxItemsFromAliasWithItemsMaxItems UniqueItemsMaxItemsAlias // want "field A.NonUniqueItemsMaxItemsFromAliasWithItemsMaxItems has multiple definitions of marker kubebuilder:validation:items:MaxItems when only a single definition should exist"

	NonUniqueItemsMaxItemsOnlyFromAliasWithItemsMaxItems NonUniqueItemsMaxItemsAlias // want "field A.NonUniqueItemsMaxItemsOnlyFromAliasWithItemsMaxItems has multiple definitions of marker kubebuilder:validation:items:MaxItems when only a single definition should exist"

	// +kubebuilder:validation:items:MaxProperties:=3
	UniqueItemsMaxProperties []map[string]string

	// +kubebuilder:validation:items:MaxProperties:=3
	// +kubebuilder:validation:items:MaxProperties:=5
	NonUniqueItemsMaxProperties []map[string]string // want "field A.NonUniqueItemsMaxProperties has multiple definitions of marker kubebuilder:validation:items:MaxProperties when only a single definition should exist"

	// +kubebuilder:validation:items:MaxProperties:=2
	NonUniqueItemsMaxPropertiesFromAliasWithItemsMaxProperties UniqueItemsMaxPropertiesAlias // want "field A.NonUniqueItemsMaxPropertiesFromAliasWithItemsMaxProperties has multiple definitions of marker kubebuilder:validation:items:MaxProperties when only a single definition should exist"

	NonUniqueItemsMaxPropertiesOnlyFromAliasWithItemsMaxProperties NonUniqueItemsMaxPropertiesAlias // want "field A.NonUniqueItemsMaxPropertiesOnlyFromAliasWithItemsMaxProperties has multiple definitions of marker kubebuilder:validation:items:MaxProperties when only a single definition should exist"

	// +kubebuilder:validation:items:Maximum:=100
	UniqueItemsMaximum []int

	// +kubebuilder:validation:items:Maximum:=100
	// +kubebuilder:validation:items:Maximum:=50
	NonUniqueItemsMaximum []int // want "field A.NonUniqueItemsMaximum has multiple definitions of marker kubebuilder:validation:items:Maximum when only a single definition should exist"

	// +kubebuilder:validation:items:Maximum:=75
	NonUniqueItemsMaximumFromAliasWithItemsMaximum UniqueItemsMaximumAlias // want "field A.NonUniqueItemsMaximumFromAliasWithItemsMaximum has multiple definitions of marker kubebuilder:validation:items:Maximum when only a single definition should exist"

	NonUniqueItemsMaximumOnlyFromAliasWithItemsMaximum NonUniqueItemsMaximumAlias // want "field A.NonUniqueItemsMaximumOnlyFromAliasWithItemsMaximum has multiple definitions of marker kubebuilder:validation:items:Maximum when only a single definition should exist"

	// +kubebuilder:validation:items:MinItems:=1
	UniqueItemsMinItems [][]string

	// +kubebuilder:validation:items:MinItems:=1
	// +kubebuilder:validation:items:MinItems:=2
	NonUniqueItemsMinItems [][]string // want "field A.NonUniqueItemsMinItems has multiple definitions of marker kubebuilder:validation:items:MinItems when only a single definition should exist"

	// +kubebuilder:validation:items:MinItems:=3
	NonUniqueItemsMinItemsFromAliasWithItemsMinItems UniqueItemsMinItemsAlias // want "field A.NonUniqueItemsMinItemsFromAliasWithItemsMinItems has multiple definitions of marker kubebuilder:validation:items:MinItems when only a single definition should exist"

	NonUniqueItemsMinItemsOnlyFromAliasWithItemsMinItems NonUniqueItemsMinItemsAlias // want "field A.NonUniqueItemsMinItemsOnlyFromAliasWithItemsMinItems has multiple definitions of marker kubebuilder:validation:items:MinItems when only a single definition should exist"

	// +kubebuilder:validation:items:MinLength:=5
	UniqueItemsMinLength []string

	// +kubebuilder:validation:items:MinLength:=5
	// +kubebuilder:validation:items:MinLength:=10
	NonUniqueItemsMinLength []string // want "field A.NonUniqueItemsMinLength has multiple definitions of marker kubebuilder:validation:items:MinLength when only a single definition should exist"

	// +kubebuilder:validation:items:MinLength:=3
	NonUniqueItemsMinLengthFromAliasWithItemsMinLength UniqueItemsMinLengthAlias // want "field A.NonUniqueItemsMinLengthFromAliasWithItemsMinLength has multiple definitions of marker kubebuilder:validation:items:MinLength when only a single definition should exist"

	NonUniqueItemsMinLengthOnlyFromAliasWithItemsMinLength NonUniqueItemsMinLengthAlias // want "field A.NonUniqueItemsMinLengthOnlyFromAliasWithItemsMinLength has multiple definitions of marker kubebuilder:validation:items:MinLength when only a single definition should exist"

	// +kubebuilder:validation:items:MinProperties:=1
	UniqueItemsMinProperties []map[string]string

	// +kubebuilder:validation:items:MinProperties:=1
	// +kubebuilder:validation:items:MinProperties:=2
	NonUniqueItemsMinProperties []map[string]string // want "field A.NonUniqueItemsMinProperties has multiple definitions of marker kubebuilder:validation:items:MinProperties when only a single definition should exist"

	// +kubebuilder:validation:items:MinProperties:=3
	NonUniqueItemsMinPropertiesFromAliasWithItemsMinProperties UniqueItemsMinPropertiesAlias // want "field A.NonUniqueItemsMinPropertiesFromAliasWithItemsMinProperties has multiple definitions of marker kubebuilder:validation:items:MinProperties when only a single definition should exist"

	NonUniqueItemsMinPropertiesOnlyFromAliasWithItemsMinProperties NonUniqueItemsMinPropertiesAlias // want "field A.NonUniqueItemsMinPropertiesOnlyFromAliasWithItemsMinProperties has multiple definitions of marker kubebuilder:validation:items:MinProperties when only a single definition should exist"

	// +kubebuilder:validation:items:Minimum:=0
	UniqueItemsMinimum []int

	// +kubebuilder:validation:items:Minimum:=0
	// +kubebuilder:validation:items:Minimum:=-5
	NonUniqueItemsMinimum []int // want "field A.NonUniqueItemsMinimum has multiple definitions of marker kubebuilder:validation:items:Minimum when only a single definition should exist"

	// +kubebuilder:validation:items:Minimum:=10
	NonUniqueItemsMinimumFromAliasWithItemsMinimum UniqueItemsMinimumAlias // want "field A.NonUniqueItemsMinimumFromAliasWithItemsMinimum has multiple definitions of marker kubebuilder:validation:items:Minimum when only a single definition should exist"

	NonUniqueItemsMinimumOnlyFromAliasWithItemsMinimum NonUniqueItemsMinimumAlias // want "field A.NonUniqueItemsMinimumOnlyFromAliasWithItemsMinimum has multiple definitions of marker kubebuilder:validation:items:Minimum when only a single definition should exist"

	// +kubebuilder:validation:items:MultipleOf:=3
	UniqueItemsMultipleOf []int

	// +kubebuilder:validation:items:MultipleOf:=3
	// +kubebuilder:validation:items:MultipleOf:=5
	NonUniqueItemsMultipleOf []int // want "field A.NonUniqueItemsMultipleOf has multiple definitions of marker kubebuilder:validation:items:MultipleOf when only a single definition should exist"

	// +kubebuilder:validation:items:MultipleOf:=2
	NonUniqueItemsMultipleOfFromAliasWithItemsMultipleOf UniqueItemsMultipleOfAlias // want "field A.NonUniqueItemsMultipleOfFromAliasWithItemsMultipleOf has multiple definitions of marker kubebuilder:validation:items:MultipleOf when only a single definition should exist"

	NonUniqueItemsMultipleOfOnlyFromAliasWithItemsMultipleOf NonUniqueItemsMultipleOfAlias // want "field A.NonUniqueItemsMultipleOfOnlyFromAliasWithItemsMultipleOf has multiple definitions of marker kubebuilder:validation:items:MultipleOf when only a single definition should exist"

	// +kubebuilder:validation:items:Pattern:="^[a-z]+$"
	UniqueItemsPattern []string

	// +kubebuilder:validation:items:Pattern:="^[a-z]+$"
	// +kubebuilder:validation:items:Pattern:="^[0-9]+$"
	NonUniqueItemsPattern []string // want "field A.NonUniqueItemsPattern has multiple definitions of marker kubebuilder:validation:items:Pattern when only a single definition should exist"

	// +kubebuilder:validation:items:Pattern:="^[A-Z]+$"
	NonUniqueItemsPatternFromAliasWithItemsPattern UniqueItemsPatternAlias // want "field A.NonUniqueItemsPatternFromAliasWithItemsPattern has multiple definitions of marker kubebuilder:validation:items:Pattern when only a single definition should exist"

	NonUniqueItemsPatternOnlyFromAliasWithItemsPattern NonUniqueItemsPatternAlias // want "field A.NonUniqueItemsPatternOnlyFromAliasWithItemsPattern has multiple definitions of marker kubebuilder:validation:items:Pattern when only a single definition should exist"

	// +kubebuilder:validation:items:Type:="string"
	UniqueItemsType []string

	// +kubebuilder:validation:items:Type:="string"
	// +kubebuilder:validation:items:Type:="date"
	NonUniqueItemsType []string // want "field A.NonUniqueItemsType has multiple definitions of marker kubebuilder:validation:items:Type when only a single definition should exist"

	// +kubebuilder:validation:items:Type:="string"
	NonUniqueItemsTypeFromAliasWithItemsType UniqueItemsTypeAlias // want "field A.NonUniqueItemsTypeFromAliasWithItemsType has multiple definitions of marker kubebuilder:validation:items:Type when only a single definition should exist"

	NonUniqueItemsTypeOnlyFromAliasWithItemsType NonUniqueItemsTypeAlias // want "field A.NonUniqueItemsTypeOnlyFromAliasWithItemsType has multiple definitions of marker kubebuilder:validation:items:Type when only a single definition should exist"

	// +kubebuilder:validation:items:UniqueItems:=true
	UniqueItemsUniqueItems [][]string

	// +kubebuilder:validation:items:UniqueItems:=true
	// +kubebuilder:validation:items:UniqueItems:=false
	NonUniqueItemsUniqueItems [][]string // want "field A.NonUniqueItemsUniqueItems has multiple definitions of marker kubebuilder:validation:items:UniqueItems when only a single definition should exist"

	// +kubebuilder:validation:items:UniqueItems:=true
	NonUniqueItemsUniqueItemsFromAliasWithItemsUniqueItems UniqueItemsUniqueItemsAlias // want "field A.NonUniqueItemsUniqueItemsFromAliasWithItemsUniqueItems has multiple definitions of marker kubebuilder:validation:items:UniqueItems when only a single definition should exist"

	NonUniqueItemsUniqueItemsOnlyFromAliasWithItemsUniqueItems NonUniqueItemsUniqueItemsAlias // want "field A.NonUniqueItemsUniqueItemsOnlyFromAliasWithItemsUniqueItems has multiple definitions of marker kubebuilder:validation:items:UniqueItems when only a single definition should exist"
}

// +default:value="bart"
type UniqueBasicDefaultAlias string

// +default:value="homer"
// +default:value="bart"
type NonUniqueBasicDefaultAlias string // want "type NonUniqueBasicDefaultAlias has multiple definitions of marker default when only a single definition should exist"

// +kubebuilder:default:="bart"
type UniqueKubebuilderDefaultAlias string

// +kubebuilder:default:="homer"
// +kubebuilder:default:="bart"
type NonUniqueKubebuilderDefaultAlias string // want "type NonUniqueKubebuilderDefaultAlias has multiple definitions of marker kubebuilder:default when only a single definition should exist"

// +kubebuilder:example:="bart"
type UniqueKubebuilderExampleAlias string

// +kubebuilder:example:="homer"
// +kubebuilder:example:="bart"
type NonUniqueKubebuilderExampleAlias string // want "type NonUniqueKubebuilderExampleAlias has multiple definitions of marker kubebuilder:example when only a single definition should exist"

// +kubebuilder:validation:MinLength:=5
type UniqueMinLengthStringAlias string

// +kubebuilder:validation:MinLength:=5
// +kubebuilder:validation:MinLength:=1
type NonUniqueMinLengthStringAlias string // want "type NonUniqueMinLengthStringAlias has multiple definitions of marker kubebuilder:validation:MinLength when only a single definition should exist"

// +kubebuilder:validation:Enum:=Baz;Qux
type UniqueEnumAlias string

// +kubebuilder:validation:Enum:=Foo;Bar
// +kubebuilder:validation:Enum=Baz;Qux
type NonUniqueEnumAlias string // want "type NonUniqueEnumAlias has multiple definitions of marker kubebuilder:validation:Enum when only a single definition should exist"

// +kubebuilder:validation:Maximum:=5
// +kubebuilder:validation:ExclusiveMaximum:=false
type UniqueExclusiveMaximumAlias int

// +kubebuilder:validation:Maximum:=5
// +kubebuilder:validation:ExclusiveMaximum:=true
// +kubebuilder:validation:ExclusiveMaximum:=false
type NonUniqueExclusiveMaximumAlias int // want "type NonUniqueExclusiveMaximumAlias has multiple definitions of marker kubebuilder:validation:ExclusiveMaximum when only a single definition should exist"

// +kubebuilder:validation:Minimum:=2
// +kubebuilder:validation:ExclusiveMinimum:=false
type UniqueExclusiveMinimumAlias int

// +kubebuilder:validation:Minimum:=2
// +kubebuilder:validation:ExclusiveMinimum:=true
// +kubebuilder:validation:ExclusiveMinimum:=false
type NonUniqueExclusiveMinimumAlias int // want "type NonUniqueExclusiveMinimumAlias has multiple definitions of marker kubebuilder:validation:ExclusiveMinimum when only a single definition should exist"

// +kubebuilder:validation:Format:="date"
type UniqueFormatAlias string

// +kubebuilder:validation:Format:="date"
// +kubebuilder:validation:Format:="password"
type NonUniqueFormatAlias string // want "type NonUniqueFormatAlias has multiple definitions of marker kubebuilder:validation:Format when only a single definition should exist"

// +kubebuilder:validation:MaxItems:=3
type UniqueMaxItemsAlias []string

// +kubebuilder:validation:MaxItems:=3
// +kubebuilder:validation:MaxItems:=7
type NonUniqueMaxItemsAlias []string // want "type NonUniqueMaxItemsAlias has multiple definitions of marker kubebuilder:validation:MaxItems when only a single definition should exist"

// +kubebuilder:validation:MaxProperties:=2
type UniqueMaxPropertiesAlias map[string]string

// +kubebuilder:validation:MaxProperties:=2
// +kubebuilder:validation:MaxProperties:=6
type NonUniqueMaxPropertiesAlias map[string]string // want "type NonUniqueMaxPropertiesAlias has multiple definitions of marker kubebuilder:validation:MaxProperties when only a single definition should exist"

// +kubebuilder:validation:Maximum:=50
type UniqueMaximumAlias int

// +kubebuilder:validation:Maximum:=50
// +kubebuilder:validation:Maximum:=75
type NonUniqueMaximumAlias int // want "type NonUniqueMaximumAlias has multiple definitions of marker kubebuilder:validation:Maximum when only a single definition should exist"

// +kubebuilder:validation:MinItems:=4
type UniqueMinItemsAlias []string

// +kubebuilder:validation:MinItems:=4
// +kubebuilder:validation:MinItems:=5
type NonUniqueMinItemsAlias []string // want "type NonUniqueMinItemsAlias has multiple definitions of marker kubebuilder:validation:MinItems when only a single definition should exist"

// +kubebuilder:validation:MinLength:=40
type UniqueMinLengthCustomAlias string

// +kubebuilder:validation:MinLength:=40
// +kubebuilder:validation:MinLength:=45
type NonUniqueMinLengthCustomAlias string // want "type NonUniqueMinLengthCustomAlias has multiple definitions of marker kubebuilder:validation:MinLength when only a single definition should exist"

// +kubebuilder:validation:MinProperties:=4
type UniqueMinPropertiesFieldAlias map[string]string

// +kubebuilder:validation:MinProperties:=4
// +kubebuilder:validation:MinProperties:=5
type NonUniqueMinPropertiesFieldAlias map[string]string // want "type NonUniqueMinPropertiesFieldAlias has multiple definitions of marker kubebuilder:validation:MinProperties when only a single definition should exist"

// +kubebuilder:validation:Minimum:=5
type UniqueMinimumAlias int

// +kubebuilder:validation:Minimum:=5
// +kubebuilder:validation:Minimum:=7
type NonUniqueMinimumAlias int // want "type NonUniqueMinimumAlias has multiple definitions of marker kubebuilder:validation:Minimum when only a single definition should exist"

// +kubebuilder:validation:MultipleOf:=7
type UniqueMultipleOfAlias int

// +kubebuilder:validation:MultipleOf:=7
// +kubebuilder:validation:MultipleOf:=11
type NonUniqueMultipleOfAlias int // want "type NonUniqueMultipleOfAlias has multiple definitions of marker kubebuilder:validation:MultipleOf when only a single definition should exist"

// +kubebuilder:validation:Pattern:="^\\w+$"
type UniquePatternAlias string

// +kubebuilder:validation:Pattern:="^\\w+$"
// +kubebuilder:validation:Pattern:="^\\d+$"
type NonUniquePatternAlias string // want "type NonUniquePatternAlias has multiple definitions of marker kubebuilder:validation:Pattern when only a single definition should exist"

// +kubebuilder:validation:Type:="string"
type UniqueTypeValidationAlias string

// +kubebuilder:validation:Type:="string"
// +kubebuilder:validation:Type:="date"
type NonUniqueTypeValidationAlias string // want "type NonUniqueTypeValidationAlias has multiple definitions of marker kubebuilder:validation:Type when only a single definition should exist"

// +kubebuilder:validation:UniqueItems:=false
type UniqueUniqueItemsAlias []string

// +kubebuilder:validation:UniqueItems:=true
// +kubebuilder:validation:UniqueItems:=false
type NonUniqueUniqueItemsAlias []string // want "type NonUniqueUniqueItemsAlias has multiple definitions of marker kubebuilder:validation:UniqueItems when only a single definition should exist"

// +kubebuilder:validation:items:Enum=Banana;Grape
type UniqueItemsEnumAlias []string

// +kubebuilder:validation:items:Enum=Apple;Orange
// +kubebuilder:validation:items:Enum=Banana;Grape
type NonUniqueItemsEnumAlias []string // want "type NonUniqueItemsEnumAlias has multiple definitions of marker kubebuilder:validation:items:Enum when only a single definition should exist"

// +kubebuilder:validation:items:Maximum:=5
// +kubebuilder:validation:items:ExclusiveMaximum:=false
type UniqueItemsExclusiveMaximumAlias []int

// +kubebuilder:validation:items:Maximum:=5
// +kubebuilder:validation:items:ExclusiveMaximum:=true
// +kubebuilder:validation:items:ExclusiveMaximum:=false
type NonUniqueItemsExclusiveMaximumAlias []int // want "type NonUniqueItemsExclusiveMaximumAlias has multiple definitions of marker kubebuilder:validation:items:ExclusiveMaximum when only a single definition should exist"

// +kubebuilder:validation:items:Minimum:=2
// +kubebuilder:validation:items:ExclusiveMinimum:=false
type UniqueItemsExclusiveMinimumAlias []int

// +kubebuilder:validation:items:Minimum:=2
// +kubebuilder:validation:items:ExclusiveMinimum:=true
// +kubebuilder:validation:items:ExclusiveMinimum:=false
type NonUniqueItemsExclusiveMinimumAlias []int // want "type NonUniqueItemsExclusiveMinimumAlias has multiple definitions of marker kubebuilder:validation:items:ExclusiveMinimum when only a single definition should exist"

// +kubebuilder:validation:items:Format:="password"
type UniqueItemsFormatAlias []string

// +kubebuilder:validation:items:Format:="date-time"
// +kubebuilder:validation:items:Format:="password"
type NonUniqueItemsFormatAlias []string // want "type NonUniqueItemsFormatAlias has multiple definitions of marker kubebuilder:validation:items:Format when only a single definition should exist"

// +kubebuilder:validation:items:MaxItems:=2
type UniqueItemsMaxItemsAlias [][]string

// +kubebuilder:validation:items:MaxItems:=2
// +kubebuilder:validation:items:MaxItems:=4
type NonUniqueItemsMaxItemsAlias [][]string // want "type NonUniqueItemsMaxItemsAlias has multiple definitions of marker kubebuilder:validation:items:MaxItems when only a single definition should exist"

// +kubebuilder:validation:items:MaxProperties:=1
type UniqueItemsMaxPropertiesAlias []map[string]string

// +kubebuilder:validation:items:MaxProperties:=1
// +kubebuilder:validation:items:MaxProperties:=4
type NonUniqueItemsMaxPropertiesAlias []map[string]string // want "type NonUniqueItemsMaxPropertiesAlias has multiple definitions of marker kubebuilder:validation:items:MaxProperties when only a single definition should exist"

// kubebuilder:validation:items:Maximum
// +kubebuilder:validation:items:Maximum:=25
type UniqueItemsMaximumAlias []int

// +kubebuilder:validation:items:Maximum:=25
// +kubebuilder:validation:items:Maximum:=10
type NonUniqueItemsMaximumAlias []int // want "type NonUniqueItemsMaximumAlias has multiple definitions of marker kubebuilder:validation:items:Maximum when only a single definition should exist"

// +kubebuilder:validation:items:MinItems:=4
type UniqueItemsMinItemsAlias [][]string

// +kubebuilder:validation:items:MinItems:=4
// +kubebuilder:validation:items:MinItems:=2
type NonUniqueItemsMinItemsAlias [][]string // want "type NonUniqueItemsMinItemsAlias has multiple definitions of marker kubebuilder:validation:items:MinItems when only a single definition should exist"

// +kubebuilder:validation:items:MinLength:=8
type UniqueItemsMinLengthAlias []string

// +kubebuilder:validation:items:MinLength:=8
// +kubebuilder:validation:items:MinLength:=2
type NonUniqueItemsMinLengthAlias []string // want "type NonUniqueItemsMinLengthAlias has multiple definitions of marker kubebuilder:validation:items:MinLength when only a single definition should exist"

// +kubebuilder:validation:items:MinProperties:=4
type UniqueItemsMinPropertiesAlias []map[string]string

// +kubebuilder:validation:items:MinProperties:=4
// +kubebuilder:validation:items:MinProperties:=2
type NonUniqueItemsMinPropertiesAlias []map[string]string // want "type NonUniqueItemsMinPropertiesAlias has multiple definitions of marker kubebuilder:validation:items:MinProperties when only a single definition should exist"

// +kubebuilder:validation:items:Minimum:=20
type UniqueItemsMinimumAlias []int

// +kubebuilder:validation:items:Minimum:=20
// +kubebuilder:validation:items:Minimum:=30
type NonUniqueItemsMinimumAlias []int // want "type NonUniqueItemsMinimumAlias has multiple definitions of marker kubebuilder:validation:items:Minimum when only a single definition should exist"

// +kubebuilder:validation:items:MultipleOf:=7
type UniqueItemsMultipleOfAlias []int

// +kubebuilder:validation:items:MultipleOf:=7
// +kubebuilder:validation:items:MultipleOf:=11
type NonUniqueItemsMultipleOfAlias []int // want "type NonUniqueItemsMultipleOfAlias has multiple definitions of marker kubebuilder:validation:items:MultipleOf when only a single definition should exist"

// kubebuilder:validation:items:Pattern
// +kubebuilder:validation:items:Pattern:="^\\w+$"
type UniqueItemsPatternAlias []string

// +kubebuilder:validation:items:Pattern:="^\\w+$"
// +kubebuilder:validation:items:Pattern:="^\\d+$"
type NonUniqueItemsPatternAlias []string // want "type NonUniqueItemsPatternAlias has multiple definitions of marker kubebuilder:validation:items:Pattern when only a single definition should exist"

// +kubebuilder:validation:items:Type:="date"
type UniqueItemsTypeAlias []string

// +kubebuilder:validation:items:Type:="string"
// +kubebuilder:validation:items:Type:="date"
type NonUniqueItemsTypeAlias []string // want "type NonUniqueItemsTypeAlias has multiple definitions of marker kubebuilder:validation:items:Type when only a single definition should exist"

// +kubebuilder:validation:items:UniqueItems:=false
type UniqueItemsUniqueItemsAlias [][]string

// +kubebuilder:validation:items:UniqueItems:=true
// +kubebuilder:validation:items:UniqueItems:=false
type NonUniqueItemsUniqueItemsAlias [][]string // want "type NonUniqueItemsUniqueItemsAlias has multiple definitions of marker kubebuilder:validation:items:UniqueItems when only a single definition should exist"

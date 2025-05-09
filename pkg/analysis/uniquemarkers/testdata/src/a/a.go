package a

// +kubebuilder:validation:MinProperties:=3
// +kubebuilder:validation:MinProperties:=5
type A struct { // want "type A has multiple definitions of marker kubebuilder:validation:MinProperties when only a single definition should exist"

	// +kubebuilder:validation:XValidation:rule='self.matches("[0-9A-Za-z]")',message='should match regex'
	// +kubebuilder:validation:XValidation:rule='isURL(self)',message='should be a URL'
	AllowedNonUniqueMarkers string

	// +kubebuilder:validation:XValidation:rule='self.matches("[0-9]")',message='should actually be a number'
	// +kubebuilder:validation:XValidation:rule='self.matches("[0-9]")',message='different message'
	NonUniqueXValidations string // want "field NonUniqueXValidations has multiple definitions of marker kubebuilder:validation:XValidation:rule='self\\.matches\\(\\\"\\[0-9]\\\"\\)' when only a single definition should exist"

	// +kubebuilder:validation:items:XValidation:rule='self.matches("[0-9]")',message='should actually be a number'
	// +kubebuilder:validation:items:XValidation:rule='self.matches("[0-9]")',message='different message'
	NonUniqueItemsXValidations []string // want "field NonUniqueItemsXValidations has multiple definitions of marker kubebuilder:validation:items:XValidation:rule='self\\.matches\\(\\\"\\[0-9]\\\"\\)' when only a single definition should exist" 

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
	NonUniqueBasicDefault string // want "field NonUniqueBasicDefault has multiple definitions of marker default when only a single definition should exist"

	// +default:value="homer"
	NonUniqueBasicDefaultFromAliasWithBasicDefault UniqueBasicDefaultAlias // want "field NonUniqueBasicDefaultFromAliasWithBasicDefault has multiple definitions of marker default when only a single definition should exist"

	NonUniqueBasicDefaultOnlyFromAliasWithBasicDefault NonUniqueBasicDefaultAlias // want "field NonUniqueBasicDefaultOnlyFromAliasWithBasicDefault has multiple definitions of marker default when only a single definition should exist"

	// +kubebuilder:default:="homer"
	UniqueKubebuilderDefault string

	// +kubebuilder:default:="homer"
	// +kubebuilder:default:="bart"
	NonUniqueKubebuilderDefault string // want "field NonUniqueKubebuilderDefault has multiple definitions of marker kubebuilder:default when only a single definition should exist"

	// +kubebuilder:default:="homer"
	NonUniqueKubebuilderDefaultFromAliasWithKubebuilderDefault UniqueKubebuilderDefaultAlias // want "field NonUniqueKubebuilderDefaultFromAliasWithKubebuilderDefault has multiple definitions of marker kubebuilder:default when only a single definition should exist"

	NonUniqueKubebuilderDefaultOnlyFromAliasWithKubebuilderDefault NonUniqueKubebuilderDefaultAlias // want "field NonUniqueKubebuilderDefaultOnlyFromAliasWithKubebuilderDefault has multiple definitions of marker kubebuilder:default when only a single definition should exist"

	// +kubebuilder:example:="homer"
	UniqueKubebuilderExample string

	// +kubebuilder:example:="homer"
	// +kubebuilder:example:="bart"
	NonUniqueKubebuilderExample string // want "field NonUniqueKubebuilderExample has multiple definitions of marker kubebuilder:example when only a single definition should exist"

	// +kubebuilder:example:="homer"
	NonUniqueKubebuilderExampleFromAliasWithKubebuilderExample UniqueKubebuilderExampleAlias // want "field NonUniqueKubebuilderExampleFromAliasWithKubebuilderExample has multiple definitions of marker kubebuilder:example when only a single definition should exist"

	NonUniqueKubebuilderExampleOnlyFromAliasWithKubebuilderExample NonUniqueKubebuilderExampleAlias // want "field NonUniqueKubebuilderExampleOnlyFromAliasWithKubebuilderExample has multiple definitions of marker kubebuilder:example when only a single definition should exist"

	// +kubebuilder:validation:MinLength:=1
	UniqueMinLength string

	// +kubebuilder:validation:MinLength:=1
	// +kubebuilder:validation:MinLength:=5
	NonUniqueMinLength string // want "field NonUniqueMinLength has multiple definitions of marker kubebuilder:validation:MinLength when only a single definition should exist"

	// +kubebuilder:validation:MinLength:=1
	NonUniqueMinLengthFromAliasWithMinLength UniqueMinLengthStringAlias // want "field NonUniqueMinLengthFromAliasWithMinLength has multiple definitions of marker kubebuilder:validation:MinLength when only a single definition should exist"

	NonUniqueMinLengthOnlyFromAliasWithMinLength NonUniqueMinLengthStringAlias // want "field NonUniqueMinLengthOnlyFromAliasWithMinLength has multiple definitions of marker kubebuilder:validation:MinLength when only a single definition should exist"

	// +kubebuilder:validation:Enum:=Foo;Bar
	UniqueEnum string

	// +kubebuilder:validation:Enum:=Foo;Bar
	// +kubebuilder:validation:Enum:=Baz;Qux
	NonUniqueEnum string // want "field NonUniqueEnum has multiple definitions of marker kubebuilder:validation:Enum when only a single definition should exist"

	// +kubebuilder:validation:Enum:=Foo;Bar
	NonUniqueEnumFromAliasWithEnum UniqueEnumAlias // want "field NonUniqueEnumFromAliasWithEnum has multiple definitions of marker kubebuilder:validation:Enum when only a single definition should exist"

	NonUniqueEnumOnlyFromAliasWithEnum NonUniqueEnumAlias // want "field NonUniqueEnumOnlyFromAliasWithEnum has multiple definitions of marker kubebuilder:validation:Enum when only a single definition should exist"

	// +kubebuilder:validation:Maximum:=10
	// +kubebuilder:validation:ExclusiveMaximum:=true
	UniqueExclusiveMaximum int

	// +kubebuilder:validation:Maximum:=10
	// +kubebuilder:validation:ExclusiveMaximum:=true
	// +kubebuilder:validation:ExclusiveMaximum:=false
	NonUniqueExclusiveMaximum int // want "field NonUniqueExclusiveMaximum has multiple definitions of marker kubebuilder:validation:ExclusiveMaximum when only a single definition should exist"

	// +kubebuilder:validation:ExclusiveMaximum:=true
	NonUniqueExclusiveMaximumFromAliasWithExclusiveMaximum UniqueExclusiveMaximumAlias // want "field NonUniqueExclusiveMaximumFromAliasWithExclusiveMaximum has multiple definitions of marker kubebuilder:validation:ExclusiveMaximum when only a single definition should exist"

	NonUniqueExclusiveMaximumOnlyFromAliasWithExclusiveMaximum NonUniqueExclusiveMaximumAlias // want "field NonUniqueExclusiveMaximumOnlyFromAliasWithExclusiveMaximum has multiple definitions of marker kubebuilder:validation:ExclusiveMaximum when only a single definition should exist"

	// +kubebuilder:validation:Minimum:=2
	// +kubebuilder:validation:ExclusiveMinimum:=true
	UniqueExclusiveMinimum int

	// +kubebuilder:validation:Minimum:=2
	// +kubebuilder:validation:ExclusiveMinimum:=true
	// +kubebuilder:validation:ExclusiveMinimum:=false
	NonUniqueExclusiveMinimum int // want "field NonUniqueExclusiveMinimum has multiple definitions of marker kubebuilder:validation:ExclusiveMinimum when only a single definition should exist"

	// +kubebuilder:validation:ExclusiveMinimum:=true
	NonUniqueExclusiveMinimumFromAliasWithExclusiveMinimum UniqueExclusiveMinimumAlias // want "field NonUniqueExclusiveMinimumFromAliasWithExclusiveMinimum has multiple definitions of marker kubebuilder:validation:ExclusiveMinimum when only a single definition should exist"

	NonUniqueExclusiveMinimumOnlyFromAliasWithExclusiveMinimum NonUniqueExclusiveMinimumAlias // want "field NonUniqueExclusiveMinimumOnlyFromAliasWithExclusiveMinimum has multiple definitions of marker kubebuilder:validation:ExclusiveMinimum when only a single definition should exist"

	// +kubebuilder:validation:Format:="date-time"
	UniqueFormat string

	// +kubebuilder:validation:Format:="date-time"
	// +kubebuilder:validation:Format:="password"
	NonUniqueFormat string // want "field NonUniqueFormat has multiple definitions of marker kubebuilder:validation:Format when only a single definition should exist"

	// +kubebuilder:validation:Format:="password"
	NonUniqueFormatFromAliasWithFormat UniqueFormatAlias // want "field NonUniqueFormatFromAliasWithFormat has multiple definitions of marker kubebuilder:validation:Format when only a single definition should exist"

	NonUniqueFormatOnlyFromAliasWithFormat NonUniqueFormatAlias // want "field NonUniqueFormatOnlyFromAliasWithFormat has multiple definitions of marker kubebuilder:validation:Format when only a single definition should exist"

	// +kubebuilder:validation:MaxItems:=10
	UniqueMaxItems []string

	// +kubebuilder:validation:MaxItems:=10
	// +kubebuilder:validation:MaxItems:=5
	NonUniqueMaxItems []string // want "field NonUniqueMaxItems has multiple definitions of marker kubebuilder:validation:MaxItems when only a single definition should exist"

	// +kubebuilder:validation:MaxItems:=8
	NonUniqueMaxItemsFromAliasWithMaxItems UniqueMaxItemsAlias // want "field NonUniqueMaxItemsFromAliasWithMaxItems has multiple definitions of marker kubebuilder:validation:MaxItems when only a single definition should exist"

	NonUniqueMaxItemsOnlyFromAliasWithMaxItems NonUniqueMaxItemsAlias // want "field NonUniqueMaxItemsOnlyFromAliasWithMaxItems has multiple definitions of marker kubebuilder:validation:MaxItems when only a single definition should exist"

	// +kubebuilder:validation:MaxProperties:=3
	UniqueMaxProperties map[string]string

	// +kubebuilder:validation:MaxProperties:=3
	// +kubebuilder:validation:MaxProperties:=5
	NonUniqueMaxProperties map[string]string // want "field NonUniqueMaxProperties has multiple definitions of marker kubebuilder:validation:MaxProperties when only a single definition should exist"

	// +kubebuilder:validation:MaxProperties:=4
	NonUniqueMaxPropertiesFromAliasWithMaxProperties UniqueMaxPropertiesAlias // want "field NonUniqueMaxPropertiesFromAliasWithMaxProperties has multiple definitions of marker kubebuilder:validation:MaxProperties when only a single definition should exist"

	NonUniqueMaxPropertiesOnlyFromAliasWithMaxProperties NonUniqueMaxPropertiesAlias // want "field NonUniqueMaxPropertiesOnlyFromAliasWithMaxProperties has multiple definitions of marker kubebuilder:validation:MaxProperties when only a single definition should exist"

	// +kubebuilder:validation:Maximum:=100
	UniqueMaximum int

	// +kubebuilder:validation:Maximum:=100
	// +kubebuilder:validation:Maximum:=200
	NonUniqueMaximum int // want "field NonUniqueMaximum has multiple definitions of marker kubebuilder:validation:Maximum when only a single definition should exist"

	// +kubebuilder:validation:Maximum:=150
	NonUniqueMaximumFromAliasWithMaximum UniqueMaximumAlias // want "field NonUniqueMaximumFromAliasWithMaximum has multiple definitions of marker kubebuilder:validation:Maximum when only a single definition should exist"

	NonUniqueMaximumOnlyFromAliasWithMaximum NonUniqueMaximumAlias // want "field NonUniqueMaximumOnlyFromAliasWithMaximum has multiple definitions of marker kubebuilder:validation:Maximum when only a single definition should exist"

	// +kubebuilder:validation:MinItems:=1
	UniqueMinItems []string

	// +kubebuilder:validation:MinItems:=1
	// +kubebuilder:validation:MinItems:=2
	NonUniqueMinItems []string // want "field NonUniqueMinItems has multiple definitions of marker kubebuilder:validation:MinItems when only a single definition should exist"

	// +kubebuilder:validation:MinItems:=3
	NonUniqueMinItemsFromAliasWithMinItems UniqueMinItemsAlias // want "field NonUniqueMinItemsFromAliasWithMinItems has multiple definitions of marker kubebuilder:validation:MinItems when only a single definition should exist"

	NonUniqueMinItemsOnlyFromAliasWithMinItems NonUniqueMinItemsAlias // want "field NonUniqueMinItemsOnlyFromAliasWithMinItems has multiple definitions of marker kubebuilder:validation:MinItems when only a single definition should exist"

	// +kubebuilder:validation:MinLength:=30
	UniqueMinLengthCustom string

	// +kubebuilder:validation:MinLength:=30
	// +kubebuilder:validation:MinLength:=35
	NonUniqueMinLengthCustom string // want "field NonUniqueMinLengthCustom has multiple definitions of marker kubebuilder:validation:MinLength when only a single definition should exist"

	// +kubebuilder:validation:MinLength:=30
	NonUniqueMinLengthCustomFromAliasWithMinLength UniqueMinLengthCustomAlias // want "field NonUniqueMinLengthCustomFromAliasWithMinLength has multiple definitions of marker kubebuilder:validation:MinLength when only a single definition should exist"

	NonUniqueMinLengthCustomOnlyFromAliasWithMinLength NonUniqueMinLengthCustomAlias // want "field NonUniqueMinLengthCustomOnlyFromAliasWithMinLength has multiple definitions of marker kubebuilder:validation:MinLength when only a single definition should exist"

	// +kubebuilder:validation:MinProperties:=1
	UniqueMinPropertiesField map[string]string

	// +kubebuilder:validation:MinProperties:=1
	// +kubebuilder:validation:MinProperties:=2
	NonUniqueMinPropertiesField map[string]string // want "field NonUniqueMinPropertiesField has multiple definitions of marker kubebuilder:validation:MinProperties when only a single definition should exist"

	// +kubebuilder:validation:MinProperties:=1
	NonUniqueMinPropertiesFieldFromAliasWithMinProperties UniqueMinPropertiesFieldAlias // want "field NonUniqueMinPropertiesFieldFromAliasWithMinProperties has multiple definitions of marker kubebuilder:validation:MinProperties when only a single definition should exist"

	NonUniqueMinPropertiesFieldOnlyFromAliasWithMinProperties NonUniqueMinPropertiesFieldAlias // want "field NonUniqueMinPropertiesFieldOnlyFromAliasWithMinProperties has multiple definitions of marker kubebuilder:validation:MinProperties when only a single definition should exist"

	// +kubebuilder:validation:Minimum:=10
	UniqueMinimum int

	// +kubebuilder:validation:Minimum:=10
	// +kubebuilder:validation:Minimum:=20
	NonUniqueMinimum int // want "field NonUniqueMinimum has multiple definitions of marker kubebuilder:validation:Minimum when only a single definition should exist"

	// +kubebuilder:validation:Minimum:=15
	NonUniqueMinimumFromAliasWithMinimum UniqueMinimumAlias // want "field NonUniqueMinimumFromAliasWithMinimum has multiple definitions of marker kubebuilder:validation:Minimum when only a single definition should exist"

	NonUniqueMinimumOnlyFromAliasWithMinimum NonUniqueMinimumAlias // want "field NonUniqueMinimumOnlyFromAliasWithMinimum has multiple definitions of marker kubebuilder:validation:Minimum when only a single definition should exist"

	// +kubebuilder:validation:MultipleOf:=3
	UniqueMultipleOf int

	// +kubebuilder:validation:MultipleOf:=3
	// +kubebuilder:validation:MultipleOf:=5
	NonUniqueMultipleOf int // want "field NonUniqueMultipleOf has multiple definitions of marker kubebuilder:validation:MultipleOf when only a single definition should exist"

	// +kubebuilder:validation:MultipleOf:=2
	NonUniqueMultipleOfFromAliasWithMultipleOf UniqueMultipleOfAlias // want "field NonUniqueMultipleOfFromAliasWithMultipleOf has multiple definitions of marker kubebuilder:validation:MultipleOf when only a single definition should exist"

	NonUniqueMultipleOfOnlyFromAliasWithMultipleOf NonUniqueMultipleOfAlias // want "field NonUniqueMultipleOfOnlyFromAliasWithMultipleOf has multiple definitions of marker kubebuilder:validation:MultipleOf when only a single definition should exist"

	// +kubebuilder:validation:Pattern:="^[a-z]+$"
	UniquePattern string

	// +kubebuilder:validation:Pattern:="^[a-z]+$"
	// +kubebuilder:validation:Pattern:="^[0-9]+$"
	NonUniquePattern string // want "field NonUniquePattern has multiple definitions of marker kubebuilder:validation:Pattern when only a single definition should exist"

	// +kubebuilder:validation:Pattern:="^[A-Z]+$"
	NonUniquePatternFromAliasWithPattern UniquePatternAlias // want "field NonUniquePatternFromAliasWithPattern has multiple definitions of marker kubebuilder:validation:Pattern when only a single definition should exist"

	NonUniquePatternOnlyFromAliasWithPattern NonUniquePatternAlias // want "field NonUniquePatternOnlyFromAliasWithPattern has multiple definitions of marker kubebuilder:validation:Pattern when only a single definition should exist"

	// +kubebuilder:validation:Type:="string"
	UniqueTypeValidationField string

	// +kubebuilder:validation:Type:="string"
	// +kubebuilder:validation:Type:="date"
	NonUniqueTypeValidationField string // want "field NonUniqueTypeValidationField has multiple definitions of marker kubebuilder:validation:Type when only a single definition should exist"

	// +kubebuilder:validation:Type:="date"
	NonUniqueTypeValidationFieldFromAliasWithType UniqueTypeValidationAlias // want "field NonUniqueTypeValidationFieldFromAliasWithType has multiple definitions of marker kubebuilder:validation:Type when only a single definition should exist"

	NonUniqueTypeValidationFieldOnlyFromAliasWithType NonUniqueTypeValidationAlias // want "field NonUniqueTypeValidationFieldOnlyFromAliasWithType has multiple definitions of marker kubebuilder:validation:Type when only a single definition should exist"

	// +kubebuilder:validation:UniqueItems:=true
	UniqueUniqueItems []string

	// +kubebuilder:validation:UniqueItems:=true
	// +kubebuilder:validation:UniqueItems:=false
	NonUniqueUniqueItems []string // want "field NonUniqueUniqueItems has multiple definitions of marker kubebuilder:validation:UniqueItems when only a single definition should exist"

	// +kubebuilder:validation:UniqueItems:=true
	NonUniqueUniqueItemsFromAliasWithUniqueItems UniqueUniqueItemsAlias // want "field NonUniqueUniqueItemsFromAliasWithUniqueItems has multiple definitions of marker kubebuilder:validation:UniqueItems when only a single definition should exist"

	NonUniqueUniqueItemsOnlyFromAliasWithUniqueItems NonUniqueUniqueItemsAlias // want "field NonUniqueUniqueItemsOnlyFromAliasWithUniqueItems has multiple definitions of marker kubebuilder:validation:UniqueItems when only a single definition should exist"

	// +kubebuilder:validation:items:Enum=Apple;Orange
	UniqueItemsEnum []string

	// +kubebuilder:validation:items:Enum=Apple;Orange
	// +kubebuilder:validation:items:Enum=Banana;Grape
	NonUniqueItemsEnum []string // want "field NonUniqueItemsEnum has multiple definitions of marker kubebuilder:validation:items:Enum when only a single definition should exist"

	// +kubebuilder:validation:items:Enum=Apple;Orange
	NonUniqueItemsEnumFromAliasWithItemsEnum UniqueItemsEnumAlias // want "field NonUniqueItemsEnumFromAliasWithItemsEnum has multiple definitions of marker kubebuilder:validation:items:Enum when only a single definition should exist"

	NonUniqueItemsEnumOnlyFromAliasWithItemsEnum NonUniqueItemsEnumAlias // want "field NonUniqueItemsEnumOnlyFromAliasWithItemsEnum has multiple definitions of marker kubebuilder:validation:items:Enum when only a single definition should exist"

	// +kubebuilder:validation:items:Maximum:=10
	// +kubebuilder:validation:items:ExclusiveMaximum:=true
	UniqueItemsExclusiveMaximum []int

	// +kubebuilder:validation:items:Maximum:=10
	// +kubebuilder:validation:items:ExclusiveMaximum:=true
	// +kubebuilder:validation:items:ExclusiveMaximum:=false
	NonUniqueItemsExclusiveMaximum []int // want "field NonUniqueItemsExclusiveMaximum has multiple definitions of marker kubebuilder:validation:items:ExclusiveMaximum when only a single definition should exist"

	// +kubebuilder:validation:items:ExclusiveMaximum:=true
	NonUniqueItemsExclusiveMaximumFromAliasWithItemsExclusiveMaximum UniqueItemsExclusiveMaximumAlias // want "field NonUniqueItemsExclusiveMaximumFromAliasWithItemsExclusiveMaximum has multiple definitions of marker kubebuilder:validation:items:ExclusiveMaximum when only a single definition should exist"

	NonUniqueItemsExclusiveMaximumOnlyFromAliasWithItemsExclusiveMaximum NonUniqueItemsExclusiveMaximumAlias // want "field NonUniqueItemsExclusiveMaximumOnlyFromAliasWithItemsExclusiveMaximum has multiple definitions of marker kubebuilder:validation:items:ExclusiveMaximum when only a single definition should exist"

	// +kubebuilder:validation:items:Minimum:=1
	// +kubebuilder:validation:items:ExclusiveMinimum:=true
	UniqueItemsExclusiveMinimum []int

	// +kubebuilder:validation:items:Minimum:=1
	// +kubebuilder:validation:items:ExclusiveMinimum:=true
	// +kubebuilder:validation:items:ExclusiveMinimum:=false
	NonUniqueItemsExclusiveMinimum []int // want "field NonUniqueItemsExclusiveMinimum has multiple definitions of marker kubebuilder:validation:items:ExclusiveMinimum when only a single definition should exist"

	// +kubebuilder:validation:items:ExclusiveMinimum:=true
	NonUniqueItemsExclusiveMinimumFromAliasWithItemsExclusiveMinimum UniqueItemsExclusiveMinimumAlias // want "field NonUniqueItemsExclusiveMinimumFromAliasWithItemsExclusiveMinimum has multiple definitions of marker kubebuilder:validation:items:ExclusiveMinimum when only a single definition should exist"

	NonUniqueItemsExclusiveMinimumOnlyFromAliasWithItemsExclusiveMinimum NonUniqueItemsExclusiveMinimumAlias // want "field NonUniqueItemsExclusiveMinimumOnlyFromAliasWithItemsExclusiveMinimum has multiple definitions of marker kubebuilder:validation:items:ExclusiveMinimum when only a single definition should exist"

	// +kubebuilder:validation:items:Format:="date-time"
	UniqueItemsFormat []string

	// +kubebuilder:validation:items:Format:="date-time"
	// +kubebuilder:validation:items:Format:="password"
	NonUniqueItemsFormat []string // want "field NonUniqueItemsFormat has multiple definitions of marker kubebuilder:validation:items:Format when only a single definition should exist"

	// +kubebuilder:validation:items:Format:="date-time"
	NonUniqueItemsFormatFromAliasWithItemsFormat UniqueItemsFormatAlias // want "field NonUniqueItemsFormatFromAliasWithItemsFormat has multiple definitions of marker kubebuilder:validation:items:Format when only a single definition should exist"

	NonUniqueItemsFormatOnlyFromAliasWithItemsFormat NonUniqueItemsFormatAlias // want "field NonUniqueItemsFormatOnlyFromAliasWithItemsFormat has multiple definitions of marker kubebuilder:validation:items:Format when only a single definition should exist"

	// +kubebuilder:validation:items:MaxItems:=5
	UniqueItemsMaxItems [][]string

	// +kubebuilder:validation:items:MaxItems:=5
	// +kubebuilder:validation:items:MaxItems:=10
	NonUniqueItemsMaxItems [][]string // want "field NonUniqueItemsMaxItems has multiple definitions of marker kubebuilder:validation:items:MaxItems when only a single definition should exist"

	// +kubebuilder:validation:items:MaxItems:=3
	NonUniqueItemsMaxItemsFromAliasWithItemsMaxItems UniqueItemsMaxItemsAlias // want "field NonUniqueItemsMaxItemsFromAliasWithItemsMaxItems has multiple definitions of marker kubebuilder:validation:items:MaxItems when only a single definition should exist"

	NonUniqueItemsMaxItemsOnlyFromAliasWithItemsMaxItems NonUniqueItemsMaxItemsAlias // want "field NonUniqueItemsMaxItemsOnlyFromAliasWithItemsMaxItems has multiple definitions of marker kubebuilder:validation:items:MaxItems when only a single definition should exist"

	// +kubebuilder:validation:items:MaxProperties:=3
	UniqueItemsMaxProperties []map[string]string

	// +kubebuilder:validation:items:MaxProperties:=3
	// +kubebuilder:validation:items:MaxProperties:=5
	NonUniqueItemsMaxProperties []map[string]string // want "field NonUniqueItemsMaxProperties has multiple definitions of marker kubebuilder:validation:items:MaxProperties when only a single definition should exist"

	// +kubebuilder:validation:items:MaxProperties:=2
	NonUniqueItemsMaxPropertiesFromAliasWithItemsMaxProperties UniqueItemsMaxPropertiesAlias // want "field NonUniqueItemsMaxPropertiesFromAliasWithItemsMaxProperties has multiple definitions of marker kubebuilder:validation:items:MaxProperties when only a single definition should exist"

	NonUniqueItemsMaxPropertiesOnlyFromAliasWithItemsMaxProperties NonUniqueItemsMaxPropertiesAlias // want "field NonUniqueItemsMaxPropertiesOnlyFromAliasWithItemsMaxProperties has multiple definitions of marker kubebuilder:validation:items:MaxProperties when only a single definition should exist"

	// +kubebuilder:validation:items:Maximum:=100
	UniqueItemsMaximum []int

	// +kubebuilder:validation:items:Maximum:=100
	// +kubebuilder:validation:items:Maximum:=50
	NonUniqueItemsMaximum []int // want "field NonUniqueItemsMaximum has multiple definitions of marker kubebuilder:validation:items:Maximum when only a single definition should exist"

	// +kubebuilder:validation:items:Maximum:=75
	NonUniqueItemsMaximumFromAliasWithItemsMaximum UniqueItemsMaximumAlias // want "field NonUniqueItemsMaximumFromAliasWithItemsMaximum has multiple definitions of marker kubebuilder:validation:items:Maximum when only a single definition should exist"

	NonUniqueItemsMaximumOnlyFromAliasWithItemsMaximum NonUniqueItemsMaximumAlias // want "field NonUniqueItemsMaximumOnlyFromAliasWithItemsMaximum has multiple definitions of marker kubebuilder:validation:items:Maximum when only a single definition should exist"

	// +kubebuilder:validation:items:MinItems:=1
	UniqueItemsMinItems [][]string

	// +kubebuilder:validation:items:MinItems:=1
	// +kubebuilder:validation:items:MinItems:=2
	NonUniqueItemsMinItems [][]string // want "field NonUniqueItemsMinItems has multiple definitions of marker kubebuilder:validation:items:MinItems when only a single definition should exist"

	// +kubebuilder:validation:items:MinItems:=3
	NonUniqueItemsMinItemsFromAliasWithItemsMinItems UniqueItemsMinItemsAlias // want "field NonUniqueItemsMinItemsFromAliasWithItemsMinItems has multiple definitions of marker kubebuilder:validation:items:MinItems when only a single definition should exist"

	NonUniqueItemsMinItemsOnlyFromAliasWithItemsMinItems NonUniqueItemsMinItemsAlias // want "field NonUniqueItemsMinItemsOnlyFromAliasWithItemsMinItems has multiple definitions of marker kubebuilder:validation:items:MinItems when only a single definition should exist"

	// +kubebuilder:validation:items:MinLength:=5
	UniqueItemsMinLength []string

	// +kubebuilder:validation:items:MinLength:=5
	// +kubebuilder:validation:items:MinLength:=10
	NonUniqueItemsMinLength []string // want "field NonUniqueItemsMinLength has multiple definitions of marker kubebuilder:validation:items:MinLength when only a single definition should exist"

	// +kubebuilder:validation:items:MinLength:=3
	NonUniqueItemsMinLengthFromAliasWithItemsMinLength UniqueItemsMinLengthAlias // want "field NonUniqueItemsMinLengthFromAliasWithItemsMinLength has multiple definitions of marker kubebuilder:validation:items:MinLength when only a single definition should exist"

	NonUniqueItemsMinLengthOnlyFromAliasWithItemsMinLength NonUniqueItemsMinLengthAlias // want "field NonUniqueItemsMinLengthOnlyFromAliasWithItemsMinLength has multiple definitions of marker kubebuilder:validation:items:MinLength when only a single definition should exist"

	// +kubebuilder:validation:items:MinProperties:=1
	UniqueItemsMinProperties []map[string]string

	// +kubebuilder:validation:items:MinProperties:=1
	// +kubebuilder:validation:items:MinProperties:=2
	NonUniqueItemsMinProperties []map[string]string // want "field NonUniqueItemsMinProperties has multiple definitions of marker kubebuilder:validation:items:MinProperties when only a single definition should exist"

	// +kubebuilder:validation:items:MinProperties:=3
	NonUniqueItemsMinPropertiesFromAliasWithItemsMinProperties UniqueItemsMinPropertiesAlias // want "field NonUniqueItemsMinPropertiesFromAliasWithItemsMinProperties has multiple definitions of marker kubebuilder:validation:items:MinProperties when only a single definition should exist"

	NonUniqueItemsMinPropertiesOnlyFromAliasWithItemsMinProperties NonUniqueItemsMinPropertiesAlias // want "field NonUniqueItemsMinPropertiesOnlyFromAliasWithItemsMinProperties has multiple definitions of marker kubebuilder:validation:items:MinProperties when only a single definition should exist"

	// +kubebuilder:validation:items:Minimum:=0
	UniqueItemsMinimum []int

	// +kubebuilder:validation:items:Minimum:=0
	// +kubebuilder:validation:items:Minimum:=-5
	NonUniqueItemsMinimum []int // want "field NonUniqueItemsMinimum has multiple definitions of marker kubebuilder:validation:items:Minimum when only a single definition should exist"

	// +kubebuilder:validation:items:Minimum:=10
	NonUniqueItemsMinimumFromAliasWithItemsMinimum UniqueItemsMinimumAlias // want "field NonUniqueItemsMinimumFromAliasWithItemsMinimum has multiple definitions of marker kubebuilder:validation:items:Minimum when only a single definition should exist"

	NonUniqueItemsMinimumOnlyFromAliasWithItemsMinimum NonUniqueItemsMinimumAlias // want "field NonUniqueItemsMinimumOnlyFromAliasWithItemsMinimum has multiple definitions of marker kubebuilder:validation:items:Minimum when only a single definition should exist"

	// +kubebuilder:validation:items:MultipleOf:=3
	UniqueItemsMultipleOf []int

	// +kubebuilder:validation:items:MultipleOf:=3
	// +kubebuilder:validation:items:MultipleOf:=5
	NonUniqueItemsMultipleOf []int // want "field NonUniqueItemsMultipleOf has multiple definitions of marker kubebuilder:validation:items:MultipleOf when only a single definition should exist"

	// +kubebuilder:validation:items:MultipleOf:=2
	NonUniqueItemsMultipleOfFromAliasWithItemsMultipleOf UniqueItemsMultipleOfAlias // want "field NonUniqueItemsMultipleOfFromAliasWithItemsMultipleOf has multiple definitions of marker kubebuilder:validation:items:MultipleOf when only a single definition should exist"

	NonUniqueItemsMultipleOfOnlyFromAliasWithItemsMultipleOf NonUniqueItemsMultipleOfAlias // want "field NonUniqueItemsMultipleOfOnlyFromAliasWithItemsMultipleOf has multiple definitions of marker kubebuilder:validation:items:MultipleOf when only a single definition should exist"

	// +kubebuilder:validation:items:Pattern:="^[a-z]+$"
	UniqueItemsPattern []string

	// +kubebuilder:validation:items:Pattern:="^[a-z]+$"
	// +kubebuilder:validation:items:Pattern:="^[0-9]+$"
	NonUniqueItemsPattern []string // want "field NonUniqueItemsPattern has multiple definitions of marker kubebuilder:validation:items:Pattern when only a single definition should exist"

	// +kubebuilder:validation:items:Pattern:="^[A-Z]+$"
	NonUniqueItemsPatternFromAliasWithItemsPattern UniqueItemsPatternAlias // want "field NonUniqueItemsPatternFromAliasWithItemsPattern has multiple definitions of marker kubebuilder:validation:items:Pattern when only a single definition should exist"

	NonUniqueItemsPatternOnlyFromAliasWithItemsPattern NonUniqueItemsPatternAlias // want "field NonUniqueItemsPatternOnlyFromAliasWithItemsPattern has multiple definitions of marker kubebuilder:validation:items:Pattern when only a single definition should exist"

	// +kubebuilder:validation:items:Type:="string"
	UniqueItemsType []string

	// +kubebuilder:validation:items:Type:="string"
	// +kubebuilder:validation:items:Type:="date"
	NonUniqueItemsType []string // want "field NonUniqueItemsType has multiple definitions of marker kubebuilder:validation:items:Type when only a single definition should exist"

	// +kubebuilder:validation:items:Type:="string"
	NonUniqueItemsTypeFromAliasWithItemsType UniqueItemsTypeAlias // want "field NonUniqueItemsTypeFromAliasWithItemsType has multiple definitions of marker kubebuilder:validation:items:Type when only a single definition should exist"

	NonUniqueItemsTypeOnlyFromAliasWithItemsType NonUniqueItemsTypeAlias // want "field NonUniqueItemsTypeOnlyFromAliasWithItemsType has multiple definitions of marker kubebuilder:validation:items:Type when only a single definition should exist"

	// +kubebuilder:validation:items:UniqueItems:=true
	UniqueItemsUniqueItems [][]string

	// +kubebuilder:validation:items:UniqueItems:=true
	// +kubebuilder:validation:items:UniqueItems:=false
	NonUniqueItemsUniqueItems [][]string // want "field NonUniqueItemsUniqueItems has multiple definitions of marker kubebuilder:validation:items:UniqueItems when only a single definition should exist"

	// +kubebuilder:validation:items:UniqueItems:=true
	NonUniqueItemsUniqueItemsFromAliasWithItemsUniqueItems UniqueItemsUniqueItemsAlias // want "field NonUniqueItemsUniqueItemsFromAliasWithItemsUniqueItems has multiple definitions of marker kubebuilder:validation:items:UniqueItems when only a single definition should exist"

	NonUniqueItemsUniqueItemsOnlyFromAliasWithItemsUniqueItems NonUniqueItemsUniqueItemsAlias // want "field NonUniqueItemsUniqueItemsOnlyFromAliasWithItemsUniqueItems has multiple definitions of marker kubebuilder:validation:items:UniqueItems when only a single definition should exist"
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

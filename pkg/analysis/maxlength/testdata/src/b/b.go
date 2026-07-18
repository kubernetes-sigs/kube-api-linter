package b

// DVPreferredMaxLength exercises the linter when configured to prefer DV markers.
// Fields that carry only DV markers should NOT lint.
// Fields that are missing markers should lint with a diagnostic citing the k8s:* form.
type DVPreferredMaxLength struct {
	// +k8s:maxLength=256
	StringWithDVMaxLength string // satisfied by DV marker — no lint

	// +kubebuilder:validation:MaxLength:=128
	StringWithKubebuilderMaxLength string // kubebuilder marker also satisfies — no lint

	StringWithoutMaxLength string // want `field DVPreferredMaxLength.StringWithoutMaxLength must have a maximum length, add k8s:maxLength marker`

	// +k8s:maxBytes=512
	ByteSliceWithDVMaxBytes []byte // satisfied by k8s:maxBytes — no lint

	ByteSliceWithoutMax []byte // want `field DVPreferredMaxLength.ByteSliceWithoutMax must have a maximum length, add k8s:maxBytes marker`

	// +k8s:maxItems=128
	ArrayWithDVMaxItems []int // satisfied by k8s:maxItems — no lint

	// +kubebuilder:validation:MaxItems:=64
	ArrayWithKubebuilderMaxItems []int // kubebuilder marker also satisfies — no lint

	ArrayWithoutMaxItems []int // want `field DVPreferredMaxLength.ArrayWithoutMaxItems must have a maximum items, add k8s:maxItems marker`

	// +k8s:maxProperties=64
	MapWithDVMaxProperties map[string]string // satisfied by k8s:maxProperties — no lint

	// +kubebuilder:validation:MaxProperties:=32
	MapWithKubebuilderMaxProperties map[string]string // kubebuilder marker also satisfies — no lint

	MapWithoutMaxProperties map[string]string // want `field DVPreferredMaxLength.MapWithoutMaxProperties must have a maximum properties, add k8s:maxProperties marker`

	// Non-string-keyed map — never linted regardless of config.
	MapWithIntKey map[int]string

	// +k8s:enum
	DVEnumString string // enum exempts from max-length — no lint

	// +kubebuilder:validation:Enum:="A";"B"
	KubebuilderEnumString string // kubebuilder enum also exempts in DV-preferred mode — no lint

	// +k8s:format=date-time
	DVDateTimeString string // format exempts from max-length — no lint

	// +k8s:format=date
	DVDateString string // format exempts from max-length — no lint

	// +k8s:format=duration
	DVDurationString string // format exempts from max-length — no lint

	// +kubebuilder:validation:Format:=date-time
	KubebuilderDateTimeString string // kubebuilder format also exempts in DV-preferred mode — no lint

	// +kubebuilder:validation:Format:=date
	KubebuilderDateString string // kubebuilder format also exempts in DV-preferred mode — no lint

	// +kubebuilder:validation:Format:=duration
	KubebuilderDurationString string // kubebuilder format also exempts in DV-preferred mode — no lint

	// k8s:maxLength=256 on a []byte field is the WRONG DV marker; k8s:maxBytes should be used.
	// The linter must still report the field as missing a valid max-length constraint.
	// +k8s:maxLength=256
	ByteSliceWithWrongDVMarker []byte // want `field DVPreferredMaxLength.ByteSliceWithWrongDVMarker must have a maximum length, add k8s:maxBytes marker`

	// Raw []string — needs both MaxItems and items:MaxLength.
	// +k8s:maxItems=16
	// +kubebuilder:validation:items:MaxLength:=64
	StringArrayWithMaxItemsAndElementLength []string // no lint

	// +k8s:maxItems=16
	StringArrayWithMaxItemsWithoutElementLength []string // want `field DVPreferredMaxLength.StringArrayWithMaxItemsWithoutElementLength array element must have a maximum length, add kubebuilder:validation:items:MaxLength`

	StringArrayWithoutMaxItemsOrElementLength []string // want `field DVPreferredMaxLength.StringArrayWithoutMaxItemsOrElementLength must have a maximum items, add k8s:maxItems marker` `field DVPreferredMaxLength.StringArrayWithoutMaxItemsOrElementLength array element must have a maximum length, add kubebuilder:validation:items:MaxLength`

	// Map with string-alias key — should lint for maxProperties.
	MapWithStringAliasKey map[StringAliasDVMaxLength]string // want `field DVPreferredMaxLength.MapWithStringAliasKey must have a maximum properties, add k8s:maxProperties marker`

	// Map with string-alias key with marker — should NOT lint.
	// +k8s:maxProperties=16
	MapWithStringAliasKeyAndMarker map[StringAliasDVMaxLength]string
}

// StringAliasDVMaxLength is a string alias carrying a DV maxLength marker.
// +k8s:maxLength=512
type StringAliasDVMaxLength string

// StringAliasNoMaxLength is a string alias without any max-length marker.
type StringAliasNoMaxLength string

// DVPreferredAliases exercises alias-level DV markers.
type DVPreferredAliases struct {
	StringWithAliasMaxLength StringAliasDVMaxLength // satisfied by alias marker — no lint

	StringWithoutMaxLength StringAliasNoMaxLength // want `field DVPreferredAliases.StringWithoutMaxLength type StringAliasNoMaxLength must have a maximum length, add k8s:maxLength marker`

	// +k8s:maxItems=64
	AliasArrayWithMaxItems []StringAliasDVMaxLength // array element max-length satisfied by alias — no lint

	AliasArrayWithoutMaxItems []StringAliasDVMaxLength // want `field DVPreferredAliases.AliasArrayWithoutMaxItems must have a maximum items, add k8s:maxItems marker`
}

package c

// Different embedding scenarios
type ResourceWithEmbeddings struct {
	Status StatusWithEmbeddings `json:"status"`
}

type StatusWithEmbeddings struct {
	// Regular inlined embed
	InlineEmbed `json:",inline"`

	// Non-inlined embed
	NonInlineEmbed `json:"nonInlineEmbed"` // want "status field \"NonInlineEmbed\" must be marked as optional"

	// Non-inlined embed with omitempty
	NonInlineOmitEmptyEmbed `json:"nonInlineOmitEmpty,omitempty"` // want "status field \"NonInlineOmitEmptyEmbed\" must be marked as optional"

	// Pointer to non-inlined embed
	*PointerEmbed `json:"pointerEmbed"` // want "status field \"PointerEmbed\" must be marked as optional"

	// Pointer to non-inlined embed with omitempty
	*PointerOmitEmptyEmbed `json:"pointerOmitEmpty,omitempty"` // want "status field \"PointerOmitEmptyEmbed\" must be marked as optional"
}

type InlineEmbed struct {
	InlineField string `json:"inlineField"` // want "status field \"InlineField\" must be marked as optional"
}

type NonInlineEmbed struct {
	NonInlineField string `json:"nonInlineField"`
}

type NonInlineOmitEmptyEmbed struct {
	NonInlineOmitEmptyField string `json:"nonInlineOmitEmptyField"`
}

type PointerEmbed struct {
	PointerField string `json:"pointerField"`
}

type PointerOmitEmptyEmbed struct {
	PointerOmitEmptyField string `json:"pointerOmitEmptyField"`
}

type NonInlineOmitZeroEmbed struct {
	NonInlineOmitZeroField string `json:"nonInlineOmitZeroField"`
}

type PointerOmitZeroEmbed struct {
	PointerOmitZeroField string `json:"pointerOmitZeroField"`
}

type ResourceWithNestedStatus struct {
	Status NestedStatusStatus `json:"status"`
}

type NestedStatusStatus struct {
	// +k8s:optional
	NestedStatus SecondLevelStatus `json:"nestedStatus"`
}

type SecondLevelStatus struct {
	// The required here is ignored because it is not the top-level status field.
	// +required
	NestedField string `json:"nestedField"`
}

type ResourceWithStatusMarkedRequired struct {
	Status StatusMarkedRequired `json:"status"`
}

type StatusMarkedRequired struct {
	// +required
	OneRequiredField string `json:"oneRequiredField"` // want "status field \"OneRequiredField\" must be marked as optional"

	// +required
	// +kubebuilder:validation:Required
	BothRequiredField string `json:"bothRequiredField"` // want "status field \"BothRequiredField\" must be marked as optional"
}

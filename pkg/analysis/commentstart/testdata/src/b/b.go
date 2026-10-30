package b

type CustomExcludeTestStruct struct {
	// Deprecated: This field is no longer used.
	OldField string `json:"oldField"`

	// TODO: fill this in later.
	IncompleteField string `json:"incompleteField"`

	// NOTE: this is a special case.
	SpecialField string `json:"specialField"`

	// WRONGSTART is a case mismatch that can be auto-fixed. // want "godoc for field CustomExcludeTestStruct.WrongStart should start with 'wrongStart ...'"
	WrongStart string `json:"wrongStart"`

	// this comment is completely wrong. // want "godoc for field CustomExcludeTestStruct.CompletelyWrong should start with 'completelyWrong ...'"
	CompletelyWrong string `json:"completelyWrong"`
}

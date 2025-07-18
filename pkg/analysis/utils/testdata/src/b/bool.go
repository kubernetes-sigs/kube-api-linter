package b

type ZeroValueTestBools struct {
	Bool bool // want "zero value is valid" "validation is complete"

	BoolPtr *bool // want "zero value is valid" "validation is complete"
}

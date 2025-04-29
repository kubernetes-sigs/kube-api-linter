package a

// IntB is being used to show a type from a different file.
type IntB int // want "type IntB should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

type InvalidSliceIntAliasB []int // want "type InvalidSliceIntAliasB array element should not use an int, int8 or int16. Use int32 or int64 depending on bounding requirements"

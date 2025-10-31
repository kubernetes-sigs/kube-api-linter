# Linters

- [ArrayOfStruct](#arrayofstruct) - Ensures arrays of structs have at least one required field
- [Conditions](#conditions) - Checks that `Conditions` fields are correctly formatted
- [CommentStart](#commentstart) - Ensures comments start with the serialized form of the type
- [ConflictingMarkers](#conflictingmarkers) - Detects mutually exclusive markers on the same field
- [DefaultOrRequired](#defaultorrequired) - Ensures fields marked as required do not have default values
- [DuplicateMarkers](#duplicatemarkers) - Checks for exact duplicates of markers
- [DependentTags](#dependenttags) - Enforces dependencies between markers
- [ForbiddenMarkers](#forbiddenmarkers) - Checks that no forbidden markers are present on types/fields.
- [Integers](#integers) - Validates usage of supported integer types
- [JSONTags](#jsontags) - Ensures proper JSON tag formatting
- [MaxLength](#maxlength) - Checks for maximum length constraints on strings and arrays
- [NamingConventions](#namingconventions) - Ensures field names adhere to user-defined naming conventions
- [NumericBounds](#numericbounds) - Validates numeric fields have appropriate bounds validation markers
- [NoBools](#nobools) - Prevents usage of boolean types
- [NoDurations](#nodurations) - Prevents usage of duration types
- [NoFloats](#nofloats) - Prevents usage of floating-point types
- [Nomaps](#nomaps) - Restricts usage of map types
- [NonPointerStructs](#nonpointerstructs) - Ensures non-pointer structs are marked correctly with required/optional markers
- [NoNullable](#nonullable) - Prevents usage of the nullable marker
- [Nophase](#nophase) - Prevents usage of 'Phase' fields
- [Notimestamp](#notimestamp) - Prevents usage of 'TimeStamp' fields
- [OptionalFields](#optionalfields) - Validates optional field conventions
- [OptionalOrRequired](#optionalorrequired) - Ensures fields are explicitly marked as optional or required
- [NoReferences](#noreferences) - Ensures field names use Ref/Refs instead of Reference/References
- [PreferredMarkers](#preferredmarkers) - Ensures preferred markers are used instead of equivalent markers
- [RequiredFields](#requiredfields) - Validates required field conventions
- [SSATags](#ssatags) - Ensures proper Server-Side Apply (SSA) tags on array fields
- [StatusOptional](#statusoptional) - Ensures status fields are marked as optional
- [StatusSubresource](#statussubresource) - Validates status subresource configuration
- [UniqueMarkers](#uniquemarkers) - Ensures unique marker definitions

## ArrayOfStruct

The `arrayofstruct` linter checks that arrays containing structs have at least one required field to prevent ambiguous YAML representations.

When an array contains structs where all fields are optional or unmarked, it becomes difficult to distinguish between different array elements in YAML configurations. This can lead to ambiguous or confusing API definitions.

The linter enforces that at least one field in the struct must be marked with one of the following required markers:
- `// +required`
- `// +kubebuilder:validation:Required`
- `// +k8s:required`

This applies to:
- Direct array fields containing struct types
- Arrays of pointers to struct types
- Arrays of inline struct definitions
- Arrays using type aliases that resolve to struct types

The linter does not check:
- Arrays of primitive types (strings, integers, etc.)
- Arrays of types from external packages (cannot inspect their fields)

## Conditions

The `conditions` linter checks that `Conditions` fields in the API types are correctly formatted.
The `Conditions` field should be a slice of `metav1.Condition` with the following tags and markers:

```go
// +listType=map
// +listMapKey=type
// +patchStrategy=merge
// +patchMergeKey=type
// +optional
Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,opt,name=conditions"`
```

Conditions are idiomatically the first field within the status struct, and the linter will highlight when the Conditions are not the first field.

Protobuf tags and patch strategy are required for in-tree API types, but not for CRDs.
When linting CRD based types, set the `useProtobuf` and `usePatchStrategy` config option to `Ignore` or `Forbid`.

### Configuration

```yaml
lintersConfig:
  conditions:
    isFirstField: Warn | Ignore # The policy for the Conditions field being the first field. Defaults to `Warn`.
    useProtobuf: SuggestFix | Warn | Ignore | Forbid # The policy for the protobuf tag on the Conditions field. Defaults to `SuggestFix`.
    usePatchStrategy: SuggestFix | Warn | Ignore | Forbid # The policy for the patchStrategy tag on the Conditions field. Defaults to `SuggestFix`.
```

### Fixes

The `conditions` linter can automatically fix the tags on the `Conditions` field.
When they do not match the expected format, the linter will suggest to update the tags to match the expected format.

For CRDs, protobuf tags and patch strategy are not expected.
By setting the `useProtobuf`/`usePatchStrategy` configuration to `Ignore`, the linter will not suggest to add the protobuf/patch strategy tag to the `Conditions` field tags.
By setting the `useProtobuf`/`usePatchStrategy` configuration to `Forbid`, the linter will suggest to remove the protobuf/patch strategy tag from the `Conditions` field tags.

The linter will also suggest to add missing markers.
If any of the 5 markers in the example above are missing, the linter will suggest to add them directly above the field.

When `usePatchStrategy` is set to `Ignore`, the linter will not suggest to add the `patchStrategy` and `patchMergeKey` tags to the `Conditions` field markers.
When `usePatchStrategy` is set to `Forbid`, the linter will suggest to remove the `patchStrategy` and `patchMergeKey` tags from the `Conditions` field markers.

## DependentTags

The `dependenttags` linter enforces dependencies between markers. This prevents API inconsistencies where one marker requires the presence of another.

The linter is configured with a main tag and a list of required dependent tags. If the main tag is present on a field, the linter checks for the presence of the dependent tags based on the `type` field:
- `All`: Ensures that **all** of the dependent tags are present.
- `Any`: Ensures that **at least one** of the dependent tags is present.

### Configuration

```yaml
lintersConfig:
  dependenttags:
    rules:
      - identifier: "k8s:unionMember"
        type: "All"
        dependents:
          - "k8s:optional"
      - identifier: "listType"
        type: "All"
        dependents:
          - "k8s:listType"
      - identifier: "example:any"
        type: "Any"
        dependents:
          - "dep1"
          - "dep2"
```

### Behavior

This linter only checks for the presence or absence of markers; it does not inspect or enforce specific values within those markers. Therefore:

- **Values:** The linter does not care about the values of the `identifier` or `dependent` markers. It only verifies if the markers themselves are present.
- **Fixes:** This linter does not provide automatic fixes. It only reports violations.
- **Same/Different Values:** Whether you want the same or different values between dependent markers is outside the scope of this linter. You would need other validation mechanisms (e.g., CEL validation) to enforce value-based dependencies.

## CommentStart

The `commentstart` linter checks that all comments in the API types start with the serialized form of the type they are commenting on.
This helps to ensure that generated documentation reflects the most common usage of the field, the serialized YAML form.

### Fixes

The `commentstart` linter can automatically fix comments that do not start with the serialized form of the type.

When the `json` tag is present, and matches the first word of the field comment in all but casing, the linter will suggest that the comment be updated to match the `json` tag.

## ConflictingMarkers

The `conflictingmarkers` linter detects and reports when mutually exclusive markers are used on the same field.
This prevents common configuration errors and unexpected behavior in Kubernetes API types.

The linter reports issues when markers from two or more sets of a conflict definition are present on the same field.
It does NOT report issues when multiple markers from the same set are present - only when markers from
different sets within the same conflict definition are found together.

The linter is configurable and allows users to define sets of conflicting markers.
Each conflict set must specify:
- A unique name for the conflict
- Multiple sets of markers that are mutually exclusive with each other (at least 2 sets)
- A description explaining why the markers conflict

### Configuration

```yaml
lintersConfig:
  conflictingmarkers:
    conflicts:
      - name: "default_vs_required"
        sets:
          - ["default", "kubebuilder:default"]
          - ["required", "kubebuilder:validation:Required", "k8s:required"]
        description: "A field with a default value cannot be required"
      - name: "three_way_conflict"
        sets:
          - ["custom:marker1", "custom:marker2"]
          - ["custom:marker3", "custom:marker4"]
          - ["custom:marker5", "custom:marker6"]
        description: "Three-way conflict between marker sets"
```

**Note**: This linter is not enabled by default and must be explicitly enabled in the configuration.

The linter does not provide automatic fixes as it cannot determine which conflicting marker should be removed.

## DefaultOrRequired

The `defaultorrequired` linter checks that fields marked as required do not have default values applied.

A field cannot be both required and have a default value, as these are conflicting concepts:
- A **required** field must be provided by the user and cannot be omitted
- A **default** value is used when a field is not provided

This linter helps prevent common configuration errors where a field is marked as required but also has a default value, which creates ambiguity about whether the field truly needs to be provided by the user.

### Example

The following will be flagged by the linter:

```go
type MyStruct struct {
	// +required
	// +default:=value
	ConflictedField string `json:"conflictedField"` // Error: field cannot have both a default value and be marked as required
}
```

The linter also detects conflicts with:
- `+required` 
- `+default`
- `+k8s:required`
- `+kubebuilder:validation:Required`
- `+kubebuilder:default`

This linter is enabled by default and helps ensure that API designs are consistent and unambiguous about whether fields are truly required or have default values.

## DuplicateMarkers

The duplicatemarkers linter checks for exact duplicates of markers for types and fields.
This means that something like:

```go
// +kubebuilder:validation:MaxLength=10
// +kubebuilder:validation:MaxLength=10
```

Will be flagged by this linter, while something like:

```go
// +kubebuilder:validation:MaxLength=10
// +kubebuilder:validation:MaxLength=11
```

will not.

### Fixes

The `duplicatemarkers` linter can automatically fix all markers that are exact match to another markers.
If there are duplicates across fields and their underlying type, the marker on the type will be preferred and the marker on the field will be removed.

## ForbiddenMarkers

The `forbiddenmarkers` linter ensures that types and fields do not contain any markers that are forbidden.

By default, `forbiddenmarkers` is not enabled.

### Configuation

It can be configured with a list of marker identifiers and optionally their attributes and values that are forbidden.

Some examples configurations explained:

**Scenario:** forbid all instances of the marker with the identifier `forbidden:marker`

```yaml
linterConfig:
  forbiddenmarkers:
    markers:
      - identifier: "forbidden:marker"
```

**Scenario:** forbid all instances of the marker with the identifier `forbidden:marker` containing the attribute `fruit`

```yaml
linterConfig:
  forbiddenmarkers:
    markers:
      - identifier: "forbidden:marker"
        ruleSets:
          - attributes:
              - name: "fruit"
```

**Scenario:** forbid all instances of the marker with the identifier `forbidden:marker` containing the `fruit` AND `color` attributes

```yaml
linterConfig:
  forbiddenmarkers:
    markers:
      - identifier: "forbidden:marker"
        ruleSets:
          - attributes:
              - name: "fruit"
              - name: "color"
```

**Scenario:** forbid all instances of the marker with the identifier `forbidden:marker` where the `fruit` attribute is set to one of `apple`, `banana`, or `orange`

```yaml
linterConfig:
  forbiddenmarkers:
    markers:
      - identifier: "forbidden:marker"
        ruleSets:
          - attributes:
              - name: "fruit"
                values:
                  - "apple"
                  - "banana"
                  - "orange"
```

**Scenario:** forbid all instances of the marker with the identifier `forbidden:marker` where the `fruit` attribute is set to one of `apple`, `banana`, or `orange` AND the `color` attribute is set to one of `red`, `green`, or `blue`

```yaml
linterConfig:
  forbiddenmarkers:
    markers:
      - identifier: "forbidden:marker"
        ruleSets:
          - attributes:
              - name: "fruit"
                values:
                  - "apple"
                  - "banana"
                  - "orange"
              - name: "color"
                values:
                  - "red"
                  - "blue"
                  - "green"
```

**Scenario:** forbid all instances of the marker with the identifier `forbidden:marker` where:

- The `fruit` attribute is set to `apple` and the `color` attribute is set to one of `blue` or `orange` (allow any other color apple)

_OR_

- The `fruit` attribute is set to `orange` and the `color` attribute is set to one of `blue`, `red`, or `green` (allow any other color orange)

_OR_

- The `fruit` attribute is set to `banana` (no bananas allowed)

```yaml
linterConfig:
  forbiddenmarkers:
    markers:
      - identifier: "forbidden:marker"
        ruleSets:
          - attributes:
              - name: "fruit"
                values:
                  - "apple"
              - name: "color"
                values:
                  - "blue"
                  - "orange"
          - attributes:
              - name: "fruit"
                values:
                  - "orange"
              - name: "color"
                values:
                  - "blue"
                  - "red"
                  - "green"
          - attributes:
              - name: "fruit"
                values:
                  - "banana"
```

### Fixes

Fixes are suggested to remove all markers that are forbidden.

## Integers

The `integers` linter checks for usage of unsupported integer types.
Only `int32` and `int64` types should be used in APIs, and other integer types, including unsigned integers are forbidden.

## JSONTags

The `jsontags` linter checks that all fields in the API types have a `json` tag, and that those tags are correctly formatted.
The `json` tag for a field within a Kubernetes API type should use a camel case version of the field name.

The `jsontags` linter checks the tag name against the regex `"^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)*$"` which allows consecutive upper case characters, to allow for acronyms, e.g. `requestTTL`.

### Configuration

```yaml
lintersConfig:
  jsontags:
    jsonTagRegex: "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)*$" # Provide a custom regex, which the json tag must match.
```

## MaxLength

The `maxlength` linter checks that string and array fields in the API are bounded by a maximum length.

For strings, this means they have a `+kubebuilder:validation:MaxLength` marker.

For arrays, this means they have a `+kubebuilder:validation:MaxItems` marker.

For arrays of strings, the array element should also have a `+kubebuilder:validation:MaxLength` marker if the array element is a type alias,
or `+kubebuilder:validation:items:MaxLenth` if the array is an element of the built-in string type.

Adding maximum lengths to strings and arrays not only ensures that the API is not abused (used to store overly large data, reduces DDOS etc.),
but also allows CEL validation cost estimations to be kept within reasonable bounds.

## NamingConventions

The `namingconventions` linter ensures that field names adhere to a set of defined naming conventions.

By default, `namingconventions` is not enabled.

When enabled, it must be configured with at least one naming convention.

### Configuration

Naming conventions must have:
- A unique human-readable name.
- A human-readable message to be included in violation errors.
- A regular expression that will match text within the field name that violates the convention.
- A defined "operation". Allowed operations are `Inform`, `Drop`, `DropField`, and `Replacement`.

The `Inform` operation will simply inform when a field name violates the naming convention.
The `Drop` operation will suggest a fix that drops violating text from the field name.
The `DropField` operation will suggest a fix that removes the field in it's entirety.
The `Replacement` operation will suggest a fix that replaces the violating text in the field name with a defined replacement value.

High-level configuration overview:
```yaml
linterConfig:
  namingconventions:
    conventions: 
      - name: {human readable string} # must be unique
        violationMatcher: {regular expression}
        operation: Inform | Drop | DropField | Replacement
        replacement: { replacement string } # required when operation is 'Replacement', forbidden otherwise
        message: {human readable string}
```

Some example configurations:

**Scenario:** Inform that any variations of the word 'fruit' in field names is not allowed
```yaml
linterConfig:
  namingconventions:
    conventions: 
      - name: nofruit
        violationMatcher: (?i)fruit
        operation: Inform
        message: fields should not contain any variation of the word 'fruit' in their names
```

**Scenario:** Drop any variations of the word 'fruit' in field names
```yaml
linterConfig:
  namingconventions:
    conventions: 
      - name: nofruit
        violationMatcher: (?i)fruit
        operation: Drop
        message: fields should not contain any variation of the word 'fruit' in their names
```

**Scenario:** Do not allow fields with any variations of the word 'fruit' in their name
```yaml
linterConfig:
  namingconventions:
    conventions: 
      - name: nofruit
        violationMatcher: (?i)fruit
        operation: DropField
        message: fields should not contain any variation of the word 'fruit' in their names
```

**Scenario:** Replace any variations of the word 'color' with 'colour' in field names
```yaml
linterConfig:
  namingconventions:
    conventions:
      - name: BritishEnglishColour 
        violationMatcher: (?i)color
        operation: Replacement
        replacement: colour
        message: prefer 'colour' over 'color' when referring to colours in field names
```

## NumericBounds

The `numericbounds` linter checks that numeric fields (`int32` and `int64`) have appropriate bounds validation markers.

According to Kubernetes API conventions, numeric fields should have bounds checking to prevent values that are too small, negative (when not intended), or too large.

This linter ensures that:
- `int32` and `int64` fields have both `+kubebuilder:validation:Minimum` and `+kubebuilder:validation:Maximum` markers
- `int64` fields with bounds outside the JavaScript safe integer range are flagged

### JavaScript Safe Integer Range

For `int64` fields, the linter checks if the bounds exceed the JavaScript safe integer range of `-(2^53)` to `(2^53)` (specifically, `-9007199254740991` to `9007199254740991`).

JavaScript represents all numbers as IEEE 754 double-precision floating-point values, which can only safely represent integers in this range. Values outside this range may lose precision when processed by JavaScript clients.

When an `int64` field has bounds that exceed this range, the linter will suggest using a string type instead to avoid precision loss.

### Examples

**Valid:** Numeric field with proper bounds markers
```go
type Example struct {
    // +kubebuilder:validation:Minimum=0
    // +kubebuilder:validation:Maximum=100
    Count int32
}
```

**Valid:** Int64 field with JavaScript-safe bounds
```go
type Example struct {
    // +kubebuilder:validation:Minimum=-9007199254740991
    // +kubebuilder:validation:Maximum=9007199254740991
    Timestamp int64
}
```

**Invalid:** Missing bounds markers
```go
type Example struct {
    Count int32 // want: should have minimum and maximum bounds validation markers
}
```

**Invalid:** Only one bound specified
```go
type Example struct {
    // +kubebuilder:validation:Minimum=0
    Count int32 // want: has minimum but is missing maximum bounds validation marker
}
```

**Invalid:** Int64 with bounds exceeding JavaScript safe range
```go
type Example struct {
    // +kubebuilder:validation:Minimum=-10000000000000000
    // +kubebuilder:validation:Maximum=10000000000000000
    LargeNumber int64 // want: bounds exceed JavaScript safe integer range
}
```

## NoBools

The `nobools` linter checks that fields in the API types do not contain a `bool` type.

Booleans are limited and do not evolve well over time.
It is recommended instead to create a string alias with meaningful values, as an enum.

## NoDurations

The `nodurations` linter checks that fields in the API types do not contain a `Duration` type ether from the `time` package or the `k8s.io/apimachinery/pkg/apis/meta/v1` package.

It is recommended to avoid the use of Duration types. Their use ties the API to Go's notion of duration parsing, which may be hard to implement in other languages.

Instead, use an integer based field with a unit in the name, e.g. `FooSeconds`.

## NoFloats

The `nofloats` linter checks that fields in the API types do not contain a `float32` or `float64` type.

Floating-point values cannot be reliably round-tripped without changing and have varying precision and representation across languages and architectures.
Their use should be avoided as much as possible.
They should never be used in spec.

## Nomaps

The `nomaps` linter checks the usage of map types.

Maps are discouraged apart from `map[string]string` which is used for labels and annotations in Kubernetes APIs since it's hard to distinguish between structs and maps in spec. Instead of plain map, lists of named subobjects are preferred.

### Configuration

```yaml
lintersConfig:
  nomaps:
    policy: Enforce | AllowStringToStringMaps | Ignore # Determines how the linter should handle maps of simple types. Defaults to AllowStringToStringMaps.
```

## NonPointerStructs

The `nonpointerstructs` linter checks that non-pointer structs that contain required fields are marked as required.
Non-pointer structs that contain no required fields are marked as optional.

This linter is important for types validated in Go as there is no way to validate the optionality of the fields at runtime,
aside from checking the fields within them.

This linter is NOT intended to be used to check for CRD types.
The advice of this linter may be applied to CRD types, but it is not necessary for CRD types due to optionality being validated by openapi and no native Go code.
For CRD types, the optionalfields and requiredfields linters should be used instead.

If a struct is marked required, this can only be validated by having a required field within it.
If there are no required fields, the struct is implicitly optional and must be marked as so.

To have an optional struct field that includes required fields, the struct must be a pointer.
To have a required struct field that includes no required fields, the struct must be a pointer.

### Configuration

```yaml
lintersConfig:
  nonpointerstructs:
    preferredRequiredMarker: required | kubebuilder:validation:Required | k8s:required # The preferred required marker to use for required fields when providing fixes. Defaults to `required`.
    preferredOptionalMarker: optional | kubebuilder:validation:Optional | k8s:optional # The preferred optional marker to use for optional fields when providing fixes. Defaults to `optional`.
```

### Fixes

The `nonpointerstructs` linter can automatically fix non-pointer struct fields that are not marked as required or optional.

It will suggest to mark the field as required or optional, depending on the fields within the non-pointer struct.

## NoNullable

The `nonullable` linter ensures that types and fields do not have the `nullable` marker.

### Fixes

Fixes are suggested to remove the `nullable` marker.

## Notimestamp

The `notimestamp` linter checks that the fields in the API are not named with the word 'Timestamp'.

The name of a field that specifies the time at which something occurs should be called `somethingTime`. It is recommended not use 'stamp' (e.g., creationTimestamp).

### Fixes

The `notimestamp` linter will automatically fix fields and json tags that are named with the word 'Timestamp'.

It will automatically replace 'Timestamp' with 'Time' and update both the field and tag name.
Example: 'FooTimestamp' will be updated to 'FooTime'.

## Nophase

The `nophase` linter checks that the fields in the API types don't contain a 'Phase', or any field which contains 'Phase' as a substring, e.g MachinePhase.

## OptionalFields

The `optionalfields` linter checks that all fields marked as optional adhere to being pointers and having either the `omitempty` or `omitzero` value in their `json` tag where appropriate.
Currently `omitzero` is handled only for fields with struct type.

If you prefer to avoid pointers where possible, the linter can be configured with the `WhenRequired` preference to determine, based on the serialization and valid values for the field, whether the field should be a pointer or not.
For example
- an optional string with a non-zero minimum length does not need to be a pointer, as the zero value is not valid, and it is safe for the Go marshaller to omit the empty value.
- an optional struct having omitzero json tag with a non-zero minimum properties does not need to be a pointer, as the zero value is not valid, and it is safe for the Go marshaller to omit the empty value.

In certain use cases, it can be desirable to not omit optional fields from the serialized form of the object.
In this case, the `omitempty` policy can be set to `Ignore`, and the linter will ensure that the zero value of the object is an acceptable value for the field.

### Configuration

```yaml
lintersConfig:
  optionalfields:
    pointers:
      preference: Always | WhenRequired # Whether to always require pointers, or only when required. Defaults to `Always`.
      policy: SuggestFix | Warn # The policy for pointers in optional fields. Defaults to `SuggestFix`.
    omitempty:
        policy: SuggestFix | Warn | Ignore # The policy for omitempty in optional fields. Defaults to `SuggestFix`.
    omitzero:
        policy: SuggestFix | Warn | Forbid # The policy for omitzero in optional fields. Defaults to `SuggestFix`.
```

### Fixes

The `optionalfields` linter can automatically fix fields that are marked as optional, that are either not pointers or do not have the `omitempty` or `omitzero` value in their `json` tag.
It will suggest to add the pointer to the field, and update the `json` tag to include the `omitempty` value or, for struct fields specifically, it will suggest to remove the pointer to the field, and update the `json` tag to include the `omitzero` value.

If you prefer not to suggest fixes for pointers in optional fields, you can change the `pointers.policy` to `Warn`.

If you prefer not to suggest fixes for `omitempty` in optional fields, you can change the `omitempty.policy` to `Warn` or `Ignore`.
If you prefer not to suggest fixes for `omitzero` in optional fields, you can change the `omitzero.policy` to `Warn` and also not to consider `omitzero` policy at all, it can be set to `Forbid`.

When the `pointers.preference` is set to `WhenRequired`, the linter will suggest to add the pointer to the field only when the field zero value is a valid value for the field.
When the field zero value is not a valid value for the field, the linter will suggest to remove the pointer from the field.
When the field zero value is not a valid value for the field of type struct, the linter will suggest to add `omitzero` json tag and to remove the pointer from the field.

When the `pointers.preference` is set to `Always`, the linter will always suggest to add the pointer to the field, regardless of the validity of the zero value of the field.

## OptionalOrRequired

The `optionalorrequired` linter checks that all fields in the API types are either optional or required, and are marked explicitly as such.

The linter expects to find a comment marker `// +optional` or `// +required` within the comment for the field.

It also supports the `// +kubebuilder:validation:Optional` and `// +kubebuilder:validation:Required` markers, but will suggest to use the `// +optional` and `// +required` markers instead.

If you prefer to use the Kubebuilder markers instead, you can change the preference in the configuration.

The `optionalorrequired` linter also checks for the presence of optional or required markers on type declarations, and forbids this pattern.

### Configuration

```yaml
lintersConfig:
  optionalorrequired:
    preferredOptionalMarker: optional | kubebuilder:validation:Optional # The preferred optional marker to use, fixes will suggest to use this marker. Defaults to `optional`.
    preferredRequiredMarker: required | kubebuilder:validation:Required # The preferred required marker to use, fixes will suggest to use this marker. Defaults to `required`.
```

### Fixes

The `optionalorrequired` linter can automatically fix fields that are using the incorrect form of either the optional or required marker.

It will also remove the secondary marker where both the preferred and secondary marker are present on a field.

## PreferredMarkers

The `preferredmarkers` linter ensures that types and fields use preferred markers instead of equivalent but different marker identifiers.

By default, `preferredmarkers` is not enabled.

This linter is useful for projects that want to enforce consistent marker usage across their codebase, especially when multiple equivalent markers exist. For example, Kubernetes has multiple ways to mark fields as optional:
- `+optional`
- `+k8s:optional`
- `+kubebuilder:validation:Optional`

The linter can be configured to enforce using one preferred marker identifier and report any equivalent markers that should be replaced.

### Configuration

The linter requires a configuration that specifies preferred markers and their equivalent identifiers.

**Scenario:** Enforce using `+k8s:optional` instead of `+kubebuilder:validation:Optional`

```yaml
linterConfig:
  preferredmarkers:
    markers:
      - preferredIdentifier: "k8s:optional"
        equivalentIdentifiers:
          - identifier: "kubebuilder:validation:Optional"
```

**Scenario:** Enforce using a custom marker instead of multiple equivalent markers

```yaml
linterConfig:
  preferredmarkers:
    markers:
      - preferredIdentifier: "custom:preferred"
        equivalentIdentifiers:
          - identifier: "custom:old:marker"
          - identifier: "custom:deprecated:marker"
          - identifier: "custom:legacy:marker"
```

**Scenario:** Multiple preferred markers with different equivalents

```yaml
linterConfig:
  preferredmarkers:
    markers:
      - preferredIdentifier: "k8s:optional"
        equivalentIdentifiers:
          - identifier: "kubebuilder:validation:Optional"
      - preferredIdentifier: "k8s:required"
        equivalentIdentifiers:
          - identifier: "kubebuilder:validation:Required"
```

The linter checks both type-level and field-level markers, including markers inherited from type aliases.

### Fixes

When one or more equivalent markers are found, the linter will:

1. Report a diagnostic message indicating which marker(s) should be preferred
2. Suggest a fix that:
   - If the preferred marker does not already exist: replaces the first equivalent marker with the preferred identifier and preserves any marker expressions (e.g., `=value` or `:key=value`)
   - If the preferred marker already exists: removes all equivalent markers to avoid duplicates
   - Removes any additional equivalent markers

**Example 1:** If both `+kubebuilder:validation:Optional` and `+custom:optional` are configured as equivalents to `+k8s:optional`, the following code:

```go
// +kubebuilder:validation:Optional
// +custom:optional
type MyType string
```

will be automatically fixed to:

```go
// +k8s:optional
type MyType string
```

**Example 2:** If the preferred marker already exists alongside equivalent markers:

```go
// +k8s:optional
// +kubebuilder:validation:Optional
type MyType string
```

will be automatically fixed to:

```go
// +k8s:optional
type MyType string
```

Marker expressions are preserved during replacement. For example, `+kubebuilder:validation:Optional:=someValue` becomes `+k8s:optional=someValue`. Note that unnamed expressions (`:=value`) are normalized to use `=value` syntax for universal compatibility across different marker systems.

## RequiredFields

The `requiredfields` linter checks that all fields marked as required adhere to having `omitempty` or `omitzero` values in their `json` tags.
Currently `omitzero` is handled only for fields with struct type.

Required fields should have omitempty tags to prevent "mess" in the encoded object. 
Fields are not typically pointers. 
A field doesn't need to be a pointer if its zero value is not a valid value, as this zero value could never be accepted. 
However, if the zero value is valid, the field should be a pointer to differentiate between an unset state and a valid zero value.

In certain use cases, it can be desirable to not omit required fields from the serialized form of the object.
In this case, the `omitempty` policy can be set to `Ignore`, and the linter will ensure that the zero value of the object is an acceptable value for the field.

### Configuration

```yaml
lintersConfig:
  requiredfields:
    pointers:
      policy: SuggestFix | Warn # The policy for pointers in required fields. Defaults to `SuggestFix`.
    omitempty:
      policy: SuggestFix | Warn | Ignore # The policy for omitempty in required fields. Defaults to `SuggestFix`.
    omitzero:
      policy: SuggestFix | Warn | Forbid # The policy for omitzero in required fields. Defaults to `SuggestFix`.
```

### Fixes

The `requiredfields` linter can automatically fix fields marked as required. It does this by checking if the field should be a pointer and if its `json` tag is set correctly with `omitempty` or `omitzero`.

If a field's zero value is valid, the linter will suggest to fix it to be a pointer type and add `omitempty` to its JSON tag.

If a field's zero value is not valid, the field doesn't need to be a pointer. However, to prevent unnecessary data from being encoded in the JSON, the linter will suggest to add `omitempty` to its JSON tag.
For a struct field with an invalid zero value, the linter will suggest to add `omitzero` to the JSON tag.

If you prefer not to suggest fixes for pointers in required fields, you can change the `pointers.policy` to `Warn`.

If you prefer not to suggest fixes for `omitempty` in required fields, you can change the `omitempty.policy` to `Warn` or `Ignore`.
If you prefer not to suggest fixes for `omitzero` in required fields, you can change the `omitzero.policy` to `Warn` and also not to consider `omitzero` policy at all, it can be set to `Forbid`.

## NoReferences

The `noreferences` linter ensures that field names use 'Ref'/'Refs' instead of 'Reference'/'References'.

By default, `noreferences` is enabled and operates in standard mode, allowing 'Ref'/'Refs' but prohibiting 'Reference'/'References' in field names.

### Configuration

```yaml
lintersConfig:
  noreferences:
    policy: PreferAbbreviatedReference | NoReferences # Defaults to `PreferAbbreviatedReference`.
```

**Default behavior (policy: PreferAbbreviatedReference):**
- Reports errors for fields containing 'Reference' or 'References' and replaces with 'Ref' or 'Refs'
- **Allows** fields containing 'Ref' or 'Refs' without reporting errors

**Strict mode (policy: NoReferences):**
- **Warns** about any reference-related words ('Ref', 'Refs', 'Reference', or 'References') in field names
- Does not provide automatic fixes - serves as an informational warning
- In this strict mode, the goal is to inform developers about reference-related words in field names

### Fixes

The `noreferences` linter can automatically fix field names in **PreferAbbreviatedReference mode**:

**PreferAbbreviatedReference mode:**
- Replaces 'Reference' with 'Ref' and 'References' with 'Refs' anywhere in field names
- Case insensitive matching
- Examples:
  - `NodeReference` → `NodeRef`
  - `ReferenceNode` → `RefNode`
  - `NodeReferences` → `NodeRefs`

**Note:** 
- The `NoReferences` mode only reports warnings without providing fixes, allowing developers to choose appropriate field names manually.

## SSATags

The `ssatags` linter ensures that array fields in Kubernetes API objects have the appropriate
listType markers (atomic, set, or map) for proper Server-Side Apply behavior.

Server-Side Apply (SSA) is a Kubernetes feature that allows multiple controllers to manage
different parts of an object. The listType markers help SSA understand how to merge arrays:

- listType=atomic: The entire list is replaced when updated
- listType=set: List elements are treated as a set (no duplicates, order doesn't matter)
- listType=map: Elements are identified by specific key fields for granular updates

**Important Note on listType=set:**
The use of listType=set is discouraged for object arrays due to Server-Side Apply
compatibility issues. When multiple controllers attempt to apply changes to an object
array with listType=set, the merge behavior can be unpredictable and may lead to
data loss or unexpected conflicts. For object arrays, use listType=atomic for simple
replacement semantics or listType=map for granular field-level merging.
listType=set is safe to use with primitive arrays (strings, integers, etc.).

The linter checks for:

1. Missing listType markers on array fields
2. Invalid listType values (must be atomic, set, or map)
3. Usage of listType=set on object arrays (discouraged due to compatibility issues)
4. Missing listMapKey markers for listType=map arrays
5. Incorrect usage of listType=map on primitive arrays

### Configuration

```yaml
lintersConfig:
  ssatags:
    listTypeSetUsage: Warn | Ignore # The policy for listType=set usage on object arrays. Defaults to `Warn`.
```

**Note:** listMapKey validation is always enforced and cannot be disabled. This ensures proper
Server-Side Apply behavior for arrays using listType=map.

## StatusOptional

The `statusoptional` linter checks that all first-level children fields within a status struct are marked as optional.

This is important because status fields should be optional to allow for partial updates and backward compatibility.
The linter ensures that all direct child fields of any status struct have either the `// +optional` or
`// +kubebuilder:validation:Optional` marker.

### Fixes

The `statusoptional` linter can automatically fix fields in status structs that are not marked as optional.

It will suggest adding the `// +optional` marker to any status field that is missing it.

## StatusSubresource

The `statussubresource` linter checks that the status subresource is configured correctly for
structs marked with the `kubebuilder:object:root:=true` marker. Correct configuration is that
when there is a status field the `kubebuilder:subresource:status` marker is present on the struct
OR when the `kubebuilder:subresource:status` marker is present on the struct there is a status field.

This linter is not enabled by default as it is only applicable to CustomResourceDefinitions.

### Fixes

In the case where there is a status field present but no `kubebuilder:subresource:status` marker, the
linter will suggest adding the comment `// +kubebuilder:subresource:status` above the struct.

## UniqueMarkers

The `uniquemarkers` linter ensures that types and fields do not contain more than a single definition of a marker that should only be present once.

Because this linter has no way of determining which marker definition was intended it does not suggest any fixes

### Configuration

It can configured to include a set of custom markers in the analysis by setting:

```yaml
lintersConfig:
  uniquemarkers:
    customMarkers:
      - identifier: custom:SomeCustomMarker
        attributes:
          - fruit
```

For each custom marker, it must specify an `identifier` and optionally some `attributes`.
As an example, take the marker definition `kubebuilder:validation:XValidation:rule='has(self.foo)',message='should have foo',fieldPath='.foo'`.
The identifier for the marker is `kubebuilder:validation:XValidation` and its attributes are `rule`, `message`, and `fieldPath`.

When specifying `attributes`, those attributes are included in the uniqueness identification of a marker definition.

Taking the example configuration from above:

- Marker definitions of `custom:SomeCustomMarker:fruit=apple,color=red` and `custom:SomeCustomMarker:fruit=apple,color=green` would violate the uniqueness requirement and be flagged.
- Marker definitions of `custom:SomeCustomMarker:fruit=apple,color=red` and `custom:SomeCustomMarker:fruit=orange,color=red` would _not_ violate the uniqueness requirement.

Each entry in `customMarkers` must have a unique `identifier`.

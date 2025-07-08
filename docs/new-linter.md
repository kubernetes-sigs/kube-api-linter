# Adding a new linter to KAL

Linters in KAL should live in their own package within the `pkg/analysis` directory.

Each linter is based on the `analysis.Analyzer` interface from the [golang.org/x/tools/go/analysis][go-analysis] package.

[go-analysis]: https://pkg.go.dev/golang.org/x/tools/go/analysis#hdr-Analyzer

The core of the linter is the `run` function, implemented with the signature:

```go
func (a *analyzer) run(pass *analysis.Pass) (interface{}, error)
```

It is recommended to implement the linter as a struct, that can contain configuration, and have the `run` function as a method on the struct.

It is also recommended to use the `inspect.Analyzer` pattern, which allows filtering the parsed syntax tree down to the types of nodes that are relevant to the linter.
This automates a lot of pre-work, and can be seen across existing linters, e.g. `jsontags` or `commentstart`.

Once you are within the `inspect.Preorder`, you can then implement the business logic of the linter, focusing on details of structs, fields, comments, etc.

## Registry

The registry in the analysis package co-ordinates the initialization of all linters.
Where linters have configuration, or are enabled/disabled by higher level configuration, the registry takes on making sure the linters are initialized correctly.

To enable the registry, each linter package must create an `Initializer` function that returns an `initializer.AnalyzerInitializer` interface (from `pkg/analysis/initializer`).

It is expected that each linter package contain a file `initializer.go`, the content of this file should be as follows:

```go
func init() {
    // Register the linter with the registry when the package is imported.
    kalanalysis.DefaultRegistry().RegisterLinter(Initializer())
}

// Initializer returns the AnalyzerInitializer for this
// Analyzer so that it can be added to the registry.
func Initializer() initializer.AnalyzerInitializer {
	return initializer.NewInitializer(
		name, // A constant containing the name of the linter. This should be lowercase.
		Analyzer, // An *analysis.Analyzer variable that is the linter.
		true, // Whether the linter is enabled by default.
	)
}
```

This pattern allows the linter to be registered with the KAL registry, and allows the linter to be initialized.

Once you have created the `initializer.go` file, you will need to import the linter package in the `pkg/analysis/registration/registration.go` file.

Once imported, the analyzer will be included in the linter builds.

## Configuration

Where the linter requires configuration, a slightly different pattern is used in `initializer.go`.
This time, use `NewConfigurableInitializer` instead of `NewInitializer` and pass in a function to validate the linter configuration.

```go

func init() {
	kalanalysis.DefaultRegistry().RegisterLinter(Initializer())
}

// Initializer returns the AnalyzerInitializer for this
// Analyzer so that it can be added to the registry.
func Initializer() initializer.AnalyzerInitializer {
	return initializer.NewConfigurableInitializer(
		name, // A constant containing the name of the linter. This should be lowercase.
		initAnalyzer, // A function that returns the initialized Analyzer.
		true, // Whether the linter is enabled by default.
		validateLintersConfig, // A function that validates the linter configuration.
	)
}

// initAnalyzer returns the intialized Analyzer.
func initAnalyzer(cfg *Config) (*analysis.Analyzer, error) {
	return newAnalyzer(cfg), nil
}

// validateLintersConfig validates the linter configuration.
func validateLintersConfig(cfg *Config, fldPath *field.Path) field.ErrorList {
	... // Validate the linter configuration.
}
```

The configuration struct should be defined within the analyzer package alongside the code.

Validation should be implemented in the `validateLintersConfig` function to ensure that the configuration is valid.
This validation function will be called before the `initAnalyzer` function is called.

## Helpers

The helpers package contains `analysis.Analyzer` implementations that can be used to source the common functionality required by linters.
For example, extracting information about `json` tags, or extracting `// +` style markers from comments.

Any new, common functionality should also be added to the helpers package.

Importantly:
* Helpers should expose a public `Analyzer` variable, of type `*analysis.Analyzer`.
* Helpers may not depend on any linter that needs to be initialized with configuration.
* Helpers themselves should not report lint issues, but should provide information to the linters that do.

In general, helpers return interfaces that can expose useful information in a simple way.
Exposing structs or maps directly as the result type of the helper means common
functions for accessing data must be implemented in each linter.

To use a helper, the helper `Analyzer` should be included in the linter `Analyzer`'s `Require` property.

```go
    return &analysis.Analyzer{
    Name:     "linterName",
    Doc:      "linter description",
    Requires: []*analysis.Analyzer{helpers.Analyzer},
    Run:      l.run,
}
```

Within the `run` function, the result of the helper can be extracted from the `*analysis.Pass` object.

```go
func (l *linter) run(pass *analysis.Pass) (interface{}, error) {
    helperResult := pass.ResultOf[helpers.Analyzer].(*helpers.ResultType)
    ...
}
```
## Marker Registry

If a linter needs to parse marker comments, it needs to register the identifier of the markers it
cares about with the `markers.MarkerRegistry` to ensure that markers are parsed in the expected way.

Marker identifiers are registered using an `init` function like so:
```go
func init() {
	markers.DefaultRegistry().Register("kubebuilder:object:root", ...)
}
```

## Tests

Basic tests can be implemented with the Go analysis test framework.

Create a `testdata` directory in the linter package and create a structure underneath.
Individual test files must be placed under `src` and then a subdirectory for each package.

If your linter has different configurations, e.g options to pass to the linter
you will need one configuration per option. Use one package per configuration
input for the linter.

```
mylinter
-- mylinter.go
-- mylinter_test.go
-- testdata
  -- src
    -- a
      -- a.go
    -- b
      -- b.go
```

The test suite can then be written using the standard go test framework.

```go
func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, commentstart.Analyzer, "a")
}
```

Each file within the test package should contain Go code to test the linter.
Typically this would mean a combination of constants, type declarations and struct declarations.

Where the linter is expected to return an issue, a comment can be added to the file to indicate the expected issue.

```go
type Foo struct {
    // this comment should be flagged by the linter // Want 'comment should start with a capital letter'
    Bar string

    foo string // want 'field is not exported'
}
```

If the expected output of the test happens to contain a regex string, then the regex within the `want` comment should be escaped.
The `jsontags` linter has an example of this pattern, which can be referred to.

### With suggested fixes

Where a linter also implements suggested fixes, the test suite can be extended to include the suggested fixes.

Replace `analysistest.Run` with `analysistest.RunWithSuggestedFixes`.

```go
func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, commentstart.Analyzer, "a")
}
```

For each file in the test packages, a corresponding `.golden` file should be created with the expected output, once the fixes have been applied.

The `commentstart` linter has an example of this pattern, which can be referred to.

## Docs

Each linter package should contain a `doc.go` file, which has a package level comment explaining what the linter is,
what it checks for, and how it can be configured, if appropriate.

The package level documentation is helpful when running `godoc` or accessing `pkg.go.dev`.

### Linters Documentation

The new linter should be added to the `docs/linters.md` file. This file contains the complete list of available linters and their detailed documentation.

Add the new linter in the correct alphabetical position in the table of contents at the top of the file, and then add a detailed section for the linter following the existing format:

````markdown
## LinterNameInPascalCase

The `linterNameInCamelCase` linter checks that [description of what the linter checks for].

### Configuration

Include an example of the configuration that can be used to configure the linter, if applicable.
```yaml
lintersConfig:
  linterNameInCamelCase:
    option: value
```

### Fixes

Include a description of what the linter fixes, if applicable.
````

The documentation should include:
- A clear description of what the linter checks for
- Configuration options with examples and descriptions
- Information about automatic fixes if the linter provides them
- Any relevant examples or edge cases

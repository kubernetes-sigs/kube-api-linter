/*
helpers contains utility functions that are used by the analysis package.
The helpers are used to extract data from the types, and provide common functionality that is used by multiple linters.

The available helpers are:
  - [extractjsontags]: Extracts JSON tags from struct fields and returns the information in a structured format.
  - [markers]: Extracts marker information from types and returns the information in a structured format.

Helpers should expose an *analysis.Analyzer as a globabl variable.
Other linters will use the `Requires` configuration to ensure that the helper is run before the linter.
The linter `Requires` relies on matching pointers to Analyzers, and therefore the helper cannot be dynamically created.
*/
package helpers

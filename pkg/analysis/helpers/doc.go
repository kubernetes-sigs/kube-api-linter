/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
helpers contains utility functions that are used by the analysis package.
The helpers are used to extract data from the types, and provide common functionality that is used by multiple linters.

The available helpers are:
  - [extractjsontags]: Extracts JSON tags from struct fields and returns the information in a structured format.
  - [markers]: Extracts marker information from types and returns the information in a structured format.

Helpers should expose an *analysis.Analyzer as a global variable.
Other linters will use the `Requires` configuration to ensure that the helper is run before the linter.
The linter `Requires` relies on matching pointers to Analyzers, and therefore the helper cannot be dynamically created.
*/
package helpers

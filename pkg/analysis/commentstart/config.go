/*
Copyright 2026 The Kubernetes Authors.

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
package commentstart

// defaultExcludePrefixes are built-in Go standard comment prefixes that are
// excluded from the "must start with JSON field name" check.
var defaultExcludePrefixes = []string{ //nolint:gochecknoglobals
	"Deprecated:",
}

// Config is the configuration for the commentstart linter.
type Config struct {
	// ExcludePrefixes is a list of additional comment prefixes that are
	// excluded from the "must start with JSON field name" check.
	// These are appended to the built-in defaults (e.g., "Deprecated:").
	// Each prefix is matched case-sensitively against the comment text
	// following the "// " prefix.
	// Must not contain duplicates or empty strings.
	ExcludePrefixes []string `json:"excludePrefixes,omitempty"`
}

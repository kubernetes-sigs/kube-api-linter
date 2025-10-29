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
package references

// Policy defines the policy for handling Ref/Refs.
type Policy string

const (
	// PolicyAllowRefAndRefs allows Ref/Refs in field names.
	PolicyAllowRefAndRefs Policy = "AllowRefAndRefs"
	// PolicyForbidRefAndRefs forbids Ref/Refs in field names.
	PolicyForbidRefAndRefs Policy = "ForbidRefAndRefs"
)

// Config represents the configuration for the references linter.
type Config struct {
	// Policy controls whether Ref/Refs are allowed or forbidden in field names.
	// When set to AllowRefAndRefs, fields containing Ref/Refs are allowed.
	// When set to ForbidRefAndRefs (default), fields containing Ref/Refs are flagged as errors.
	Policy Policy `json:"policy,omitempty"`
}

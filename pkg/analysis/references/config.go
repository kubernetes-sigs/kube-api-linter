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

// Config represents the configuration for the references linter.
type Config struct {
	// AllowRefAndRefs, when set to true, allows fields ending with "Ref" or "Refs".
	// This is useful for OpenShift compatibility where Ref/Refs fields are used instead of Reference/References.
	// When false (default), the linter will report errors for all fields ending with "Ref" or "Refs".
	AllowRefAndRefs bool `json:"allowRefAndRefs,omitempty"`
}

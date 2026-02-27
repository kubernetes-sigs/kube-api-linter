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

package discriminatedunions

// NonMemberFieldsPolicy controls how non-member fields are handled on union types.
type NonMemberFieldsPolicy string

const (
	// NonMemberFieldsForbid reports non-member fields on union types.
	NonMemberFieldsForbid NonMemberFieldsPolicy = "Forbid"
	// NonMemberFieldsAllow allows non-member fields on union types.
	NonMemberFieldsAllow NonMemberFieldsPolicy = "Allow"
)

// Config contains the configuration for the discriminatedunions linter.
type Config struct {
	// NonMemberFields defines how fields that are neither discriminator nor member are handled.
	NonMemberFields NonMemberFieldsPolicy `json:"nonMemberFields"`
}

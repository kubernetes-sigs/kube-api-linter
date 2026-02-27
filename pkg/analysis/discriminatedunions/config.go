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

// NonMemberFieldsPolicy controls how fields with no union marker are handled on union types.
type NonMemberFieldsPolicy string

const (
	// NonMemberFieldsForbid reports fields that are not union discriminators or union members.
	NonMemberFieldsForbid NonMemberFieldsPolicy = "Forbid"

	// NonMemberFieldsAllow permits fields that are not union discriminators or union members.
	NonMemberFieldsAllow NonMemberFieldsPolicy = "Allow"
)

// Config contains the configuration for the discriminatedunions linter.
type Config struct {
	// NonMemberFields defines how fields that are neither discriminator nor member are handled.
	// Valid values are "Forbid" and "Allow".
	// When set to "Forbid", the linter reports fields without union markers on union types.
	// When set to "Allow", non-member fields are permitted.
	// When otherwise not specified, the default value is "Forbid".
	NonMemberFields NonMemberFieldsPolicy `json:"nonMemberFields"`

	// PreferredOptionalMarker is the marker this linter recommends when union members are not marked optional.
	// Valid values are "optional", "kubebuilder:validation:Optional", and "k8s:optional".
	// When otherwise not specified, the default value is "optional".
	PreferredOptionalMarker string `json:"preferredOptionalMarker"`

	// EnforceCEL requires type-level XValidation rules for union membership on types
	// reachable from a kubebuilder object root in the same package.
	// When omitted or false, only structural checks are performed.
	EnforceCEL bool `json:"enforceCel"`
}

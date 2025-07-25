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
package serialization

// PointersPolicy is the policy for pointers.
// SuggestFix will suggest a fix for the field.
// Warn will warn about the field.
// Ignore will ignore the field.
type PointersPolicy string

const (
	// PointersPolicySuggestFix will suggest a fix for the field.
	PointersPolicySuggestFix PointersPolicy = "SuggestFix"

	// PointersPolicyWarn will warn about the field.
	PointersPolicyWarn PointersPolicy = "Warn"
)

// PointersPreference is the preference for pointers.
// Always will always suggest a fix for the field.
// WhenRequired will only suggest a fix for the field when it is required.
type PointersPreference string

const (
	// PointersPreferenceAlways will always suggest a pointer.
	PointersPreferenceAlways PointersPreference = "Always"

	// PointersPreferenceWhenRequired will only suggest a pointer for the field when it is required.
	PointersPreferenceWhenRequired PointersPreference = "WhenRequired"
)

// OmitEmptyPolicy is the policy for omitempty.
// SuggestFix will suggest a fix for the field.
// Warn will warn about the field.
// Ignore will ignore the field.
type OmitEmptyPolicy string

const (
	// OmitEmptyPolicySuggestFix will suggest a fix for the field.
	OmitEmptyPolicySuggestFix OmitEmptyPolicy = "SuggestFix"

	// OmitEmptyPolicyWarn will warn about the field.
	OmitEmptyPolicyWarn OmitEmptyPolicy = "Warn"

	// OmitEmptyPolicyIgnore will ignore the field.
	OmitEmptyPolicyIgnore OmitEmptyPolicy = "Ignore"
)

// Config is the configuration for the serialization check.
type Config struct {
	// Pointers is the configuration for pointers.
	Pointers PointersConfig

	// OmitEmpty is the configuration for omitempty.
	OmitEmpty OmitEmptyConfig
}

// PointersConfig is the configuration for pointers.
type PointersConfig struct {
	// Policy is the policy for pointers.
	Policy PointersPolicy

	// Preference is the preference for pointers.
	Preference PointersPreference
}

// OmitEmptyConfig is the configuration for omitempty.
type OmitEmptyConfig struct {
	// Policy is the policy for omitempty.
	Policy OmitEmptyPolicy
}

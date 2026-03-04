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
package enums

// KubebuilderEnumPolicy controls whether plain string fields with enum markers are allowed.
type KubebuilderEnumPolicy string

const (
	// KubebuilderEnumPolicyRequireTypeAlias enforces that string fields representing enums
	// must use type aliases instead of plain string types.
	KubebuilderEnumPolicyRequireTypeAlias KubebuilderEnumPolicy = "RequireTypeAlias"
	// KubebuilderEnumPolicyAllowPlainString allows plain string types for enum-validated fields.
	KubebuilderEnumPolicyAllowPlainString KubebuilderEnumPolicy = "AllowPlainString"
)

// Config is the configuration for the enums linter.
type Config struct {
	// Allowlist contains values that are exempt from PascalCase validation.
	// This is useful for command-line executable names like "kubectl", "docker", etc.
	Allowlist []string `yaml:"allowlist" json:"allowlist"`

	// KubebuilderEnumPolicy controls whether string fields with enum validation must use type aliases.
	// RequireTypeAlias (default): enforce type aliases; AllowPlainString: allow plain strings.
	KubebuilderEnumPolicy KubebuilderEnumPolicy `yaml:"kubebuilderEnumPolicy" json:"kubebuilderEnumPolicy"`
}

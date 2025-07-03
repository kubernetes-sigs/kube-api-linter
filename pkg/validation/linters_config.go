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
package validation

import (
	"sigs.k8s.io/kube-api-linter/pkg/analysis"
	"sigs.k8s.io/kube-api-linter/pkg/config"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateLintersConfig is used to validate the configuration in the config.LintersConfig struct.
func ValidateLintersConfig(l config.Linters, lc config.LintersConfig, fldPath *field.Path) field.ErrorList {
	return analysis.NewRegistry().ValidateLintersConfig(l, lc, fldPath)
}

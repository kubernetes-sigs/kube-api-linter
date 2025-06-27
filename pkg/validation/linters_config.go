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
	"sigs.k8s.io/kube-api-linter/pkg/analysis/conditions"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/jsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/nomaps"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/optionalfields"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/optionalorrequired"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/requiredfields"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/ssatags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/statusoptional"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/uniquemarkers"
	"sigs.k8s.io/kube-api-linter/pkg/config"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateLintersConfig is used to validate the configuration in the config.LintersConfig struct.
func ValidateLintersConfig(lc config.LintersConfig, fldPath *field.Path) field.ErrorList {
	fieldErrors := field.ErrorList{}

	fieldErrors = append(fieldErrors, conditions.Initializer().ValidateConfig(lc.Conditions, fldPath.Child("conditions"))...)
	fieldErrors = append(fieldErrors, jsontags.Initializer().ValidateConfig(lc.JSONTags, fldPath.Child("jsonTags"))...)
	fieldErrors = append(fieldErrors, nomaps.Initializer().ValidateConfig(lc.NoMaps, fldPath.Child("nomaps"))...)
	fieldErrors = append(fieldErrors, optionalfields.Initializer().ValidateConfig(lc.OptionalFields, fldPath.Child("optionalFields"))...)
	fieldErrors = append(fieldErrors, optionalorrequired.Initializer().ValidateConfig(lc.OptionalOrRequired, fldPath.Child("optionalOrRequired"))...)
	fieldErrors = append(fieldErrors, requiredfields.Initializer().ValidateConfig(lc.RequiredFields, fldPath.Child("requiredFields"))...)
	fieldErrors = append(fieldErrors, ssatags.Initializer().ValidateConfig(lc.SSATags, fldPath.Child("ssatags"))...)
	fieldErrors = append(fieldErrors, statusoptional.Initializer().ValidateConfig(lc.StatusOptional, fldPath.Child("statusOptional"))...)
	fieldErrors = append(fieldErrors, uniquemarkers.Initializer().ValidateConfig(lc.UniqueMarkers, fldPath.Child("uniqueMarkers"))...)

	return fieldErrors
}

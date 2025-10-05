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

// Package markerscope provides a linter that validates markers are applied in the correct scope.
//
// Some markers are only valid when applied to specific Go constructs:
// - Field-only markers: optional, required, nullable
// - Type/Struct-only markers: MinProperties, MaxProperties, kubebuilder:object:root, kubebuilder:subresource:status
// - Field or Type markers: default, MinLength, MaxLength, etc.
//
// This linter ensures markers are applied in their appropriate contexts to prevent
// configuration errors and improve API consistency.
package markerscope
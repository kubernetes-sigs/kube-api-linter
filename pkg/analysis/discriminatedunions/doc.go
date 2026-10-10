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

// Package discriminatedunions validates discriminated-union marker usage.
//
// This linter detects union types using either legacy markers (`+union`,
// `+unionDiscriminator`, `+unionMember`) or declarative markers
// (`+k8s:unionDiscriminator`, `+k8s:unionMember`). Union detection is triggered
// by either a type-level `+union` marker, legacy union field markers or a
// declarative discriminator marker. Named declarative unions are checked
// separately; member-only declarative unions are outside this linter's scope.
//
// It enforces:
//   - Exactly one discriminator field per union.
//   - Discriminator fields are required.
//   - Member fields are marked optional with an optional marker.
//   - Optional forbidding of non-member fields on union types.
//   - Opt-in CEL membership rules for unions reachable from CRD roots.
//
// The legacy `+unionMember,optional` marker is treated as union membership
// metadata and does not replace field optionality markers such as `+optional`
// or `+k8s:optional`.
//
// With enforceCel enabled, the linter follows same-package types from
// `+kubebuilder:object:root=true` roots and requires a type-level XValidation
// rule for each member. It recognizes the required and optional membership
// ternary patterns, tolerating formatting differences but not arbitrary
// logically equivalent expressions. It does not provide automatic fixes.
package discriminatedunions

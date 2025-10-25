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

/*
markertypos is an analyzer that checks for common typos and syntax issues in marker comments.

This linter validates three main categories of marker issues:

 1. Spacing Issues:
    Detects and fixes incorrect spacing after the '+' symbol in markers.

    For example, this would be reported:
    // + kubebuilder:validation:MaxLength:=256
    type Foo string

    And should be fixed to:
    // +kubebuilder:validation:MaxLength:=256
    type Foo string

 2. Syntax Issues:
    Validates that kubebuilder markers use ':=' syntax while non-kubebuilder markers use '=' syntax.

    Kubebuilder markers should use ':=' syntax:
    // +kubebuilder:validation:MaxLength:=256  (correct)
    // +kubebuilder:validation:MaxLength=256   (incorrect, will be reported)

    Non-kubebuilder markers should use '=' syntax:
    // +default:value="test"   (correct)
    // +default:value:="test"  (incorrect, will be reported)

 3. Common Typos:
    Detects and suggests corrections for frequently misspelled marker identifiers.

    Examples of typos that would be reported:
    // +kubebuidler:validation:Required  → should be 'kubebuilder'
    // +optinal                          → should be 'optional'
    // +requied                          → should be 'required'
    // +kubebuilder:validaton:MaxLength  → should be 'validation'

This linter provides automatic fixes for all detected issues, making it easy to maintain
consistent and correct marker syntax across Kubernetes API definitions.
*/
package markertypos

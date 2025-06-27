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
* The `forbiddenmarkers` linter ensures that types and fields do not contain any markers
* that are forbidden.
*
* By default, `forbiddenmarkers` is not enabled.
*
* It can be configured with a list of marker identifiers that are forbidden
* ```yaml
* lintersConfig:
*   forbiddenMarkers:
*     markers:
*       - some:forbidden:marker
*       - anotherforbidden
* ```
*
* Fixes are suggested to remove all markers that are forbidden.
*
 */
package forbiddenmarkers

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
conflictingmarkers is a linter that detects and reports when mutually exclusive markers are used on the same field.
This prevents common configuration errors and unexpected behavior in Kubernetes API types.

The linter reports issues when markers from both sets of a conflict pair are present on the same field.
It does NOT report issues when multiple markers from the same set are present - only when markers from
different sets within the same conflict definition are found together.

The linter includes built-in checks for the following conflicts:

1. Optional vs Required:
  - Set A: optional, +kubebuilder:validation:Optional
  - Set B: required, +kubebuilder:validation:Required
  - A field cannot be both optional and required
  - Note: Multiple optional markers or multiple required markers are allowed

2. Default vs Required:
  - Set A: default, +kubebuilder:default
  - Set B: required, +kubebuilder:validation:Required
  - A field with a default value cannot be required
  - Note: Multiple default markers or multiple required markers are allowed

The linter is configurable and allows users to define their own custom sets of conflicting markers.
Each custom conflict set must specify a name, two sets of markers, and a description of why they conflict.

Example configuration:
```yaml
lintersConfig:

	conflictingmarkers:
	  customConflicts:
	    - name: "my_custom_conflict"
	      setA: ["custom:marker1", "custom:marker2"]
	      setB: ["custom:marker3", "custom:marker4"]
	      description: "These markers conflict because..."

```

The linter does not provide automatic fixes as it cannot determine which conflicting marker should be removed.
*/
package conflictingmarkers

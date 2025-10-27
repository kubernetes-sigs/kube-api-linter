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
The `references` linter ensures that field names use 'Ref'/'Refs' suffixes and not 'Reference'/'References'.

By default, `references` is enabled and enforces this naming convention.

The linter checks that fields ending with 'Reference' are suggested to use 'Ref' instead.
Similarly, fields ending with 'References' are suggested to use 'Refs' instead.

Example configuration:

**Default behavior (report errors for Reference/References):**
```yaml
lintersConfig:

	references: {}

```

**For OpenShift compatibility (allow Ref/Refs):**
```yaml
lintersConfig:

	references:
	  allowRefAndRefs: true

```

When `allowRefAndRefs` is set to false (the default), fields ending with 'Ref' or 'Refs' (other than those matched by the above rules) will also be reported as errors.
This is useful to ensure consistency across the codebase. However, for OpenShift compatibility, this option can be set to true to allow such field names.
*/
package references

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
The `references` linter ensures that field names use 'Ref'/'Refs' instead of 'Reference'/'References'.

By default, `references` is enabled and enforces this naming convention.

The linter checks that 'Reference' anywhere in field names (beginning, middle, or end) is replaced with 'Ref'.
Similarly, 'References' anywhere in field names is replaced with 'Refs'.

Example configuration:

**Default behavior (forbid Ref/Refs in field names):**
```yaml
lintersConfig:

	references:
	  policy: ForbidRefAndRefs

```

**For compatibility (allow Ref/Refs in field names):**
```yaml
lintersConfig:

	references:
	  policy: AllowRefAndRefs

```

When `policy` is set to `ForbidRefAndRefs` (the default), fields containing 'Ref' or 'Refs' anywhere in their names will be reported as errors.
This is useful to ensure consistency across the codebase. The policy can be set to `AllowRefAndRefs` to allow such field names.
*/
package references

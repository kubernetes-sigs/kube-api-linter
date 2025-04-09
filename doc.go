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
KAL is a linter for Kubernetes API types, that implements API conventions and best practices.

This package provides a GolangCI-Lint plugin that can be used to build a custom linter for Kubernetes API types.
The custom golangci-lint binary can be built by checking out the KAL repository and running `make build-golangci`.
This will generate a custom golangci-lint binary in the `bin` directory.

The custom golangci-lint binary can be run with the `run` command, and the KAL linters can be enabled by setting the `kal` linter in the `.golangci.yml` configuration file.

Example `.golangci.yml` configuration file:

	linters-settings:
	custom:
	  kal:
	  type: "module"
	  description: KAL is the Kube-API-Linter and lints Kube like APIs based on API conventions and best practices.
	  settings:
	    linters:
	      enabled: []
	      disabled: []
	    lintersConfig:
	      jsonTags:
	        jsonTagRegex: ""
	      optionalOrRequired:
	        preferredOptionalMarker: ""
	        preferredRequiredMarker: ""
	linters:
	  disable-all: true
	  enable:
	    - kal

New linters can be added in the [sigs.k8s.io/kube-api-linter/pkg/analysis] package.
*/
package kal

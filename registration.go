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
// This file is intended to allow folks forking KAL to customise what is/isn't
// enabled in their own fork.
// This should be the only file to modify, by commenting out the registration import
// and importing the linters they want to enable.
package kubeapilinter

import (
	// Importing the registration package enables all of the default linters for KAL.
	// DO NOT ADD DIRECTLY TO THIS FILE.
	_ "sigs.k8s.io/kube-api-linter/pkg/registration"
)

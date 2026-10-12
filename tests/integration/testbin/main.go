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

// testbin is a golangci-lint binary that includes the standard KAL linters
// plus a custom test preset, simulating a fork maintainer's custom binary.
package main

import (
	"fmt"
	"os"

	"github.com/golangci/golangci-lint/v2/pkg/commands"
	"github.com/golangci/golangci-lint/v2/pkg/exitcodes"

	// Standard KAL registration (linters + built-in presets).
	_ "sigs.k8s.io/kube-api-linter"

	// Custom test preset — simulates a fork maintainer's custom preset package.
	_ "sigs.k8s.io/kube-api-linter/tests/integration/testpreset"
)

func main() {
	info := commands.BuildInfo{
		Version: "test",
	}

	if err := commands.Execute(info); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(exitcodes.Failure)
	}
}

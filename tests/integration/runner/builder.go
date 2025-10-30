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
package runner

// RunnerBuilder is used to configure a Runner.
type RunnerBuilder struct {
	binPath  string
	exitCode int
}

// NewRunnerBuilder creates a new builder.
func NewRunnerBuilder() RunnerBuilder {
	return RunnerBuilder{}
}

// Runner builds the test Runner.
func (r RunnerBuilder) Runner() Runner {
	return &runner{
		binPath:  r.binPath,
		exitCode: r.exitCode,
	}
}

// WithBinPath sets the path to the golangci-lint binary including kubeapilinter.
func (r RunnerBuilder) WithBinPath(binPath string) RunnerBuilder {
	r.binPath = binPath
	return r
}

// WithExitCode sets the expected exit code of golangci-lint.
func (r RunnerBuilder) WithExitCode(exitCode int) RunnerBuilder {
	r.exitCode = exitCode
	return r
}

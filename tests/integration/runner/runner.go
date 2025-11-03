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

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/golangci/golangci-lint/v2/test/testshared"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	configYAML = "config.yaml"
	runCommand = "run"

	baseGolangCILintConfigTmpl = `version: "2"
linters:
  default: none
  enable:
  - kubeapilinter
  settings:
    custom:
      kubeapilinter:
        type: "module"
        description: kubeapilinter is the Kube-API-Linter and lints Kube like APIs based on API conventions and best practices.
        settings:
          {{ .cfg }}
`
)

// Runner is used to run tests.
type Runner interface {
	RunTestsFromDir(dir string)
}

type runner struct {
	binPath  string
	exitCode int
}

// RunTestsFromDir executes golangci-lint with kubeapilinter
// against the Go source files in the given directory.
// It expects a config.yaml file to configure the KAL settings.
// Use `want` style comments to annotate test files with any
// expected linter issues.
func (r *runner) RunTestsFromDir(dir string) {
	configData, err := os.ReadFile(filepath.Join(dir, configYAML))
	if err != nil && !os.IsNotExist(err) {
		ExpectWithOffset(1, err).ToNot(HaveOccurred())
	}

	tmpl, err := template.New("config").Parse(baseGolangCILintConfigTmpl)
	Expect(err).ToNot(HaveOccurred())

	buf := bytes.NewBuffer(nil)
	Expect(tmpl.Execute(buf, map[string]string{"cfg": string(configData)})).To(Succeed())

	args := []string{
		"--default=none",
		"--show-stats=false",
		"--output.json.path=stdout",
		"--max-same-issues=1000",
		"--max-issues-per-linter=1000",
		"--uniq-by-line=false",
	}

	golangciRunner := testshared.NewRunnerBuilder(ginkgo.GinkgoTB()).
		WithBinPath(r.binPath).
		WithCommand(runCommand).
		WithArgs(args...).
		WithConfig(buf.String()).
		WithTargetPath(dir).
		Runner().
		Command()

	output, err := golangciRunner.CombinedOutput()

	// The returned error will be nil if the test file does not have any issues
	// and thus the linter exits with exit code 0.
	// So perform the additional assertions only if the error is non-nil.
	if err != nil {
		var exitErr *exec.ExitError

		Expect(errors.As(err, &exitErr)).To(BeTrue())
	}

	Expect(golangciRunner.ProcessState.ExitCode()).To(Equal(r.exitCode), fmt.Sprintf("expected golangci-lint to run correctly: %s, config: %s", string(output), buf.String()))

	analyze(dir, output)
}

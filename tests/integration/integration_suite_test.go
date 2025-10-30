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
package integration

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	ctx     = context.Background()
	binPath string
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration")
}

var _ = BeforeSuite(func() {
	tempDir, err := os.MkdirTemp("", "kube-api-linter-integration-")
	Expect(err).ToNot(HaveOccurred())

	binPath = filepath.Join(tempDir, "golangci-lint")

	_, err = exec.CommandContext(ctx, "go", "build", "-o", binPath, "../../cmd/golangci-lint-kube-api-linter/").CombinedOutput()
	Expect(err).ToNot(HaveOccurred())
})

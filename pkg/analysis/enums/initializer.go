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
package enums

import (
	"golang.org/x/tools/go/analysis"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/registry"
)

func init() {
	registry.DefaultRegistry().RegisterLinter(Initializer())
}

// Initializer returns the AnalyzerInitializer for this
// Analyzer so that it can be added to the registry.
func Initializer() initializer.AnalyzerInitializer {
	return initializer.NewConfigurableInitializer(
		name,
		initAnalyzer,
		// Enabled by default: validates string type aliases with constants have enum markers
		true,
		validateConfig,
	)
}

func initAnalyzer(cfg *Config) (*analysis.Analyzer, error) {
	if cfg == nil {
		cfg = &Config{}
	}

	if cfg.KubebuilderEnumPolicy == "" {
		cfg.KubebuilderEnumPolicy = KubebuilderEnumPolicyRequireTypeAlias
	}

	return newAnalyzer(cfg), nil
}

// validateConfig implements validation of the enums linter config.
func validateConfig(cfg *Config, fldPath *field.Path) field.ErrorList {
	var errs field.ErrorList
	if cfg == nil {
		return errs
	}

	allowlistPath := fldPath.Child("allowlist")
	seen := make(map[string]bool, len(cfg.Allowlist))

	for i, v := range cfg.Allowlist {
		if seen[v] {
			errs = append(errs, field.Duplicate(allowlistPath.Index(i), v))
		}

		seen[v] = true
	}

	return errs
}

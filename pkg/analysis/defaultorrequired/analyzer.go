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
package defaultorrequired

import (
	"golang.org/x/tools/go/analysis"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/conflictingmarkers"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
)

const name = "defaultorrequired"

// newAnalyzer creates a new analyzer that wraps conflictingmarkers with a predefined configuration
// for checking default and required marker conflicts.
func newAnalyzer() *analysis.Analyzer {
	cfg := &conflictingmarkers.ConflictingMarkersConfig{
		Conflicts: []conflictingmarkers.ConflictSet{
			{
				Name: "default_or_required",
				Sets: [][]string{
					{markers.DefaultMarker, markers.KubebuilderDefaultMarker},
					{markers.RequiredMarker, markers.KubebuilderRequiredMarker, markers.K8sRequiredMarker},
				},
				Description: "A field with a default value cannot be required. A required field must be provided by the user, so a default value is not meaningful.",
			},
		},
	}

	// Create the underlying conflicting markers analyzer
	analyzer := conflictingmarkers.NewAnalyzer(cfg)

	// Override the name to match this linter
	analyzer.Name = name
	analyzer.Doc = "Checks that fields marked as required do not have default values applied"

	return analyzer
}

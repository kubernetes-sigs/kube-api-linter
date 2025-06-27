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
package nonullable

import (
	"golang.org/x/tools/go/analysis"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/forbiddenmarkers"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

const name = "nonullable"
const doc = "Check that nullable marker is not present on any types or fields."

func newAnalyzer() *analysis.Analyzer {
	analyzer := forbiddenmarkers.NewAnalyzer(config.ForbiddenMarkersConfig{
		Markers: []config.ForbiddenMarker{
			{
				Identifier: "nullable",
			},
		},
	}, forbiddenmarkers.WithName(name), forbiddenmarkers.WithDoc(doc))

	return analyzer
}

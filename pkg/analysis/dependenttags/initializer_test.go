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

package dependenttags_test

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/dependenttags"
)

func TestInitializer(t *testing.T) {
	cases := []struct {
		name      string
		cfg       *dependenttags.Config
		expectErr bool
	}{
		{
			name: "valid config",
			cfg: &dependenttags.Config{
				Rules: []dependenttags.Rule{
					{
						Identifier: "k8s:unionMember",
						Type:       dependenttags.DependencyTypeAll,
						Dependents: []string{"k8s:optional"},
					},
				},
			},
			expectErr: false,
		},
		{
			name: "missing type",
			cfg: &dependenttags.Config{
				Rules: []dependenttags.Rule{
					{
						Identifier: "k8s:unionMember",
						Dependents: []string{"k8s:optional"},
					},
				},
			},
			expectErr: true,
		},
		{
			name: "empty rules",
			cfg: &dependenttags.Config{
				Rules: []dependenttags.Rule{},
			},
			expectErr: true,
		},
		{
			name: "missing identifier",
			cfg: &dependenttags.Config{
				Rules: []dependenttags.Rule{
					{
						Type:       dependenttags.DependencyTypeAll,
						Dependents: []string{"k8s:optional"},
					},
				},
			},
			expectErr: true,
		},
		{
			name: "missing dependents",
			cfg: &dependenttags.Config{
				Rules: []dependenttags.Rule{
					{
						Identifier: "k8s:unionMember",
						Type:       dependenttags.DependencyTypeAll,
					},
				},
			},
			expectErr: true,
		},
		{
			name: "invalid type",
			cfg: &dependenttags.Config{
				Rules: []dependenttags.Rule{
					{
						Identifier: "k8s:unionMember",
						Type:       "invalid",
						Dependents: []string{"k8s:optional"},
					},
				},
			},
			expectErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			initializer := dependenttags.Initializer()
			errs := initializer.ValidateConfig(tc.cfg, field.NewPath(""))

			if tc.expectErr && len(errs) == 0 {
				t.Errorf("expected validation errors, but got none")
			}

			if !tc.expectErr && len(errs) > 0 {
				t.Errorf("unexpected validation errors: %v", errs)
			}

			if !tc.expectErr {
				if _, err := initializer.Init(tc.cfg); err != nil {
					t.Errorf("unexpected error initializing analyzer: %v", err)
				}
			}
		})
	}
}

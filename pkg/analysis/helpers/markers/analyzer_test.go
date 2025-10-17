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
package markers

import (
	"go/ast"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	. "github.com/onsi/gomega"
)

func TestExtractMarker(t *testing.T) {
	type testcase struct {
		name     string
		comment  *ast.Comment
		expected Marker
	}

	testcases := []testcase{
		// Kubebuilder-style markers
		{
			name:    "non-namespaced marker",
			comment: &ast.Comment{Text: "// +required"},
			expected: Marker{
				Type:       MarkerTypeKubebuilder,
				Identifier: "required",
				Arguments:  make(map[string]string),
			},
		},
		{
			name:    "non-namespaced marker with payload value",
			comment: &ast.Comment{Text: "// +listType=atomic"},
			expected: Marker{
				Type:       MarkerTypeKubebuilder,
				Identifier: "listType",
				Arguments:  make(map[string]string),
				Payload: Payload{
					Value: "atomic",
				},
			},
		},
		{
			name:    "kubebuilder marker with single unnamed expression using '='",
			comment: &ast.Comment{Text: "// +kubebuilder:object:root=true"},
			expected: Marker{
				Type:       MarkerTypeKubebuilder,
				Identifier: "kubebuilder:object:root",
				Arguments:  make(map[string]string),
				Payload: Payload{
					Value: "true",
				},
			},
		},
		{
			name:    "kubebuilder marker with single unnamed expression using ':='",
			comment: &ast.Comment{Text: "// +kubebuilder:object:root:=true"},
			expected: Marker{
				Type:       MarkerTypeKubebuilder,
				Identifier: "kubebuilder:object:root",
				Arguments:  make(map[string]string),
				Payload: Payload{
					Value: "true",
				},
			},
		},
		{
			name:    "kubebuilder marker with no expressions",
			comment: &ast.Comment{Text: "// +kubebuilder:validation:Required"},
			expected: Marker{
				Type:       MarkerTypeKubebuilder,
				Identifier: "kubebuilder:validation:Required",
				Arguments:  make(map[string]string),
			},
		},
		{
			name:    "kubebuilder marker with multiple named expressions",
			comment: &ast.Comment{Text: "// +kubebuilder:validation:XValidation:rule='has(self.field)',message='must have field!'"},
			expected: Marker{
				Type:       MarkerTypeKubebuilder,
				Identifier: "kubebuilder:validation:XValidation",
				Arguments: map[string]string{
					"rule":    "'has(self.field)'",
					"message": "'must have field!'",
				},
			},
		},
		{
			name:    "other namespaced marker in kubebuilder-style with expression wrapped in double quotes (\")",
			comment: &ast.Comment{Text: "// +foo:bar:rule=\"foo\""},
			expected: Marker{
				Type:       MarkerTypeKubebuilder,
				Identifier: "foo:bar:rule",
				Arguments:  make(map[string]string),
				Payload: Payload{
					Value: "\"foo\"",
				},
			},
		},
		{
			name:    "kubebuilder marker with expression with a comma in its value",
			comment: &ast.Comment{Text: `// +kubebuilder:validation:XValidation:rule='self.map(a, a == "someValue")',message='must have field!'`},
			expected: Marker{
				Type:       MarkerTypeKubebuilder,
				Identifier: "kubebuilder:validation:XValidation",
				Arguments: map[string]string{
					"rule":    `'self.map(a, a == "someValue")'`,
					"message": "'must have field!'",
				},
			},
		},
		{
			name:    "kubebuilder marker with expression with a comma in its value with double quotes",
			comment: &ast.Comment{Text: `// +kubebuilder:validation:XValidation:rule="self.map(a, a == \"someValue\")",message="must have field!"`},
			expected: Marker{
				Type:       MarkerTypeKubebuilder,
				Identifier: "kubebuilder:validation:XValidation",
				Arguments: map[string]string{
					"rule":    `"self.map(a, a == \"someValue\")"`,
					"message": `"must have field!"`,
				},
			},
		},
		{
			name:    "kubebuilder marker with expression ending in a valid double quote",
			comment: &ast.Comment{Text: `// +kubebuilder:validation:Enum:=foo;bar;baz;""`},
			expected: Marker{
				Type:       MarkerTypeKubebuilder,
				Identifier: "kubebuilder:validation:Enum",
				Arguments:  make(map[string]string),
				Payload: Payload{
					Value: "foo;bar;baz;\"\"",
				},
			},
		},
		{
			name:    "other namespaced kubebuilder-style marker with chained expressions without quotes",
			comment: &ast.Comment{Text: `// +custom:marker:fruit=apple,color=blue,country=UK`},
			expected: Marker{
				Type:       MarkerTypeKubebuilder,
				Identifier: "custom:marker",
				Arguments: map[string]string{
					"fruit":   "apple",
					"color":   "blue",
					"country": "UK",
				},
			},
		},
		{
			name:    "kubebuilder marker with numeric value",
			comment: &ast.Comment{Text: `// +kubebuilder:validation:Minimum=10`},
			expected: Marker{
				Type:       MarkerTypeKubebuilder,
				Identifier: "kubebuilder:validation:Minimum",
				Arguments:  make(map[string]string),
				Payload: Payload{
					Value: "10",
				},
			},
		},
		{
			name:    "kubebuilder marker with negative numeric value",
			comment: &ast.Comment{Text: `// +kubebuilder:validation:Minimum=-10`},
			expected: Marker{
				Type:       MarkerTypeKubebuilder,
				Identifier: "kubebuilder:validation:Minimum",
				Arguments:  make(map[string]string),
				Payload: Payload{
					Value: "-10",
				},
			},
		},
		{
			name:    "kubebuilder marker with named expression using backtick ('`') for strings",
			comment: &ast.Comment{Text: "// +kubebuilder:validation:XValidation:rule=`has(self.field)`,message=`must have field!`"},
			expected: Marker{
				Type:       MarkerTypeKubebuilder,
				Identifier: "kubebuilder:validation:XValidation",
				Arguments: map[string]string{
					"rule":    "`has(self.field)`",
					"message": "`must have field!`",
				},
			},
		},

		//  Not actually markers
		{
			name:     "invalid marker - markdown table border",
			comment:  &ast.Comment{Text: "// +-------+-------+-------+"},
			expected: Marker{},
		},
		{
			name:     "invalid marker - markdown table border without pipes",
			comment:  &ast.Comment{Text: "// +----------"},
			expected: Marker{},
		},
		{
			name:     "invalid marker - starts with special characters",
			comment:  &ast.Comment{Text: "// +!*@(#&KSDJUF:A"},
			expected: Marker{},
		},
		{
			name:     "regular comment - no plus sign",
			comment:  &ast.Comment{Text: "// This is a regular comment"},
			expected: Marker{},
		},

		// Declarative Validation Tag Parsing
		{
			name: "simple declarative validation marker",
			comment: &ast.Comment{
				Text: "// +k8s:required",
			},
			expected: Marker{
				Type:       MarkerTypeDeclarativeValidation,
				Identifier: "k8s:required",
				Arguments:  make(map[string]string),
			},
		},
		{
			name: "declarative validation marker with a value",
			comment: &ast.Comment{
				Text: "// +k8s:maxLength=10",
			},
			expected: Marker{
				Type:       MarkerTypeDeclarativeValidation,
				Identifier: "k8s:maxLength",
				Arguments:  make(map[string]string),
				Payload: Payload{
					Value: "10",
				},
			},
		},
		{
			name: "declarative validation marker with named argument",
			comment: &ast.Comment{
				Text: "// +k8s:unionMember(union: \"union1\")",
			},
			expected: Marker{
				Type:       MarkerTypeDeclarativeValidation,
				Identifier: "k8s:unionMember",
				Arguments: map[string]string{
					"union": "union1",
				},
			},
		},
		{
			name:    "declarative validation marker with multiple named arguments",
			comment: &ast.Comment{Text: "// +k8s:someThing(one: \"a\", two: \"b\")=+k8s:required"},
			expected: Marker{
				Type:       MarkerTypeDeclarativeValidation,
				Identifier: "k8s:someThing",
				Arguments: map[string]string{
					"one": "a",
					"two": "b",
				},
				Payload: Payload{
					Marker: &Marker{
						Type:       MarkerTypeDeclarativeValidation,
						Identifier: "k8s:required",
						Arguments:  make(map[string]string),
					},
				},
			},
		},
		{
			name: "declarative validation marker with unnamed argument",
			comment: &ast.Comment{
				Text: "// +k8s:doesWork(100)", // not a real DV marker, but AFAIK nothing stops something like this from coming up
			},
			expected: Marker{
				Type:       MarkerTypeDeclarativeValidation,
				Identifier: "k8s:doesWork",
				Arguments: map[string]string{
					"": "100",
				},
			},
		},
		{
			name:    "declarative validation marker with unnamed argument in backticks",
			comment: &ast.Comment{Text: "// +k8s:doesWork(`foo`)"}, // not a real DV marker, but AFAIK nothing stops something like this from coming up
			expected: Marker{
				Type:       MarkerTypeDeclarativeValidation,
				Identifier: "k8s:doesWork",
				Arguments: map[string]string{
					"": "foo",
				},
			},
		},
		{
			name: "declarative validation marker with unnamed argument and simple validation tag payload",
			comment: &ast.Comment{
				Text: "// +k8s:ifEnabled(\"my-feature\")=+k8s:required",
			},
			expected: Marker{
				Type:       MarkerTypeDeclarativeValidation,
				Identifier: "k8s:ifEnabled",
				Arguments: map[string]string{
					"": "my-feature",
				},
				Payload: Payload{
					Marker: &Marker{
						Type:       MarkerTypeDeclarativeValidation,
						Identifier: "k8s:required",
						Arguments:  make(map[string]string),
					},
				},
			},
		},
		{
			name: "declarative validation marker with deeper chained validation tags",
			comment: &ast.Comment{
				Text: "// +k8s:ifEnabled(\"my-feature\")=+k8s:item(type: \"Approved\")=+k8s:zeroOrOneOfMember",
			},
			expected: Marker{
				Type:       MarkerTypeDeclarativeValidation,
				Identifier: "k8s:ifEnabled",
				Arguments: map[string]string{
					"": "my-feature",
				},
				Payload: Payload{
					Marker: &Marker{
						Type:       MarkerTypeDeclarativeValidation,
						Identifier: "k8s:item",
						Arguments: map[string]string{
							"type": "Approved",
						},
						Payload: Payload{
							Marker: &Marker{
								Type:       MarkerTypeDeclarativeValidation,
								Identifier: "k8s:zeroOrOneOfMember",
								Arguments:  make(map[string]string),
							},
						},
					},
				},
			},
		},
		{
			name: "declarative validation marker parsing error",
			comment: &ast.Comment{
				Text: "// +k8s:ifEnabled(\"my-feature\")=a,b,c,d", // DV tags do not allow comma separated payloads
			},
			expected: Marker{},
		},
		{
			name: "declarative validation ':='",
			comment: &ast.Comment{
				Text: "// +k8s:format:=password", // DV tags parse out ':=' weirdly. This is a typo, DV tags are supposed to use '='.
			},
			expected: Marker{
				Type:       MarkerTypeDeclarativeValidation,
				Identifier: "k8s:format:",
				Arguments:  make(map[string]string),
				Payload: Payload{
					Value: "password",
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			marker := extractMarker(tc.comment)

			diff := cmp.Diff(marker, tc.expected, cmpopts.IgnoreFields(Marker{}, "RawComment", "End", "Pos"))
			g.Expect(diff).To(BeEmpty())
		})
	}
}

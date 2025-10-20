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

	. "github.com/onsi/gomega"
)

func TestExtractMarkerIdAndExpressions(t *testing.T) {
	type testcase struct {
		name                string
		marker              string
		expectedID          string
		expectedExpressions map[string]string
	}

	testcases := []testcase{
		{
			name:       "registered marker with single unnamed expression using '='",
			marker:     "kubebuilder:object:root=true",
			expectedID: "kubebuilder:object:root",
			expectedExpressions: map[string]string{
				"": "true",
			},
		},
		{
			name:       "registered marker with single unnamed expression using ':='",
			marker:     "kubebuilder:object:root:=true",
			expectedID: "kubebuilder:object:root",
			expectedExpressions: map[string]string{
				"": "true",
			},
		},
		{
			name:                "registered marker with no expressions",
			marker:              "required",
			expectedID:          "required",
			expectedExpressions: map[string]string{},
		},
		{
			name:       "registered marker with multiple named expressions",
			marker:     "kubebuilder:validation:XValidation:rule='has(self.field)',message='must have field!'",
			expectedID: "kubebuilder:validation:XValidation",
			expectedExpressions: map[string]string{
				"rule":    "'has(self.field)'",
				"message": "'must have field!'",
			},
		},
		{
			name:       " unregistered marker with expression wrapped in double quotes (\")",
			marker:     "foo:bar:rule=\"foo\"",
			expectedID: "foo:bar:rule",
			expectedExpressions: map[string]string{
				"": "\"foo\"",
			},
		},
		{
			name:       "registered marker with expression with a comma in its value",
			marker:     `kubebuilder:validation:XValidation:rule='self.map(a, a == "someValue")',message='must have field!'`,
			expectedID: "kubebuilder:validation:XValidation",
			expectedExpressions: map[string]string{
				"rule":    `'self.map(a, a == "someValue")'`,
				"message": "'must have field!'",
			},
		},
		{
			name:       "registered marker with expression with a comma in its value with double quotes",
			marker:     `kubebuilder:validation:XValidation:rule="self.map(a, a == \"someValue\")",message="must have field!"`,
			expectedID: "kubebuilder:validation:XValidation",
			expectedExpressions: map[string]string{
				"rule":    `"self.map(a, a == \"someValue\")"`,
				"message": `"must have field!"`,
			},
		},
		{
			name:       "registered marker with expression ending in a valid double quote",
			marker:     `kubebuilder:validation:Enum:=foo;bar;baz;""`,
			expectedID: "kubebuilder:validation:Enum",
			expectedExpressions: map[string]string{
				"": `foo;bar;baz;""`,
			},
		},
		{
			name:       "registered marker with chained expressions without quotes",
			marker:     `custom:marker:fruit=apple,color=blue,country=UK`,
			expectedID: "custom:marker",
			expectedExpressions: map[string]string{
				"fruit":   "apple",
				"color":   "blue",
				"country": "UK",
			},
		},
		{
			name:       "registered marker with numeric value",
			marker:     `kubebuilder:validation:Minimum=10`,
			expectedID: "kubebuilder:validation:Minimum",
			expectedExpressions: map[string]string{
				"": "10",
			},
		},
		{
			name:       "registered marker with negative numeric value",
			marker:     `kubebuilder:validation:Minimum=-10`,
			expectedID: "kubebuilder:validation:Minimum",
			expectedExpressions: map[string]string{
				"": "-10",
			},
		},
		{
			name:       "registered marker with named expression using backtick ('`') for strings",
			marker:     "kubebuilder:validation:XValidation:rule=`has(self.field)`,message=`must have field!`",
			expectedID: "kubebuilder:validation:XValidation",
			expectedExpressions: map[string]string{
				"rule":    "`has(self.field)`",
				"message": "`must have field!`",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			reg := NewRegistry()
			reg.Register(tc.expectedID)

			id, expressions := extractMarkerIDAndExpressions(reg, tc.marker)

			g.Expect(id).To(Equal(tc.expectedID), "marker", tc.marker)
			g.Expect(expressions).To(Equal(tc.expectedExpressions), "marker", tc.marker)
		})
	}
}

func TestExtractMarker(t *testing.T) {
	type testcase struct {
		name           string
		comment        string
		shouldBeMarker bool
		expectedID     string
	}

	testcases := []testcase{
		{
			name:           "valid marker - required",
			comment:        "// +required",
			shouldBeMarker: true,
			expectedID:     "required",
		},
		{
			name:           "valid marker - kubebuilder:validation:Required",
			comment:        "// +kubebuilder:validation:Required",
			shouldBeMarker: true,
			expectedID:     "kubebuilder:validation:Required",
		},
		{
			name:           "valid marker - with expressions",
			comment:        "// +kubebuilder:validation:XValidation:rule=\"something\",message=\"haha\"",
			shouldBeMarker: true,
			expectedID:     "kubebuilder:validation:XValidation",
		},
		{
			name:           "valid marker - with parentheses",
			comment:        "// +k8s:ifEnabled(\"foo\")=+k8s:required",
			shouldBeMarker: true,
			expectedID:     "k8s:ifEnabled(\"foo\")",
		},
		{
			name:           "valid marker - with single quotes and parentheses",
			comment:        "// +k8s:ifEnabled('foo')=+k8s:required",
			shouldBeMarker: true,
			expectedID:     "k8s:ifEnabled('foo')",
		},
		{
			name:           "valid marker - with backtickets and parentheses",
			comment:        "// +k8s:ifEnabled(`foo`)=+k8s:required",
			shouldBeMarker: true,
			expectedID:     "k8s:ifEnabled(`foo`)",
		},
		{
			name:           "invalid marker - markdown table border",
			comment:        "// +-------+-------+-------+",
			shouldBeMarker: false,
			expectedID:     "",
		},
		{
			name:           "invalid marker - markdown table border without pipes",
			comment:        "// +----------",
			shouldBeMarker: false,
			expectedID:     "",
		},
		{
			name:           "invalid marker - starts with special characters",
			comment:        "// +!*@(#&KSDJUF:A",
			shouldBeMarker: false,
			expectedID:     "",
		},
		{
			name:           "regular comment - no plus sign",
			comment:        "// This is a regular comment",
			shouldBeMarker: false,
			expectedID:     "",
		},
		{
			name:           "valid marker - complex nested expression",
			comment:        "// +k8s:someThing(one: \"a\", two: \"b\")=+k8s:required",
			shouldBeMarker: true,
			expectedID:     "k8s:someThing(one: \"a\", two: \"b\")",
		},
		{
			name:           "valid marker - complex nested expression with '",
			comment:        "// +k8s:someThing(one: 'a', two: 'b')=+k8s:required",
			shouldBeMarker: true,
			expectedID:     "k8s:someThing(one: 'a', two: 'b')",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			// Create a mock comment
			comment := &ast.Comment{
				Text: tc.comment,
			}

			marker := extractMarker(comment)

			if tc.shouldBeMarker {
				g.Expect(marker.Identifier).To(Equal(tc.expectedID), "comment", tc.comment)
			} else {
				g.Expect(marker.Identifier).To(BeEmpty(), "comment", tc.comment)
			}
		})
	}
}

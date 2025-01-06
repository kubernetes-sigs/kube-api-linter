package markers

import (
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
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			reg := NewRegistry()
			reg.Register("kubebuilder:object:root", "required", "kubebuilder:validation:XValidation")

			id, expressions := extractMarkerIDAndExpressions(reg, tc.marker)

			g.Expect(id).To(Equal(tc.expectedID), "marker", tc.marker)
			g.Expect(expressions).To(Equal(tc.expectedExpressions), "marker", tc.marker)
		})
	}
}

package markers

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestRegistryMatch(t *testing.T) {
	type testcase struct {
		name              string
		marker            string
		registeredMarkers []string
		expectedID        string
		expectedMatch     bool
	}

	testcases := []testcase{
		{
			name: "one marker registered, marker matches",
			registeredMarkers: []string{
				"kubebuilder:object:root",
			},
			marker:        "kubebuilder:object:root=true",
			expectedID:    "kubebuilder:object:root",
			expectedMatch: true,
		},
		{
			name: "multiple markers registered, matches longest registered entry",
			registeredMarkers: []string{
				"kubebuilder:validation:XValidation",
				"kubebuilder:validation",
			},
			marker:        "kubebuilder:validation:XValidation:rule='foo'",
			expectedID:    "kubebuilder:validation:XValidation",
			expectedMatch: true,
		},
		{
			name: "multiple markers registered, no matches",
			registeredMarkers: []string{
				"kubebuilder:validation:XValidation",
				"kubebuilder:validation",
			},
			marker:        "kubebuilder:notreal:foo",
			expectedID:    "",
			expectedMatch: false,
		},
		{
			name: "marker registered, exact match",
			registeredMarkers: []string{
				"optional",
			},
			marker:        "optional",
			expectedID:    "optional",
			expectedMatch: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)
			registry := NewRegistry()
			registry.Register(tc.registeredMarkers...)

			id, ok := registry.Match(tc.marker)

			g.Expect(id).To(Equal(tc.expectedID))
			g.Expect(ok).To(Equal(tc.expectedMatch))
		})
	}
}

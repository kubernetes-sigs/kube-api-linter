/*
Copyright 2026 The Kubernetes Authors.

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

package discriminatedunions

import "testing"

func TestCELRuleMatching(t *testing.T) {
	matcher, err := newCELRuleMatcher()
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		name     string
		rule     string
		optional bool
		matches  bool
	}{
		{
			name:    "required member",
			rule:    `has(self.kind) && self.kind == 'Payload' ? has(self.payload) : !has(self.payload)`,
			matches: true,
		},
		{
			name:     "optional member",
			rule:     `has(self.kind) && self.kind == 'Payload' ? true : !has(self.payload)`,
			optional: true,
			matches:  true,
		},
		{
			name:    "whitespace quotes and redundant parentheses",
			rule:    ` ( has ( self.kind ) && ( self.kind == "Payload" ) ) ? (has(self.payload)) : (!has(self.payload)) `,
			matches: true,
		},
		{
			name: "reversed branches",
			rule: `has(self.kind) && self.kind == 'Payload' ? !has(self.payload) : has(self.payload)`,
		},
		{
			name:     "reversed optional branches",
			rule:     `has(self.kind) && self.kind == 'Payload' ? !has(self.payload) : true`,
			optional: true,
		},
		{
			name: "required member allowed to be absent",
			rule: `has(self.kind) && self.kind == 'Payload' ? true : !has(self.payload)`,
		},
		{
			name:     "optional member incorrectly required",
			rule:     `has(self.kind) && self.kind == 'Payload' ? has(self.payload) : !has(self.payload)`,
			optional: true,
		},
		{
			name: "wrong discriminator",
			rule: `has(self.other) && self.other == 'Payload' ? has(self.payload) : !has(self.payload)`,
		},
		{
			name: "wrong discriminator value",
			rule: `has(self.kind) && self.kind == 'Other' ? has(self.payload) : !has(self.payload)`,
		},
		{
			name: "whitespace inside literal is significant",
			rule: `has(self.kind) && self.kind == 'Pay load' ? has(self.payload) : !has(self.payload)`,
		},
		{
			name: "wrong member",
			rule: `has(self.kind) && self.kind == 'Payload' ? has(self.other) : !has(self.other)`,
		},
		{
			name: "missing presence guard",
			rule: `self.kind == 'Payload' ? has(self.payload) : !has(self.payload)`,
		},
		{
			name: "valid pattern inside string",
			rule: `"has(self.kind) && self.kind == 'Payload' ? has(self.payload) : !has(self.payload)" == ""`,
		},
		{
			name: "valid pattern inside comment",
			rule: "true // has(self.kind) && self.kind == 'Payload' ? has(self.payload) : !has(self.payload)",
		},
		{
			name: "bypassed gating",
			rule: `(has(self.kind) && self.kind == 'Payload' ? has(self.payload) : !has(self.payload)) || true`,
		},
		{
			name: "malformed rule",
			rule: `has(self.kind) && self.kind == 'Payload' ? has(self.payload) : !has(self.payload`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			expected, ok := matcher.normalize(membershipRule("kind", "payload", "Payload", tc.optional))
			if !ok {
				t.Fatal("generated membership rule did not parse")
			}

			actual, valid := matcher.normalize(tc.rule)
			if matches := valid && actual == expected; matches != tc.matches {
				t.Errorf("matches = %v, want %v; normalized rule %q", matches, tc.matches, actual)
			}
		})
	}
}

func TestCELFieldName(t *testing.T) {
	for name, expected := range map[string]string{
		"field":      "field",
		"namespace":  "__namespace__",
		"a__b":       "a__underscores__b",
		"a.b/c-d":    "a__dot__b__slash__c__dash__d",
		"__has-dash": "__underscores__has__dash__dash",
	} {
		t.Run(name, func(t *testing.T) {
			if actual := celFieldName(name); actual != expected {
				t.Errorf("celFieldName(%q) = %q, want %q", name, actual, expected)
			}
		})
	}
}

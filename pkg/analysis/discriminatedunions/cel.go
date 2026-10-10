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

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/google/cel-go/common"
	"github.com/google/cel-go/parser"
	"golang.org/x/tools/go/analysis"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	markersconsts "sigs.k8s.io/kube-api-linter/pkg/markers"
)

type celRuleMatcher struct {
	parser *parser.Parser
}

func newCELRuleMatcher() (*celRuleMatcher, error) {
	// Keep macros unexpanded so has(...) remains part of the matched syntax.
	p, err := parser.NewParser()
	if err != nil {
		return nil, fmt.Errorf("creating CEL parser: %w", err)
	}

	return &celRuleMatcher{parser: p}, nil
}

// normalize preserves expression structure while ignoring quotes, whitespace and
// redundant parentheses. It deliberately does not prove logical equivalence.
func (m *celRuleMatcher) normalize(rule string) (string, bool) {
	expr, errors := m.parser.Parse(common.NewTextSource(rule))
	if len(errors.GetErrors()) > 0 {
		return "", false
	}

	normalized, err := parser.Unparse(expr.Expr(), expr.SourceInfo())

	return normalized, err == nil
}

func (m *celRuleMatcher) rules(markers markershelper.MarkerSet) []string {
	var rules []string

	for _, marker := range markers.Get(markersconsts.KubebuilderXValidationMarker) {
		rule, err := strconv.Unquote(marker.Arguments["rule"])
		if err != nil {
			continue
		}

		if normalized, ok := m.normalize(rule); ok {
			rules = append(rules, normalized)
		}
	}

	return rules
}

func (m *celRuleMatcher) reportViolations(pass *analysis.Pass, union *unionType, rules []string) {
	if len(union.discriminatorFields) != 1 || !union.discriminatorFields[0].required {
		return
	}

	discriminator := union.discriminatorFields[0]

	for _, member := range union.memberFields {
		if !member.optional {
			continue
		}

		value, optional := memberSelection(member, union.key)
		expected := membershipRule(discriminator.jsonName, member.jsonName, value, optional)

		normalized, ok := m.normalize(expected)
		if ok && slices.Contains(rules, normalized) {
			continue
		}

		pass.Reportf(member.field.Pos(), "union member field %s in %s requires a recognized +%s rule; expected %s",
			member.qualifiedName, describeUnion(union), markersconsts.KubebuilderXValidationMarker, expected)
	}
}

func memberSelection(member unionField, key string) (string, bool) {
	value := utils.FieldName(member.field)

	optional := false

	if key == defaultUnionKey {
		for _, marker := range member.markers.Get(markersconsts.UnionMemberMarker) {
			optional = optional || marker.Arguments[markershelper.UnnamedArgument] == "optional"
		}
	}

	for _, marker := range member.markers.Get(markersconsts.K8sUnionMemberMarker) {
		if memberName, ok := marker.Arguments["memberName"]; ok && unionKey(marker) == key {
			value = memberName
		}
	}

	return value, optional
}

func membershipRule(discriminator, member, value string, optional bool) string {
	discriminator = celFieldName(discriminator)
	member = celFieldName(member)

	selected := "has(self." + member + ")"
	if optional {
		selected = "true"
	}

	return fmt.Sprintf("has(self.%s) && self.%s == %q ? %s : !has(self.%s)", discriminator, discriminator, value, selected, member)
}

// Kubernetes escapes JSON property names before exposing them as CEL fields.
func celFieldName(name string) string {
	switch name {
	case "true", "false", "null", "in", "as", "break", "const", "continue", "else", "for", "function", "if", "import", "let", "loop", "package", "namespace", "return", "var", "void", "while":
		return "__" + name + "__"
	default:
		return strings.NewReplacer("__", "__underscores__", ".", "__dot__", "-", "__dash__", "/", "__slash__").Replace(name)
	}
}

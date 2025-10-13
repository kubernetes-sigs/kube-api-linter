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
package markerscope

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
)

const (
	name = "markerscope"
)

func init() {
	// Register all markers we want to validate scope for
	defaults := DefaultMarkerRules()
	markers := make([]string, 0, len(defaults))
	for marker := range defaults {
		markers = append(markers, marker)
	}
	markershelper.DefaultRegistry().Register(markers...)
}

type analyzer struct {
	markerRules map[string]MarkerScopeRule
	policy      MarkerScopePolicy
}

// newAnalyzer creates a new analyzer.
func newAnalyzer(cfg *MarkerScopeConfig) *analysis.Analyzer {
	if cfg == nil {
		cfg = &MarkerScopeConfig{}
	}

	a := &analyzer{
		markerRules: DefaultMarkerRules(),
		policy:      cfg.Policy,
	}

	// Override with custom rules if provided
	if cfg.MarkerRules != nil {
		for marker, rule := range cfg.MarkerRules {
			a.markerRules[marker] = rule
		}
	}

	// Set default policy if not specified
	if a.policy == "" {
		a.policy = MarkerScopePolicyWarn
	}

	return &analysis.Analyzer{
		Name:             name,
		Doc:              "Validates that markers are applied in the correct scope.",
		Run:              a.run,
		Requires:         []*analysis.Analyzer{inspect.Analyzer, markershelper.Analyzer},
		RunDespiteErrors: true,
	}
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	markersAccess, ok := pass.ResultOf[markershelper.Analyzer].(markershelper.Markers)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetMarkers
	}

	// Check field markers and type markers
	nodeFilter := []ast.Node{
		(*ast.Field)(nil),
		(*ast.GenDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch node := n.(type) {
		case *ast.Field:
			a.checkFieldMarkers(pass, node, markersAccess)
		case *ast.GenDecl:
			a.checkTypeMarkers(pass, node, markersAccess)
		}
	})

	return nil, nil
}

// reportScopeViolation reports a scope violation error
func (a *analyzer) reportScopeViolation(pass *analysis.Pass, marker markershelper.Marker, rule MarkerScopeRule) {
	var allowedScopes []string
	if rule.Scope&FieldScope != 0 {
		allowedScopes = append(allowedScopes, "fields")
	}
	if rule.Scope&TypeScope != 0 {
		allowedScopes = append(allowedScopes, "types")
	}

	scopeMsg := strings.Join(allowedScopes, " or ")
	if len(allowedScopes) == 0 {
		scopeMsg = "unknown scope"
	}

	pass.Report(analysis.Diagnostic{
		Pos:     marker.Pos,
		End:     marker.End,
		Message: fmt.Sprintf("marker %q can only be applied to %s", marker.Identifier, scopeMsg),
	})
}

// checkFieldMarkers checks markers on fields for violations
func (a *analyzer) checkFieldMarkers(pass *analysis.Pass, field *ast.Field, markersAccess markershelper.Markers) {
	fieldMarkers := markersAccess.FieldMarkers(field)

	for _, marker := range fieldMarkers.UnsortedList() {
		rule, ok := a.markerRules[marker.Identifier]
		if !ok {
			// No rule defined for this marker, skip validation
			continue
		}

		// Check if FieldScope is allowed
		if !rule.Scope.Allows(FieldScope) {
			a.reportScopeViolation(pass, marker, rule)
		}
		// TODO: Add type constraint validation here
	}
}

// checkTypeMarkers checks markers on types for violations
func (a *analyzer) checkTypeMarkers(pass *analysis.Pass, genDecl *ast.GenDecl, markersAccess markershelper.Markers) {
	if len(genDecl.Specs) == 0 {
		return
	}

	for i := range genDecl.Specs {
		typeSpec, ok := genDecl.Specs[i].(*ast.TypeSpec)
		if !ok {
			continue
		}

		typeMarkers := markersAccess.TypeMarkers(typeSpec)

		for _, marker := range typeMarkers.UnsortedList() {
			rule, ok := a.markerRules[marker.Identifier]
			if !ok {
				// No rule defined for this marker, skip validation
				continue
			}

			// Check if TypeScope is allowed
			if !rule.Scope.Allows(TypeScope) {
				a.reportScopeViolation(pass, marker, rule)
			}
			// TODO: Add type constraint validation here
		}
	}
}


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
package markertypos

import (
	"fmt"
	"go/ast"
	"regexp"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	helper_inspector "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
)

const (
	name = "markertypos"
)

// Analyzer is the analyzer for the markertypos package.
// It checks for common typos and syntax issues in marker comments.
var Analyzer = &analysis.Analyzer{
	Name:     name,
	Doc:      "Check for common typos and syntax issues in marker comments.",
	Run:      run,
	Requires: []*analysis.Analyzer{helper_inspector.Analyzer, inspect.Analyzer},
}

// Regular expressions for marker validation.
var (
	// Matches markers that start with + followed by optional space.
	markerWithSpaceRegex = regexp.MustCompile(`^\s*//\s*\+\s+\w+`)

	// Matches markers missing space after // (e.g., //+marker instead of // +marker).
	markerMissingSpaceAfterSlashRegex = regexp.MustCompile(`^\s*//\+\S+`)
)

func run(pass *analysis.Pass) (any, error) {
	helperInspect, ok := pass.ResultOf[helper_inspector.Analyzer].(helper_inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	// Regular marker analysis for well-formed markers
	helperInspect.InspectFields(func(field *ast.Field, _ []ast.Node, _ extractjsontags.FieldTagInfo, markersAccess markers.Markers) {
		checkFieldMarkers(pass, field, markersAccess)
	})

	helperInspect.InspectTypeSpec(func(typeSpec *ast.TypeSpec, markersAccess markers.Markers) {
		checkTypeSpecMarkers(pass, typeSpec, markersAccess)
	})

	// Additional analysis for malformed markers that aren't picked up by the marker parser
	astInspector, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	// Scan all comments for malformed markers
	astInspector.Preorder([]ast.Node{(*ast.GenDecl)(nil), (*ast.Field)(nil)}, func(n ast.Node) {
		switch node := n.(type) {
		case *ast.GenDecl:
			if node.Doc != nil {
				for _, comment := range node.Doc.List {
					checkMalformedMarker(pass, comment)
				}
			}
		case *ast.Field:
			if node.Doc != nil {
				for _, comment := range node.Doc.List {
					checkMalformedMarker(pass, comment)
				}
			}
		}
	})

	return nil, nil //nolint:nilnil
}

func checkFieldMarkers(pass *analysis.Pass, field *ast.Field, markersAccess markers.Markers) {
	if field == nil || len(field.Names) == 0 {
		return
	}

	fieldMarkers := markersAccess.FieldMarkers(field)
	for _, marker := range fieldMarkers.UnsortedList() {
		checkMarkerSyntax(pass, marker)
	}
}

func checkTypeSpecMarkers(pass *analysis.Pass, typeSpec *ast.TypeSpec, markersAccess markers.Markers) {
	if typeSpec == nil {
		return
	}

	typeMarkers := markersAccess.TypeMarkers(typeSpec)
	for _, marker := range typeMarkers.UnsortedList() {
		checkMarkerSyntax(pass, marker)
	}
}

func checkMalformedMarker(pass *analysis.Pass, comment *ast.Comment) {
	// Check for markers missing space after //
	if markerMissingSpaceAfterSlashRegex.MatchString(comment.Text) {
		// Create a pseudo-marker for reporting
		marker := markers.Marker{
			RawComment: comment.Text,
			Pos:        comment.Pos(),
			End:        comment.End(),
		}
		reportMissingSpaceAfterSlashIssue(pass, marker)
		// Also check for typos in malformed markers
		checkCommonTypos(pass, marker)
	}
}

func checkMarkerSyntax(pass *analysis.Pass, marker markers.Marker) {
	rawComment := marker.RawComment

	// Check for missing space after //
	if markerMissingSpaceAfterSlashRegex.MatchString(rawComment) {
		reportMissingSpaceAfterSlashIssue(pass, marker)
	}

	// Check for space after +
	if markerWithSpaceRegex.MatchString(rawComment) {
		reportSpacingIssue(pass, marker)
	}

	// Check for common typos
	checkCommonTypos(pass, marker)
}

func reportMissingSpaceAfterSlashIssue(pass *analysis.Pass, marker markers.Marker) {
	pass.Report(analysis.Diagnostic{
		Pos:     marker.Pos,
		Message: "marker should have space after '//' comment prefix",
	})
}

func reportSpacingIssue(pass *analysis.Pass, marker markers.Marker) {
	pass.Report(analysis.Diagnostic{
		Pos:     marker.Pos,
		Message: "marker should not have space after '+' symbol",
	})
}

func checkCommonTypos(pass *analysis.Pass, marker markers.Marker) {
	rawComment := marker.RawComment
	foundTypos := make(map[string]string)

	// Common marker typos.
	commonTypos := map[string]string{
		"kubebuidler": "kubebuilder",
		"kubebuiler":  "kubebuilder",
		"kubebulider": "kubebuilder",
		"kubbuilder":  "kubebuilder",
		"kubebulder":  "kubebuilder",
		"optinal":     "optional",
		"requied":     "required",
		"requird":     "required",
		"nullabel":    "nullable",
		"validaton":   "validation",
		"valdiation":  "validation",
		"defualt":     "default", //nolint:misspell
		"defult":      "default",
		"exampl":      "example",
		"examle":      "example",
	}

	// Collect all typos found in this marker
	for typo, correction := range commonTypos {
		typoRegex := regexp.MustCompile(`\b` + regexp.QuoteMeta(typo) + `\b`)
		if typoRegex.MatchString(rawComment) {
			foundTypos[typo] = correction
		}
	}

	// Report each typo separately (but with combined fix if multiple typos exist)
	if len(foundTypos) > 0 {
		reportTypos(pass, marker, foundTypos)
	}
}

func reportTypos(pass *analysis.Pass, marker markers.Marker, foundTypos map[string]string) {
	// Report each typo as a separate diagnostic
	for typo, correction := range foundTypos {
		pass.Report(analysis.Diagnostic{
			Pos:     marker.Pos,
			Message: fmt.Sprintf("possible typo: '%s' should be '%s'", typo, correction),
		})
	}
}

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
package jsontags

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"

	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	markersconsts "sigs.k8s.io/kube-api-linter/pkg/markers"

	"golang.org/x/tools/go/analysis"
	astinspector "golang.org/x/tools/go/ast/inspector"
)

const (
	// camelCaseRegex is a regular expression that matches camel case strings.
	camelCaseRegex = "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)*$"

	name = "jsontags"
)

type analyzer struct {
	jsonTagRegex *regexp.Regexp
}

// newAnalyzer creates a new analyzer with the given json tag regex.
func newAnalyzer(cfg *JSONTagsConfig) (*analysis.Analyzer, error) {
	if cfg == nil {
		cfg = &JSONTagsConfig{}
	}

	defaultConfig(cfg)

	jsonTagRegex, err := regexp.Compile(cfg.JSONTagRegex)
	if err != nil {
		return nil, fmt.Errorf("could not compile json tag regex: %w", err)
	}

	a := &analyzer{
		jsonTagRegex: jsonTagRegex,
	}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      "Check that all struct fields in an API are tagged with json tags",
		Run:      a.run,
		Requires: []*analysis.Analyzer{extractjsontags.Analyzer, markers.Analyzer},
	}, nil
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	astInspector := astinspector.New(pass.Files)

	jsonTags, ok := pass.ResultOf[extractjsontags.Analyzer].(extractjsontags.StructFieldTags)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetJSONTags
	}

	markersAccess, ok := pass.ResultOf[markers.Analyzer].(markers.Markers)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetMarkers
	}

	// Custom field iteration that includes list types
	a.inspectFieldsIncludingListTypes(astInspector, jsonTags, markersAccess, pass)

	return nil, nil //nolint:nilnil
}

// inspectFieldsIncludingListTypes iterates over fields in structs, including list types.
// This is a custom implementation that bypasses the inspector's InspectFields helper
// to ensure list types are also linted for proper JSON tags.
func (a *analyzer) inspectFieldsIncludingListTypes(
	inspector *astinspector.Inspector,
	jsonTags extractjsontags.StructFieldTags,
	markersAccess markers.Markers,
	pass *analysis.Pass,
) {
	nodeFilter := []ast.Node{
		(*ast.Field)(nil),
	}

	inspector.WithStack(nodeFilter, func(n ast.Node, push bool, stack []ast.Node) (proceed bool) {
		if !push {
			return false
		}

		field, ok := n.(*ast.Field)
		if !ok || !a.shouldProcessField(stack) {
			return ok
		}

		if a.shouldSkipField(field, jsonTags, markersAccess) {
			return false
		}

		tagInfo := jsonTags.FieldTags(field)
		a.checkField(pass, field, tagInfo)

		return true
	})
}

// shouldProcessField checks if the field should be processed.
// This is similar to the inspector's version but does NOT skip list types.
func (a *analyzer) shouldProcessField(stack []ast.Node) bool {
	if len(stack) < 3 {
		return false
	}

	// The 0th node in the stack is the *ast.File.
	// The 1st node in the stack is the *ast.GenDecl.
	decl, ok := stack[1].(*ast.GenDecl)
	if !ok || decl.Tok != token.TYPE {
		// Make sure that we don't inspect structs within a function or non-type declarations.
		return false
	}

	_, ok = stack[len(stack)-3].(*ast.StructType)
	if !ok {
		// Not in a struct.
		return false
	}

	return true
}

// shouldSkipField checks if a field should be skipped.
func (a *analyzer) shouldSkipField(field *ast.Field, jsonTags extractjsontags.StructFieldTags, markersAccess markers.Markers) bool {
	tagInfo := jsonTags.FieldTags(field)
	if tagInfo.Ignored {
		return true
	}

	markerSet := markersAccess.FieldMarkers(field)

	return isSchemalessType(markerSet)
}

func isSchemalessType(markerSet markers.MarkerSet) bool {
	// Check if the field is marked as schemaless.
	schemalessMarker := markerSet.Get(markersconsts.KubebuilderSchemaLessMarker)
	return len(schemalessMarker) > 0
}

func (a *analyzer) checkField(pass *analysis.Pass, field *ast.Field, tagInfo extractjsontags.FieldTagInfo) {
	prefix := "field %s"
	if len(field.Names) == 0 || field.Names[0] == nil {
		prefix = "embedded field %s"
	}

	prefix = fmt.Sprintf(prefix, utils.FieldName(field))

	if tagInfo.Missing {
		pass.Reportf(field.Pos(), "%s is missing json tag", prefix)
		return
	}

	if tagInfo.Inline {
		return
	}

	if tagInfo.Name == "" {
		pass.Reportf(field.Pos(), "%s has empty json tag", prefix)
		return
	}

	matched := a.jsonTagRegex.Match([]byte(tagInfo.Name))
	if !matched {
		pass.Reportf(field.Pos(), "%s json tag does not match pattern %q: %s", prefix, a.jsonTagRegex.String(), tagInfo.Name)
	}
}

func defaultConfig(cfg *JSONTagsConfig) {
	if cfg.JSONTagRegex == "" {
		cfg.JSONTagRegex = camelCaseRegex
	}
}

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

package discriminatedunions

import (
	"go/ast"
	"slices"
	"strings"

	"golang.org/x/tools/go/analysis"

	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	inspectorhelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	markersconsts "sigs.k8s.io/kube-api-linter/pkg/markers"
)

const name = "discriminatedunions"

func init() {
	markershelper.DefaultRegistry().Register(
		markersconsts.UnionMarker,
		markersconsts.UnionDiscriminatorMarker,
		markersconsts.UnionMemberMarker,
		markersconsts.K8sUnionDiscriminatorMarker,
		markersconsts.K8sUnionMemberMarker,
		markersconsts.OptionalMarker,
		markersconsts.KubebuilderOptionalMarker,
		markersconsts.K8sOptionalMarker,
		markersconsts.RequiredMarker,
		markersconsts.KubebuilderRequiredMarker,
		markersconsts.K8sRequiredMarker,
	)
}

type analyzer struct {
	nonMemberFields NonMemberFieldsPolicy
}

type unionType struct {
	typeSpec *ast.TypeSpec
	name     string

	hasUnionMarker bool

	discriminatorFields []*unionField
	memberFields        []*unionField
	nonMemberFields     []*unionField
}

type unionField struct {
	field         *ast.Field
	qualifiedName string

	required bool
	optional bool

	isDiscriminator      bool
	isMember             bool
	memberOptionalMarker bool
}

type unionFieldClassification struct {
	isDiscriminator  bool
	isMember         bool
	isMemberOptional bool
}

func newAnalyzer(cfg *Config) *analysis.Analyzer {
	if cfg == nil {
		cfg = &Config{}
	}

	defaultConfig(cfg)

	a := &analyzer{
		nonMemberFields: cfg.NonMemberFields,
	}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      "Validates discriminated-union marker structure.",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspectorhelper.Analyzer},
	}
}

func defaultConfig(cfg *Config) {
	if cfg.NonMemberFields == "" {
		cfg.NonMemberFields = NonMemberFieldsForbid
	}
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspectorhelper.Analyzer].(inspectorhelper.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	fieldInfos := make(map[*ast.Field]*unionField)

	inspect.InspectFields(func(field *ast.Field, _ extractjsontags.FieldTagInfo, markersAccess markershelper.Markers, qualifiedFieldName string) {
		fieldInfos[field] = buildUnionFieldInfo(field, markersAccess, qualifiedFieldName)
	})

	inspect.InspectTypeSpec(func(typeSpec *ast.TypeSpec, markersAccess markershelper.Markers) {
		union := buildUnionType(typeSpec, markersAccess, fieldInfos)
		if union == nil {
			return
		}

		a.reportStructureViolations(pass, union)
	})

	return nil, nil //nolint:nilnil
}

func buildUnionType(typeSpec *ast.TypeSpec, markersAccess markershelper.Markers, fieldInfos map[*ast.Field]*unionField) *unionType {
	if typeSpec == nil || typeSpec.Name == nil {
		return nil
	}

	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok || structType.Fields == nil {
		return nil
	}

	structMarkers := markersAccess.StructMarkers(structType)

	union := &unionType{
		typeSpec:       typeSpec,
		name:           typeSpec.Name.Name,
		hasUnionMarker: structMarkers.Has(markersconsts.UnionMarker),
	}

	for _, field := range structType.Fields.List {
		unionFieldInfo, ok := fieldInfos[field]
		if !ok {
			continue
		}

		addUnionField(union, unionFieldInfo)
	}

	if !union.hasUnionMarker && len(union.discriminatorFields) == 0 && len(union.memberFields) == 0 {
		return nil
	}

	return union
}

func (a *analyzer) reportStructureViolations(pass *analysis.Pass, union *unionType) {
	if union == nil {
		return
	}

	reportDiscriminatorViolations(pass, union)
	reportMissingMemberViolations(pass, union)
	reportMemberOptionalityViolations(pass, union)
	a.reportNonMemberFieldViolations(pass, union)
}

func buildUnionFieldInfo(field *ast.Field, markersAccess markershelper.Markers, qualifiedFieldName string) *unionField {
	if field == nil {
		return nil
	}

	fieldMarkers := markersAccess.FieldMarkers(field)
	classification := classifyUnionField(fieldMarkers)

	return &unionField{
		field:                field,
		qualifiedName:        qualifiedFieldName,
		required:             utils.IsFieldRequired(field, markersAccess),
		optional:             utils.IsFieldOptional(field, markersAccess),
		isDiscriminator:      classification.isDiscriminator,
		isMember:             classification.isMember,
		memberOptionalMarker: classification.isMemberOptional,
	}
}

func addUnionField(union *unionType, field *unionField) {
	if union == nil || field == nil {
		return
	}

	if field.isDiscriminator {
		union.discriminatorFields = append(union.discriminatorFields, field)
	}

	if field.isMember {
		union.memberFields = append(union.memberFields, field)
	}

	if !field.isDiscriminator && !field.isMember {
		union.nonMemberFields = append(union.nonMemberFields, field)
	}
}

func reportDiscriminatorViolations(pass *analysis.Pass, union *unionType) {
	switch len(union.discriminatorFields) {
	case 0:
		pass.Reportf(
			union.typeSpec.Pos(),
			"type %s is marked as a discriminated union but has no discriminator field; expected exactly one field with +%s or +%s",
			union.name,
			markersconsts.UnionDiscriminatorMarker,
			markersconsts.K8sUnionDiscriminatorMarker,
		)
	case 1:
		discriminator := union.discriminatorFields[0]
		if !discriminator.required {
			pass.Reportf(discriminator.field.Pos(), "discriminator field %s must be marked as required", discriminator.qualifiedName)
		}
	default:
		discriminatorNames := make([]string, 0, len(union.discriminatorFields))
		for _, field := range union.discriminatorFields {
			discriminatorNames = append(discriminatorNames, field.qualifiedName)
		}

		pass.Reportf(
			union.typeSpec.Pos(),
			"type %s is marked as a discriminated union but has %d discriminator fields; expected exactly one: %s",
			union.name,
			len(union.discriminatorFields),
			strings.Join(discriminatorNames, ", "),
		)
	}
}

func reportMissingMemberViolations(pass *analysis.Pass, union *unionType) {
	if len(union.memberFields) == 0 {
		pass.Reportf(union.typeSpec.Pos(), "type %s is marked as a discriminated union but has no union member fields", union.name)
	}
}

func reportMemberOptionalityViolations(pass *analysis.Pass, union *unionType) {
	for _, member := range union.memberFields {
		if member.optional || member.memberOptionalMarker {
			continue
		}

		pass.Reportf(
			member.field.Pos(),
			"union member field %s must be marked as optional (use +optional/+k8s:optional or +%s,optional)",
			member.qualifiedName,
			markersconsts.UnionMemberMarker,
		)
	}
}

func (a *analyzer) reportNonMemberFieldViolations(pass *analysis.Pass, union *unionType) {
	if a.nonMemberFields != NonMemberFieldsForbid {
		return
	}

	for _, field := range union.nonMemberFields {
		pass.Reportf(
			field.field.Pos(),
			"field %s is not a union discriminator/member in union type %s (non-member fields are forbidden)",
			field.qualifiedName,
			union.name,
		)
	}
}

func classifyUnionField(fieldMarkers markershelper.MarkerSet) unionFieldClassification {
	classification := unionFieldClassification{
		isDiscriminator: fieldMarkers.Has(markersconsts.UnionDiscriminatorMarker) ||
			fieldMarkers.Has(markersconsts.K8sUnionDiscriminatorMarker),
		isMember: fieldMarkers.Has(markersconsts.UnionMemberMarker) || fieldMarkers.Has(markersconsts.K8sUnionMemberMarker),
	}

	if !classification.isMember {
		return classification
	}

	classification.isMemberOptional = slices.ContainsFunc(fieldMarkers.Get(markersconsts.UnionMemberMarker), markerSpecifiesOptionalMember) ||
		slices.ContainsFunc(fieldMarkers.Get(markersconsts.K8sUnionMemberMarker), markerSpecifiesOptionalMember)

	return classification
}

func markerSpecifiesOptionalMember(marker markershelper.Marker) bool {
	value, ok := marker.Arguments[markershelper.UnnamedArgument]
	if !ok {
		return false
	}

	return strings.TrimSpace(strings.Trim(value, `"'`)) == markersconsts.OptionalMarker
}

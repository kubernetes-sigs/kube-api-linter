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
	"fmt"
	"go/ast"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"

	kalerrors "sigs.k8s.io/kube-api-linter/pkg/analysis/errors"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	inspectorhelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/inspector"
	markershelper "sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
	markersconsts "sigs.k8s.io/kube-api-linter/pkg/markers"
)

const (
	name = "discriminatedunions"

	defaultUnionKey = ""
	unionArgument   = "union"
)

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
		markersconsts.KubebuilderRootMarker,
		markersconsts.KubebuilderXValidationMarker,
	)
}

type analyzer struct {
	nonMemberFields         NonMemberFieldsPolicy
	preferredOptionalMarker string
	enforceCEL              bool
}

type unionType struct {
	typeSpec *ast.TypeSpec
	name     string
	key      string

	hasUnionMarker bool

	discriminatorFields []unionField
	memberFields        []unionField
	nonMemberFields     []unionField
}

type unionField struct {
	field         *ast.Field
	qualifiedName string
	jsonName      string
	markers       markershelper.MarkerSet

	required bool
	optional bool

	discriminatorKeys []string
	memberKeys        []string
}

type unionFieldClassification struct {
	discriminatorKeys []string
	memberKeys        []string
}

func newAnalyzer(cfg *Config) *analysis.Analyzer {
	if cfg == nil {
		cfg = &Config{}
	}

	defaultConfig(cfg)

	a := &analyzer{
		nonMemberFields:         cfg.NonMemberFields,
		preferredOptionalMarker: cfg.PreferredOptionalMarker,
		enforceCEL:              cfg.EnforceCEL,
	}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      "Validates discriminator and member structure, with configurable non-member field policy and CEL membership checks for CRD types.",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspectorhelper.Analyzer},
	}
}

func defaultConfig(cfg *Config) {
	if cfg.NonMemberFields == "" {
		cfg.NonMemberFields = NonMemberFieldsForbid
	}

	if cfg.PreferredOptionalMarker == "" {
		cfg.PreferredOptionalMarker = markersconsts.OptionalMarker
	}
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspectorhelper.Analyzer].(inspectorhelper.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	fieldInfos := make(map[*ast.Field]unionField)

	for f := range inspect.FieldsIncludingListTypes() {
		fieldInfos[f.Field] = buildUnionFieldInfo(f.Field, f.JSONTagInfo, f.Markers, f.QualifiedFieldName)
	}

	var (
		reachable map[*ast.StructType]bool
		matcher   *celRuleMatcher
	)

	if a.enforceCEL {
		var err error

		matcher, err = newCELRuleMatcher()
		if err != nil {
			return nil, fmt.Errorf("creating CEL rule matcher: %w", err)
		}

		reachable = crdReachableStructs(pass, inspect, fieldInfos)
	}

	for ts := range inspect.TypeSpecs() {
		unions := buildUnionTypes(ts.TypeSpec, ts.Markers, fieldInfos)
		if len(unions) == 0 {
			continue
		}

		structType, _ := ts.TypeSpec.Type.(*ast.StructType)

		var rules []string

		if reachable[structType] {
			rules = matcher.rules(ts.Markers.StructMarkers(structType))
		}

		for _, union := range unions {
			a.reportStructureViolations(pass, &union)

			if reachable[structType] {
				matcher.reportViolations(pass, &union, rules)
			}
		}
	}

	return nil, nil //nolint:nilnil
}

func buildUnionTypes(typeSpec *ast.TypeSpec, markersAccess markershelper.Markers, fieldInfos map[*ast.Field]unionField) []unionType {
	if typeSpec == nil || typeSpec.Name == nil {
		return nil
	}

	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok || structType.Fields == nil {
		return nil
	}

	structMarkers := markersAccess.StructMarkers(structType)
	hasUnionMarker := structMarkers.Has(markersconsts.UnionMarker)
	unionsByKey, getUnion := newUnionTypeCollector(typeSpec, hasUnionMarker)
	nonMemberFields := collectUnionFields(structType, fieldInfos, getUnion)

	if !hasUnionMarker && len(unionsByKey) == 0 {
		return nil
	}

	if len(unionsByKey) == 0 {
		getUnion(defaultUnionKey)
	}

	return sortedUnionTypes(unionsByKey, nonMemberFields)
}

func newUnionTypeCollector(typeSpec *ast.TypeSpec, hasUnionMarker bool) (map[string]*unionType, func(string) *unionType) {
	unionsByKey := map[string]*unionType{}

	return unionsByKey, func(key string) *unionType {
		if union, ok := unionsByKey[key]; ok {
			return union
		}

		union := &unionType{
			typeSpec:       typeSpec,
			name:           typeSpec.Name.Name,
			key:            key,
			hasUnionMarker: hasUnionMarker,
		}

		unionsByKey[key] = union

		return union
	}
}

func collectUnionFields(structType *ast.StructType, fieldInfos map[*ast.Field]unionField, getUnion func(string) *unionType) []unionField {
	nonMemberFields := []unionField{}

	for _, field := range structType.Fields.List {
		unionFieldInfo, ok := fieldInfos[field]
		if !ok {
			continue
		}

		addUnionField(getUnion, unionFieldInfo)

		if !unionFieldInfo.hasUnionRole() {
			nonMemberFields = append(nonMemberFields, unionFieldInfo)
		}
	}

	return nonMemberFields
}

func sortedUnionTypes(unionsByKey map[string]*unionType, nonMemberFields []unionField) []unionType {
	unionKeys := make([]string, 0, len(unionsByKey))
	for key := range unionsByKey {
		unionKeys = append(unionKeys, key)
	}

	slices.Sort(unionKeys)

	unions := make([]unionType, 0, len(unionKeys))

	for _, key := range unionKeys {
		union := *unionsByKey[key]
		// Declarative member markers also describe valid undiscriminated unions.
		if !union.hasUnionMarker && len(union.discriminatorFields) == 0 {
			continue
		}

		if len(unions) == 0 {
			union.nonMemberFields = nonMemberFields
		}

		unions = append(unions, union)
	}

	return unions
}

func (a *analyzer) reportStructureViolations(pass *analysis.Pass, union *unionType) {
	if union == nil {
		return
	}

	reportDiscriminatorViolations(pass, union)
	a.reportMemberOptionalityViolations(pass, union)
	a.reportNonMemberFieldViolations(pass, union)
}

func buildUnionFieldInfo(field *ast.Field, tag extractjsontags.FieldTagInfo, markersAccess markershelper.Markers, qualifiedFieldName string) unionField {
	if field == nil {
		return unionField{}
	}

	fieldMarkers := markersAccess.FieldMarkers(field)
	classification := classifyUnionField(fieldMarkers)

	jsonName := tag.Name
	if jsonName == "" && !tag.Inline {
		jsonName = utils.FieldName(field)
	}

	return unionField{
		field:             field,
		qualifiedName:     qualifiedFieldName,
		jsonName:          jsonName,
		markers:           fieldMarkers,
		required:          utils.IsFieldRequired(field, markersAccess),
		optional:          utils.IsFieldOptional(field, markersAccess),
		discriminatorKeys: classification.discriminatorKeys,
		memberKeys:        classification.memberKeys,
	}
}

func addUnionField(getUnion func(string) *unionType, field unionField) {
	if field.field == nil {
		return
	}

	for _, key := range field.discriminatorKeys {
		union := getUnion(key)
		union.discriminatorFields = append(union.discriminatorFields, field)
	}

	for _, key := range field.memberKeys {
		union := getUnion(key)
		union.memberFields = append(union.memberFields, field)

		if key == defaultUnionKey && field.markers.Has(markersconsts.UnionMemberMarker) {
			union.hasUnionMarker = true
		}
	}
}

func reportDiscriminatorViolations(pass *analysis.Pass, union *unionType) {
	switch len(union.discriminatorFields) {
	case 0:
		pass.Reportf(
			union.typeSpec.Pos(),
			"%s is marked as a discriminated union but has no discriminator field; expected exactly one field with +%s or +%s",
			describeUnion(union),
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
			"%s is marked as a discriminated union but has %d discriminator fields; expected exactly one: %s",
			describeUnion(union),
			len(union.discriminatorFields),
			strings.Join(discriminatorNames, ", "),
		)
	}
}

func (a *analyzer) reportMemberOptionalityViolations(pass *analysis.Pass, union *unionType) {
	for _, member := range union.memberFields {
		if member.optional {
			continue
		}

		pass.Reportf(
			member.field.Pos(),
			"union member field %s must be marked as optional (use +%s)",
			member.qualifiedName,
			a.preferredOptionalMarker,
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
	classification := unionFieldClassification{}

	if fieldMarkers.Has(markersconsts.UnionDiscriminatorMarker) {
		classification.discriminatorKeys = append(classification.discriminatorKeys, defaultUnionKey)
	}

	for _, marker := range fieldMarkers.Get(markersconsts.K8sUnionDiscriminatorMarker) {
		classification.discriminatorKeys = appendUnionKey(classification.discriminatorKeys, unionKey(marker))
	}

	if fieldMarkers.Has(markersconsts.UnionMemberMarker) {
		classification.memberKeys = append(classification.memberKeys, defaultUnionKey)
	}

	for _, marker := range fieldMarkers.Get(markersconsts.K8sUnionMemberMarker) {
		classification.memberKeys = appendUnionKey(classification.memberKeys, unionKey(marker))
	}

	return classification
}

func (f unionField) hasUnionRole() bool {
	return len(f.discriminatorKeys) > 0 || len(f.memberKeys) > 0
}

func unionKey(marker markershelper.Marker) string {
	return marker.Arguments[unionArgument]
}

func appendUnionKey(keys []string, key string) []string {
	if slices.Contains(keys, key) {
		return keys
	}

	return append(keys, key)
}

func describeUnion(union *unionType) string {
	if union.key == defaultUnionKey {
		return "type " + union.name
	}

	return "union " + strconv.Quote(union.key) + " in type " + union.name
}

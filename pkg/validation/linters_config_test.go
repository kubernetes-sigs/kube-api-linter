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
package validation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"

	"sigs.k8s.io/kube-api-linter/pkg/analysis/conditions"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/jsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/nomaps"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/optionalfields"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/optionalorrequired"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/requiredfields"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/ssatags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/statusoptional"
	"sigs.k8s.io/kube-api-linter/pkg/config"
	"sigs.k8s.io/kube-api-linter/pkg/markers"
	"sigs.k8s.io/kube-api-linter/pkg/validation"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

var _ = Describe("LintersConfig", func() {
	type validateLintersConfigTableInput struct {
		linters     config.Linters
		config      config.LintersConfig
		expectedErr string
	}

	DescribeTable("Validate Linters Configuration", func(in validateLintersConfigTableInput) {
		errs := validation.ValidateLintersConfig(in.linters, in.config, field.NewPath("lintersConfig"))
		if len(in.expectedErr) > 0 {
			Expect(errs.ToAggregate()).To(MatchError(in.expectedErr))
		} else {
			Expect(errs).To(HaveLen(0), "No errors were expected")
		}

	},
		Entry("Empty config", validateLintersConfigTableInput{
			config:      config.LintersConfig{},
			expectedErr: "",
		}),

		// ConditionsConfig validation
		Entry("With a valid ConditionsConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"conditions": toYaml(conditions.ConditionsConfig{
					IsFirstField: "",
					UseProtobuf:  "",
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig IsFirstField: Warn", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"conditions": toYaml(conditions.ConditionsConfig{
					IsFirstField: conditions.ConditionsFirstFieldWarn,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig IsFirstField: Ignore", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"conditions": toYaml(conditions.ConditionsConfig{
					IsFirstField: conditions.ConditionsFirstFieldIgnore,
				}),
			},
			expectedErr: "",
		}),
		Entry("With an invalid ConditionsConfig IsFirstField", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"conditions": toYaml(conditions.ConditionsConfig{
					IsFirstField: "invalid",
				}),
			},
			expectedErr: "lintersConfig.conditions.isFirstField: Invalid value: \"invalid\": invalid value, must be one of \"Warn\", \"Ignore\" or omitted",
		}),
		Entry("With a valid ConditionsConfig UseProtobuf: SuggestFix", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"conditions": toYaml(conditions.ConditionsConfig{
					UseProtobuf: conditions.ConditionsUseProtobufSuggestFix,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig UseProtobuf: Warn", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"conditions": toYaml(conditions.ConditionsConfig{
					UseProtobuf: conditions.ConditionsUseProtobufWarn,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig UseProtobuf: Ignore", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"conditions": toYaml(conditions.ConditionsConfig{
					UseProtobuf: conditions.ConditionsUseProtobufIgnore,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig UseProtobuf: Forbid", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"conditions": toYaml(conditions.ConditionsConfig{
					UseProtobuf: conditions.ConditionsUseProtobufForbid,
				}),
			},
			expectedErr: "",
		}),
		Entry("With an invalid ConditionsConfig UseProtobuf", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"conditions": toYaml(conditions.ConditionsConfig{
					UseProtobuf: "invalid",
				}),
			},
			expectedErr: "lintersConfig.conditions.useProtobuf: Invalid value: \"invalid\": invalid value, must be one of \"SuggestFix\", \"Warn\", \"Ignore\", \"Forbid\" or omitted",
		}),
		Entry("With a valid ConditionsConfig UsePatchStrategy: SuggestFix", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"conditions": toYaml(conditions.ConditionsConfig{
					UsePatchStrategy: conditions.ConditionsUsePatchStrategySuggestFix,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig UsePatchStrategy: Warn", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"conditions": toYaml(conditions.ConditionsConfig{
					UsePatchStrategy: conditions.ConditionsUsePatchStrategyWarn,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig UsePatchStrategy: Ignore", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"conditions": toYaml(conditions.ConditionsConfig{
					UsePatchStrategy: conditions.ConditionsUsePatchStrategyIgnore,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid ConditionsConfig UsePatchStrategy: Forbid", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"conditions": toYaml(conditions.ConditionsConfig{
					UsePatchStrategy: conditions.ConditionsUsePatchStrategyForbid,
				}),
			},
			expectedErr: "",
		}),
		Entry("With an invalid ConditionsConfig UsePatchStrategy", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"conditions": toYaml(conditions.ConditionsConfig{
					UsePatchStrategy: "invalid",
				}),
			},
			expectedErr: "lintersConfig.conditions.usePatchStrategy: Invalid value: \"invalid\": invalid value, must be one of \"SuggestFix\", \"Warn\", \"Ignore\", \"Forbid\" or omitted",
		}),

		// JSONTagsConfig validation (with legacy field name)
		Entry("With a valid JSONTagsConfig JSONTagRegex", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"jsonTags": toYaml(jsontags.JSONTagsConfig{
					JSONTagRegex: "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)*$",
				}),
			},
			expectedErr: "",
		}),
		Entry("With an invalid JSONTagsConfig JSONTagRegex", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"jsonTags": toYaml(jsontags.JSONTagsConfig{
					JSONTagRegex: "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*",
				}),
			},
			expectedErr: "lintersConfig.jsontags.jsonTagRegex: Invalid value: \"^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*\": invalid regex: error parsing regexp: missing closing ): `^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*`",
		}),

		// JSONTagsConfig validation
		Entry("With a valid JSONTagsConfig JSONTagRegex", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"jsontags": toYaml(jsontags.JSONTagsConfig{
					JSONTagRegex: "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)*$",
				}),
			},
			expectedErr: "",
		}),
		Entry("With an invalid JSONTagsConfig JSONTagRegex", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"jsontags": toYaml(jsontags.JSONTagsConfig{
					JSONTagRegex: "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*",
				}),
			},
			expectedErr: "lintersConfig.jsontags.jsonTagRegex: Invalid value: \"^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*\": invalid regex: error parsing regexp: missing closing ): `^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*`",
		}),

		// NoMapsConfig validation
		Entry("With a valid NoMapsConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"nomaps": toYaml(nomaps.NoMapsConfig{
					Policy: "",
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid NoMapsConfig: enforce is specified", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"nomaps": toYaml(nomaps.NoMapsConfig{
					Policy: nomaps.NoMapsEnforce,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid NoMapsConfig: allowStringToStringMaps is specified", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"nomaps": toYaml(nomaps.NoMapsConfig{
					Policy: nomaps.NoMapsAllowStringToStringMaps,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid NoMapsConfig: ignore is specified", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"nomaps": toYaml(nomaps.NoMapsConfig{
					Policy: nomaps.NoMapsIgnore,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a invalid NoMapsConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"nomaps": toYaml(nomaps.NoMapsConfig{
					Policy: "invalid",
				}),
			},
			expectedErr: `lintersConfig.nomaps.policy: Invalid value: "invalid": invalid value, must be one of "Enforce", "AllowStringToStringMaps", "Ignore" or omitted`,
		}),

		// OptionalFieldsConfig validation (with legacy field name)
		Entry("With a valid OptionalFieldsConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalFields": toYaml(optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Preference: "",
						Policy:     "",
					},
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: "",
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid OptionalFieldsConfig: Pointer Preference Always", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalFields": toYaml(optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Preference: optionalfields.OptionalFieldsPointerPreferenceAlways,
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid OptionalFieldsConfig: Pointer Preference WhenRequired", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalFields": toYaml(optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Preference: optionalfields.OptionalFieldsPointerPreferenceWhenRequired,
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With an invalid OptionalFieldsConfig: Pointer Preference", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalFields": toYaml(optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Preference: "invalid",
					},
				}),
			},
			expectedErr: "lintersConfig.optionalfields.pointers.preference: Invalid value: \"invalid\": invalid value, must be one of \"Always\", \"WhenRequired\" or omitted",
		}),
		Entry("With a valid OptionalFieldsConfig: Pointer Policy SuggestFix", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalFields": toYaml(optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Policy: optionalfields.OptionalFieldsPointerPolicySuggestFix,
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid OptionalFieldsConfig: Pointer Policy Warn", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalFields": toYaml(optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Policy: optionalfields.OptionalFieldsPointerPolicyWarn,
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With an invalid OptionalFieldsConfig: Pointer Policy", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalFields": toYaml(optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Policy: "invalid",
					},
				}),
			},
			expectedErr: "lintersConfig.optionalfields.pointers.policy: Invalid value: \"invalid\": invalid value, must be one of \"SuggestFix\", \"Warn\" or omitted",
		}),
		Entry("With a valid OptionalFieldsConfig: OmitEmpty Policy Ignore", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalFields": toYaml(optionalfields.OptionalFieldsConfig{
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: optionalfields.OptionalFieldsOmitEmptyPolicyIgnore,
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid OptionalFieldsConfig: OmitEmpty Policy Warn", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalFields": toYaml(optionalfields.OptionalFieldsConfig{
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: optionalfields.OptionalFieldsOmitEmptyPolicyWarn,
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid OptionalFieldsConfig: OmitEmpty Policy SuggestFix", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalFields": toYaml(optionalfields.OptionalFieldsConfig{
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: optionalfields.OptionalFieldsOmitEmptyPolicySuggestFix,
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With an invalid OptionalFieldsConfig: OmitEmpty Policy", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalFields": toYaml(optionalfields.OptionalFieldsConfig{
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: "invalid",
					},
				}),
			},
			expectedErr: "lintersConfig.optionalfields.omitEmpty.policy: Invalid value: \"invalid\": invalid value, must be one of \"Ignore\", \"Warn\", \"SuggestFix\" or omitted",
		}),

		// OptionalFieldsConfig validation
		Entry("With a valid OptionalFieldsConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalfields": toYaml(optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Preference: "",
						Policy:     "",
					},
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: "",
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid OptionalFieldsConfig: Pointer Preference Always", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalfields": toYaml(optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Preference: optionalfields.OptionalFieldsPointerPreferenceAlways,
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid OptionalFieldsConfig: Pointer Preference WhenRequired", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalfields": toYaml(optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Preference: optionalfields.OptionalFieldsPointerPreferenceWhenRequired,
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With an invalid OptionalFieldsConfig: Pointer Preference", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalfields": toYaml(optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Preference: "invalid",
					},
				}),
			},
			expectedErr: "lintersConfig.optionalfields.pointers.preference: Invalid value: \"invalid\": invalid value, must be one of \"Always\", \"WhenRequired\" or omitted",
		}),
		Entry("With a valid OptionalFieldsConfig: Pointer Policy SuggestFix", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalfields": toYaml(optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Policy: optionalfields.OptionalFieldsPointerPolicySuggestFix,
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid OptionalFieldsConfig: Pointer Policy Warn", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalfields": toYaml(optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Policy: optionalfields.OptionalFieldsPointerPolicyWarn,
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With an invalid OptionalFieldsConfig: Pointer Policy", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalfields": toYaml(optionalfields.OptionalFieldsConfig{
					Pointers: optionalfields.OptionalFieldsPointers{
						Policy: "invalid",
					},
				}),
			},
			expectedErr: "lintersConfig.optionalfields.pointers.policy: Invalid value: \"invalid\": invalid value, must be one of \"SuggestFix\", \"Warn\" or omitted",
		}),
		Entry("With a valid OptionalFieldsConfig: OmitEmpty Policy Ignore", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalfields": toYaml(optionalfields.OptionalFieldsConfig{
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: optionalfields.OptionalFieldsOmitEmptyPolicyIgnore,
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid OptionalFieldsConfig: OmitEmpty Policy Warn", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalfields": toYaml(optionalfields.OptionalFieldsConfig{
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: optionalfields.OptionalFieldsOmitEmptyPolicyWarn,
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid OptionalFieldsConfig: OmitEmpty Policy SuggestFix", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalfields": toYaml(optionalfields.OptionalFieldsConfig{
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: optionalfields.OptionalFieldsOmitEmptyPolicySuggestFix,
					},
				}),
			},
			expectedErr: "",
		}),
		Entry("With an invalid OptionalFieldsConfig: OmitEmpty Policy", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalfields": toYaml(optionalfields.OptionalFieldsConfig{
					OmitEmpty: optionalfields.OptionalFieldsOmitEmpty{
						Policy: "invalid",
					},
				}),
			},
			expectedErr: "lintersConfig.optionalfields.omitEmpty.policy: Invalid value: \"invalid\": invalid value, must be one of \"Ignore\", \"Warn\", \"SuggestFix\" or omitted",
		}),

		// OptionalOrRequiredConfig validation (with legacy field name)
		Entry("With a valid OptionalOrRequiredConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalOrRequired": toYaml(optionalorrequired.OptionalOrRequiredConfig{
					PreferredOptionalMarker: markers.OptionalMarker,
					PreferredRequiredMarker: markers.RequiredMarker,
				}),
			},
			expectedErr: "",
		}),
		Entry("With kubebuilder preferred markers", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalOrRequired": toYaml(optionalorrequired.OptionalOrRequiredConfig{
					PreferredOptionalMarker: markers.KubebuilderOptionalMarker,
					PreferredRequiredMarker: markers.KubebuilderRequiredMarker,
				}),
			},
			expectedErr: "",
		}),
		Entry("With invalid preferred optional marker", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalOrRequired": toYaml(optionalorrequired.OptionalOrRequiredConfig{
					PreferredOptionalMarker: "invalid",
				}),
			},
			expectedErr: "lintersConfig.optionalorrequired.preferredOptionalMarker: Invalid value: \"invalid\": invalid value, must be one of \"optional\", \"kubebuilder:validation:Optional\" or omitted",
		}),
		Entry("With invalid preferred required marker", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalOrRequired": toYaml(optionalorrequired.OptionalOrRequiredConfig{
					PreferredRequiredMarker: "invalid",
				}),
			},
			expectedErr: "lintersConfig.optionalorrequired.preferredRequiredMarker: Invalid value: \"invalid\": invalid value, must be one of \"required\", \"kubebuilder:validation:Required\" or omitted",
		}),

		// OptionalOrRequiredConfig validation
		Entry("With a valid OptionalOrRequiredConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalorrequired": toYaml(optionalorrequired.OptionalOrRequiredConfig{
					PreferredOptionalMarker: markers.OptionalMarker,
					PreferredRequiredMarker: markers.RequiredMarker,
				}),
			},
			expectedErr: "",
		}),
		Entry("With kubebuilder preferred markers", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalorrequired": toYaml(optionalorrequired.OptionalOrRequiredConfig{
					PreferredOptionalMarker: markers.KubebuilderOptionalMarker,
					PreferredRequiredMarker: markers.KubebuilderRequiredMarker,
				}),
			},
			expectedErr: "",
		}),
		Entry("With invalid preferred optional marker", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalorrequired": toYaml(optionalorrequired.OptionalOrRequiredConfig{
					PreferredOptionalMarker: "invalid",
				}),
			},
			expectedErr: "lintersConfig.optionalorrequired.preferredOptionalMarker: Invalid value: \"invalid\": invalid value, must be one of \"optional\", \"kubebuilder:validation:Optional\" or omitted",
		}),
		Entry("With invalid preferred required marker", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"optionalorrequired": toYaml(optionalorrequired.OptionalOrRequiredConfig{
					PreferredRequiredMarker: "invalid",
				}),
			},
			expectedErr: "lintersConfig.optionalorrequired.preferredRequiredMarker: Invalid value: \"invalid\": invalid value, must be one of \"required\", \"kubebuilder:validation:Required\" or omitted",
		}),

		// RequiredFieldsConfig validation (with legacy field name)
		Entry("With a valid RequiredFieldsConfig: omitted", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"requiredFields": toYaml(requiredfields.RequiredFieldsConfig{
					PointerPolicy: "",
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid RequiredFieldsConfig: SuggestFix", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"requiredFields": toYaml(requiredfields.RequiredFieldsConfig{
					PointerPolicy: requiredfields.RequiredFieldPointerSuggestFix,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid RequiredFieldsConfig: Warn", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"requiredFields": toYaml(requiredfields.RequiredFieldsConfig{
					PointerPolicy: requiredfields.RequiredFieldPointerWarn,
				}),
			},
			expectedErr: "",
		}),
		Entry("With an invalid RequiredFieldsConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"requiredFields": toYaml(requiredfields.RequiredFieldsConfig{
					PointerPolicy: "invalid",
				}),
			},
			expectedErr: "lintersConfig.requiredfields.pointerPolicy: Invalid value: \"invalid\": invalid value, must be one of \"Warn\", \"SuggestFix\" or omitted",
		}),

		// SSATagsConfig validation
		Entry("With a valid SSATagsConfig: Warn", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"ssatags": toYaml(ssatags.SSATagsConfig{
					ListTypeSetUsage: ssatags.SSATagsListTypeSetUsageWarn,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid SSATagsConfig: Ignore", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"ssatags": toYaml(ssatags.SSATagsConfig{
					ListTypeSetUsage: ssatags.SSATagsListTypeSetUsageIgnore,
				}),
			},
			expectedErr: "",
		}),
		Entry("With an invalid SSATagsConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"ssatags": toYaml(ssatags.SSATagsConfig{
					ListTypeSetUsage: "invalid",
				}),
			},
			expectedErr: "lintersConfig.ssatags.listTypeSetUsage: Invalid value: \"invalid\": invalid value, must be one of \"Warn\", \"Ignore\" or omitted",
		}),

		// RequiredFieldsConfig validation
		Entry("With a valid RequiredFieldsConfig: omitted", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"requiredfields": toYaml(requiredfields.RequiredFieldsConfig{
					PointerPolicy: "",
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid RequiredFieldsConfig: SuggestFix", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"requiredfields": toYaml(requiredfields.RequiredFieldsConfig{
					PointerPolicy: requiredfields.RequiredFieldPointerSuggestFix,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid RequiredFieldsConfig: Warn", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"requiredfields": toYaml(requiredfields.RequiredFieldsConfig{
					PointerPolicy: requiredfields.RequiredFieldPointerWarn,
				}),
			},
			expectedErr: "",
		}),
		Entry("With an invalid RequiredFieldsConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"requiredfields": toYaml(requiredfields.RequiredFieldsConfig{
					PointerPolicy: "invalid",
				}),
			},
			expectedErr: "lintersConfig.requiredfields.pointerPolicy: Invalid value: \"invalid\": invalid value, must be one of \"Warn\", \"SuggestFix\" or omitted",
		}),

		// StatusOptionalConfig validation (with legacy field name)
		Entry("With a valid StatusOptionalConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"statusOptional": toYaml(statusoptional.StatusOptionalConfig{
					PreferredOptionalMarker: markers.OptionalMarker,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid StatusOptionalConfig: k8s optional marker", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"statusOptional": toYaml(statusoptional.StatusOptionalConfig{
					PreferredOptionalMarker: markers.K8sOptionalMarker,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid StatusOptionalConfig: kubebuilder optional marker", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"statusOptional": toYaml(statusoptional.StatusOptionalConfig{
					PreferredOptionalMarker: markers.KubebuilderOptionalMarker,
				}),
			},
		}),
		Entry("With an invalid StatusOptionalConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"statusOptional": toYaml(statusoptional.StatusOptionalConfig{
					PreferredOptionalMarker: "invalid",
				}),
			},
			expectedErr: "lintersConfig.statusoptional.preferredOptionalMarker: Invalid value: \"invalid\": invalid value, must be one of \"optional\", \"kubebuilder:validation:Optional\", \"k8s:optional\" or omitted",
		}),

		// StatusOptionalConfig validation
		Entry("With a valid StatusOptionalConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"statusoptional": toYaml(statusoptional.StatusOptionalConfig{
					PreferredOptionalMarker: markers.OptionalMarker,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid StatusOptionalConfig: k8s optional marker", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"statusoptional": toYaml(statusoptional.StatusOptionalConfig{
					PreferredOptionalMarker: markers.K8sOptionalMarker,
				}),
			},
			expectedErr: "",
		}),
		Entry("With a valid StatusOptionalConfig: kubebuilder optional marker", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"statusoptional": toYaml(statusoptional.StatusOptionalConfig{
					PreferredOptionalMarker: markers.KubebuilderOptionalMarker,
				}),
			},
		}),
		Entry("With an invalid StatusOptionalConfig", validateLintersConfigTableInput{
			config: config.LintersConfig{
				"statusoptional": toYaml(statusoptional.StatusOptionalConfig{
					PreferredOptionalMarker: "invalid",
				}),
			},
			expectedErr: "lintersConfig.statusoptional.preferredOptionalMarker: Invalid value: \"invalid\": invalid value, must be one of \"optional\", \"kubebuilder:validation:Optional\", \"k8s:optional\" or omitted",
		}),
	)
})

func toYaml(v any) []byte {
	yaml, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}

	return yaml
}

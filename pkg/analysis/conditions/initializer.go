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
package conditions

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/initializer"
	"sigs.k8s.io/kube-api-linter/pkg/config"
)

// Initializer returns the AnalyzerInitializer for this
// Analyzer so that it can be added to the registry.
func Initializer() initializer.AnalyzerInitializer {
	return initializer.NewConfigurableInitializer(
		name,
		initAnalyzer,
		true,
		validateConfig,
	)
}

func initAnalyzer(cfg config.LintersConfig) (*analysis.Analyzer, error) {
	return newAnalyzer(cfg.Conditions), nil
}

// validateConfig implements validation of the conditions linter config.
func validateConfig(cfg any, fldPath *field.Path) field.ErrorList {
	cc, ok := cfg.(config.ConditionsConfig)
	if !ok {
		return field.ErrorList{field.InternalError(fldPath, initializer.NewIncorrectTypeError(cfg))}
	}

	fieldErrors := field.ErrorList{}

	switch cc.IsFirstField {
	case "", config.ConditionsFirstFieldWarn, config.ConditionsFirstFieldIgnore:
	default:
		fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("isFirstField"), cc.IsFirstField, fmt.Sprintf("invalid value, must be one of %q, %q or omitted", config.ConditionsFirstFieldWarn, config.ConditionsFirstFieldIgnore)))
	}

	switch cc.UseProtobuf {
	case "", config.ConditionsUseProtobufSuggestFix, config.ConditionsUseProtobufWarn, config.ConditionsUseProtobufIgnore, config.ConditionsUseProtobufForbid:
	default:
		fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("useProtobuf"), cc.UseProtobuf, fmt.Sprintf("invalid value, must be one of %q, %q, %q, %q or omitted", config.ConditionsUseProtobufSuggestFix, config.ConditionsUseProtobufWarn, config.ConditionsUseProtobufIgnore, config.ConditionsUseProtobufForbid)))
	}

	switch cc.UsePatchStrategy {
	case "", config.ConditionsUsePatchStrategySuggestFix, config.ConditionsUsePatchStrategyWarn, config.ConditionsUsePatchStrategyIgnore, config.ConditionsUsePatchStrategyForbid:
	default:
		fieldErrors = append(fieldErrors, field.Invalid(fldPath.Child("usePatchStrategy"), cc.UsePatchStrategy, fmt.Sprintf("invalid value, must be one of %q, %q, %q, %q or omitted", config.ConditionsUsePatchStrategySuggestFix, config.ConditionsUsePatchStrategyWarn, config.ConditionsUsePatchStrategyIgnore, config.ConditionsUsePatchStrategyForbid)))
	}

	return fieldErrors
}

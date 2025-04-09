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
package inspector

import (
	"errors"
	"reflect"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	astinspector "golang.org/x/tools/go/ast/inspector"

	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/extractjsontags"
	"sigs.k8s.io/kube-api-linter/pkg/analysis/helpers/markers"
)

const name = "inspector"

var (
	errCouldNotGetInspector = errors.New("could not get inspector")
	errCouldNotGetJSONTags  = errors.New("could not get json tags")
	errCouldNotGetMarkers   = errors.New("could not get markers")
)

// Analyzer is the analyzer for the inspector package.
// It provides common functionality for analyzers that need to inspect fields and struct.
// Abstracting away filtering of fields that the analyzers should and shouldn't be worrying about.
var Analyzer = &analysis.Analyzer{
	Name:       name,
	Doc:        "Provides common functionality for analyzers that need to inspect fields and struct",
	Run:        run,
	Requires:   []*analysis.Analyzer{inspect.Analyzer, extractjsontags.Analyzer, markers.Analyzer},
	ResultType: reflect.TypeOf(newInspector(nil, nil, nil)),
}

func run(pass *analysis.Pass) (interface{}, error) {
	astinspector, ok := pass.ResultOf[inspect.Analyzer].(*astinspector.Inspector)
	if !ok {
		return nil, errCouldNotGetInspector
	}

	jsonTags, ok := pass.ResultOf[extractjsontags.Analyzer].(extractjsontags.StructFieldTags)
	if !ok {
		return nil, errCouldNotGetJSONTags
	}

	markersAccess, ok := pass.ResultOf[markers.Analyzer].(markers.Markers)
	if !ok {
		return nil, errCouldNotGetMarkers
	}

	return newInspector(astinspector, jsonTags, markersAccess), nil
}

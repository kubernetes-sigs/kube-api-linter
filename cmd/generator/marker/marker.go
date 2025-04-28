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
package main

import (
	"bytes"
	_ "embed"
	"errors"
	"go/format"
	"html/template"
	"os"
	"path/filepath"
	"slices"
	"sort"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	crdmarkers "sigs.k8s.io/controller-tools/pkg/crd/markers"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

//go:embed markers.tmpl
var tmpl []byte

// Marker is a elements for generating marker definitions.
type Marker struct {
	// This field is used for struct name.
	StructName string

	Name           string
	RequiredFields []string
}

// Input is the input for the template.
type Input struct {
	Markers []Marker
}

func structName(name string) string {
	words := bytes.Split([]byte(name), []byte(":"))
	for i, word := range words {
		words[i] = cases.Title(language.Und).Bytes(word)
	}

	return string(bytes.Join(words, nil))
}

func makeDefinitions() []markers.Definition {
	allDefinitions := make([]markers.Definition, 0, len(crdmarkers.AllDefinitions))
	for _, v := range crdmarkers.AllDefinitions {
		allDefinitions = append(allDefinitions, *v.Definition)
	}

	sort.Slice(allDefinitions, func(i, j int) bool {
		return allDefinitions[i].Name < allDefinitions[j].Name
	})

	return slices.CompactFunc(allDefinitions, func(a, b markers.Definition) bool {
		return a.Name == b.Name
	})
}

func makeMarkers() []Marker {
	definitions := makeDefinitions()
	markers := make([]Marker, 0, len(definitions))

	for _, v := range definitions {
		marker := Marker{
			StructName: structName(v.Name),
			Name:       v.Name,
		}

		requiredField := make([]string, 0)

		for k, v := range v.Fields {
			if v.Optional || k == "" {
				continue
			}

			requiredField = append(requiredField, k)
		}

		sort.Strings(requiredField)
		marker.RequiredFields = requiredField
		markers = append(markers, marker)
	}

	return markers
}

func generateMarkers() []byte {
	markers := makeMarkers()

	tmpl, err := template.New("marker-gen").Parse(string(tmpl))
	if err != nil {
		panic(err)
	}

	input := Input{
		Markers: markers,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, input); err != nil {
		panic(err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	return formatted
}

func write(b []byte) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filePath := filepath.Join(wd, "pkg", "analysis", "helpers", "markers", "marker.go")
	if err := os.Remove(filePath); err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}

	f, err := os.Create(filepath.Clean(filePath))
	if err != nil && !errors.Is(err, os.ErrExist) {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := f.Write(b); err != nil {
		panic(err)
	}
}

func main() {
	markers := generateMarkers()
	write(markers)
}

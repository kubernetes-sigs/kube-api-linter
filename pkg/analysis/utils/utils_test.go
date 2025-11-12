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
package utils_test

import (
	"go/ast"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
)

var _ = Describe("FieldName", func() {
	type fieldNameInput struct {
		field *ast.Field
		want  string
	}

	DescribeTable("Should extract the field name", func(in fieldNameInput) {
		Expect(utils.FieldName(in.field)).To(Equal(in.want), "expect to match the extracted field name")
	},
		Entry("field has Names", fieldNameInput{
			field: &ast.Field{
				Names: []*ast.Ident{
					{
						Name: "foo",
					},
				},
			},
			want: "foo",
		}),
		Entry("field has no Names, but is an Ident", fieldNameInput{
			field: &ast.Field{
				Type: &ast.Ident{
					Name: "foo",
				},
			},
			want: "foo",
		}),
		Entry("field has no Names, but is a StarExpr with an Ident", fieldNameInput{
			field: &ast.Field{
				Type: &ast.StarExpr{
					X: &ast.Ident{
						Name: "foo",
					},
				},
			},
			want: "foo",
		}),
		Entry("field has no Names, and is not an Ident or StarExpr", fieldNameInput{
			field: &ast.Field{
				Type: &ast.ArrayType{
					Elt: &ast.Ident{
						Name: "foo",
					},
				},
			},
			want: "",
		}),
	)
})

func TestGetStructName(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, structNameAnalyzer(), "getstructname")
}

func structNameAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "test",
		Doc:      "test",
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run: func(pass *analysis.Pass) (any, error) {
			inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
			if !ok {
				return nil, errCouldNotGetInspector
			}

			// Filter to structs so that we can iterate over fields in a struct.
			nodeFilter := []ast.Node{
				(*ast.Field)(nil),
			}

			inspect.Preorder(nodeFilter, func(n ast.Node) {
				field, ok := n.(*ast.Field)
				if !ok {
					return
				}

				fieldName := utils.FieldName(field)
				structName := utils.GetStructName(pass, field)

				pass.Reportf(field.Pos(), "field %s is in struct %s", fieldName, structName)
			})

			return nil, nil
		},
	}
}

package utils_test

import (
	"go/ast"
	"testing"

	"sigs.k8s.io/kube-api-linter/pkg/analysis/utils"
)

func TestFieldName(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		field *ast.Field
		want  string
	}{
		"field has Names": {
			field: &ast.Field{
				Names: []*ast.Ident{
					{
						Name: "foo",
					},
				},
			},
			want: "foo",
		},
		"filed has no Names, but is an Ident": {
			field: &ast.Field{
				Type: &ast.Ident{
					Name: "foo",
				},
			},
			want: "foo",
		},
		"field has no Names, but is a StarExpr with an Ident": {
			field: &ast.Field{
				Type: &ast.StarExpr{
					X: &ast.Ident{
						Name: "foo",
					},
				},
			},
			want: "foo",
		},
		"field has no Names, and is not an Ident or StarExpr": {
			field: &ast.Field{
				Type: &ast.ArrayType{
					Elt: &ast.Ident{
						Name: "foo",
					},
				},
			},
			want: "",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := utils.FieldName(tc.field)
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

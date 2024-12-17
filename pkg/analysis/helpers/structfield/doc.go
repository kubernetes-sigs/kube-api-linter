/*
structfield is a helper used to map an *ast.Field to the parent *ast.StructType.

The package returns a [StructField] interface, which can be used to access the parent struct of a field.

Example:

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	structField := pass.ResultOf[structfield.Analyzer].(structfield.StructField)

	nodeFilter := []ast.Node{
		(*ast.Field)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		field, ok := n.(*ast.FieldType)
		if !ok {
			return
		}

		if structType := structField.StructForField(field) == nil {
			// This is not a field in a struct.
			return
		}

		...
	})
*/
package structfield

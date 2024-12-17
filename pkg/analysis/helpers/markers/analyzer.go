package markers

import (
	"errors"
	"go/ast"
	"go/token"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var (
	errCouldNotGetInspector  = errors.New("could not get inspector")
	errCouldNotCreateMarkers = errors.New("could not create markers")
)

// Markers allows access to markers extracted from the
// go types.
type Markers interface {
	// FieldMarkers returns markers associated to the field.
	FieldMarkers(*ast.Field) MarkerSet

	// StructMarkers returns markers associated to the given sturct.
	StructMarkers(*ast.StructType) MarkerSet

	// TypeMarkers returns markers associated to the given type.
	TypeMarkers(*ast.TypeSpec) MarkerSet
}

func newMarkers() Markers {
	return &markers{
		fieldMarkers:  make(map[*ast.Field]MarkerSet),
		structMarkers: make(map[*ast.StructType]MarkerSet),
		typeMarkers:   make(map[*ast.TypeSpec]MarkerSet),
	}
}

// markers implements the storage for the implementation of the Markers interface.
type markers struct {
	fieldMarkers  map[*ast.Field]MarkerSet
	structMarkers map[*ast.StructType]MarkerSet
	typeMarkers   map[*ast.TypeSpec]MarkerSet
}

// FieldMarkers return the appropriate MarkerSet for the field,
// or an empty MarkerSet if the appropriate MarkerSet isn't found.
func (m *markers) FieldMarkers(field *ast.Field) MarkerSet {
	fMarkers, ok := m.fieldMarkers[field]
	if !ok {
		return NewMarkerSet()
	}

	return fMarkers
}

// StructMarkers returns the appropriate MarkerSet if found, else
// it returns an empty MarkerSet.
func (m *markers) StructMarkers(sTyp *ast.StructType) MarkerSet {
	sMarkers, ok := m.structMarkers[sTyp]
	if !ok {
		return NewMarkerSet()
	}

	return sMarkers
}

// TypeMarkers return the appropriate MarkerSet for the type,
// or an empty MarkerSet if the appropriate MarkerSet isn't found.
func (m *markers) TypeMarkers(typ *ast.TypeSpec) MarkerSet {
	tMarkers, ok := m.typeMarkers[typ]
	if !ok {
		return NewMarkerSet()
	}

	return tMarkers
}

func (m *markers) insertFieldMarkers(field *ast.Field, ms MarkerSet) {
	m.fieldMarkers[field] = ms
}

func (m *markers) insertStructMarkers(sTyp *ast.StructType, ms MarkerSet) {
	m.structMarkers[sTyp] = ms
}

func (m *markers) insertTypeMarkers(typ *ast.TypeSpec, ms MarkerSet) {
	m.typeMarkers[typ] = ms
}

// Analyzer is the analyzer for the markers package.
// It iterates over declarations within a package and parses the comments to extract markers.
var Analyzer = &analysis.Analyzer{
	Name:       "markers",
	Doc:        "Iterates over declarations within a package and parses the comments to extract markers",
	Run:        run,
	Requires:   []*analysis.Analyzer{inspect.Analyzer},
	ResultType: reflect.TypeOf(newMarkers()),
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errCouldNotGetInspector
	}

	nodeFilter := []ast.Node{
		(*ast.TypeSpec)(nil),
		(*ast.Field)(nil),
	}

	results, ok := newMarkers().(*markers)
	if !ok {
		return nil, errCouldNotCreateMarkers
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch typ := n.(type) {
		case *ast.TypeSpec:
			extractTypeSpecMarkers(typ, results)
		case *ast.Field:
			extractFieldMarkers(typ, results)
		}
	})

	return results, nil
}

func extractTypeSpecMarkers(typ *ast.TypeSpec, results *markers) {
	typeMarkers := NewMarkerSet()

	if typ.Doc != nil {
		for _, comment := range typ.Doc.List {
			if marker := extractMarker(comment); marker.Value != "" {
				typeMarkers.Insert(marker)
			}
		}
	}

	results.insertTypeMarkers(typ, typeMarkers)

	if uTyp, ok := typ.Type.(*ast.StructType); ok {
		results.insertStructMarkers(uTyp, typeMarkers)
	}
}

func extractFieldMarkers(field *ast.Field, results *markers) {
	if field == nil || field.Doc == nil {
		return
	}

	fieldMarkers := NewMarkerSet()

	for _, comment := range field.Doc.List {
		if marker := extractMarker(comment); marker.Value != "" {
			fieldMarkers.Insert(marker)
		}
	}

	results.insertFieldMarkers(field, fieldMarkers)
}

func extractMarker(comment *ast.Comment) Marker {
	if !strings.HasPrefix(comment.Text, "// +") {
		return Marker{}
	}

	return Marker{
		Value:      strings.TrimPrefix(comment.Text, "// +"),
		RawComment: comment.Text,
		Pos:        comment.Pos(),
		End:        comment.End(),
	}
}

// Marker represents a marker extracted from a comment on a declaration.
type Marker struct {
	// Value is the value of the marker once the leading comment and '+' are trimmed.
	Value string

	// RawComment is the raw comment line, unfiltered.
	RawComment string

	// Pos is the starting position in the file for the comment line containing the marker.
	Pos token.Pos

	// End is the ending position in the file for the coment line containing the marker.
	End token.Pos
}

// MarkerSet is a set implementation for Markers that uses
// the Marker value as the key, but returns the full Marker
// as the result.
type MarkerSet map[string]Marker

// NewMarkerSet initialises a new MarkerSet with the provided values.
// If any markers have the same value, the latter marker in the list
// will take precedence, no duplication checks are implemented.
func NewMarkerSet(markers ...Marker) MarkerSet {
	ms := make(MarkerSet)

	ms.Insert(markers...)

	return ms
}

// Insert add the given markers to the MarkerSet.
// If any markers have the same value, the latter marker in the list
// will take precedence, no duplication checks are implemented.
func (ms MarkerSet) Insert(markers ...Marker) {
	for _, marker := range markers {
		ms[marker.Value] = marker
	}
}

// Has returns whether a marker with the value given is present in the
// MarkerSet.
func (ms MarkerSet) Has(value string) bool {
	_, ok := ms[value]
	return ok
}

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
package markerscope

import (
	"errors"
	"fmt"
)

var (
	errScopeRequired = errors.New("scope is required")
)

// InvalidScopeConstraintError represents an error when a scope constraint is invalid.
type InvalidScopeConstraintError struct {
	Scope string
}

func (e *InvalidScopeConstraintError) Error() string {
	return fmt.Sprintf("invalid scope: %q (must be one of: Field, Type, Any)", e.Scope)
}

// InvalidSchemaTypeError represents an error when a schema type is invalid.
type InvalidSchemaTypeError struct {
	SchemaType string
}

func (e *InvalidSchemaTypeError) Error() string {
	return fmt.Sprintf("invalid schema type: %q", e.SchemaType)
}

// InvalidTypeConstraintError represents an error when a type constraint is invalid.
type InvalidTypeConstraintError struct {
	Err error
}

func (e *InvalidTypeConstraintError) Error() string {
	return fmt.Sprintf("invalid type constraint: %v", e.Err)
}

// InvalidElementConstraintError represents an error when an element constraint is invalid.
type InvalidElementConstraintError struct {
	Err error
}

func (e *InvalidElementConstraintError) Error() string {
	return fmt.Sprintf("array element: %v", e.Err)
}

// MarkerShouldBeOnTypeDefinitionError represents an error when a marker should be declared on the type definition.
type MarkerShouldBeOnTypeDefinitionError struct {
	TypeName string
}

func (e *MarkerShouldBeOnTypeDefinitionError) Error() string {
	return fmt.Sprintf("marker should be declared on the type definition of %s instead of the field", e.TypeName)
}

// DengerousTypeError represents an error when a dangerous type is used.
type DengerousTypeError struct {
	Type string
}

func (e *DengerousTypeError) Error() string {
	return fmt.Sprintf("type %s is dangerous and not allowed (set allowDangerousTypes to true to permit)", e.Type)
}

// TypeNotAllowedError represents an error when a type is not allowed.
type TypeNotAllowedError struct {
	Type         SchemaType
	AllowedTypes []SchemaType
}

func (e *TypeNotAllowedError) Error() string {
	return fmt.Sprintf("type %s is not allowed (expected one of: %v)", e.Type, e.AllowedTypes)
}

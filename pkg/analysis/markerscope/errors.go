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

type invalidScopeConstraintError struct {
	scope string
}

func (e *invalidScopeConstraintError) Error() string {
	return fmt.Sprintf("invalid scope: %q (must be one of: Field, Type, Any)", e.scope)
}

type invalidNamedTypeConstraintError struct {
	constraint string
}

func (e *invalidNamedTypeConstraintError) Error() string {
	return fmt.Sprintf("invalid namedTypeConstraint: %q (must be one of: AllowTypeOrField, OnTypeOnly, or empty)", e.constraint)
}

type invalidSchemaTypeError struct {
	schemaType string
}

func (e *invalidSchemaTypeError) Error() string {
	return fmt.Sprintf("invalid schema type: %q", e.schemaType)
}

type invalidTypeConstraintError struct {
	err error
}

func (e *invalidTypeConstraintError) Error() string {
	return fmt.Sprintf("invalid type constraint: %v", e.err)
}

type invalidElementConstraintError struct {
	err error
}

func (e *invalidElementConstraintError) Error() string {
	return fmt.Sprintf("array element: %v", e.err)
}

type markerShouldBeOnTypeDefinitionError struct {
	typeName string
}

func (e *markerShouldBeOnTypeDefinitionError) Error() string {
	return fmt.Sprintf("marker should be declared on the type definition of %s instead of the field", e.typeName)
}

// Is implements error matching for markerShouldBeOnTypeDefinitionError.
func (e *markerShouldBeOnTypeDefinitionError) Is(target error) bool {
	_, ok := target.(*markerShouldBeOnTypeDefinitionError)
	return ok
}

type typeNotAllowedError struct {
	schemaType   SchemaType
	allowedTypes []SchemaType
}

func (e *typeNotAllowedError) Error() string {
	return fmt.Sprintf("type %s is not allowed (expected one of: %v)", e.schemaType, e.allowedTypes)
}

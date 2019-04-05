package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/NAKKA-K/learn-interpreter-in-go/ast"
)

// ObjectType is type string identifier
type ObjectType string

// Object is all value in program
type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE_OBJ"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
)

// Integer is from IntegerLiteral
type Integer struct {
	Value int64
}

// Inspect return "1"
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// Boolean is from ast.Boolean
type Boolean struct {
	Value bool
}

// Inspect return "true"
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// Null is null
type Null struct{}

// Inspect return "null"
func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

// ReturnValue is object wrapper
type ReturnValue struct {
	Value Object
}

// Inspect return Inspect() of object
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

// Error include error message
type Error struct {
	Message string
}

// Inspect return "ERROR: ~"
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
func (e *Error) Type() ObjectType { return ERROR_OBJ }

// Function is from ast.FunctionLiteral
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Inspect return "fn(...params) { 'BlockStatement.String()' }"
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n}")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

// String is from ast.StringLiteral
type String struct {
	Value string
}

// Inspect return "<string>"
func (s *String) Inspect() string  { return s.Value }
func (s *String) Type() ObjectType { return STRING_OBJ }

// BuiltinFunction is builtin function type
type BuiltinFunction func(args ...Object) Object

// Builtin is builtin function type
type Builtin struct {
	Fn BuiltinFunction
}

// Inspect return "builtin function"
func (b *Builtin) Inspect() string  { return "builtin function" }
func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }

// Array is from ast.ArrayLiteral
type Array struct {
	Elements []Object
}

// Inspect return "[...<elements>]"
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
func (ao *Array) Type() ObjectType { return ARRAY_OBJ }

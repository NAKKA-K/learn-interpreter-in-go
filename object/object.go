package object

import "fmt"

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

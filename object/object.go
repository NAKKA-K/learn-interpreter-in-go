package object

// ObjectType is type string identifier
type ObjectType string

// Object is all value in program
type Object interface {
	Type() ObjectType
	Inspect() string
}

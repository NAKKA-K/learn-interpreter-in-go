package object

// NewEnviroment return Environment object
func NewEnviroment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment is map of identifier
type Environment struct {
	store map[string]Object
}

// Get identifier from environment map
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set identifier to environment map
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

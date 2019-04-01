package object

// NewEnclosedEnvironment is
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment return Environment object
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment is map of identifier
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get identifier from environment map
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set identifier to environment map
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

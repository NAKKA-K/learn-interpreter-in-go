package evaluator

import (
	"github.com/NAKKA-K/learn-interpreter-in-go/ast"
	"github.com/NAKKA-K/learn-interpreter-in-go/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}

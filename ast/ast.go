package ast

import "github.com/NAKKA-K/learn-interpreter-in-go/token"

// Node interface must be implemented to all nodes of AST
type Node interface {
	TokenLiteral() string
}

// Statement are sentences in program
type Statement interface {
	Node
	statementNode() // statementNode is dummy function for type verification
}

// Expression are the interface of things that generate values
type Expression interface {
	Node
	expressionNode() // expressionNode is dummy function for type verification
}

// Program is root node of AST
type Program struct {
	Statements []Statement
}

// TokenLiteral return literal of first statement in program
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement is for "let" statement
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

// TokenLiteral return "let"
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) statementNode()       {}

// Identifier is for identifer expression(x, tmp, etc...)
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

// TokenLiteral return variable name(x, tmp, etc...)
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) expressionNode()      {}


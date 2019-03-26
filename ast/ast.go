package ast

import (
	"bytes"

	"github.com/NAKKA-K/learn-interpreter-in-go/token"
)

// Node interface must be implemented to all nodes of AST
type Node interface {
	TokenLiteral() string
	String() string
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

// String return string view of all statements
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement is for "let" statement
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

// TokenLiteral return "let"
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil { // HACK: nil check
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
func (ls *LetStatement) statementNode() {}

// Identifier is for identifer expression(x, tmp, etc...)
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

// TokenLiteral return variable name(x, tmp, etc...)
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }
func (i *Identifier) expressionNode()      {}

// ReturnStatement for "return" statement
type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

// TokenLiteral return "return"
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil { // HACK: nil check
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
func (rs *ReturnStatement) statementNode() {}

// ExpressionStatement for expression statement
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

// TokenLiteral return a token of first of expression
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil { // HACK: nil check
		return es.Expression.String()
	}
	return ""
}
func (es *ExpressionStatement) statementNode() {}

// IntegerLiteral for integer literal
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// TokenLiteral return integer
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }
func (il *IntegerLiteral) expressionNode()      {}

// PrefixExpression for prefix of expression
type PrefixExpression struct {
	Token    token.Token // For example, !, -, etc...
	Operator string
	Right    Expression
}

// TokenLiteral return prefix of expression string
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
func (pe *PrefixExpression) expressionNode() {}

// InfixExpression for 'expression operator expression'
type InfixExpression struct {
	Token    token.Token // For example, +, !=, etc...
	Left     Expression
	Operator string
	Right    Expression
}

// TokenLiteral return  +, !=, etc...
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}
func (oe *InfixExpression) expressionNode() {}

// Boolean for true or false
type Boolean struct {
	Token token.Token
	Value bool
}

// TokenLiteral return 'true' or 'false'
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }
func (b *Boolean) expressionNode()      {}

package ast

import (
	"bytes"
	"strings"

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

// IfExpression for 'if' and 'else' expression
type IfExpression struct {
	Token       token.Token // 'if'
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

// TokenLiteral return 'if'
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}
func (ie *IfExpression) expressionNode() {}

// BlockStatement for '{'
type BlockStatement struct {
	Token      token.Token // '{'
	Statements []Statement
}

// TokenLiteral return '{'
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
func (bs *BlockStatement) statementNode() {}

// FunctionLiteral for 'fn'
type FunctionLiteral struct {
	Token      token.Token // 'fn'
	Parameters []*Identifier
	Body       *BlockStatement
}

// TokenLiteral return 'fn'
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}
func (fl *FunctionLiteral) expressionNode() {}

// CallExpression for call function (in identifier)
type CallExpression struct {
	Token     token.Token // '('
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

// TokenLiteral return '('
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
func (ce *CallExpression) expressionNode() {}

// StringLiteral for string
type StringLiteral struct {
	Token token.Token
	Value string
}

// TokenLiteral return <string>
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }
func (sl *StringLiteral) expressionNode()      {}

// ArrayLiteral for array
type ArrayLiteral struct {
	Token    token.Token // '['
	Elements []Expression
}

// TokenLiteral return '['
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
func (al *ArrayLiteral) expressionNode() {}

// IndexExpression for call of array
type IndexExpression struct {
	Token token.Token // '['
	Left  Expression
	Index Expression
}

// TokenLiteral return '['
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}
func (ie *IndexExpression) expressionNode() {}

// HashLiteral for hash
type HashLiteral struct {
	Token token.Token // '{'
	Pairs map[Expression]Expression
}

// TokenLiteral return '{'
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
func (hl *HashLiteral) expressionNode() {}

// MacroLiteral for macro
type MacroLiteral struct {
	Token      token.Token // 'macro'
	Parameters []*Identifier
	Body       *BlockStatement
}

// TokenLiteral return 'macro'
func (ml *MacroLiteral) TokenLiteral() string { return ml.Token.Literal }
func (ml *MacroLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range ml.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(ml.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(ml.Body.String())

	return out.String()
}
func (ml *MacroLiteral) expressionNode() {}

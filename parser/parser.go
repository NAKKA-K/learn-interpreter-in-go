package parser

import (
	"github.com/NAKKA-K/learn-interpreter-in-go/ast"
	"github.com/NAKKA-K/learn-interpreter-in-go/lexer"
	"github.com/NAKKA-K/learn-interpreter-in-go/token"
)

// Parser manage tokens
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New generated parser be returned
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Set token to curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram generate ast.Program, and set statements to it
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

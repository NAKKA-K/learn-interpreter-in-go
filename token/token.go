package token

// TokenType is data type
type TokenType string

// Token is any token of program
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifier + literal
	IDENT = "IDENT"
	INT   = "INT"

	// operator
	ASSIGN = "="
	PLUS   = "+"

	// delimiter
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBARAN = "{"
	RBARAN = "}"

	// keyword
	FUCNTION = "FUNCTION"
	LET      = "LET"
)

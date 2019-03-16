package main

// TokenType is data type
type TokenType string

// Token is any token of program
type Token struct {
	Type    TokenType
	Literal string
}


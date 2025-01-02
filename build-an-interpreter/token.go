package main

// Token and Lexer (From Step 2)
type TokenType string

const (
	// EOF is short for End of line
	EOF     TokenType = "EOF"
	IDENT   TokenType = "IDENT"
	NUMBER  TokenType = "NUMBER"
	ASSIGN  TokenType = "="
	PLUS    TokenType = "+"
	MINUS   TokenType = "-"
	IF      TokenType = "IF"
	WHILE   TokenType = "WHILE"
	ILLEGAL TokenType = "ILLEGAL"
)

type Token struct {
	Type  TokenType
	Value string
}

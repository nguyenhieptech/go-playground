package main

import (
	"unicode"
)

type Lexer struct {
	input       string
	position    int
	currentRune rune
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readRune()
	return lexer
}

func (lexer *Lexer) readRune() {
	if lexer.position >= len(lexer.input) {
		lexer.currentRune = 0 // EOF
	} else {
		lexer.currentRune = rune(lexer.input[lexer.position])
	}
	lexer.position++
}

func (lexer *Lexer) readIdentifier() Token {
	start := lexer.position - 1
	for unicode.IsLetter(lexer.currentRune) {
		lexer.readRune()
	}
	value := lexer.input[start : lexer.position-1]

	switch value {
	case "if":
		return Token{Type: IF, Value: value}
	case "while":
		return Token{Type: WHILE, Value: value}
	default:
		return Token{Type: IDENT, Value: value}
	}
}

func (lexer *Lexer) readNumber() Token {
	start := lexer.position - 1
	for unicode.IsDigit(lexer.currentRune) {
		lexer.readRune()
	}
	return Token{Type: NUMBER, Value: lexer.input[start : lexer.position-1]}
}

func (lexer *Lexer) skipWhitespace() {
	for unicode.IsSpace(lexer.currentRune) {
		lexer.readRune()
	}
}

func (lexer *Lexer) NextToken() Token {
	lexer.skipWhitespace()

	switch lexer.currentRune {
	case '=':
		lexer.readRune()
		return Token{Type: ASSIGN, Value: "="}
	case '+':
		lexer.readRune()
		return Token{Type: PLUS, Value: "+"}
	case '-':
		lexer.readRune()
		return Token{Type: MINUS, Value: "-"}
	case 0:
		return Token{Type: EOF, Value: ""}
	default:
		if unicode.IsLetter(lexer.currentRune) {
			return lexer.readIdentifier()
		} else if unicode.IsDigit(lexer.currentRune) {
			return lexer.readNumber()
		}
	}
	return Token{Type: ILLEGAL, Value: string(lexer.currentRune)}
}

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// Token and Lexer (From Step 2)
type TokenType string

const (
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

type Lexer struct {
	input       string
	position    int
	currentRune rune
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readRune()
	return l
}

func (l *Lexer) readRune() {
	if l.position >= len(l.input) {
		l.currentRune = 0 // EOF
	} else {
		l.currentRune = rune(l.input[l.position])
	}
	l.position++
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	switch l.currentRune {
	case '=':
		l.readRune()
		return Token{Type: ASSIGN, Value: "="}
	case '+':
		l.readRune()
		return Token{Type: PLUS, Value: "+"}
	case '-':
		l.readRune()
		return Token{Type: MINUS, Value: "-"}
	case 0:
		return Token{Type: EOF, Value: ""}
	default:
		if unicode.IsLetter(l.currentRune) {
			return l.readIdentifier()
		} else if unicode.IsDigit(l.currentRune) {
			return l.readNumber()
		}
	}
	return Token{Type: ILLEGAL, Value: string(l.currentRune)}
}

func (l *Lexer) readIdentifier() Token {
	start := l.position - 1
	for unicode.IsLetter(l.currentRune) {
		l.readRune()
	}
	value := l.input[start : l.position-1]

	switch value {
	case "if":
		return Token{Type: IF, Value: value}
	case "while":
		return Token{Type: WHILE, Value: value}
	default:
		return Token{Type: IDENT, Value: value}
	}
}

func (l *Lexer) readNumber() Token {
	start := l.position - 1
	for unicode.IsDigit(l.currentRune) {
		l.readRune()
	}
	return Token{Type: NUMBER, Value: l.input[start : l.position-1]}
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.currentRune) {
		l.readRune()
	}
}

// AST and Parser (From Step 3)
type ASTNode struct {
	Token Token
	Left  *ASTNode
	Right *ASTNode
}

type Parser struct {
	lexer        *Lexer
	currentToken Token
}

func NewParser(lexer *Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.lexer.NextToken()
}

func (p *Parser) Parse() *ASTNode {
	return p.parseExpression()
}

func (p *Parser) parseExpression() *ASTNode {
	node := p.parseTerm()
	for p.currentToken.Type == PLUS || p.currentToken.Type == MINUS {
		token := p.currentToken
		p.nextToken()
		node = &ASTNode{
			Token: token,
			Left:  node,
			Right: p.parseTerm(),
		}
	}
	return node
}

func (p *Parser) parseTerm() *ASTNode {
	token := p.currentToken
	if token.Type == NUMBER {
		p.nextToken()
		return &ASTNode{Token: token}
	}
	return nil
}

// Helper function to print the AST
func printAST(node *ASTNode, indent int) {
	if node == nil {
		return
	}
	fmt.Printf("%sNode: %s\n", indentSpaces(indent), node.Token.Value)
	printAST(node.Left, indent+2)
	printAST(node.Right, indent+2)
}

func indentSpaces(n int) string {
	return strconv.Itoa(n)
}

func main() {
	// Step 1: Lexing the input
	input := "if x = 10 + 5 - 3"
	lexer := NewLexer(input)
	fmt.Println("Lexing Tokens:")
	for {
		tok := lexer.NextToken()
		fmt.Printf("Token: %+v\n", tok)
		if tok.Type == EOF {
			break
		}
	}

	// Step 2: Parsing the input and generating AST
	parser := NewParser(lexer)
	fmt.Println("\nParsing and generating AST:")
	ast := parser.Parse()
	printAST(ast, 0)
}

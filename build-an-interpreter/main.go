package main

import (
	"fmt"
)

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

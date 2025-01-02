package main

import (
	"fmt"
	"strconv"
)

// Abstract Syntax Tree
type ASTNode struct {
	Token Token
	Left  *ASTNode
	Right *ASTNode
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

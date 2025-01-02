package main

type Parser struct {
	lexer        *Lexer
	currentToken Token
}

func NewParser(lexer *Lexer) *Parser {
	parser := &Parser{lexer: lexer}
	parser.nextToken()
	return parser
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.lexer.NextToken()
}

func (parser *Parser) Parse() *ASTNode {
	return parser.parseExpression()
}

func (parser *Parser) parseExpression() *ASTNode {
	node := parser.parseTerm()
	for parser.currentToken.Type == PLUS || parser.currentToken.Type == MINUS {
		token := parser.currentToken
		parser.nextToken()
		node = &ASTNode{
			Token: token,
			Left:  node,
			Right: parser.parseTerm(),
		}
	}
	return node
}

func (parser *Parser) parseTerm() *ASTNode {
	token := parser.currentToken
	if token.Type == NUMBER {
		parser.nextToken()
		return &ASTNode{Token: token}
	}
	return nil
}

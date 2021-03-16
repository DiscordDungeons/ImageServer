package iql

import (
	"errors"
	"fmt"

	iqlSchema "discorddungeons.me/imageserver/iql/schema"
	iqlSyntax "discorddungeons.me/imageserver/iql/syntax"
)

type Parser struct {
	tokens  []iqlSchema.Token
	current int
}

func NewParser(tokens []iqlSchema.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (parser *Parser) Parse() []iqlSyntax.Stmt {
	fmt.Println("Start parse")

	statements := []iqlSyntax.Stmt{}

	for !parser.isAtEnd() {
		fmt.Println("Not at end")

		statement := parser.declaration()

		if statement != nil {
			statements = append(statements, statement)
		}
	}

	fmt.Println("Return statements")

	return statements
}

func (parser *Parser) expression() iqlSyntax.Expr {
	return parser.assignment()
}

func (parser *Parser) assignment() iqlSyntax.Expr {
	return iqlSyntax.NewVariable(iqlSchema.Token{TokenType: iqlSchema.TokenTypes.LOAD})
}

func (parser *Parser) declaration() iqlSyntax.Stmt {
	// if parser.Match(TokenTypes.FN) {
	// 	return parser.Func("function")
	// }

	return parser.statement()
}

func (parser *Parser) statement() iqlSyntax.Stmt {
	if parser.match(iqlSchema.TokenTypes.LOAD) {
		return parser.loadStatement()
	}

	return parser.expressionStatement()
}

func (parser *Parser) loadStatement() iqlSyntax.Stmt {
	//value := parser.expression()

	parser.consume(iqlSchema.TokenTypes.IMAGE, "expect image")
	parser.consume(iqlSchema.TokenTypes.FROM, "expect from")
	parser.consume(iqlSchema.TokenTypes.URL, "expect url")

	imageUrl := parser.advance()

	parser.consume(iqlSchema.TokenTypes.AS, "expect as")

	imageName := parser.advance()

	return iqlSyntax.NewLoadStmt(imageUrl.Literal.StrVal, imageName.Literal.StrVal)
}

func (parser *Parser) expressionStatement() iqlSyntax.Stmt {
	expr := parser.expression()

	return iqlSyntax.NewExpressionStmt(expr)
}

func (parser *Parser) match(types ...iqlSchema.TokenType) bool {
	for _, t := range types {
		if parser.check(t) {
			parser.advance()
			return true
		}
	}

	return false
}

func (parser *Parser) consume(tokenType iqlSchema.TokenType, message string) (iqlSchema.Token, error) {
	if parser.check(tokenType) {
		return parser.advance(), nil
	}

	return iqlSchema.Token{}, errors.New(message)
}

func (parser *Parser) check(t iqlSchema.TokenType) bool {
	if parser.isAtEnd() {
		return false
	}

	return parser.peek().TokenType == t
}

func (parser *Parser) advance() iqlSchema.Token {
	if !parser.isAtEnd() {
		parser.current++
	}

	return parser.previous()
}

func (parser *Parser) isAtEnd() bool {
	token := parser.peek()
	fmt.Printf("Token: %s\n", token)

	return token.TokenType == iqlSchema.TokenTypes.EOF

}

func (parser *Parser) peek() iqlSchema.Token {
	return parser.tokens[parser.current]
}

func (parser *Parser) previous() iqlSchema.Token {
	return parser.tokens[parser.current-1]
}

func (parser *Parser) primary() (iqlSyntax.Expr, error) {
	if parser.match(iqlSchema.TokenTypes.NUMBER, iqlSchema.TokenTypes.STRING) {
		return iqlSyntax.NewLiteralExpr(parser.previous().Literal), nil
	}

	if parser.match(iqlSchema.TokenTypes.IDENTIFIER) {
		return iqlSyntax.NewVariable(parser.previous()), nil
	}

	// if parser.match(iqlSchema.TokenTypes.LEFT_PAREN) {
	// 	expr := parser.expression()
	// 	parser.consume(iqlSchema.TokenTypes.RIGHT_PAREN, `Expect ')' after expression.`)

	// 	return iqlSyntax.Grouping(expr), nil
	// }

	return nil, errors.New("expected expression")
}

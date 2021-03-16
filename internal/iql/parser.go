package iql

import (
	"errors"

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
	statements := []iql.Stmt{}

	for !parser.isAtEnd() {
		statement := parser.declaration()

		if statement != nil {
			statements = append(statements, statement)
		}
	}

	return statements
}

func (parser *Parser) expression() Expr {
	return parser.assignment()
}

func (parser *Parser) declaration() iql.Stmt {
	// if parser.Match(TokenTypes.FN) {
	// 	return parser.Func("function")
	// }

	return parser.statement()
}

func (parser *Parser) statement() iql.Stmt {
	// if parser.Match(TokenTypes.FOR) {
	// 	return parser.ForStatement()
	// }
	// if parser.Match(TokenTypes.IF) {
	// 	return parser.IfStatement()
	// }
	// if parser.Match(TokenTypes.PRINT) {
	// 	return parser.PrintStatement()
	// }
	// if parser.Match(TokenTypes.RETURN) {
	// 	return parser.ReturnStatement()
	// }
	// if parser.Match(TokenTypes.WHILE) {
	// 	return parser.WhileStatement()
	// }
	// if parser.Match(TokenTypes.LEFT_BRACE) {
	// 	return iql.Block(parser.block())
	// }

	//return parser.ExpressionStatement()

	return parser.expressionStatement()
}

func (parser *Parser) expressionStatement() iql.Stmt {
	expr := parser.expression()

	return NewExpression(expr)
}

func (parser *Parser) match(types ...TokenType) bool {
	for _, t := range types {
		if parser.check(t) {
			parser.advance()
			return true
		}
	}

	return false
}

func (parser *Parser) consume(tokenType TokenType, message string) (Token, error) {
	if parser.check(tokenType) {
		return parser.advance(), nil
	}

	return Token{}, errors.New(message)
}

func (parser *Parser) check(t TokenType) bool {
	if parser.isAtEnd() {
		return false
	}

	return parser.peek().tokenType == t
}

func (parser *Parser) advance() Token {
	if !parser.isAtEnd() {
		parser.current++
	}

	return parser.previous()
}

func (parser *Parser) isAtEnd() bool {
	return parser.peek().tokenType == TokenTypes.EOF

}

func (parser *Parser) peek() Token {
	return parser.tokens[parser.current]
}

func (parser *Parser) previous() Token {
	return parser.tokens[parser.current-1]
}

func (parser *Parser) primary() (iql.Expr, error) {
	if parser.match(TokenTypes.NUMBER, TokenTypes.STRING) {
		return iql.NewLiteralExpr(parser.previous().literal), nil
	}

	if parser.match(TokenTypes.IDENTIFIER) {
		return iql.Variable(parser.previous())
	}

	if parser.match(TokenTypes.LEFT_PAREN) {
		expr := parser.expression()
		parser.consume(TokenTypes.RIGHT_PAREN, `Expect ')' after expression.`)

		return iql.Grouping(expr)
	}

	return nil, errors.New("expected expression")
}

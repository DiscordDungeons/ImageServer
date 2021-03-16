package iqlSchema

import "fmt"

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   Literal
	Line      int
}

func (token Token) String() string {
	return fmt.Sprintf("[%d] %s %s", token.TokenType, token.Lexeme, token.Literal)
}

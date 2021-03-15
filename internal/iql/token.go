package iql

import "fmt"

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   Literal
	line      int
}

func (token Token) String() string {
	return fmt.Sprintf("%d %s %s", token.tokenType, token.lexeme, token.literal)
}

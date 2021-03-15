package iql

type TokenType = int

type tokenList struct {
	// Single char tokens
	LEFT_PAREN  TokenType
	RIGHT_PAREN TokenType
	LEFT_BRACE  TokenType
	RIGHT_BRACE TokenType

	COMMA     TokenType
	DOT       TokenType
	MINUS     TokenType
	PLUS      TokenType
	SEMICOLON TokenType
	SLASH     TokenType
	STAR      TokenType
	DOLLAR    TokenType

	// One or two character tokens

	BANG       TokenType
	BANG_EQUAL TokenType

	EQUAL       TokenType
	EQUAL_EQUAL TokenType

	GREATER       TokenType
	GREATER_EQUAL TokenType

	LESS       TokenType
	LESS_EQUAL TokenType

	// Literals

	IDENTIFIER TokenType
	STRING     TokenType
	NUMBER     TokenType

	// Keywords

	AND   TokenType
	LOAD  TokenType
	IMAGE TokenType

	FROM TokenType
	URL  TokenType
	AS   TokenType

	GENERATE TokenType

	WITH TokenType
	DO   TokenType

	SET       TokenType
	GRAYSCALE TokenType
	TO        TokenType

	EOF TokenType
}

var TokenTypes = &tokenList{
	// Single char tokens
	LEFT_PAREN:  0,
	RIGHT_PAREN: 1,
	LEFT_BRACE:  2,
	RIGHT_BRACE: 3,

	COMMA:     4,
	DOT:       5,
	MINUS:     6,
	PLUS:      7,
	SEMICOLON: 8,
	SLASH:     9,
	STAR:      10,
	DOLLAR:    11,

	// One or two character tokens

	BANG:       12,
	BANG_EQUAL: 13,
	
	EQUAL:       14,
	EQUAL_EQUAL: 15,

	GREATER:       16,
	GREATER_EQUAL: 17,

	LESS:       18,
	LESS_EQUAL: 19,

	// Literals

	IDENTIFIER: 20,
	STRING:     21,
	NUMBER:     22,

	// Keywords

	AND:   23,
	LOAD:  24,
	IMAGE: 25,

	FROM: 26,
	URL:  27,
	AS:   28,

	GENERATE: 29,

	WITH: 30,
	DO:   31,

	SET:       32,
	GRAYSCALE: 33,
	TO:        34,

	EOF: 35,
}

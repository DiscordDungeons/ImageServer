package iql

type TokenType int

const (
	// Single char tokens

	LEFT_PAREN TokenType = iota
	RIGHT_PAREN, LEFT_BRACE, RIGHT_BRACE,
	COMMA, DOT, MINUS, PLUS, SEMICOLON, SLASH, STAR,
	DOLLAR,

	// One or two character tokens

	BANG, BANG_EQUAL,
	EQUAL, EQUAL_EQUAL,
	GREATER, GREATER_EQUAL,
	LESS, LESS_EQUAL

	// Literals

	IDENTIFIER, STRING, NUMBER

	// Keywords

	AND, LOAD, IMAGE,
	
	FROM, URL, AS
	
	GENERATE
	
	WITH, DO
	
	SET, GRAYSCALE, TO

	STRING,
)

package iql

import (
	"errors"
	"fmt"
	"strconv"
)

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int

	keywords map[string]TokenType
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		keywords: map[string]TokenType{
			"LOAD":      TokenTypes.LOAD,
			"IMAGE":     TokenTypes.IMAGE,
			"FROM":      TokenTypes.FROM,
			"URL":       TokenTypes.URL,
			"AS":        TokenTypes.AS,
			"GENERATE":  TokenTypes.GENERATE,
			"WITH":      TokenTypes.WITH,
			"SET":       TokenTypes.SET,
			"GRAYSCALE": TokenTypes.GRAYSCALE,
		},
	}
}

func (scanner *Scanner) IsAtEnd() bool {
	return scanner.current >= len(scanner.source)
}

func (scanner *Scanner) Advance() rune {
	scanner.current++

	return []rune(scanner.source)[scanner.current-1]
}

func (scanner *Scanner) AddToken(tokenType TokenType, literal Literal) {
	text := scanner.source[scanner.start:scanner.current]

	scanner.tokens = append(scanner.tokens, Token{tokenType: tokenType, lexeme: text, literal: literal, line: scanner.line})
}

func (scanner *Scanner) Match(expected rune) bool {
	if scanner.IsAtEnd() {
		return false
	}

	currentChar := []rune(scanner.source)[scanner.current]

	if currentChar != expected {
		return false
	}

	scanner.current++

	return true
}

func (scanner *Scanner) Peek() rune {
	if scanner.IsAtEnd() {
		return rune(0)
	}

	return []rune(scanner.source)[scanner.current]
}

func (scanner *Scanner) String() error {
	for scanner.Peek() != '"' && !scanner.IsAtEnd() {
		if scanner.Peek() == '\n' {
			scanner.line++
		}
		scanner.Advance()
	}

	if scanner.IsAtEnd() {
		return errors.New("unterminated string")
	}

	// The closing "
	scanner.Advance()

	// Trim surrounding quotes
	value := string([]rune(scanner.source)[scanner.start+1 : scanner.current-1])

	newLiteral := Literal{strVal: value, lType: LiteralTypes.STRING_LITERAL}

	scanner.AddToken(TokenTypes.STRING, newLiteral)

	return nil
}

func (scanner *Scanner) IsDigit(c string) bool {
	n, err := strconv.Atoi(c)

	if err != nil {
		return false
	}

	return n >= 0 && n <= 9
}

func (scanner *Scanner) IsRuneDigit(c rune) bool {
	n := int(c)

	return n >= 0 && n <= 9
}

func (scanner *Scanner) PeekNext() rune {
	if scanner.current+1 > len(scanner.source) {
		return rune(0)
	}

	return []rune(scanner.source)[scanner.current+1]
}

func (scanner *Scanner) Number() error {
	for scanner.IsRuneDigit(scanner.Peek()) {
		scanner.Advance()
	}

	// Look for a fractional part
	if scanner.Peek() == '.' && scanner.IsRuneDigit(scanner.PeekNext()) {
		// Consume the "."

		scanner.Advance()

		for scanner.IsRuneDigit(scanner.Peek()) {
			scanner.Advance()
		}
	}

	f, err := strconv.ParseFloat(scanner.source[scanner.start:scanner.current], 64)

	if err != nil {
		return errors.New("invalid float")
	}

	newLiteral := Literal{float64Val: f, lType: LiteralTypes.FLOAT64_LITERAL}

	scanner.AddToken(TokenTypes.NUMBER, newLiteral)

	return nil
}

func (scanner *Scanner) IsAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func (scanner *Scanner) IsAlphaNumberic(c rune) bool {
	return scanner.IsAlpha(c) || scanner.IsRuneDigit(c)
}

func (scanner *Scanner) Identifier() {
	for scanner.IsAlphaNumberic(scanner.Peek()) {
		scanner.Advance()
	}

	println("scanner.source: " + scanner.source)

	text := scanner.source[scanner.start:scanner.current]

	fmt.Println(fmt.Printf("Text is %s | string: %s", text, string(text)))

	fmt.Println("Keywords ", scanner.keywords)

	tokenType, ok := scanner.keywords[text]

	fmt.Printf("tokenType: %d OK: %t\n", tokenType, ok)

	if !ok {
		tokenType = TokenTypes.IDENTIFIER
	}

	scanner.AddToken(tokenType, Literal{})
}

func (scanner *Scanner) ScanToken() error {
	c := scanner.Advance()

	fmt.Println("C: " + string(c))

	switch c {
	case '(':
		scanner.AddToken(TokenTypes.LEFT_PAREN, Literal{})
	case ')':
		scanner.AddToken(TokenTypes.RIGHT_PAREN, Literal{})
	case '{':
		scanner.AddToken(TokenTypes.LEFT_BRACE, Literal{})
	case '}':
		scanner.AddToken(TokenTypes.RIGHT_BRACE, Literal{})
	case ',':
		scanner.AddToken(TokenTypes.COMMA, Literal{})
	case '.':
		scanner.AddToken(TokenTypes.DOT, Literal{})
	case '-':
		scanner.AddToken(TokenTypes.MINUS, Literal{})
	case '+':
		scanner.AddToken(TokenTypes.PLUS, Literal{})
	case ';':
		scanner.AddToken(TokenTypes.SEMICOLON, Literal{})
	case '*':
		scanner.AddToken(TokenTypes.STAR, Literal{})
	case '$':
		scanner.AddToken(TokenTypes.DOLLAR, Literal{})
	case '!':
		scanner.AddToken((map[bool]TokenType{true: TokenTypes.BANG_EQUAL, false: TokenTypes.BANG})[scanner.Match('=')], Literal{})
	case '=':
		scanner.AddToken((map[bool]TokenType{true: TokenTypes.EQUAL_EQUAL, false: TokenTypes.EQUAL})[scanner.Match('=')], Literal{})
	case '<':
		scanner.AddToken((map[bool]TokenType{true: TokenTypes.LESS_EQUAL, false: TokenTypes.LESS})[scanner.Match('=')], Literal{})
	case '>':
		scanner.AddToken((map[bool]TokenType{true: TokenTypes.GREATER_EQUAL, false: TokenTypes.GREATER})[scanner.Match('=')], Literal{})
	case '/':
		if scanner.Match('/') {
			// A comment goes until the end of the line.
			for scanner.Peek() != '\n' && !scanner.IsAtEnd() {
				scanner.Advance()
			}
		} else {
			scanner.AddToken(TokenTypes.SLASH, Literal{})
		}

	case ' ':
	case '\r':
	case '\t':
		// Ignore whitespace.

	case '\n':
		scanner.line++

	case '"':
		scanner.String()

	default:
		if scanner.IsRuneDigit(c) {
			scanner.Number()
		} else if scanner.IsAlpha(c) {
			scanner.Identifier()
		} else {
			return errors.New("unexpected character")
		}
	}

	return nil
}

func (scanner *Scanner) ScanTokens() []Token {
	for !scanner.IsAtEnd() {
		scanner.start = scanner.current
		scanner.ScanToken()
	}

	scanner.tokens = append(scanner.tokens, Token{
		tokenType: TokenTypes.EOF,
		lexeme:    "",
		literal:   Literal{},
		line:      scanner.line,
	})

	return scanner.tokens
}

package iql

import (
	"errors"
	"strconv"

	iqlSchema "discorddungeons.me/imageserver/iql/schema"
)

type Scanner struct {
	source  string
	tokens  []iqlSchema.Token
	start   int
	current int
	line    int

	keywords map[string]iqlSchema.TokenType
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		keywords: map[string]iqlSchema.TokenType{
			"LOAD":      iqlSchema.TokenTypes.LOAD,
			"IMAGE":     iqlSchema.TokenTypes.IMAGE,
			"FROM":      iqlSchema.TokenTypes.FROM,
			"URL":       iqlSchema.TokenTypes.URL,
			"AS":        iqlSchema.TokenTypes.AS,
			"GENERATE":  iqlSchema.TokenTypes.GENERATE,
			"WITH":      iqlSchema.TokenTypes.WITH,
			"SET":       iqlSchema.TokenTypes.SET,
			"GRAYSCALE": iqlSchema.TokenTypes.GRAYSCALE,
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

func (scanner *Scanner) AddToken(tokenType iqlSchema.TokenType, literal iqlSchema.Literal) {
	text := scanner.source[scanner.start:scanner.current]

	scanner.tokens = append(scanner.tokens, iqlSchema.Token{TokenType: tokenType, Lexeme: text, Literal: literal, Line: scanner.line})
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

	newLiteral := iqlSchema.Literal{StrVal: value, LType: iqlSchema.LiteralTypes.STRING_LITERAL}

	scanner.AddToken(iqlSchema.TokenTypes.STRING, newLiteral)

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

	newLiteral := iqlSchema.Literal{Float64Val: f, LType: iqlSchema.LiteralTypes.FLOAT64_LITERAL}

	scanner.AddToken(iqlSchema.TokenTypes.NUMBER, newLiteral)

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

	text := scanner.source[scanner.start:scanner.current]

	tokenType, ok := scanner.keywords[text]

	if !ok {
		tokenType = iqlSchema.TokenTypes.IDENTIFIER
	}

	scanner.AddToken(tokenType, iqlSchema.Literal{})
}

func (scanner *Scanner) ScanToken() error {
	c := scanner.Advance()

	switch c {
	case '(':
		scanner.AddToken(iqlSchema.TokenTypes.LEFT_PAREN, iqlSchema.Literal{})
	case ')':
		scanner.AddToken(iqlSchema.TokenTypes.RIGHT_PAREN, iqlSchema.Literal{})
	case '{':
		scanner.AddToken(iqlSchema.TokenTypes.LEFT_BRACE, iqlSchema.Literal{})
	case '}':
		scanner.AddToken(iqlSchema.TokenTypes.RIGHT_BRACE, iqlSchema.Literal{})
	case ',':
		scanner.AddToken(iqlSchema.TokenTypes.COMMA, iqlSchema.Literal{})
	case '.':
		scanner.AddToken(iqlSchema.TokenTypes.DOT, iqlSchema.Literal{})
	case '-':
		scanner.AddToken(iqlSchema.TokenTypes.MINUS, iqlSchema.Literal{})
	case '+':
		scanner.AddToken(iqlSchema.TokenTypes.PLUS, iqlSchema.Literal{})
	case ';':
		scanner.AddToken(iqlSchema.TokenTypes.SEMICOLON, iqlSchema.Literal{})
	case '*':
		scanner.AddToken(iqlSchema.TokenTypes.STAR, iqlSchema.Literal{})
	case '$':
		scanner.AddToken(iqlSchema.TokenTypes.DOLLAR, iqlSchema.Literal{})
	case '!':
		scanner.AddToken((map[bool]iqlSchema.TokenType{true: iqlSchema.TokenTypes.BANG_EQUAL, false: iqlSchema.TokenTypes.BANG})[scanner.Match('=')], iqlSchema.Literal{})
	case '=':
		scanner.AddToken((map[bool]iqlSchema.TokenType{true: iqlSchema.TokenTypes.EQUAL_EQUAL, false: iqlSchema.TokenTypes.EQUAL})[scanner.Match('=')], iqlSchema.Literal{})
	case '<':
		scanner.AddToken((map[bool]iqlSchema.TokenType{true: iqlSchema.TokenTypes.LESS_EQUAL, false: iqlSchema.TokenTypes.LESS})[scanner.Match('=')], iqlSchema.Literal{})
	case '>':
		scanner.AddToken((map[bool]iqlSchema.TokenType{true: iqlSchema.TokenTypes.GREATER_EQUAL, false: iqlSchema.TokenTypes.GREATER})[scanner.Match('=')], iqlSchema.Literal{})
	case '/':
		if scanner.Match('/') {
			// A comment goes until the end of the line.
			for scanner.Peek() != '\n' && !scanner.IsAtEnd() {
				scanner.Advance()
			}
		} else {
			scanner.AddToken(iqlSchema.TokenTypes.SLASH, iqlSchema.Literal{})
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

func (scanner *Scanner) ScanTokens() []iqlSchema.Token {
	for !scanner.IsAtEnd() {
		scanner.start = scanner.current
		scanner.ScanToken()
	}

	scanner.tokens = append(scanner.tokens, iqlSchema.Token{
		TokenType: iqlSchema.TokenTypes.EOF,
		Lexeme:    "",
		Literal:   iqlSchema.Literal{},
		Line:      scanner.line,
	})

	return scanner.tokens
}

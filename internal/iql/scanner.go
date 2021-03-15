package iql

import (
	"errors"
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

func New(source string) Scanner {
	return Scanner{
		source: source,
		keywords: map[string]TokenType{
			"as": AS,
		},
	}
}

func (scanner *Scanner) isAtEnd() bool {
	return scanner.current >= len(scanner.source)
}

func (scanner *Scanner) advance() rune {
	scanner.current++

	return []rune(scanner.source)[scanner.current-1]
}

func (scanner *Scanner) addToken(tokenType TokenType, literal Literal) {
	text := scanner.source[scanner.start:scanner.current]

	scanner.tokens = append(scanner.tokens, Token{tokenType: tokenType, lexeme: text, literal: literal, line: scanner.line})
}

func (scanner *Scanner) match(expected string) bool {
	if scanner.isAtEnd() {
		return false
	}

	currentChar := string([]rune(scanner.source)[scanner.current])

	if currentChar != expected {
		return false
	}

	scanner.current++

	return true
}

func (scanner *Scanner) peek() rune {
	if scanner.isAtEnd() {
		return rune(0)
	}

	return []rune(scanner.source)[scanner.current]
}

func (scanner *Scanner) string() error {
	for scanner.peek() != '"' && !scanner.isAtEnd() {
		if scanner.peek() == '\n' {
			scanner.line++
		}
		scanner.advance()
	}

	if scanner.isAtEnd() {
		return errors.New("Unterminated string.")
	}

	// The closing "
	scanner.advance()

	// Trim surrounding quotes
	value := string([]rune(scanner.source)[scanner.start+1 : scanner.current-1])

	newLiteral := Literal{strVal: value}

	scanner.addToken(STRING, newLiteral)

	return nil
}

func (scanner *Scanner) isDigit(c string) bool {
	n, err := strconv.Atoi(c)

	if err != nil {
		return false
	}

	return n >= 0 && n <= 9
}

func (scanner *Scanner) isRuneDigit(c rune) bool {
	n := int(c)

	return n >= 0 && n <= 9
}

func (scanner *Scanner) peekNext() rune {
	if scanner.current+1 > len(scanner.source) {
		return rune(0)
	}

	return []rune(scanner.source)[scanner.current+1]
}

func (scanner *Scanner) number() error {
	for scanner.isRuneDigit(scanner.peek()) {
		scanner.advance()
	}

	// Look for a fractional part
	if scanner.peek() == '.' && scanner.isRuneDigit(scanner.peekNext()) {
		// Consume the "."

		scanner.advance()

		for scanner.isRuneDigit(scanner.peek()) {
			scanner.advance()
		}
	}

	f, err := strconv.ParseFloat(scanner.source[scanner.start:scanner.current], 64)

	if err != nil {
		return errors.New("Invalid float.")
	}

	newLiteral := Literal{float64Val: f}

	scanner.addToken(NUMBER, newLiteral)

	return nil
}

func (scanner *Scanner) isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func (scanner *Scanner) isAlphaNumberic(c rune) bool {
	return scanner.isAlpha(c) || scanner.isRuneDigit(c)
}

func (scanner *Scanner) identifier() {
	for scanner.isAlphaNumberic(scanner.peek()) {
		scanner.advance()
	}

	text := scanner.source[scanner.start:scanner.current]

	tokenType, ok := scanner.keywords[text]

	if !ok {
		tokenType = IDENTIFIER
	}

	scanner.addToken(tokenType, Literal{})
}

func (scanner *Scanner) scanToken () {
	c := scanner.advance()

	switch (c) {
		case '(': scanner.addToken(LEFT_PAREN, null); break
		case ')': scanner.addToken(RIGHT_PAREN, null); break
		case '{': scanner.addToken(LEFT_BRACE, null); break
		case '}': scanner.addToken(RIGHT_BRACE, null); break
		case ',': scanner.addToken(COMMA, null); break
		case '.': scanner.addToken(DOT, null); break
		case '-': scanner.addToken(MINUS, null); break
		case '+': scanner.addToken(PLUS, null); break
		case ';': scanner.addToken(SEMICOLON, null); break
		case '*': scanner.addToken(STAR, null); break
		case '$': scanner.addToken(DOLLAR, null); break
		case '!':
			scanner.addToken(scanner.match('=') ? BANG_EQUAL : BANG, null);
			break
		case '=':
			scanner.addToken(scanner.match('=') ? TokenType.EQUAL_EQUAL : TokenType.EQUAL, null);
			break
		case '<':
			scanner.addToken(scanner.match('=') ? TokenType.LESS_EQUAL : TokenType.LESS, null);
			break
		case '>':
			scanner.addToken(scanner.match('=') ? TokenType.GREATER_EQUAL : TokenType.GREATER, null);
			break;
		case '/':
			if (scanner.match('/')) {
				// A comment goes until the end of the line.
				while (scanner.peek() != '\n' && !scanner.isAtEnd) scanner.advance()
			} else {
				scanner.addToken(TokenType.SLASH, null)
			}
			break
		case ' ':
		case '\r':
		case '\t':
			// Ignore whitespace.
			break
	
		case '\n':
			scanner.line++
			break
		
		case '"':
			scanner.string()
			break
	
		default:
			if (scanner.isDigit(c)) {
				scanner.number()
			} else if (scanner.isAlpha(c)) {
				scanner.identifier()
			} else {
				TagScript.error(scanner.line, "Unexpected character")
			}
	}
}
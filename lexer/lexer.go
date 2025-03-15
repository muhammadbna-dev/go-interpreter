package lexer

import "monkey/token"

type Lexer struct {
	input           string
	currentPosition int
	readPosition    int
	currentChar     byte
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.readPosition]
	}
	l.currentPosition = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.currentChar {
	case '=':
		if l.peekChar() == '=' {
			currentChar := l.currentChar
			l.readChar()
			literal := string(currentChar) + string(l.currentChar)
			tok = token.Token{Literal: literal, Type: token.EQ}
		} else {
			tok = newToken(token.ASSIGN, l.currentChar)
		}
	case '+':
		tok = newToken(token.PLUS, l.currentChar)
	case '-':
		tok = newToken(token.MINUS, l.currentChar)
	case '!':
		if l.peekChar() == '=' {
			currentChar := l.currentChar
			l.readChar()
			literal := string(currentChar) + string(l.currentChar)
			tok = token.Token{Literal: literal, Type: token.NOT_EQ}
		} else {
			tok = newToken(token.BANG, l.currentChar)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.currentChar)
	case '/':
		tok = newToken(token.SLASH, l.currentChar)
	case '<':
		tok = newToken(token.LT, l.currentChar)
	case '>':
		tok = newToken(token.GT, l.currentChar)
	case ',':
		tok = newToken(token.COMMA, l.currentChar)
	case ';':
		tok = newToken(token.SEMICOLON, l.currentChar)
	case '(':
		tok = newToken(token.LPAREN, l.currentChar)
	case ')':
		tok = newToken(token.RPAREN, l.currentChar)
	case '{':
		tok = newToken(token.LBRACE, l.currentChar)
	case '}':
		tok = newToken(token.RBRACE, l.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.currentChar) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.currentChar)
		}
	}
	l.readChar()
	return tok
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.currentPosition
	for isLetter(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.currentPosition]
}

func (l *Lexer) readNumber() string {
	position := l.currentPosition
	for isDigit(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.currentPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

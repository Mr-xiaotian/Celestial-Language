package lexer

import "github.com/Mr-xiaotian/Celestial-Language/Monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}

}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch {
	case l.ch == '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: "=="}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case l.ch == '+':
		tok = newToken(token.PLUS, l.ch)
	case l.ch == '-':
		tok = newToken(token.MINUS, l.ch)
	case l.ch == '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case l.ch == '/':
		tok = newToken(token.SLASH, l.ch)
	case l.ch == '*':
		tok = newToken(token.ASTERISK, l.ch)
	case l.ch == '<':
		tok = newToken(token.LT, l.ch)
	case l.ch == '>':
		tok = newToken(token.RT, l.ch)
	case l.ch == ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case l.ch == ',':
		tok = newToken(token.COMMA, l.ch)
	case l.ch == '(':
		tok = newToken(token.LPAREN, l.ch)
	case l.ch == ')':
		tok = newToken(token.RPAREN, l.ch)
	case l.ch == '{':
		tok = newToken(token.LBRACE, l.ch)
	case l.ch == '}':
		tok = newToken(token.RBRACE, l.ch)
	case l.ch == 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case isDigit(l.ch):
		tok.Type = token.INT
		tok.Literal = l.readNumber()
		return tok
	case isLetter(l.ch):
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
		return tok
	default:
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

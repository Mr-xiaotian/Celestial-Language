package lexer

import "github.com/Mr-xiaotian/Celestial-Language/Monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	Line         int
	Column       int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	l.Line = 1
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
	l.Column += 1
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
		colunm := l.Column
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.NewStringToken(token.EQ, "==", l.Line, colunm)
		} else {
			tok = token.NewByteToken(token.ASSIGN, l.ch, l.Line, l.Column)
		}
	case l.ch == '+':
		tok = token.NewByteToken(token.PLUS, l.ch, l.Line, l.Column)
	case l.ch == '-':
		tok = token.NewByteToken(token.MINUS, l.ch, l.Line, l.Column)
	case l.ch == '!':
		colunm := l.Column
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.NewStringToken(token.NOT_EQ, "!=", l.Line, colunm)
		} else {
			tok = token.NewByteToken(token.BANG, l.ch, l.Line, l.Column)
		}
	case l.ch == '/':
		tok = token.NewByteToken(token.SLASH, l.ch, l.Line, l.Column)
	case l.ch == '*':
		tok = token.NewByteToken(token.ASTERISK, l.ch, l.Line, l.Column)
	case l.ch == '<':
		tok = token.NewByteToken(token.LT, l.ch, l.Line, l.Column)
	case l.ch == '>':
		tok = token.NewByteToken(token.RT, l.ch, l.Line, l.Column)
	case l.ch == ';':
		tok = token.NewByteToken(token.SEMICOLON, l.ch, l.Line, l.Column)
	case l.ch == ',':
		tok = token.NewByteToken(token.COMMA, l.ch, l.Line, l.Column)
	case l.ch == '(':
		tok = token.NewByteToken(token.LPAREN, l.ch, l.Line, l.Column)
	case l.ch == ')':
		tok = token.NewByteToken(token.RPAREN, l.ch, l.Line, l.Column)
	case l.ch == '{':
		tok = token.NewByteToken(token.LBRACE, l.ch, l.Line, l.Column)
	case l.ch == '}':
		tok = token.NewByteToken(token.RBRACE, l.ch, l.Line, l.Column)
	case l.ch == '"':
		colunm := l.Column
		tokLiteral := l.readString()
		tok = token.NewStringToken(token.STRING, tokLiteral, l.Line, colunm)
		return tok
	case l.ch == 0:
		tok = token.NewStringToken(token.EOF, "", l.Line, l.Column)
	case isDigit(l.ch):
		colunm := l.Column
		tok = token.NewStringToken(token.INT, l.readNumber(), l.Line, colunm)
		return tok
	case isLetter(l.ch):
		colunm := l.Column
		tokLiteral := l.readIdentifier()
		tokType := token.LookupIdent(tokLiteral)
		tok = token.NewStringToken(tokType, tokLiteral, l.Line, colunm)
		return tok
	default:
		tok = token.NewByteToken(token.ILLEGAL, l.ch, l.Line, l.Column)
	}

	l.readChar()
	return tok
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

func (l *Lexer) readString() string {
	start := l.position + 1 // 字符串内容从起始引号后一位开始
	l.readChar()            // 跳过起始引号

	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}

	str := l.input[start:l.position]

	if l.ch == '"' {
		l.readChar() // 跳过结束引号
	}

	return str
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
		if l.ch == '\n' {
			l.Line += 1
			l.Column = 1
		}
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

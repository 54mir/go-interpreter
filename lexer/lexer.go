package lexer

import "go-interpreter/token"

type Lexer struct {
	input           string
	currentPosition int
	readPosition    int
	ch              byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.EQ, string(ch)+string(l.ch))
		} else {
			tok = token.New(token.ASSIGN, string(l.ch))
		}
	case ';':
		tok = token.New(token.SEMICOLON, string(l.ch))
	case '(':
		tok = token.New(token.LPAREN, string(l.ch))
	case ')':
		tok = token.New(token.RPAREN, string(l.ch))
	case ',':
		tok = token.New(token.COMMA, string(l.ch))
	case '+':
		tok = token.New(token.PLUS, string(l.ch))
	case '-':
		tok = token.New(token.MINUS, string(l.ch))
	case '{':
		tok = token.New(token.LBRACE, string(l.ch))
	case '}':
		tok = token.New(token.RBRACE, string(l.ch))
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.New(token.NOT_EQ, string(ch)+string(l.ch))
		} else {
			tok = token.New(token.BANG, string(l.ch))
		}
	case '*':
		tok = token.New(token.ASTERISK, string(l.ch))
	case '/':
		tok = token.New(token.SLASH, string(l.ch))
	case '<':
		tok = token.New(token.LT, string(l.ch))
	case '>':
		tok = token.New(token.GT, string(l.ch))
	case 0:
		tok = token.New(token.EOF, "")
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.New(token.ILLEGAL, string(l.ch))
		}

	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.currentPosition
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.currentPosition]
}

func (l *Lexer) readNumber() string {
	position := l.currentPosition
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.currentPosition]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.currentPosition = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

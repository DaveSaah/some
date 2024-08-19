// package lexer takes source code as input and outputs tokens that represents the
// source code
package lexer

import "github.com/DaveSaah/some/token"

type Lexer struct {
	input        string
	position     int  // current char position in input
	readPosition int  // next char position after current char in input
	ch           byte // current char under examination
}

// New creates a new lexer from an input
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // set l.ch to the first character
	return l
}

// newToken creates a new token from a token type and a character
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

// readChar reads a character from the lexer input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // set to ASCII NUL
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken returns the next token from a lexer input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return (ch >= '0' && ch <= '9')
}

// readIdentifier reads an identifier from a lexer's input string
func (l *Lexer) readIdentifier() string {
	position := l.position // track current position

	// check if the first character is a letter
	if isLetter(l.ch) {
		l.readChar()
	}

	// check if the identifier satisfies its rule:
	// Can have a letter or digit
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// eatWhitespace removes all whitespaces in lexer's input
func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumberToken() token.Token {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()

		if isLetter(l.input[l.readPosition]) {
			return token.Token{
				Type:    token.ILLEGAL,
				Literal: l.input[position:l.readPosition],
			}
		}
	}

	return token.Token{
		Type:    token.INT,
		Literal: l.input[position:l.position],
	}
}

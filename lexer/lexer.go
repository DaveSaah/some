// package lexer takes source code as input and outputs tokens that represents the
// source code
package lexer

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


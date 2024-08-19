// package token defines small categorisable data structures that are
// fed into a parser
package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // signifies an unknown token
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"

	// Data types
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

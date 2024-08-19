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
	MINUS  = "-"

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

// define keyword symbol table
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LoopupIdentifier checks whether the identiier is a
// variable or a keyword
func LoopupIdentifier(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}

	return IDENTIFIER
}

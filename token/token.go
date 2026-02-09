// package token defines small categorisable data structures that are fed into
// a parser
package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // signifies an unknown token
	EOF     = "EOF"     // End of File

	IDENTIFIER = "IDENTIFIER"

	// Data types
	INT = "INT"

	// Operators
	ASSIGN       = "="
	PLUS         = "+"
	MINUS        = "-"
	BANG         = "!"
	ASTERISK     = "*"
	SLASH        = "/"
	LESS_THAN    = "<"
	GREATER_THAN = ">"
	EQUALS       = "=="
	NOT_EQUALS   = "!="

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
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

// define keyword symbol table
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

// LoopupIdentifier checks whether the identifier is a variable or a keyword
func LoopupIdentifier(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}

	return IDENTIFIER
}

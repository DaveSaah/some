package lexer_test

import (
	"testing"

	"github.com/DaveSaah/some/lexer"
	"github.com/DaveSaah/some/token"
)

func TestNextTokenBasic(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		switch {
		case tok.Type != tt.expectedType:
			t.Fatalf(
				"test[%d]: token type wrong. expected='%s', got='%s'",
				i,
				tt.expectedType, tok.Type,
			)

		case tok.Literal != tt.expectedLiteral:
			t.Fatalf(
				"test[%d]: token literal wrong. expected='%s', got='%s'",
				i,
				tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenComplex(t *testing.T) {
	input := `let five = 5;
  let ten = 10;

  let add = fn(x, y) {
    x + y;
  };

  let result = add(five, ten);
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		switch {
		case tok.Type != tt.expectedType:
			t.Fatalf(
				"test[%d]: token type wrong. expected='%s', got='%s'",
				i,
				tt.expectedType, tok.Type,
			)

		case tok.Literal != tt.expectedLiteral:
			t.Fatalf(
				"test[%d]: token literal wrong. expected='%s', got='%s'",
				i,
				tt.expectedLiteral, tok.Literal,
			)
		}
	}
}

func TestFailedToken(t *testing.T) {
	input := `let some_one = 5;
  let test1 = -20;
  let 1five = 5;
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "some_one"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "test1"},
		{token.ASSIGN, "="},
		{token.MINUS, "-"},
		{token.INT, "20"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.ILLEGAL, "1f"},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		switch {
		case tok.Type != tt.expectedType:
			t.Fatalf(
				"test[%d]: token type wrong. expected='%s', got='%s'",
				i,
				tt.expectedType, tok.Type,
			)

		case tok.Literal != tt.expectedLiteral:
			t.Fatalf(
				"test[%d]: token literal wrong. expected='%s', got='%s'",
				i,
				tt.expectedLiteral, tok.Literal,
			)
		}
	}
}

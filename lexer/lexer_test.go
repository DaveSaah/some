package lexer_test

import (
	"testing"

	"github.com/DaveSaah/some/lexer"
	"github.com/DaveSaah/some/token"
)

func TestNextTokenBasic(t *testing.T) {
	input := "=+(){},;-!*/<>"

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
		{token.MINUS, "-"},
		{token.BANG, "!"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.LESS_THAN, "<"},
		{token.GREATER_THAN, ">"},
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

  !-/*5;
  5 < 10 > 5;

  if (5 < 10) {
    return true;
  } else {
    return false;
  }

  10 == 10;
  10 != 9;
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
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LESS_THAN, "<"},
		{token.INT, "10"},
		{token.GREATER_THAN, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LESS_THAN, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQUALS, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQUALS, "!="},
		{token.INT, "9"},
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

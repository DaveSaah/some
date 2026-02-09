package parser

import (
	"testing"

	"github.com/DaveSaah/some/ast"
	"github.com/DaveSaah/some/lexer"
)

func TestLetStmts(t *testing.T) {
	input := `let x = 5;
	let y = 10;
	let fizzbuzz = 16;`

	// input := `let x 5;
	// let = 10;
	// let 16;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Error("ParseProgram returned nil")
		return
	}

	checkParserErrors(t, p)

	if len(program.Stmts) != 3 {
		t.Errorf("Expected 3 statements. got=%d", len(program.Stmts))
		return
	}

	tests := []struct {
		expectedIdent string
	}{
		{"x"},
		{"y"},
		{"fizzbuzz"},
	}

	for i, tt := range tests {
		stmt := program.Stmts[i]
		if !testLetStmt(t, stmt, tt.expectedIdent) {
			return
		}
	}
}

func testLetStmt(t *testing.T, stmt ast.Stmt, ident string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStmt)
	if !ok {
		t.Errorf("stmt not '*ast.LetStmt'. got=%T", stmt)
		return false
	}

	if letStmt.Var.Value != ident {
		t.Errorf("letStmt.Var.Value not '%s'. got=%s", ident, letStmt.Var.Value)
		return false
	}

	if letStmt.Var.TokenLiteral() != ident {
		t.Errorf("letStmt.Var not '%s'. got=%s", ident, letStmt.Var)
		return false
	}

	return true
}

func TestReturnStmts(t *testing.T) {
	input := `return 5;
	return 10;
	return 16;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Error("ParseProgram returned nil")
		return
	}

	checkParserErrors(t, p)

	if len(program.Stmts) != 3 {
		t.Errorf("Expected 3 statements. got=%d", len(program.Stmts))
		return
	}

	for _, stmt := range program.Stmts {
		returnStmt, ok := stmt.(*ast.ReturnStmt)
		if !ok {
			t.Errorf("stmt not return stmt. got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf(
				"token literal is not 'return'. got=%q",
				returnStmt.TokenLiteral(),
			)
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

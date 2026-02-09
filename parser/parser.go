// package parser parses lexer tokens to confirm if they are valid language
// grammar
package parser

import (
	"fmt"

	"github.com/DaveSaah/some/ast"
	"github.com/DaveSaah/some/lexer"
	"github.com/DaveSaah/some/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	// read two tokens so curToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Stmts = []ast.Stmt{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStmt()
		if stmt != nil {
			program.Stmts = append(program.Stmts, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) parseStmt() ast.Stmt {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStmt()
	case token.RETURN:
		return p.parseReturnStmt()
	default:
		return nil
	}
}

func (p *Parser) tokenTypeErr(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be '%s', got '%s' instead",
		t, p.peekToken.Type,
	)

	p.errors = append(p.errors, msg)
}

func (p *Parser) parseLetStmt() ast.Stmt {
	stmt := &ast.LetStmt{Token: p.curToken}

	if p.peekToken.Type != token.IDENTIFIER {
		p.tokenTypeErr(token.IDENTIFIER)
		return nil
	}

	p.nextToken() // consume 'let' token

	stmt.Var = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekToken.Type != token.ASSIGN {
		p.tokenTypeErr(token.ASSIGN)
		return nil
	}

	p.nextToken() // consume 'assign' token

	// skip expressions
	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStmt() ast.Stmt {
	stmt := &ast.ReturnStmt{Token: p.curToken}

	p.nextToken() // consume 'return' token

	// skip expressions
	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

// package ast defines an abstract syntax tree
package ast

import (
	"github.com/DaveSaah/some/token"
)

type Node interface {
	TokenLiteral() string
}

type Stmt interface {
	Node
	stmtNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Stmts []Stmt
}

func (p *Program) TokenLiteral() string {
	if len(p.Stmts) > 0 {
		return p.Stmts[0].TokenLiteral()
	}

	return ""
}

type Identifier struct {
	Token token.Token // token.Identifier
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type LetStmt struct {
	Token token.Token // token.LET
	Var   *Identifier // variable
	Value Expression  // expression
}

func (ls *LetStmt) stmtNode() {}
func (ls *LetStmt) TokenLiteral() string {
	return ls.Token.Literal
}

type ReturnStmt struct {
	Token token.Token // token.RETURN
	Value Expression  // expression
}

func (rs *ReturnStmt) stmtNode() {}
func (rs *ReturnStmt) TokenLiteral() string {
	return rs.Token.Literal
}

package ast

import (
	"GoParser2/lex"
)

type ForClause struct {
	// forClause: initStmt = simpleStmt? eos expression? eos postStmt = simpleStmt?;
	initStmt   SimpleStmt
	expression *Expression
	postStmt   SimpleStmt
}

func (a *ForClause) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendTreeNode(a.initStmt).AppendString(";")
	cb.AppendTreeNode(a.expression).AppendString(";")
	cb.AppendTreeNode(a.postStmt)
	return cb
}

func (a *ForClause) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ForClause)(nil)

func VisitForClause(lexer *lex.Lexer) *ForClause {
	initStmt := VisitSimpleStmt(lexer)
	VisitEos(lexer)
	expression := VisitExpression(lexer)
	VisitEos(lexer)
	postStmt := VisitSimpleStmt(lexer)

	return &ForClause{initStmt: initStmt, expression: expression, postStmt: postStmt}
}

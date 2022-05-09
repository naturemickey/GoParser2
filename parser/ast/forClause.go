package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type ForClause struct {
	// forClause: initStmt = simpleStmt? eos expression? eos postStmt = simpleStmt?;
	initStmt   SimpleStmt
	expression *Expression
	postStmt   SimpleStmt
}

func (a *ForClause) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*ForClause)(nil)

func VisitForClause(lexer *lex.Lexer) *ForClause {
	initStmt := VisitSimpleStmt(lexer)
	VisitEos(lexer)
	expression := VisitExpression(lexer)
	VisitEos(lexer)
	postStmt := VisitSimpleStmt(lexer)

	return &ForClause{initStmt: initStmt, expression: expression, postStmt: postStmt}
}

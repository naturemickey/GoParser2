package ast

import "GoParser2/lex"

type GoStmt struct {
}

func (g GoStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*GoStmt)(nil)

func VisitGoStmt(lexer *lex.Lexer) *GoStmt {
	panic("todo")
}

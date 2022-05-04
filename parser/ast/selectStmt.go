package ast

import "GoParser2/lex"

type SelectStmt struct {
}

func (s SelectStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*SelectStmt)(nil)

func VisitSelectStmt(lexer *lex.Lexer) *SelectStmt {
	panic("todo")
}

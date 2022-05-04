package ast

import "GoParser2/lex"

type ReturnStmt struct {
}

func (r ReturnStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*ReturnStmt)(nil)

func VisitReturnStmt(lexer *lex.Lexer) *ReturnStmt {
	panic("todo")
}

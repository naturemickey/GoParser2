package ast

import "GoParser2/lex"

type IncDecStmt struct {
}

func (i IncDecStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

func (i IncDecStmt) __SimpleStmt__() {
	//TODO implement me
	panic("implement me")
}

var _ SimpleStmt = (*IncDecStmt)(nil)

func VisitIncDecStmt(lexer *lex.Lexer) *IncDecStmt {
	panic("todo")
}

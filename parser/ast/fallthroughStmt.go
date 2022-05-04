package ast

import "GoParser2/lex"

type FallthroughStmt struct {
}

func (f FallthroughStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*FallthroughStmt)(nil)

func VisitFallthroughStmt(lexer *lex.Lexer) *FallthroughStmt {
	panic("todo")
}

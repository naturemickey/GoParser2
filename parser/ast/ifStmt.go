package ast

import "GoParser2/lex"

type IfStmt struct {
}

func (i IfStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*IfStmt)(nil)

func VisitIfStmt(lexer *lex.Lexer) *IfStmt {
	panic("todo")
}

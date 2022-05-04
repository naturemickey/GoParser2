package ast

import "GoParser2/lex"

type Assignment struct {
}

func (a Assignment) __Statement__() {
	//TODO implement me
	panic("implement me")
}

func (a Assignment) __SimpleStmt__() {
	//TODO implement me
	panic("implement me")
}

var _ SimpleStmt = (*Assignment)(nil)

func VisitAssignment(lexer *lex.Lexer) *Assignment {
	panic("todo")
}

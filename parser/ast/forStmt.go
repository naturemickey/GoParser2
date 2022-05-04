package ast

import "GoParser2/lex"

type ForStmt struct {
}

func (f ForStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*ForStmt)(nil)

func VisitForStmt(lexer *lex.Lexer) *ForStmt {
	panic("todo")
}

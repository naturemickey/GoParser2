package ast

import "GoParser2/lex"

type ContinueStmt struct {
}

func (c ContinueStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*ContinueStmt)(nil)

func VisitContinueStmt(lexer *lex.Lexer) *ContinueStmt {
	panic("todo")
}

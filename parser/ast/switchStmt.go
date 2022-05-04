package ast

import "GoParser2/lex"

type SwitchStmt struct {
}

func (s SwitchStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*SwitchStmt)(nil)

func VisitSwitchStmt(lexer *lex.Lexer) *SwitchStmt {
	panic("todo")
}

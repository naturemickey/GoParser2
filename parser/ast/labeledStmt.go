package ast

import "GoParser2/lex"

type LabeledStmt struct {
}

func (l LabeledStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*LabeledStmt)(nil)

func VisitLabeledStmt(lexer *lex.Lexer) *LabeledStmt {
	panic("todo")
}

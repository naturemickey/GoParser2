package ast

import "GoParser2/lex"

type BreakStmt struct {
}

func (b *BreakStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*BreakStmt)(nil)

func VisitBreakStmt(lexer *lex.Lexer) *BreakStmt {
	panic("todo")
}

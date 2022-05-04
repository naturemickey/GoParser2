package ast

import "GoParser2/lex"

type DeferStmt struct {
}

func (d DeferStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*DeferStmt)(nil)

func VisitDeferStmt(lexer *lex.Lexer) *DeferStmt {
	panic("todo")
}

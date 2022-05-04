package ast

import "GoParser2/lex"

type GotoStmt struct {
}

func (g GotoStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*GotoStmt)(nil)

func VisitGotoStmt(lexer *lex.Lexer) *GotoStmt {
	panic("todo")
}

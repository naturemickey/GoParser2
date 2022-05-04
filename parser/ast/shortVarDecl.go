package ast

import "GoParser2/lex"

type ShortVarDecl struct {
}

func (s ShortVarDecl) __Statement__() {
	//TODO implement me
	panic("implement me")
}

func (s ShortVarDecl) __SimpleStmt__() {
	//TODO implement me
	panic("implement me")
}

var _ SimpleStmt = (*ShortVarDecl)(nil)

func VisitShortVarDecl(lexer *lex.Lexer) *ShortVarDecl {
	panic("todo")
}

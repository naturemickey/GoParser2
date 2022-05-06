package ast

import "GoParser2/lex"

type CompositeLit struct {
}

func (c CompositeLit) __Literal__() {
	//TODO implement me
	panic("implement me")
}

var _ Literal = (*CompositeLit)(nil)

func VisitCompositeLit(lexer *lex.Lexer) *CompositeLit {
	panic("todo")
}

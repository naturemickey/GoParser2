package ast

import "GoParser2/lex"

type FunctionLit struct {
}

func (f FunctionLit) __Literal__() {
	//TODO implement me
	panic("implement me")
}

var _ Literal = (*FunctionLit)(nil)

func VisitFunctionLit(lexer *lex.Lexer) *FunctionLit {
	panic("todo")
}

package ast

import "GoParser2/lex"

type FunctionType struct {
}

func (f FunctionType) __TypeLit__() {
	//TODO implement me
	panic("implement me")
}

var _ TypeLit = (*FunctionType)(nil)

func VisitFunctionType(lexer *lex.Lexer) *FunctionType {
	panic("todo")
}

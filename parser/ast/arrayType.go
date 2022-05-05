package ast

import "GoParser2/lex"

type ArrayType struct {
}

func (a ArrayType) __TypeLit__() {
	//TODO implement me
	panic("implement me")
}

var _ TypeLit = (*ArrayType)(nil)

func VisitArrayType(lexer *lex.Lexer) *ArrayType {
	panic("todo")
}

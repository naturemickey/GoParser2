package ast

import "GoParser2/lex"

type StructType struct {
}

func (s StructType) __TypeLit__() {
	//TODO implement me
	panic("implement me")
}

var _ TypeLit = (*StructType)(nil)

func VisitStructType(lexer *lex.Lexer) *StructType {
	panic("todo")
}

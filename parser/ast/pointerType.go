package ast

import "GoParser2/lex"

type PointerType struct {
}

func (p PointerType) __TypeLit__() {
	//TODO implement me
	panic("implement me")
}

var _ TypeLit = (*PointerType)(nil)

func VisitPointerType(lexer *lex.Lexer) *PointerType {
	panic("todo")
}

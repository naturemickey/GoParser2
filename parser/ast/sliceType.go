package ast

import "GoParser2/lex"

type SliceType struct {
	// sliceType: L_BRACKET R_BRACKET elementType;
}

func (s SliceType) __TypeLit__() {
	//TODO implement me
	panic("implement me")
}

var _ TypeLit = (*SliceType)(nil)

func VisitSliceType(lexer *lex.Lexer) *SliceType {
	panic("todo")
}

package ast

import "GoParser2/lex"

type SliceType struct {
}

func (s SliceType) __TypeLit__() {
	//TODO implement me
	panic("implement me")
}

var _ TypeLit = (*SliceType)(nil)

func VisitSliceType(lexer *lex.Lexer) *SliceType {
	panic("todo")
}

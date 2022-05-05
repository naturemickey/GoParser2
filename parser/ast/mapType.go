package ast

import "GoParser2/lex"

type MapType struct {
}

func (m MapType) __TypeLit__() {
	//TODO implement me
	panic("implement me")
}

var _ TypeLit = (*MapType)(nil)

func VisitMapType(lexer *lex.Lexer) *MapType {
	panic("todo")
}

package ast

import "GoParser2/lex"

type InterfaceType struct {
}

func (i InterfaceType) __TypeLit__() {
	//TODO implement me
	panic("implement me")
}

var _ TypeLit = (*InterfaceType)(nil)

func VisitInterfaceType(lexer *lex.Lexer) *InterfaceType {
	panic("todo")
}

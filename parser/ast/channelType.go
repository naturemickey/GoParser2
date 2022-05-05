package ast

import "GoParser2/lex"

type ChannelType struct {
}

func (c ChannelType) __TypeLit__() {
	//TODO implement me
	panic("implement me")
}

var _ TypeLit = (*ChannelType)(nil)

func VisitChannelType(lexer *lex.Lexer) *ChannelType {
	panic("todo")
}

package ast

import "GoParser2/lex"

type ChannelType struct {
	// channelType: (ch=CHAN | ch_re=CHAN RECEIVE | re_ch=RECEIVE CHAN) elementType;
	// elementType: type_;
}

func (c ChannelType) __TypeLit__() {
	//TODO implement me
	panic("implement me")
}

var _ TypeLit = (*ChannelType)(nil)

func VisitChannelType(lexer *lex.Lexer) *ChannelType {
	panic("todo")
}

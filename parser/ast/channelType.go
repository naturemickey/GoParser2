package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"fmt"
)

type ChannelType struct {
	// channelType: (ch=CHAN | ch_re=CHAN RECEIVE | re_ch=RECEIVE CHAN) elementType;
	// elementType: type_;
	chan_        *lex.Token
	chan_receive *chanReceivePair
	receive_chan *receiveChanPair
	elementType  *Type_
}

func (a *ChannelType) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*ChannelType)(nil)

type chanReceivePair struct {
	chan_   *lex.Token
	receive *lex.Token
}

type receiveChanPair struct {
	receive *lex.Token
	chan_   *lex.Token
}

func (c ChannelType) __TypeLit__() {
	panic("imposible")
}

var _ TypeLit = (*ChannelType)(nil)

func VisitChannelType(lexer *lex.Lexer) *ChannelType {
	clone := lexer.Clone()

	var chan_ *lex.Token
	var chan_receive *chanReceivePair
	var receive_chan *receiveChanPair

	la := lexer.LA()
	if la.Type_() == lex.GoLexerCHAN { // CHAN | CHAN RECEIVE
		lexer.Pop() // chan
		receive := lexer.LA()
		if receive.Type_() == lex.GoLexerRECEIVE {
			lexer.Pop() // receive
			chan_receive = &chanReceivePair{chan_: la, receive: receive}
		} else {
			chan_ = la
		}
	} else if la.Type_() == lex.GoLexerRECEIVE { // RECEIVE CHAN
		chan_ := lexer.LA1()
		if chan_.Type_() == lex.GoLexerCHAN {
			receive_chan = &receiveChanPair{receive: la, chan_: chan_}
			lexer.Pop() // receive <-
			lexer.Pop() // chan
		} else {
			fmt.Printf("'<-'后面应该有一个'chan'。%s\n", la.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
	} else {
		lexer.Recover(clone)
		return nil
	}

	elementType := VisitType_(lexer)
	if elementType == nil {
		fmt.Printf("此处应该是一个类型描述。%s\n", lexer.LA().ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &ChannelType{chan_: chan_, chan_receive: chan_receive, receive_chan: receive_chan, elementType: elementType}
}

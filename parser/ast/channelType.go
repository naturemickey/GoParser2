package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type ChannelType struct {
	// channelType: (ch=CHAN | ch_re=CHAN RECEIVE | re_ch=RECEIVE CHAN) elementType;
	// elementType: type_;
	chan_        *lex.Token
	chan_receive *chanReceivePair
	receive_chan *receiveChanPair
	elementType  *Type_
}

func (a *ChannelType) Chan_() *lex.Token {
	return a.chan_
}

func (a *ChannelType) SetChan_(chan_ *lex.Token) {
	a.chan_ = chan_
}

func (a *ChannelType) Chan_receive() *chanReceivePair {
	return a.chan_receive
}

func (a *ChannelType) SetChan_receive(chan_receive *chanReceivePair) {
	a.chan_receive = chan_receive
}

func (a *ChannelType) Receive_chan() *receiveChanPair {
	return a.receive_chan
}

func (a *ChannelType) SetReceive_chan(receive_chan *receiveChanPair) {
	a.receive_chan = receive_chan
}

func (a *ChannelType) ElementType() *Type_ {
	return a.elementType
}

func (a *ChannelType) SetElementType(elementType *Type_) {
	a.elementType = elementType
}

func (a *ChannelType) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	if a.chan_ != nil {
		cb.AppendToken(a.chan_)
	} else if a.chan_receive != nil {
		cb.AppendToken(a.chan_receive.chan_)
		cb.AppendToken(a.chan_receive.receive)
	} else if a.receive_chan != nil {
		cb.AppendToken(a.receive_chan.receive)
		cb.AppendToken(a.receive_chan.chan_)
	}
	cb.AppendTreeNode(a.elementType)
	return cb
}

func (a *ChannelType) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ChannelType)(nil)

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
			fmt.Printf("channelType,'<-'后面应该有一个'chan'。%s\n", la.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
	} else {
		lexer.Recover(clone)
		return nil
	}

	elementType := VisitType_(lexer)
	if elementType == nil {
		fmt.Printf("channelType,此处应该是一个类型描述。%s\n", lexer.LA().ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &ChannelType{chan_: chan_, chan_receive: chan_receive, receive_chan: receive_chan, elementType: elementType}
}

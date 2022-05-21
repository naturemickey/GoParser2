package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type SendStmt struct {
	// sendStmt: channel = expression RECEIVE expression;
	channel    *Expression
	receive    *lex.Token
	expression *Expression
}

func (a *SendStmt) Channel() *Expression {
	return a.channel
}

func (a *SendStmt) SetChannel(channel *Expression) {
	a.channel = channel
}

func (a *SendStmt) Receive() *lex.Token {
	return a.receive
}

func (a *SendStmt) SetReceive(receive *lex.Token) {
	a.receive = receive
}

func (a *SendStmt) Expression() *Expression {
	return a.expression
}

func (a *SendStmt) SetExpression(expression *Expression) {
	a.expression = expression
}

func (a *SendStmt) CodeBuilder() *CodeBuilder {
	return NewCB().AppendTreeNode(a.channel).AppendToken(a.receive).AppendTreeNode(a.expression)
}

func (a *SendStmt) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*SendStmt)(nil)

func (s SendStmt) __Statement__() {
	panic("imposible")
}

func (s SendStmt) __SimpleStmt__() {
	panic("imposible")
}

var _ SimpleStmt = (*SendStmt)(nil)

func VisitSendStmt(lexer *lex.Lexer) *SendStmt {
	clone := lexer.Clone()

	channel := VisitExpression(lexer)
	if channel == nil {
		lexer.Recover(clone)
		return nil
	}
	receive := lexer.LA()
	if receive == nil || receive.Type_() != lex.GoLexerRECEIVE {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop()
	expression := VisitExpression(lexer)
	if expression == nil {
		lexer.Recover(clone)
		return nil
	}

	return &SendStmt{channel: channel, receive: receive, expression: expression}
}

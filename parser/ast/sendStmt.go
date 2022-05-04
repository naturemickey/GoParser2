package ast

import "GoParser2/lex"

type SendStmt struct {
	// sendStmt: channel = expression RECEIVE expression;
	channel    *Expression
	receive    *lex.Token
	expression *Expression
}

func (s SendStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

func (s SendStmt) __SimpleStmt__() {
	//TODO implement me
	panic("implement me")
}

var _ SimpleStmt = (*SendStmt)(nil)

func VisitSendStmt(lexer *lex.Lexer) *SendStmt {
	panic("todo")
}

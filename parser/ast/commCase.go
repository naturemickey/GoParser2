package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type CommCase struct {
	// commCase: CASE (sendStmt | recvStmt) | DEFAULT;
	case_    *lex.Token
	sendStmt *SendStmt
	recvStmt *RecvStmt
	default_ *lex.Token
}

func (a *CommCase) Case_() *lex.Token {
	return a.case_
}

func (a *CommCase) SetCase_(case_ *lex.Token) {
	a.case_ = case_
}

func (a *CommCase) SendStmt() *SendStmt {
	return a.sendStmt
}

func (a *CommCase) SetSendStmt(sendStmt *SendStmt) {
	a.sendStmt = sendStmt
}

func (a *CommCase) RecvStmt() *RecvStmt {
	return a.recvStmt
}

func (a *CommCase) SetRecvStmt(recvStmt *RecvStmt) {
	a.recvStmt = recvStmt
}

func (a *CommCase) Default_() *lex.Token {
	return a.default_
}

func (a *CommCase) SetDefault_(default_ *lex.Token) {
	a.default_ = default_
}

func (a *CommCase) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.case_)
	cb.AppendTreeNode(a.sendStmt)
	cb.AppendTreeNode(a.recvStmt)
	cb.AppendToken(a.default_)
	return cb
}

func (a *CommCase) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*CommCase)(nil)

func VisitCommCase(lexer *lex.Lexer) *CommCase {
	clone := lexer.Clone()

	la := lexer.LA()
	if la.Type_() == lex.GoLexerCASE {
		case_ := la
		lexer.Pop() // case_

		sendStmt := VisitSendStmt(lexer)
		var recvStmt *RecvStmt
		if sendStmt == nil {
			recvStmt = VisitRecvStmt(lexer)
			if recvStmt == nil {
				fmt.Printf("commCase,case后面要么是一个send语句，要么是一个recv语句，现在都不是。%s\n", case_.ErrorMsg())
				lexer.Recover(clone)
				return nil
			}
		}
		return &CommCase{case_: case_, sendStmt: sendStmt, recvStmt: recvStmt}
	} else if la.Type_() == lex.GoLexerDEFAULT {
		default_ := la
		lexer.Pop() // default_
		return &CommCase{default_: default_}
	} else {
		// fmt.Printf("commCase,这里要么是个case，要么是个default。%s\n", la.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
}

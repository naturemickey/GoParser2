package ast

import (
	"GoParser2/lex"
)

type FallthroughStmt struct {
	// fallthroughStmt: FALLTHROUGH;
	fallthrough_ *lex.Token
}

func (a *FallthroughStmt) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.fallthrough_)
}

func (a *FallthroughStmt) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*FallthroughStmt)(nil)

func (f FallthroughStmt) __Statement__() {
	panic("imposible")
}

var _ Statement = (*FallthroughStmt)(nil)

func VisitFallthroughStmt(lexer *lex.Lexer) *FallthroughStmt {
	if lexer.LA() == nil { // 文件结束
		return nil
	}
	fallthrough_ := lexer.LA()
	if fallthrough_.Type_() != lex.GoLexerFALLTHROUGH {
		return nil
	}
	lexer.Pop()

	return &FallthroughStmt{fallthrough_: fallthrough_}
}

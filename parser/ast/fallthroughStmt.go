package ast

import "GoParser2/lex"

type FallthroughStmt struct {
	// fallthroughStmt: FALLTHROUGH;
	fallthrough_ *lex.Token
}

func (f FallthroughStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*FallthroughStmt)(nil)

func VisitFallthroughStmt(lexer *lex.Lexer) *FallthroughStmt {
	fallthrough_ := lexer.LA()
	if fallthrough_.Type_() != lex.GoLexerFALLTHROUGH {
		return nil
	}
	lexer.Pop()

	return &FallthroughStmt{fallthrough_: fallthrough_}
}

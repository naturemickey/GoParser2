package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type ContinueStmt struct {
	// continueStmt: CONTINUE IDENTIFIER?;
	continue_  *lex.Token
	identifier *lex.Token
}

func (a *ContinueStmt) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*ContinueStmt)(nil)

func (c ContinueStmt) __Statement__() {
	panic("imposible")
}

var _ Statement = (*ContinueStmt)(nil)

func VisitContinueStmt(lexer *lex.Lexer) *ContinueStmt {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	continue_ := lexer.LA()
	if continue_.Type_() != lex.GoLexerCONTINUE {
		return nil
	}
	lexer.Pop() // continue_

	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		identifier = nil
	} else {
		lexer.Pop() // identifier
	}

	return &ContinueStmt{continue_: continue_, identifier: identifier}
}

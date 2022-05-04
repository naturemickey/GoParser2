package ast

import (
	"GoParser2/lex"
	"fmt"
)

type DeferStmt struct {
	// deferStmt: DEFER expression;
	defer_     *lex.Token
	expression *Expression
}

func (d DeferStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*DeferStmt)(nil)

func VisitDeferStmt(lexer *lex.Lexer) *DeferStmt {
	defer_ := lexer.LA()
	if defer_.Type_() != lex.GoLexerDEFER {
		return nil
	}
	lexer.Pop()

	expression := VisitExpression(lexer)
	if expression == nil {
		fmt.Printf("defer后面需要是一个'表达式'。%s\n", defer_.ErrorMsg())
		return nil
	}

	return &DeferStmt{defer_: defer_, expression: expression}
}

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
	panic("imposible")
}

var _ Statement = (*DeferStmt)(nil)

func VisitDeferStmt(lexer *lex.Lexer) *DeferStmt {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	defer_ := lexer.LA()
	if defer_.Type_() != lex.GoLexerDEFER {
		return nil
	}
	lexer.Pop() // defer_

	expression := VisitExpression(lexer)
	if expression == nil {
		fmt.Printf("defer后面需要是一个'表达式'。%s\n", defer_.ErrorMsg())
		return nil
	}

	return &DeferStmt{defer_: defer_, expression: expression}
}

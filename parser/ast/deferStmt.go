package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
	"fmt"
)

type DeferStmt struct {
	// deferStmt: DEFER expression;
	defer_     *lex.Token
	expression *Expression
}

func (a *DeferStmt) CodeBuilder() *util.CodeBuilder {
	cb := util.NewCB()
	cb.AppendToken(a.defer_)
	cb.AppendTreeNode(a.expression)
	return cb
}

func (a *DeferStmt) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*DeferStmt)(nil)

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

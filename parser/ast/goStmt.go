package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type GoStmt struct {
	// goStmt: GO expression;
	go_        *lex.Token
	expression *Expression
}

func (a *GoStmt) Go_() *lex.Token {
	return a.go_
}

func (a *GoStmt) SetGo_(go_ *lex.Token) {
	a.go_ = go_
}

func (a *GoStmt) Expression() *Expression {
	return a.expression
}

func (a *GoStmt) SetExpression(expression *Expression) {
	a.expression = expression
}

func (a *GoStmt) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.go_).AppendTreeNode(a.expression)
}

func (a *GoStmt) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*GoStmt)(nil)

func (g GoStmt) __Statement__() {
	panic("imposible")
}

var _ Statement = (*GoStmt)(nil)

func VisitGoStmt(lexer *lex.Lexer) *GoStmt {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	clone := lexer.Clone()

	go_ := lexer.LA()
	if go_.Type_() != lex.GoLexerGO {
		return nil
	}
	lexer.Pop() // go_

	expression := VisitExpression(lexer)
	if expression == nil {
		fmt.Printf("goStmt,go后面必须是一个表达式。%s\n", go_.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &GoStmt{go_: go_, expression: expression}
}

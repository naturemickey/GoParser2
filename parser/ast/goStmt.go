package ast

import (
	"GoParser2/lex"
	"fmt"
)

type GoStmt struct {
	// goStmt: GO expression;
	go_        *lex.Token
	expression *Expression
}

func (g GoStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*GoStmt)(nil)

func VisitGoStmt(lexer *lex.Lexer) *GoStmt {
	clone := lexer.Clone()

	go_ := lexer.LA()
	if go_.Type_() != lex.GoLexerGO {
		return nil
	}
	lexer.Pop() // go_

	expression := VisitExpression(lexer)
	if expression == nil {
		fmt.Printf("go后面必须是一个表达式。%s\n", go_.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &GoStmt{go_: go_, expression: expression}
}

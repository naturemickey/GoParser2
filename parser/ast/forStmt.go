package ast

import (
	"GoParser2/lex"
)

type ForStmt struct {
	// forStmt: FOR (expression | forClause | rangeClause)? block;
	for_        *lex.Token
	expression  *Expression
	forClause   *ForClause
	rangeClause *RangeClause
	block       *Block
}

func (a *ForStmt) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.for_)
	cb.AppendTreeNode(a.expression)
	cb.AppendTreeNode(a.forClause)
	cb.AppendTreeNode(a.rangeClause)
	cb.AppendTreeNode(a.block)
	return cb
}

func (a *ForStmt) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ForStmt)(nil)

func (f ForStmt) __Statement__() {
	panic("imposible")
}

var _ Statement = (*ForStmt)(nil)

func VisitForStmt(lexer *lex.Lexer) *ForStmt {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	for_ := lexer.LA()
	if for_.Type_() != lex.GoLexerFOR {
		return nil
	}
	lexer.Pop() // for_

	var expression *Expression
	var forClause *ForClause
	var rangeClause *RangeClause

	rangeClause = VisitRangeClause(lexer)
	if rangeClause == nil {
		forClause = VisitForClause(lexer)
		if forClause == nil {
			expression = VisitExpression(lexer)
		}
	}

	block := VisitBlock(lexer)

	return &ForStmt{for_: for_, expression: expression, forClause: forClause, rangeClause: rangeClause, block: block}
}

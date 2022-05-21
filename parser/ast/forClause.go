package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type ForClause struct {
	// forClause: initStmt = simpleStmt? eos expression? eos postStmt = simpleStmt?;
	initStmt   SimpleStmt
	eos1       *Eos
	expression *Expression
	eos2       *Eos
	postStmt   SimpleStmt
}

func (a *ForClause) InitStmt() SimpleStmt {
	return a.initStmt
}

func (a *ForClause) SetInitStmt(initStmt SimpleStmt) {
	a.initStmt = initStmt
}

func (a *ForClause) Eos1() *Eos {
	return a.eos1
}

func (a *ForClause) SetEos1(eos1 *Eos) {
	a.eos1 = eos1
}

func (a *ForClause) Expression() *Expression {
	return a.expression
}

func (a *ForClause) SetExpression(expression *Expression) {
	a.expression = expression
}

func (a *ForClause) Eos2() *Eos {
	return a.eos2
}

func (a *ForClause) SetEos2(eos2 *Eos) {
	a.eos2 = eos2
}

func (a *ForClause) PostStmt() SimpleStmt {
	return a.postStmt
}

func (a *ForClause) SetPostStmt(postStmt SimpleStmt) {
	a.postStmt = postStmt
}

func (a *ForClause) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendTreeNode(a.initStmt).AppendTreeNode(a.eos1)
	cb.AppendTreeNode(a.expression).AppendTreeNode(a.eos2)
	cb.AppendTreeNode(a.postStmt)
	return cb
}

func (a *ForClause) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ForClause)(nil)

func VisitForClause(lexer *lex.Lexer) *ForClause {
	clone := lexer.Clone()

	var initStmt SimpleStmt = VisitSimpleStmt(lexer)
	var eos1 *Eos = VisitEos(lexer)
	var expression *Expression = VisitExpression(lexer)
	var eos2 *Eos = VisitEos(lexer)
	var postStmt SimpleStmt = VisitSimpleStmt(lexer)

	if initStmt != nil && eos1 == nil && eos2 == nil {
		expCode := initStmt.String()
		expression = VisitExpression(lex.NewLexerWithCode(expCode))
		if expression == nil {
			lexer.Recover(clone)
			return nil
		}
		initStmt = nil
	}

	return &ForClause{initStmt: initStmt, eos1: eos1, expression: expression, eos2: eos2, postStmt: postStmt}
}

package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type ExprSwitchStmt struct {
	// todo 我理解下面几个问号都是可以去掉的
	// exprSwitchStmt:
	//	SWITCH (expression?
	//					| simpleStmt? eos expression?
	//					) L_CURLY exprCaseClause* R_CURLY;

	switch_ *lex.Token

	expression *Expression
	simpleStmt SimpleStmt
	eos        *Eos

	lCurly          *lex.Token
	exprCaseClauses []*ExprCaseClause
	rCurly          *lex.Token
}

func (a *ExprSwitchStmt) Switch_() *lex.Token {
	return a.switch_
}

func (a *ExprSwitchStmt) SetSwitch_(switch_ *lex.Token) {
	a.switch_ = switch_
}

func (a *ExprSwitchStmt) Expression() *Expression {
	return a.expression
}

func (a *ExprSwitchStmt) SetExpression(expression *Expression) {
	a.expression = expression
}

func (a *ExprSwitchStmt) SimpleStmt() SimpleStmt {
	return a.simpleStmt
}

func (a *ExprSwitchStmt) SetSimpleStmt(simpleStmt SimpleStmt) {
	a.simpleStmt = simpleStmt
}

func (a *ExprSwitchStmt) Eos() *Eos {
	return a.eos
}

func (a *ExprSwitchStmt) SetEos(eos *Eos) {
	a.eos = eos
}

func (a *ExprSwitchStmt) LCurly() *lex.Token {
	return a.lCurly
}

func (a *ExprSwitchStmt) SetLCurly(lCurly *lex.Token) {
	a.lCurly = lCurly
}

func (a *ExprSwitchStmt) ExprCaseClauses() []*ExprCaseClause {
	return a.exprCaseClauses
}

func (a *ExprSwitchStmt) SetExprCaseClauses(exprCaseClauses []*ExprCaseClause) {
	a.exprCaseClauses = exprCaseClauses
}

func (a *ExprSwitchStmt) RCurly() *lex.Token {
	return a.rCurly
}

func (a *ExprSwitchStmt) SetRCurly(rCurly *lex.Token) {
	a.rCurly = rCurly
}

func (a *ExprSwitchStmt) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.switch_)
	if a.simpleStmt != nil {
		cb.AppendTreeNode(a.simpleStmt)
		cb.AppendString(";")
		cb.AppendTreeNode(a.expression)
	} else if a.eos != nil {
		cb.AppendString(";")
		cb.AppendTreeNode(a.expression)
	} else {
		cb.AppendTreeNode(a.expression)
	}
	cb.AppendToken(a.lCurly)
	if len(a.exprCaseClauses) > 0 {
		cb.Newline()
		for _, clause := range a.exprCaseClauses {
			cb.AppendTreeNode(clause).Newline()
		}
	}
	cb.AppendToken(a.rCurly)
	return cb
}

func (a *ExprSwitchStmt) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ExprSwitchStmt)(nil)

func (e ExprSwitchStmt) __Statement__() {
	panic("imposible")
}

func (e ExprSwitchStmt) __SwitchStmt__() {
	panic("imposible")
}

var _ SwitchStmt = (*ExprSwitchStmt)(nil)

func VisitExprSwitchStmt(lexer *lex.Lexer) *ExprSwitchStmt {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	clone := lexer.Clone()

	switch_ := lexer.LA()
	if switch_.Type_() != lex.GoLexerSWITCH {
		return nil
	}
	lexer.Pop() // switch_

	var expression *Expression
	var simpleStmt SimpleStmt
	var eos *Eos

	clone1 := lexer.Clone()
	expression = VisitExpression(lexer)
	if expression != nil {
		lCurly := lexer.LA()
		if lCurly.Type_() != lex.GoLexerL_CURLY {
			expression = nil
			lexer.Recover(clone1)
		}
	}
	if expression == nil {
		simpleStmt = VisitSimpleStmt(lexer)

		eos = VisitEos(lexer)
		if eos != nil {
			expression = VisitExpression(lexer)
		}
	}

	lCurly := lexer.LA()
	if lCurly.Type_() != lex.GoLexerL_CURLY {
		//fmt.Printf("exprSwitchStmt,此处应该是一个左花括号才对。%s\n", lCurly.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lCurly

	var exprCaseClauses []*ExprCaseClause
	for true {
		exprCaseClause := VisitExprCaseClause(lexer)
		if exprCaseClause != nil {
			exprCaseClauses = append(exprCaseClauses, exprCaseClause)
		} else {
			break
		}
	}

	rCurly := lexer.LA()
	if rCurly.Type_() != lex.GoLexerR_CURLY {
		//fmt.Printf("exprSwitchStmt,此处应该是一个右花括号才对。%s\n", lCurly.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rCurly

	return &ExprSwitchStmt{switch_: switch_, expression: expression, simpleStmt: simpleStmt, eos: eos, lCurly: lCurly, exprCaseClauses: exprCaseClauses, rCurly: rCurly}
}

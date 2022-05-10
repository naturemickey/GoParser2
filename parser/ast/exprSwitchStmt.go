package ast

import (
	"GoParser2/lex"
	"fmt"
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

	lCurly          *lex.Token
	exprCaseClauses []*ExprCaseClause
	rCurly          *lex.Token
}

func (a *ExprSwitchStmt) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.switch_)
	if a.simpleStmt != nil {
		cb.AppendTreeNode(a.simpleStmt)
		cb.AppendString(";")
		cb.AppendTreeNode(a.expression)
	} else {
		cb.AppendTreeNode(a.expression)
	}
	cb.AppendToken(a.lCurly).Newline()
	for _, clause := range a.exprCaseClauses {
		cb.AppendTreeNode(clause).Newline()
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

	expression := VisitExpression(lexer)
	var simpleStmt SimpleStmt
	if expression == nil {
		simpleStmt = VisitSimpleStmt(lexer)
		if simpleStmt == nil {
			fmt.Printf("switch关键字后面需要有一个表达式。%s\n", switch_.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		eos := VisitEos(lexer)
		if eos == nil {
			fmt.Printf("此处应该有一个分号。%s\n", lexer.LA().ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		expression = VisitExpression(lexer)
		if expression == nil {
			fmt.Printf("分号后面需要有一个表达式。%s\n", switch_.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
	}
	lCurly := lexer.LA()
	if lCurly.Type_() != lex.GoLexerL_CURLY {
		fmt.Printf("此处应该是一个左花括号才对。%s\n", lCurly.ErrorMsg())
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
		fmt.Printf("此处应该是一个右花括号才对。%s\n", lCurly.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rCurly

	return &ExprSwitchStmt{switch_: switch_, expression: expression, simpleStmt: simpleStmt, lCurly: lCurly, exprCaseClauses: exprCaseClauses, rCurly: rCurly}
}

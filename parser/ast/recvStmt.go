package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type RecvStmt struct {
	// recvStmt: (expressionList ASSIGN | identifierList DECLARE_ASSIGN)? recvExpr = expression;
	expressionList *ExpressionList // todo 这里要研究一下等号后面是哪些类型的表达式（应该是有限制的）
	assign         *lex.Token
	identifierList *IdentifierList
	declare_assign *lex.Token
	recvExpr       *Expression
}

func (a *RecvStmt) ExpressionList() *ExpressionList {
	return a.expressionList
}

func (a *RecvStmt) SetExpressionList(expressionList *ExpressionList) {
	a.expressionList = expressionList
}

func (a *RecvStmt) Assign() *lex.Token {
	return a.assign
}

func (a *RecvStmt) SetAssign(assign *lex.Token) {
	a.assign = assign
}

func (a *RecvStmt) IdentifierList() *IdentifierList {
	return a.identifierList
}

func (a *RecvStmt) SetIdentifierList(identifierList *IdentifierList) {
	a.identifierList = identifierList
}

func (a *RecvStmt) Declare_assign() *lex.Token {
	return a.declare_assign
}

func (a *RecvStmt) SetDeclare_assign(declare_assign *lex.Token) {
	a.declare_assign = declare_assign
}

func (a *RecvStmt) RecvExpr() *Expression {
	return a.recvExpr
}

func (a *RecvStmt) SetRecvExpr(recvExpr *Expression) {
	a.recvExpr = recvExpr
}

func (a *RecvStmt) CodeBuilder() *CodeBuilder {
	return NewCB().AppendTreeNode(a.expressionList).AppendToken(a.assign).
		AppendTreeNode(a.identifierList).AppendToken(a.declare_assign).AppendTreeNode(a.recvExpr)
}

func (a *RecvStmt) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*RecvStmt)(nil)

func VisitRecvStmt(lexer *lex.Lexer) *RecvStmt {
	clone := lexer.Clone()

	var identifierList *IdentifierList
	var expressionList *ExpressionList
	var assign *lex.Token
	var declare_assign *lex.Token

	identifierList = VisitIdentifierList(lexer)
	if identifierList != nil {
		declare_assign = lexer.LA()
		if declare_assign.Type_() != lex.GoLexerDECLARE_ASSIGN {
			identifierList = nil
			declare_assign = nil
			lexer.Recover(clone)
		} else {
			lexer.Pop() // declare_assign
		}
	}
	if identifierList == nil {
		expressionList = VisitExpressionList(lexer)
		if expressionList != nil {
			assign = lexer.LA()
			if assign.Type_() != lex.GoLexerASSIGN {
				expressionList = nil
				assign = nil
				lexer.Recover(clone)
			} else {
				lexer.Pop() // assign
			}
		}
	}

	recvExpr := VisitExpression(lexer)
	if recvExpr == nil {
		lexer.Recover(clone)
		return nil
	}

	return &RecvStmt{expressionList: expressionList, assign: assign, identifierList: identifierList, declare_assign: declare_assign, recvExpr: recvExpr}
}

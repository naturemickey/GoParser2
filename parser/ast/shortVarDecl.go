package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type ShortVarDecl struct {
	// shortVarDecl: identifierList DECLARE_ASSIGN expressionList;
	identifierList *IdentifierList
	declare_assign *lex.Token
	expressionList *ExpressionList
}

func (a *ShortVarDecl) IdentifierList() *IdentifierList {
	return a.identifierList
}

func (a *ShortVarDecl) SetIdentifierList(identifierList *IdentifierList) {
	a.identifierList = identifierList
}

func (a *ShortVarDecl) Declare_assign() *lex.Token {
	return a.declare_assign
}

func (a *ShortVarDecl) SetDeclare_assign(declare_assign *lex.Token) {
	a.declare_assign = declare_assign
}

func (a *ShortVarDecl) ExpressionList() *ExpressionList {
	return a.expressionList
}

func (a *ShortVarDecl) SetExpressionList(expressionList *ExpressionList) {
	a.expressionList = expressionList
}

func (a *ShortVarDecl) CodeBuilder() *CodeBuilder {
	return NewCB().AppendTreeNode(a.identifierList).AppendToken(a.declare_assign).AppendTreeNode(a.expressionList)
}

func (a *ShortVarDecl) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ShortVarDecl)(nil)

func (s ShortVarDecl) __Statement__() {
	panic("imposible")
}

func (s ShortVarDecl) __SimpleStmt__() {
	panic("imposible")
}

var _ SimpleStmt = (*ShortVarDecl)(nil)

func VisitShortVarDecl(lexer *lex.Lexer) *ShortVarDecl {
	clone := lexer.Clone()
	identifierList := VisitIdentifierList(lexer)
	if identifierList == nil {
		return nil
	}

	declare_assign := lexer.LA()
	if declare_assign.Type_() != lex.GoLexerDECLARE_ASSIGN {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // declare_assign

	expressionList := VisitExpressionList(lexer)
	if expressionList == nil {
		lexer.Recover(clone)
		return nil
	}
	return &ShortVarDecl{identifierList: identifierList, declare_assign: declare_assign, expressionList: expressionList}
}

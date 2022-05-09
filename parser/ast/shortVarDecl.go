package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type ShortVarDecl struct {
	// shortVarDecl: identifierList DECLARE_ASSIGN expressionList;
	identifierList *IdentifierList
	declare_assign *lex.Token
	expressionList *ExpressionList
}

func (a *ShortVarDecl) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*ShortVarDecl)(nil)

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

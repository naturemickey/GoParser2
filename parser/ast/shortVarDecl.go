package ast

import "GoParser2/lex"

type ShortVarDecl struct {
	// shortVarDecl: identifierList DECLARE_ASSIGN expressionList;
	identifierList *IdentifierList
	declare_assign *lex.Token
	expressionList *ExpressionList
}

func (s ShortVarDecl) __Statement__() {
	//TODO implement me
	panic("implement me")
}

func (s ShortVarDecl) __SimpleStmt__() {
	//TODO implement me
	panic("implement me")
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

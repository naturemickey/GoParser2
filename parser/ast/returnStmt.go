package ast

import "GoParser2/lex"

type ReturnStmt struct {
	// returnStmt: RETURN expressionList?;
	return_        *lex.Token
	expressionList *ExpressionList
}

func (r ReturnStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*ReturnStmt)(nil)

func VisitReturnStmt(lexer *lex.Lexer) *ReturnStmt {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	return_ := lexer.LA()
	if return_.Type_() != lex.GoLexerRETURN {
		return nil
	}
	lexer.Pop() // return_

	expressionList := VisitExpressionList(lexer)

	return &ReturnStmt{return_: return_, expressionList: expressionList}
}

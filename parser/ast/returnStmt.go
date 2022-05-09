package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
)

type ReturnStmt struct {
	// returnStmt: RETURN expressionList?;
	return_        *lex.Token
	expressionList *ExpressionList
}

func (a *ReturnStmt) CodeBuilder() *util.CodeBuilder {
	return util.NewCB().AppendToken(a.return_).AppendTreeNode(a.expressionList)
}

func (a *ReturnStmt) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*ReturnStmt)(nil)

func (r ReturnStmt) __Statement__() {
	panic("imposible")
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

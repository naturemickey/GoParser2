package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type ExprCaseClause struct {
	// exprCaseClause: exprSwitchCase COLON statementList?;
	exprSwitchCase *ExprSwitchCase
	colon          *lex.Token
	statementList  *StatementList
}

func (a *ExprCaseClause) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendTreeNode(a.exprSwitchCase).AppendToken(a.colon).AppendTreeNode(a.statementList)
	return cb
}

func (a *ExprCaseClause) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ExprCaseClause)(nil)

func VisitExprCaseClause(lexer *lex.Lexer) *ExprCaseClause {
	clone := lexer.Clone()

	exprSwitchCase := VisitExprSwitchCase(lexer)
	if exprSwitchCase == nil {
		return nil
	}

	colon := lexer.LA()
	if colon.Type_() != lex.GoLexerCOLON {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // colon

	statementList := VisitStatementList(lexer)

	return &ExprCaseClause{exprSwitchCase: exprSwitchCase, colon: colon, statementList: statementList}
}

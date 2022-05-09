package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type ExprCaseClause struct {
	// exprCaseClause: exprSwitchCase COLON statementList?;
	exprSwitchCase *ExprSwitchCase
	colon          *lex.Token
	statementList  *StatementList
}

func (a *ExprCaseClause) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*ExprCaseClause)(nil)

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

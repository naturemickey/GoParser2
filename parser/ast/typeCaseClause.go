package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type TypeCaseClause struct {
	// typeCaseClause: typeSwitchCase COLON statementList?;
	typeSwitchCase *TypeSwitchCase
	colon          *lex.Token
	statementList  *StatementList
}

func (a *TypeCaseClause) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*TypeCaseClause)(nil)

func VisitTypeCaseClause(lexer *lex.Lexer) *TypeCaseClause {
	clone := lexer.Clone()

	typeSwitchCase := VisitTypeSwitchCase(lexer)
	if typeSwitchCase == nil {
		return nil
	}

	colon := lexer.LA()
	if colon.Type_() != lex.GoLexerCOLON {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // colon

	statementList := VisitStatementList(lexer)

	return &TypeCaseClause{typeSwitchCase: typeSwitchCase, colon: colon, statementList: statementList}
}

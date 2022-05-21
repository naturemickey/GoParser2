package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type TypeCaseClause struct {
	// typeCaseClause: typeSwitchCase COLON statementList?;
	typeSwitchCase *TypeSwitchCase
	colon          *lex.Token
	statementList  *StatementList
}

func (a *TypeCaseClause) TypeSwitchCase() *TypeSwitchCase {
	return a.typeSwitchCase
}

func (a *TypeCaseClause) SetTypeSwitchCase(typeSwitchCase *TypeSwitchCase) {
	a.typeSwitchCase = typeSwitchCase
}

func (a *TypeCaseClause) Colon() *lex.Token {
	return a.colon
}

func (a *TypeCaseClause) SetColon(colon *lex.Token) {
	a.colon = colon
}

func (a *TypeCaseClause) StatementList() *StatementList {
	return a.statementList
}

func (a *TypeCaseClause) SetStatementList(statementList *StatementList) {
	a.statementList = statementList
}

func (a *TypeCaseClause) CodeBuilder() *CodeBuilder {
	return NewCB().AppendTreeNode(a.typeSwitchCase).AppendToken(a.colon).AppendTreeNode(a.statementList)
}

func (a *TypeCaseClause) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*TypeCaseClause)(nil)

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

package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type CommClause struct {
	// commClause: commCase COLON statementList?;
	commCase      *CommCase
	colon         *lex.Token
	statementList *StatementList
}

func (a *CommClause) CommCase() *CommCase {
	return a.commCase
}

func (a *CommClause) SetCommCase(commCase *CommCase) {
	a.commCase = commCase
}

func (a *CommClause) Colon() *lex.Token {
	return a.colon
}

func (a *CommClause) SetColon(colon *lex.Token) {
	a.colon = colon
}

func (a *CommClause) StatementList() *StatementList {
	return a.statementList
}

func (a *CommClause) SetStatementList(statementList *StatementList) {
	a.statementList = statementList
}

func (a *CommClause) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendTreeNode(a.commCase)
	cb.AppendToken(a.colon)
	if a.statementList != nil {
		cb.AppendTreeNode(a.statementList)
	}
	return cb
}

func (a *CommClause) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*CommClause)(nil)

func VisitCommClause(lexer *lex.Lexer) *CommClause {
	clone := lexer.Clone()

	commCase := VisitCommCase(lexer)
	if commCase == nil {
		lexer.Recover(clone)
		return nil
	}

	colon := lexer.LA()
	if colon.Type_() != lex.GoLexerCOLON {
		fmt.Printf("commClause,冒号在哪里？%s\n", colon.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop()

	statementList := VisitStatementList(lexer)

	return &CommClause{commCase, colon, statementList}
}

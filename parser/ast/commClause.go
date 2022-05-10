package ast

import (
	"GoParser2/lex"
	"fmt"
)

type CommClause struct {
	// commClause: commCase COLON statementList?;
	commCase      *CommCase
	colon         *lex.Token
	statementList *StatementList
}

func (a *CommClause) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendTreeNode(a.commCase)
	cb.AppendToken(a.colon).Newline()
	cb.AppendTreeNode(a.statementList)
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

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

func VisitCommClause(lexer *lex.Lexer) *CommClause {
	clone := lexer.Clone()

	commCase := VisitCommCase(lexer)
	if commCase == nil {
		lexer.Recover(clone)
		return nil
	}

	colon := lexer.LA()
	if colon.Type_() != lex.GoLexerCOLON {
		fmt.Printf("冒号在哪里？%s\n", colon.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop()

	statementList := VisitStatementList(lexer)

	return &CommClause{commCase, colon, statementList}
}

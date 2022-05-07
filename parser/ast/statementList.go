package ast

import (
	"GoParser2/lex"
)

type StatementList struct {
	// statementList: (eos? statement eos)+;
	statements []Statement
}

func VisitStatementList(lexer *lex.Lexer) *StatementList {
	VisitEos(lexer)

	var statements []Statement

	for true {
		statement := VisitStatement(lexer)
		if statement != nil {
			statements = append(statements, statement)
		} else {
			break
		}
	}
	VisitEos(lexer)

	return &StatementList{statements: statements}
}

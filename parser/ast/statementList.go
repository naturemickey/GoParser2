package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"reflect"
)

type StatementList struct {
	// statementList: (eos? statement eos)+;
	statements []Statement
}

func (a *StatementList) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*StatementList)(nil)

func VisitStatementList(lexer *lex.Lexer) *StatementList {
	VisitEos(lexer)

	var statements []Statement

	for true {
		VisitEos(lexer)
		statement := VisitStatement(lexer)
		if statement != nil && !reflect.ValueOf(statement).IsNil() {
			statements = append(statements, statement)
			VisitEos(lexer)
		} else {
			break
		}
	}
	VisitEos(lexer)

	return &StatementList{statements: statements}
}

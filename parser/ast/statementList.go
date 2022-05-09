package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
	"reflect"
)

type StatementList struct {
	// statementList: (eos? statement eos)+;
	statements []Statement
}

func (a *StatementList) CodeBuilder() *util.CodeBuilder {
	cb := util.NewCB()
	for _, statement := range a.statements {
		cb.AppendTreeNode(statement).Newline()
	}
	return cb
}

func (a *StatementList) String() string {
	return a.CodeBuilder().String()
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

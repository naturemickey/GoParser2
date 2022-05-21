package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type StatementList struct {
	// statementList: (eos? statement eos)+;
	statements []Statement
}

func (a *StatementList) Statements() []Statement {
	return a.statements
}

func (a *StatementList) SetStatements(statements []Statement) {
	a.statements = statements
}

func NewStatementList() *StatementList {
	return &StatementList{}
}

func (a *StatementList) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	for _, statement := range a.statements {
		cb.AppendTreeNode(statement).Newline()
	}
	cb.popLast()
	return cb
}

func (a *StatementList) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*StatementList)(nil)

func VisitStatementList(lexer *lex.Lexer) *StatementList {
	VisitEos(lexer)

	var statements []Statement

	for true {
		VisitEos(lexer)
		statement := VisitStatement(lexer)
		if statement != nil {
			statements = append(statements, statement)
			VisitEos(lexer)
		} else {
			break
		}
	}
	VisitEos(lexer)

	return &StatementList{statements: statements}
}

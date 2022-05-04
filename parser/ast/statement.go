package ast

import "GoParser2/lex"

type Statement interface {
}

func VisitStatement(lexer *lex.Lexer) *Statement {
	panic("todo")
}

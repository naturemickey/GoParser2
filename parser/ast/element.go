package ast

import "GoParser2/lex"

type Element interface {
	// element: expression | literalValue;
}

func VisitElement(lexer *lex.Lexer) Element {
	panic("todo")
}

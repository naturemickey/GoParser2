package ast

import "GoParser2/lex"

type KeyedElement struct {
	// keyedElement: (key COLON)? element;
}

func VisitKeyedElement(lexer *lex.Lexer) *KeyedElement {
	panic("todo")
}

package ast

import "GoParser2/lex"

type Key interface {
	// key: expression | literalValue;
}

func VisitKey(lexer *lex.Lexer) Key {
	panic("todo")
}

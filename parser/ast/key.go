package ast

import "GoParser2/lex"

type Key interface {
	// key: expression | literalValue;
	__Key__()
}

func VisitKey(lexer *lex.Lexer) Key {
	expression := VisitExpression(lexer)
	if expression != nil {
		return expression
	}
	return VisitLiteralValue(lexer)
}

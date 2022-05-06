package ast

import "GoParser2/lex"

type Element interface {
	// element: expression | literalValue;
	__Element__()
}

func VisitElement(lexer *lex.Lexer) Element {
	expression := VisitExpression(lexer)
	if expression != nil {
		return expression
	}
	return VisitLiteralValue(lexer)
}

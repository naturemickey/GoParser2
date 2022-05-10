package ast

import (
	"GoParser2/lex"
)

type Element interface {
	ITreeNode
	// element: expression | literalValue;
	__Element__()
}

func VisitElement(lexer *lex.Lexer) Element {
	expression := VisitExpression(lexer)
	if expression != nil {
		return expression
	}

	literalValue := VisitLiteralValue(lexer)
	if literalValue != nil {
		return literalValue
	}

	return nil
}

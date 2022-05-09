package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type Element interface {
	parser.ITreeNode
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

package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type Key interface {
	parser.ITreeNode
	// key: expression | literalValue;
	__Key__()
}

func VisitKey(lexer *lex.Lexer) Key {
	expression := VisitExpression(lexer)
	if expression != nil {
		return expression
	}

	literalValue := VisitLiteralValue(lexer)
	if literalValue == nil {
		return nil
	}
	return literalValue
}

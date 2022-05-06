package ast

import "GoParser2/lex"

type CompositeLit struct {
	// compositeLit: literalType literalValue;
	literalType  *LiteralType
	literalValue *LiteralValue
}

func (c CompositeLit) __Literal__() {
	//TODO implement me
	panic("implement me")
}

var _ Literal = (*CompositeLit)(nil)

func VisitCompositeLit(lexer *lex.Lexer) *CompositeLit {
	clone := lexer.Clone()

	literalType := VisitLiteralType(lexer)
	if literalType == nil {
		lexer.Recover(clone)
		return nil
	}

	literalValue := VisitLiteralValue(lexer)
	if literalValue == nil {
		lexer.Recover(clone)
		return nil
	}

	return &CompositeLit{literalType: literalType, literalValue: literalValue}
}

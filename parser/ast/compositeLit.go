package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type CompositeLit struct {
	// compositeLit: literalType literalValue;
	literalType  *LiteralType
	literalValue *LiteralValue
}

func (a *CompositeLit) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*CompositeLit)(nil)

func (c CompositeLit) __Literal__() {
	panic("imposible")
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

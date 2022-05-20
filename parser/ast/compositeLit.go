package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type CompositeLit struct {
	// compositeLit: literalType literalValue;
	literalType  *LiteralType
	literalValue *LiteralValue
}

func (a *CompositeLit) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendTreeNode(a.literalType)
	cb.AppendTreeNode(a.literalValue)
	return cb
}

func (a *CompositeLit) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*CompositeLit)(nil)

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

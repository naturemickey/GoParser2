package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type Literal interface {
	parser.ITreeNode
	// literal: basicLit | compositeLit | functionLit;

	__Literal__()
}

func VisitLiteral(lexer *lex.Lexer) Literal {
	basicLit := VisitBasicLit(lexer)
	if basicLit != nil {
		return basicLit
	}

	functionLit := VisitFunctionLit(lexer)
	if functionLit != nil {
		return functionLit
	}

	compositeLit := VisitCompositeLit(lexer)
	if compositeLit != nil {
		return compositeLit
	}
	return nil
}

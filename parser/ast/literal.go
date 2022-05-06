package ast

import "GoParser2/lex"

type Literal interface {
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

	return VisitCompositeLit(lexer)
}

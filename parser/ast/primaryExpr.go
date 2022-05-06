package ast

import "GoParser2/lex"

type PrimaryExpr struct {
	// primaryExpr:
	//	operand
	//	| conversion
	//	| methodExpr
	//	| primaryExpr (
	//		(DOT IDENTIFIER)
	//		| index
	//		| slice
	//		| typeAssertion
	//		| arguments
	//	);
}

func VisitPrimaryExpr(lexer *lex.Lexer) *PrimaryExpr {
	panic("todo")
}

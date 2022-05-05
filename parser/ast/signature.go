package ast

import "GoParser2/lex"

type Signature struct {
	// signature:
	//	parameters result
	//	| parameters;
	parameters *Parameters
	result     *Result
}

func VisitSignature(lexer *lex.Lexer) *Signature {
	panic("todo")
}

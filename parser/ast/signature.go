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
	parameters := VisitParameters(lexer)
	if parameters == nil {
		return nil
	}

	result := VisitResult(lexer)

	return &Signature{parameters: parameters, result: result}
}

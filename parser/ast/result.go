package ast

import "GoParser2/lex"

type Result struct {
	// result: parameters | type_;
	parameters *Parameters
	type_      *Type_
}

func VisitResult(lexer *lex.Lexer) *Result {
	parameters := VisitParameters(lexer)
	if parameters == nil {
		type_ := VisitType_(lexer)
		if type_ == nil {
			return nil
		}
		return &Result{type_: type_}
	}
	return &Result{parameters: parameters}
}

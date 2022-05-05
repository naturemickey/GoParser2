package ast

import (
	"GoParser2/lex"
	"fmt"
)

type MethodSpec struct {
	// methodSpec:
	//	IDENTIFIER parameters result
	//	| IDENTIFIER parameters;
	identifier *lex.Token
	parameters *Parameters
	result     *Result
}

func (m MethodSpec) __IMethodspecOrTypename__() {
	//TODO implement me
	panic("implement me")
}

var _ IMethodspecOrTypename = (*MethodSpec)(nil)

func VisitMethodSpec(lexer *lex.Lexer) *MethodSpec {
	clone := lexer.Clone()

	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		return nil
	}

	parameters := VisitParameters(lexer)
	if parameters == nil {
		fmt.Printf("方法名后面找不到参数列表。%s\n", identifier.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	result := VisitResult(lexer)

	return &MethodSpec{identifier: identifier, parameters: parameters, result: result}
}

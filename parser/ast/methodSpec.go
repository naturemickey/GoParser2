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
	result     Result
}

func (a *MethodSpec) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.identifier).AppendTreeNode(a.parameters).AppendTreeNode(a.result)
}

func (a *MethodSpec) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*MethodSpec)(nil)

func (m MethodSpec) __IMethodspecOrTypename__() {
	panic("imposible")
}

var _ IMethodspecOrTypename = (*MethodSpec)(nil)

func VisitMethodSpec(lexer *lex.Lexer) *MethodSpec {
	clone := lexer.Clone()

	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		return nil
	}
	lexer.Pop() // identifier

	parameters := VisitParameters(lexer)
	if parameters == nil {
		fmt.Printf("methodSpec,方法名后面找不到参数列表。%s\n", identifier.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	result := VisitResult(lexer)

	return &MethodSpec{identifier: identifier, parameters: parameters, result: result}
}

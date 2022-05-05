package ast

import (
	"GoParser2/lex"
	"fmt"
)

type InterfaceType struct {
	// interfaceType: INTERFACE L_CURLY ((methodSpec | typeName) eos)* R_CURLY;
	interface_     *lex.Token
	lCurly         *lex.Token
	methodOrType_s []IMethodspecOrTypename
	rCurly         *lex.Token
}

func (i InterfaceType) __TypeLit__() {
	//TODO implement me
	panic("implement me")
}

var _ TypeLit = (*InterfaceType)(nil)

func VisitInterfaceType(lexer *lex.Lexer) *InterfaceType {
	clone := lexer.Clone()

	interface_ := lexer.LA()
	if interface_.Type_() != lex.GoLexerINTERFACE {
		return nil
	}
	lexer.Pop() // interface_

	lCurly := lexer.LA()
	if lCurly.Type_() != lex.GoLexerL_CURLY {
		fmt.Printf("此处应该是一个'{'。%s\n", lCurly.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	var methodOrType_s []IMethodspecOrTypename
	for {
		methodSpec := VisitMethodSpec(lexer)
		if methodSpec != nil {
			methodOrType_s = append(methodOrType_s, methodSpec)
		} else {
			typeName := VisitTypeName(lexer)
			if typeName != nil {
				methodOrType_s = append(methodOrType_s, typeName)
			} else {
				break
			}
		}
	}

	rCurly := lexer.LA()
	if rCurly.Type_() != lex.GoLexerR_CURLY {
		fmt.Printf("此处应该有一个'}'才对。%s\n", rCurly.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &InterfaceType{interface_: interface_, lCurly: lCurly, methodOrType_s: methodOrType_s, rCurly: rCurly}
}

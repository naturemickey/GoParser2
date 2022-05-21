package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type InterfaceType struct {
	// interfaceType: INTERFACE L_CURLY ((methodSpec | typeName) eos)* R_CURLY;
	interface_     *lex.Token
	lCurly         *lex.Token
	methodOrType_s []IMethodspecOrTypename
	rCurly         *lex.Token
}

func (a *InterfaceType) Interface_() *lex.Token {
	return a.interface_
}

func (a *InterfaceType) SetInterface_(interface_ *lex.Token) {
	a.interface_ = interface_
}

func (a *InterfaceType) LCurly() *lex.Token {
	return a.lCurly
}

func (a *InterfaceType) SetLCurly(lCurly *lex.Token) {
	a.lCurly = lCurly
}

func (a *InterfaceType) MethodOrType_s() []IMethodspecOrTypename {
	return a.methodOrType_s
}

func (a *InterfaceType) SetMethodOrType_s(methodOrType_s []IMethodspecOrTypename) {
	a.methodOrType_s = methodOrType_s
}

func (a *InterfaceType) RCurly() *lex.Token {
	return a.rCurly
}

func (a *InterfaceType) SetRCurly(rCurly *lex.Token) {
	a.rCurly = rCurly
}

func (a *InterfaceType) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.interface_).AppendToken(a.lCurly)
	if len(a.methodOrType_s) != 0 {
		cb.Newline()
		for _, mt := range a.methodOrType_s {
			cb.AppendTreeNode(mt).Newline()
		}
	}
	cb.AppendToken(a.rCurly)
	return cb
}

func (a *InterfaceType) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*InterfaceType)(nil)

func (i InterfaceType) __TypeLit__() {
	panic("imposible")
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
		fmt.Printf("interfaceType,此处应该是一个'{'。%s\n", lCurly.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lCurly

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
		fmt.Printf("interfaceType,此处应该有一个'}'才对。%s\n", rCurly.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rCurly

	return &InterfaceType{interface_: interface_, lCurly: lCurly, methodOrType_s: methodOrType_s, rCurly: rCurly}
}

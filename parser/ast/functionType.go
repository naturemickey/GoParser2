package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type FunctionType struct {
	// functionType: FUNC signature;
	func_     *lex.Token
	signature *Signature
}

func (a *FunctionType) Func_() *lex.Token {
	return a.func_
}

func (a *FunctionType) SetFunc_(func_ *lex.Token) {
	a.func_ = func_
}

func (a *FunctionType) Signature() *Signature {
	return a.signature
}

func (a *FunctionType) SetSignature(signature *Signature) {
	a.signature = signature
}

func (a *FunctionType) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.func_).AppendTreeNode(a.signature)
}

func (a *FunctionType) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*FunctionType)(nil)

func (f FunctionType) __TypeLit__() {
	panic("imposible")
}

var _ TypeLit = (*FunctionType)(nil)

func VisitFunctionType(lexer *lex.Lexer) *FunctionType {
	clone := lexer.Clone()

	func_ := lexer.LA()
	if func_.Type_() != lex.GoLexerFUNC {
		return nil
	}
	lexer.Pop() // func_

	signature := VisitSignature(lexer)
	if signature == nil {
		fmt.Printf("functionType,func后面找不到函数的签名。%s\n", func_.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &FunctionType{func_: func_, signature: signature}
}

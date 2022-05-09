package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"fmt"
)

type FunctionType struct {
	// functionType: FUNC signature;
	func_     *lex.Token
	signature *Signature
}

func (a *FunctionType) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*FunctionType)(nil)

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
		fmt.Printf("func后面找不到函数的签名。%s\n", func_.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &FunctionType{func_: func_, signature: signature}
}

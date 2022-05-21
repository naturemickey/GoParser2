package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type FunctionLit struct {
	// functionLit: FUNC signature block; // function
	func_     *lex.Token
	signature *Signature
	block     *Block
}

func (a *FunctionLit) Func_() *lex.Token {
	return a.func_
}

func (a *FunctionLit) SetFunc_(func_ *lex.Token) {
	a.func_ = func_
}

func (a *FunctionLit) Signature() *Signature {
	return a.signature
}

func (a *FunctionLit) SetSignature(signature *Signature) {
	a.signature = signature
}

func (a *FunctionLit) Block() *Block {
	return a.block
}

func (a *FunctionLit) SetBlock(block *Block) {
	a.block = block
}

func (a *FunctionLit) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.func_).AppendTreeNode(a.signature).AppendTreeNode(a.block)
	return cb
}

func (a *FunctionLit) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*FunctionLit)(nil)

func (f FunctionLit) __Literal__() {
	panic("imposible")
}

var _ Literal = (*FunctionLit)(nil)

func VisitFunctionLit(lexer *lex.Lexer) *FunctionLit {
	clone := lexer.Clone()

	func_ := lexer.LA()
	if func_.Type_() != lex.GoLexerFUNC {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // func_

	signature := VisitSignature(lexer)
	if signature == nil {
		lexer.Recover(clone)
		return nil
	}

	block := VisitBlock(lexer)
	if block == nil {
		lexer.Recover(clone)
		return nil
	}

	return &FunctionLit{func_: func_, signature: signature, block: block}
}

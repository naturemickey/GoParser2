package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
)

type FunctionLit struct {
	// functionLit: FUNC signature block; // function
	func_     *lex.Token
	signature *Signature
	block     *Block
}

func (a *FunctionLit) CodeBuilder() *util.CodeBuilder {
	cb := util.NewCB()
	cb.AppendToken(a.func_).AppendTreeNode(a.signature).AppendTreeNode(a.block)
	return cb
}

func (a *FunctionLit) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*FunctionLit)(nil)

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

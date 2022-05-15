package ast

import (
	"GoParser2/lex"
	"fmt"
)

type FunctionDecl struct {
	// functionDecl: FUNC annotation=ANNOTATION_COMMENT IDENTIFIER (signature block?);
	func_      *lex.Token
	annotation *Annotation
	identifier *lex.Token
	signature  *Signature
	block      *Block
}

func (this *FunctionDecl) Signature() *Signature {
	return this.signature
}

func (this *FunctionDecl) Block() *Block {
	return this.block
}

func (this *FunctionDecl) Annotation() *Annotation {
	return this.annotation
}

func (this *FunctionDecl) Name() string {
	return this.identifier.Literal()
}

func (a *FunctionDecl) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.func_).AppendToken(a.identifier)
	cb.AppendTreeNode(a.signature).AppendTreeNode(a.block)
	return cb
}

func (a *FunctionDecl) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*FunctionDecl)(nil)

func (f FunctionDecl) __IFunctionMethodDeclaration__() {
	panic("imposible")
}

var _ IFunctionMethodDeclaration = (*FunctionDecl)(nil)

func VisitFunctionDecl(lexer *lex.Lexer) *FunctionDecl {
	clone := lexer.Clone()

	func_ := lexer.LA()
	if func_.Type_() != lex.GoLexerFUNC {
		return nil
	}
	lexer.Pop() // func_

	annotation := VisitAnnotation(lexer)

	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // identifier

	signature := VisitSignature(lexer)
	if signature == nil {
		fmt.Printf("functionDecl,没看到参数的部分。%s\n", identifier.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	block := VisitBlock(lexer)
	// todo 看一下定义上为什么block可以为nil的？

	return &FunctionDecl{func_: func_, annotation: annotation, identifier: identifier, signature: signature, block: block}
}

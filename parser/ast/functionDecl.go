package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type FunctionDecl struct {
	// functionDecl: FUNC annotationList IDENTIFIER (signature block?);
	func_          *lex.Token
	annotationList *AnnotationList
	identifier     *lex.Token
	signature      *Signature
	block          *Block
}

func (this *FunctionDecl) Func_() *lex.Token {
	return this.func_
}

func (this *FunctionDecl) SetFunc_(func_ *lex.Token) {
	this.func_ = func_
}

func (this *FunctionDecl) SetAnnotationList(annotationList *AnnotationList) {
	this.annotationList = annotationList
}

func (this *FunctionDecl) Identifier() *lex.Token {
	return this.identifier
}

func (this *FunctionDecl) SetIdentifier(identifier *lex.Token) {
	this.identifier = identifier
}

func (this *FunctionDecl) SetSignature(signature *Signature) {
	this.signature = signature
}

func (this *FunctionDecl) SetBlock(block *Block) {
	this.block = block
}

func (this *FunctionDecl) AnnotationList() *AnnotationList {
	return this.annotationList
}

func (this *FunctionDecl) Signature() *Signature {
	return this.signature
}

func (this *FunctionDecl) Block() *Block {
	return this.block
}

func (this *FunctionDecl) Name() string {
	return this.identifier.Literal()
}

func (a *FunctionDecl) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.func_).AppendTreeNode(a.annotationList).AppendToken(a.identifier)
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

	annotationList := VisitAnnotationList(lexer)

	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // identifier

	signature := VisitSignature(lexer)
	if signature == nil {
		fmt.Printf("functionDecl,???????????????????????????%s\n", identifier.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	block := VisitBlock(lexer)
	// todo ???????????????????????????block?????????nil??????

	return &FunctionDecl{func_: func_, annotationList: annotationList, identifier: identifier, signature: signature, block: block}
}

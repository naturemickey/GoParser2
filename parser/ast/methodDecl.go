package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type MethodDecl struct {
	// methodDecl: FUNC annotationList receiver IDENTIFIER ( signature block?);
	func_          *lex.Token
	annotationList *AnnotationList
	receiver       *Receiver
	identifier     *lex.Token
	signature      *Signature
	block          *Block
}

func (this *MethodDecl) Func_() *lex.Token {
	return this.func_
}

func (this *MethodDecl) SetFunc_(func_ *lex.Token) {
	this.func_ = func_
}

func (this *MethodDecl) SetAnnotationList(annotationList *AnnotationList) {
	this.annotationList = annotationList
}

func (this *MethodDecl) SetReceiver(receiver *Receiver) {
	this.receiver = receiver
}

func (this *MethodDecl) Identifier() *lex.Token {
	return this.identifier
}

func (this *MethodDecl) SetIdentifier(identifier *lex.Token) {
	this.identifier = identifier
}

func (this *MethodDecl) Signature() *Signature {
	return this.signature
}

func (this *MethodDecl) SetSignature(signature *Signature) {
	this.signature = signature
}

func (this *MethodDecl) Block() *Block {
	return this.block
}

func (this *MethodDecl) SetBlock(block *Block) {
	this.block = block
}

func (this *MethodDecl) AnnotationList() *AnnotationList {
	return this.annotationList
}

func (this *MethodDecl) Receiver() *Receiver {
	return this.receiver
}

func (this *MethodDecl) Name() string {
	return this.identifier.Literal()
}

func (a *MethodDecl) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.func_).AppendTreeNode(a.annotationList).AppendTreeNode(a.receiver).AppendToken(a.identifier).AppendTreeNode(a.signature).AppendTreeNode(a.block)
}

func (a *MethodDecl) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*MethodDecl)(nil)

func (m MethodDecl) __IFunctionMethodDeclaration__() {
	panic("imposible")
}

var _ IFunctionMethodDeclaration = (*MethodDecl)(nil)

func VisitMethodDecl(lexer *lex.Lexer) *MethodDecl {
	clone := lexer.Clone()

	func_ := lexer.LA()
	if func_.Type_() != lex.GoLexerFUNC {
		return nil
	}
	lexer.Pop()

	annotationList := VisitAnnotationList(lexer)

	receiver := VisitReceiver(lexer)
	if receiver == nil {
		fmt.Printf("methodDecl,func???????????????receiver?????????%s\n", func_.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // identifier

	signature := VisitSignature(lexer)
	if signature == nil {
		fmt.Printf("methodDecl,???????????????????????????%s\n", identifier.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	block := VisitBlock(lexer)
	// todo ???????????????????????????block?????????nil??????

	return &MethodDecl{func_: func_, annotationList: annotationList, receiver: receiver, identifier: identifier, signature: signature, block: block}
}

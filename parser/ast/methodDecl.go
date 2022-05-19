package ast

import (
	"GoParser2/lex"
	"fmt"
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
	return NewCB().AppendToken(a.func_).AppendTreeNode(a.receiver).AppendToken(a.identifier).AppendTreeNode(a.signature).AppendTreeNode(a.block)
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
		fmt.Printf("methodDecl,func后面没看到receiver定义。%s\n", func_.ErrorMsg())
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
		fmt.Printf("methodDecl,没看到参数的部分。%s\n", identifier.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	block := VisitBlock(lexer)
	// todo 看一下定义上为什么block可以为nil的？

	return &MethodDecl{func_: func_, annotationList: annotationList, receiver: receiver, identifier: identifier, signature: signature, block: block}
}

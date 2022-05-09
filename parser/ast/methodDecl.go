package ast

import (
	"GoParser2/lex"
	"fmt"
)

type MethodDecl struct {
	// methodDecl: FUNC receiver IDENTIFIER ( signature block?);
	func_      *lex.Token
	receiver   *Receiver
	identifier *lex.Token
	signature  *Signature
	block      *Block
}

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

	receiver := VisitReceiver(lexer)
	if receiver == nil {
		fmt.Printf("func后面没看到receiver定义。%s\n", func_.ErrorMsg())
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
		fmt.Printf("没看到参数的部分。%s\n", identifier.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	block := VisitBlock(lexer)
	// todo 看一下定义上为什么block可以为nil的？

	return &MethodDecl{func_: func_, receiver: receiver, identifier: identifier, signature: signature, block: block}
}

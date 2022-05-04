package ast

import (
	"GoParser2/lex"
	"fmt"
)

type MethodDecl struct {
	// methodDecl: FUNC receiver IDENTIFIER ( signature block?);
	funcToken  *lex.Token
	receiver   *Receiver
	identifier *lex.Token
	signature  *Signature
	block      *Block
}

func (m MethodDecl) __IFunctionMethodDeclaration__() {
	//TODO implement me
	panic("implement me")
}

var _ IFunctionMethodDeclaration = (*MethodDecl)(nil)

func VisitMethodDecl(lexer *lex.Lexer) *MethodDecl {
	clone := lexer.Clone()

	funcToken := lexer.LA()
	if funcToken.Type_() != lex.GoLexerFUNC {
		return nil
	}
	lexer.Pop()

	receiver := VisitReceiver(lexer)
	if receiver == nil {
		fmt.Printf("func后面没看到receiver定义。%s\n", funcToken.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop()

	signature := VisitSignature(lexer)
	if signature == nil {
		fmt.Printf("没看到参数的部分。%s\n", identifier.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	block := VisitBlock(lexer)
	// todo 看一下定义上为什么block可以为nil的？

	return &MethodDecl{funcToken: funcToken, receiver: receiver, identifier: identifier, signature: signature, block: block}
}

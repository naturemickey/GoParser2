package ast

import (
	"GoParser2/lex"
	"fmt"
)

type FunctionDecl struct {
	// functionDecl: FUNC IDENTIFIER (signature block?);
	funcToken  *lex.Token
	identifier *lex.Token
	signature  *Signature
	block      *Block
}

func (f FunctionDecl) __IFunctionMethodDeclaration__() {
	//TODO implement me
	panic("implement me")
}

var _ IFunctionMethodDeclaration = (*FunctionDecl)(nil)

func VisitFunctionDecl(lexer *lex.Lexer) *FunctionDecl {
	clone := lexer.Clone()

	funcToken := lexer.LA()
	if funcToken.Type_() != lex.GoLexerFUNC {
		return nil
	}
	lexer.Pop()

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

	return &FunctionDecl{funcToken: funcToken, identifier: identifier, signature: signature, block: block}
}

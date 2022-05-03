package ast

import (
	"GoParser2/lex"
	"fmt"
)

type Declaration interface {
	IFunctionMethodDeclaration
	__Declaration__()

	// declaration: constDecl | typeDecl | varDecl;
}

func VisitDeclaration(lexer *lex.Lexer) Declaration {
	la := lexer.LA()

	switch la.Type_() {
	case lex.GoLexerCONST:
		return VisitConstDecl(lexer)
	case lex.GoLexerTYPE:
		return VisitTypeDecl(lexer)
	case lex.GoLexerVAR:
		return VisitVarDecl(lexer)
	default:
		fmt.Printf("或许不可能，但的确发生了：在读Declaration的时候应该先判断是否是const/type/var，parser里面有一个bug。%s\n", la.ErrorMsg())
		return nil
	}
}

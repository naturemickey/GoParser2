package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type Declaration interface {
	ITreeNode
	IFunctionMethodDeclaration
	Statement
	__Declaration__()

	// declaration: constDecl | typeDecl | varDecl;
}

func VisitDeclaration(lexer *lex.Lexer) Declaration {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	la := lexer.LA()

	switch la.Type_() {
	case lex.GoLexerCONST:
		constDecl := VisitConstDecl(lexer)
		if constDecl != nil {
			return constDecl
		}
		return nil
	case lex.GoLexerTYPE:
		typeDecl := VisitTypeDecl(lexer)
		if typeDecl != nil {
			return typeDecl
		}
		return nil
	case lex.GoLexerVAR:
		varDecl := VisitVarDecl(lexer)
		if varDecl != nil {
			return varDecl
		}
		return nil
	default:
		return nil
	}
}

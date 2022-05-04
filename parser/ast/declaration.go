package ast

import (
	"GoParser2/lex"
)

type Declaration interface {
	IFunctionMethodDeclaration
	Statement
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
		return nil
	}
}

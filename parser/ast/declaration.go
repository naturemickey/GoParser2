package ast

import "GoParser2/lex"

type Declaration interface {
	IFunctionMethodDeclaration
	__Declaration__()

	// declaration: constDecl | typeDecl | varDecl;
}

func VisitDeclaration(lexer *lex.Lexer) Declaration {
	panic("todo")
}

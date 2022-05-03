package ast

import "GoParser2/lex"

type MethodDecl struct {
}

func (m MethodDecl) __IFunctionMethodDeclaration__() {
	//TODO implement me
	panic("implement me")
}

var _ IFunctionMethodDeclaration = (*MethodDecl)(nil)

func VisitMethodDecl(lexer *lex.Lexer) *MethodDecl {
	panic("todo")
}

package ast

import "GoParser2/lex"

type FunctionDecl struct {
}

func (f FunctionDecl) __IFunctionMethodDeclaration__() {
	//TODO implement me
	panic("implement me")
}

var _ IFunctionMethodDeclaration = (*FunctionDecl)(nil)

func VisitFunctionDecl(lexer *lex.Lexer) *FunctionDecl {
	panic("todo")
}

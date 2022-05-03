package ast

import "GoParser2/lex"

type VarDecl struct {
}

func (v VarDecl) __IFunctionMethodDeclaration__() {
	//TODO implement me
	panic("implement me")
}

func (v VarDecl) __Declaration__() {
	//TODO implement me
	panic("implement me")
}

var _ Declaration = (*VarDecl)(nil)

func VisitVarDecl(lexer *lex.Lexer) *VarDecl {
	panic("todo")
}

package ast

import "GoParser2/lex"

type TypeDecl struct {
	// typeDecl: TYPE (typeSpec | L_PAREN (typeSpec eos)* R_PAREN);
	typeToken *lex.Token
}

func (t TypeDecl) __IFunctionMethodDeclaration__() {
	//TODO implement me
	panic("implement me")
}

func (t TypeDecl) __Declaration__() {
	//TODO implement me
	panic("implement me")
}

var _ Declaration = (*TypeDecl)(nil)

func VisitTypeDecl(lexer *lex.Lexer) *TypeDecl {
	panic("todo")
}

package ast

import "GoParser2/lex"

type MethodSpec struct {
}

func (m MethodSpec) __IMethodspecOrTypename__() {
	//TODO implement me
	panic("implement me")
}

var _ IMethodspecOrTypename = (*MethodSpec)(nil)

func VisitMethodSpec(lexer *lex.Lexer) *MethodSpec {
	panic("todo")
}

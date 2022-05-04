package ast

import "GoParser2/lex"

type Expression struct {
}

func (e Expression) __Statement__() {
	//TODO implement me
	panic("implement me")
}

func (e Expression) __SimpleStmt__() {
	//TODO implement me
	panic("implement me")
}

func (e Expression) __ExpressionStmt__() {
	//TODO implement me
	panic("implement me")
}

var _ ExpressionStmt = (*Expression)(nil)

func VisitExpression(lexer *lex.Lexer) *Expression {
	// clone := lexer.Clone()
	panic("todo")
}

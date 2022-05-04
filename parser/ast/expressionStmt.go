package ast

import "GoParser2/lex"

type ExpressionStmt interface {
	SimpleStmt
	__ExpressionStmt__()

	// expressionStmt: expression;
}

func VisitExpressionStmt(lexer *lex.Lexer) ExpressionStmt {
	return VisitExpression(lexer)
}

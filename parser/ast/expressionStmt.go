package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type ExpressionStmt interface {
	parser.ITreeNode
	SimpleStmt
	__ExpressionStmt__()

	// expressionStmt: expression;
}

func VisitExpressionStmt(lexer *lex.Lexer) ExpressionStmt {
	expression := VisitExpression(lexer)
	if expression == nil {
		return nil
	}
	return expression
}

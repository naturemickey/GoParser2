package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type ExpressionStmt interface {
	ITreeNode
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

package ast

func _getPrimaryFromExpression(expression *Expression) *PrimaryExpr {
	if expression.exp2s != nil && len(expression.exp2s) > 0 {
		return _getPrimaryFromExpression(expression.exp2s[len(expression.exp2s)-1].expression2)
	} else if expression.expression != nil {
		return _getPrimaryFromExpression(expression.expression)
	} else {
		return expression.primaryExpr
	}
}

func _getPrimaryFromExpressionList(expressionList *ExpressionList) *PrimaryExpr {
	if expressionList == nil {
		return nil
	}
	if expressionList.expressions == nil {
		return nil
	}
	if len(expressionList.expressions) == 0 {
		return nil
	}
	return _getPrimaryFromExpression(expressionList.expressions[len(expressionList.expressions)-1])
}

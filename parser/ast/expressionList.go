package ast

import (
	"GoParser2/lex"
	"fmt"
)

type ExpressionList struct {
	// expressionList: expression (COMMA expression)*;
	expressions []*Expression
}

func VisitExpressionList(lexer *lex.Lexer) *ExpressionList {
	expression := VisitExpression(lexer)
	if expression == nil {
		return nil
	}
	var expressions []*Expression
	expressions = append(expressions, expression)

	for true {
		la := lexer.LA()
		if la.Type_() == lex.GoLexerCOMMA {
			expression := VisitExpression(lexer)
			if expression == nil {
				fmt.Printf("逗号后面要跟着一个表达式才对。%s\n", la.ErrorMsg())
				return nil
			}
			expressions = append(expressions, expression)
		} else {
			break
		}
	}
	return &ExpressionList{expressions: expressions}
}

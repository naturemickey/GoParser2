package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"fmt"
)

type ExpressionList struct {
	// expressionList: expression (COMMA expression)*;
	expressions []*Expression
}

func (a *ExpressionList) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*ExpressionList)(nil)

func VisitExpressionList(lexer *lex.Lexer) *ExpressionList {
	expression := VisitExpression(lexer)
	if expression == nil {
		return nil
	}
	var expressions []*Expression
	expressions = append(expressions, expression)

	for true {
		comma := lexer.LA()
		if comma.Type_() == lex.GoLexerCOMMA {
			lexer.Pop() // comma
			expression := VisitExpression(lexer)
			if expression == nil {
				fmt.Printf("逗号后面要跟着一个表达式才对。%s\n", comma.ErrorMsg())
				return nil
			}
			expressions = append(expressions, expression)
		} else {
			break
		}
	}
	return &ExpressionList{expressions: expressions}
}

package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type ExpressionList struct {
	// expressionList: expression (COMMA expression)*;
	expressions []*Expression
}

func (a *ExpressionList) Expressions() []*Expression {
	return a.expressions
}

func (a *ExpressionList) SetExpressions(expressions []*Expression) {
	a.expressions = expressions
}

func (a *ExpressionList) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	for i, expression := range a.expressions {
		if i == 0 {
			cb.AppendTreeNode(expression)
		} else {
			cb.AppendString(",").AppendTreeNode(expression)
		}
	}
	return cb
}

func (a *ExpressionList) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ExpressionList)(nil)

func VisitExpressionList(lexer *lex.Lexer) *ExpressionList {
	expression := VisitExpression(lexer)
	if expression == nil {
		return nil
	}
	var expressions []*Expression
	expressions = append(expressions, expression)

	for true {
		comma := lexer.LA()
		if comma != nil && comma.Type_() == lex.GoLexerCOMMA {
			lexer.Pop() // comma
			expression := VisitExpression(lexer)
			if expression == nil {
				//fmt.Printf("expressionList,逗号后面要跟着一个表达式才对。%s\n", comma.ErrorMsg())
				//return nil
				break
			}
			expressions = append(expressions, expression)
		} else {
			break
		}
	}
	return &ExpressionList{expressions: expressions}
}

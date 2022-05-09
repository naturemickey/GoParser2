package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type Assignment struct {
	// assignment: expressionList assign_op expressionList;
	lExpressionList *ExpressionList
	assign_op       *Assign_op
	rExpressionList *ExpressionList
}

func (a *Assignment) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*Assignment)(nil)

func (a Assignment) __Statement__() {
	panic("imposible")
}

func (a Assignment) __SimpleStmt__() {
	panic("imposible")
}

var _ SimpleStmt = (*Assignment)(nil)

func VisitAssignment(lexer *lex.Lexer) *Assignment {
	clone := lexer.Clone()
	// todo 研究一下为什么左边是expressionList，而不是identifierList？

	lExpressionList := VisitExpressionList(lexer)
	if lExpressionList == nil {
		lexer.Recover(clone)
		return nil
	}
	assign_op := VisitAssign_op(lexer)
	if assign_op == nil {
		lexer.Recover(clone)
		return nil
	}
	rExpressionList := VisitExpressionList(lexer)
	if rExpressionList == nil {
		lexer.Recover(clone)
		return nil
	}
	return &Assignment{lExpressionList: lExpressionList, assign_op: assign_op, rExpressionList: rExpressionList}
}

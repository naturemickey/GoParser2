package ast

import "GoParser2/lex"

type Assignment struct {
	// assignment: expressionList assign_op expressionList;
	lExpressionList *ExpressionList
	assign_op       *Assign_op
	rExpressionList *ExpressionList
}

func (a Assignment) __Statement__() {
	//TODO implement me
	panic("implement me")
}

func (a Assignment) __SimpleStmt__() {
	//TODO implement me
	panic("implement me")
}

var _ SimpleStmt = (*Assignment)(nil)

func VisitAssignment(lexer *lex.Lexer) *Assignment {
	// todo 研究一下为什么左边是expressionList，而不是identifierList？

	lExpressionList := VisitExpressionList(lexer)
	if lExpressionList == nil {
		return nil
	}
	assign_op := VisitAssign_op(lexer)
	if assign_op == nil {
		return nil
	}
	rExpressionList := VisitExpressionList(lexer)
	if rExpressionList == nil {
		return nil
	}
	return &Assignment{lExpressionList: lExpressionList, assign_op: assign_op, rExpressionList: rExpressionList}
}

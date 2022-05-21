package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type Assignment struct {
	// assignment: expressionList assign_op expressionList;
	lExpressionList *ExpressionList
	assign_op       *Assign_op
	rExpressionList *ExpressionList
}

func (a *Assignment) LExpressionList() *ExpressionList {
	return a.lExpressionList
}

func (a *Assignment) SetLExpressionList(lExpressionList *ExpressionList) {
	a.lExpressionList = lExpressionList
}

func (a *Assignment) Assign_op() *Assign_op {
	return a.assign_op
}

func (a *Assignment) SetAssign_op(assign_op *Assign_op) {
	a.assign_op = assign_op
}

func (a *Assignment) RExpressionList() *ExpressionList {
	return a.rExpressionList
}

func (a *Assignment) SetRExpressionList(rExpressionList *ExpressionList) {
	a.rExpressionList = rExpressionList
}

func (a *Assignment) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendTreeNode(a.lExpressionList)
	cb.AppendTreeNode(a.assign_op)
	cb.AppendTreeNode(a.rExpressionList)
	return cb
}

func (a *Assignment) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*Assignment)(nil)

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

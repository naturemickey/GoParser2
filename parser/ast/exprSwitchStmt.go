package ast

import "GoParser2/lex"

type ExprSwitchStmt struct {
	// exprSwitchStmt:
	//	SWITCH (expression?
	//					| simpleStmt? eos expression?
	//					) L_CURLY exprCaseClause* R_CURLY;
}

func (e ExprSwitchStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

func (e ExprSwitchStmt) __SwitchStmt__() {
	//TODO implement me
	panic("implement me")
}

var _ SwitchStmt = (*ExprSwitchStmt)(nil)

func VisitExprSwitchStmt(lexer *lex.Lexer) *ExprSwitchStmt {
	panic("todo")
}

package ast

import "GoParser2/lex"

type TypeSwitchStmt struct {
	// typeSwitchStmt:
	//	SWITCH ( typeSwitchGuard
	//					| eos typeSwitchGuard
	//					| simpleStmt eos typeSwitchGuard)
	//					 L_CURLY typeCaseClause* R_CURLY;
}

func (t TypeSwitchStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

func (t TypeSwitchStmt) __SwitchStmt__() {
	//TODO implement me
	panic("implement me")
}

var _ SwitchStmt = (*TypeSwitchStmt)(nil)

func VisitTypeSwitchStmt(lexer *lex.Lexer) *TypeSwitchStmt {
	panic("todo")
}

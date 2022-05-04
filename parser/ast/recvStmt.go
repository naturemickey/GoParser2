package ast

import "GoParser2/lex"

type RecvStmt struct {
	// recvStmt: (expressionList ASSIGN | identifierList DECLARE_ASSIGN)? recvExpr = expression;
}

func VisitRecvStmt(lexer *lex.Lexer) *RecvStmt {
	panic("todo")
}

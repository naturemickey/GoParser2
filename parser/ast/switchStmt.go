package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type SwitchStmt interface {
	parser.ITreeNode
	Statement
	__SwitchStmt__()

	// switchStmt: exprSwitchStmt | typeSwitchStmt;
}

func VisitSwitchStmt(lexer *lex.Lexer) SwitchStmt {
	exprSwitchStmt := VisitExprSwitchStmt(lexer)
	if exprSwitchStmt != nil {
		return exprSwitchStmt
	}
	typeSwitchStmt := VisitTypeSwitchStmt(lexer)
	if typeSwitchStmt != nil {
		return typeSwitchStmt
	}
	return nil
}

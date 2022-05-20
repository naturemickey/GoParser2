package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type Statement interface {
	ITreeNode
	// statement:
	//	declaration
	//	| labeledStmt
	//	| simpleStmt
	//	| goStmt
	//	| returnStmt
	//	| breakStmt
	//	| continueStmt
	//	| gotoStmt
	//	| fallthroughStmt
	//	| block
	//	| ifStmt
	//	| switchStmt
	//	| selectStmt
	//	| forStmt
	//	| deferStmt;
	__Statement__()
}

func VisitStatement(lexer *lex.Lexer) Statement {
	declaration := VisitDeclaration(lexer)
	if declaration != nil {
		return declaration
	}

	labeledStmt := VisitLabeledStmt(lexer)
	if labeledStmt != nil {
		return labeledStmt
	}
	simpleStmt := VisitSimpleStmt(lexer)
	if simpleStmt != nil {
		return simpleStmt
	}
	goStmt := VisitGoStmt(lexer)
	if goStmt != nil {
		return goStmt
	}
	returnStmt := VisitReturnStmt(lexer)
	if returnStmt != nil {
		return returnStmt
	}
	breakStmt := VisitBreakStmt(lexer)
	if breakStmt != nil {
		return breakStmt
	}
	continueStmt := VisitContinueStmt(lexer)
	if continueStmt != nil {
		return continueStmt
	}
	gotoStmt := VisitGotoStmt(lexer)
	if gotoStmt != nil {
		return gotoStmt
	}
	fallthroughStmt := VisitFallthroughStmt(lexer)
	if fallthroughStmt != nil {
		return fallthroughStmt
	}
	block := VisitBlock(lexer)
	if block != nil {
		return block
	}
	ifStmt := VisitIfStmt(lexer)
	if ifStmt != nil {
		return ifStmt
	}
	switchStmt := VisitSwitchStmt(lexer)
	if switchStmt != nil {
		return switchStmt
	}
	selectStmt := VisitSelectStmt(lexer)
	if selectStmt != nil {
		return selectStmt
	}
	forStmt := VisitForStmt(lexer)
	if forStmt != nil {
		return forStmt
	}
	deferStmt := VisitDeferStmt(lexer)
	if deferStmt != nil {
		return deferStmt
	}
	return nil
}

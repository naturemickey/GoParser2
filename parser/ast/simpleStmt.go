package ast

import "GoParser2/lex"

type SimpleStmt interface {
	Statement
	__SimpleStmt__()

	// simpleStmt:
	//	sendStmt
	//	| incDecStmt
	//	| assignment
	//	| expressionStmt
	//	| shortVarDecl;
}

func VisitSimpleStmt(lexer *lex.Lexer) SimpleStmt {
	sendStmt := VisitSendStmt(lexer)
	if sendStmt != nil {
		return sendStmt
	}
	incDecStmt := VisitIncDecStmt(lexer)
	if incDecStmt != nil {
		return incDecStmt
	}
	assignment := VisitAssignment(lexer)
	if assignment != nil {
		return assignment
	}
	expressionStmt := VisitExpressionStmt(lexer)
	if expressionStmt != nil {
		return expressionStmt
	}
	shortVarDecl := VisitShortVarDecl(lexer)
	if shortVarDecl != nil {
		return shortVarDecl
	}
	return nil
}

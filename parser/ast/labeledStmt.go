package ast

import "GoParser2/lex"

type LabeledStmt struct {
	// labeledStmt: IDENTIFIER COLON statement?;
	identifier *lex.Token
	colon      *lex.Token
	statement  Statement
}

func (l LabeledStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*LabeledStmt)(nil)

func VisitLabeledStmt(lexer *lex.Lexer) *LabeledStmt {
	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		return nil
	}
	lexer.Pop() // identifier

	colon := lexer.LA()
	if colon.Type_() != lex.GoLexerCOLON {
		return nil
	}
	lexer.Pop() // colon

	// todo 冒号后面有一个statement是什么语法？
	statement := VisitStatement(lexer)

	return &LabeledStmt{identifier: identifier, colon: colon, statement: statement}
}

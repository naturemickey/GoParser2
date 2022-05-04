package ast

import "GoParser2/lex"

type GotoStmt struct {
	// gotoStmt: GOTO IDENTIFIER;
	goto_      *lex.Token
	identifier *lex.Token
}

func (g GotoStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*GotoStmt)(nil)

func VisitGotoStmt(lexer *lex.Lexer) *GotoStmt {
	clone := lexer.Clone()

	goto_ := lexer.LA()
	if goto_.Type_() != lex.GoLexerGOTO {
		return nil
	}
	lexer.Pop() // goto_

	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // identifier

	return &GotoStmt{goto_: goto_, identifier: identifier}
}

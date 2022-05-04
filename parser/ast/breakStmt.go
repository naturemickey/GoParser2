package ast

import "GoParser2/lex"

type BreakStmt struct {
	// breakStmt: BREAK IDENTIFIER?;
	break_     *lex.Token
	identifier *lex.Token
}

func (b *BreakStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*BreakStmt)(nil)

func VisitBreakStmt(lexer *lex.Lexer) *BreakStmt {
	break_ := lexer.LA()
	if break_.Type_() != lex.GoLexerBREAK {
		return nil
	}
	lexer.Pop() // break_

	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		identifier = nil
	} else {
		lexer.Pop()
	}

	return &BreakStmt{break_: break_, identifier: identifier}
}

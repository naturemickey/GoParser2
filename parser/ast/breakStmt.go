package ast

import (
	"GoParser2/lex"
)

type BreakStmt struct {
	// breakStmt: BREAK IDENTIFIER?;
	break_     *lex.Token
	identifier *lex.Token
}

func (a *BreakStmt) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.break_)
	cb.AppendToken(a.identifier)
	return cb
}

func (a *BreakStmt) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*BreakStmt)(nil)

func (b *BreakStmt) __Statement__() {
	panic("imposible")
}

var _ Statement = (*BreakStmt)(nil)

func VisitBreakStmt(lexer *lex.Lexer) *BreakStmt {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

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

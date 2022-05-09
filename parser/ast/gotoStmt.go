package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
)

type GotoStmt struct {
	// gotoStmt: GOTO IDENTIFIER;
	goto_      *lex.Token
	identifier *lex.Token
}

func (a *GotoStmt) CodeBuilder() *util.CodeBuilder {
	return util.NewCB().AppendToken(a.goto_).AppendToken(a.identifier)
}

func (a *GotoStmt) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*GotoStmt)(nil)

func (g GotoStmt) __Statement__() {
	panic("imposible")
}

var _ Statement = (*GotoStmt)(nil)

func VisitGotoStmt(lexer *lex.Lexer) *GotoStmt {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

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

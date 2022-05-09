package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
)

type LabeledStmt struct {
	// labeledStmt: IDENTIFIER COLON statement?;
	identifier *lex.Token
	colon      *lex.Token
	statement  Statement
}

func (a *LabeledStmt) CodeBuilder() *util.CodeBuilder {
	return util.NewCB().AppendToken(a.identifier).AppendToken(a.colon).AppendTreeNode(a.statement)
}

func (a *LabeledStmt) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*LabeledStmt)(nil)

func (l LabeledStmt) __Statement__() {
	panic("imposible")
}

var _ Statement = (*LabeledStmt)(nil)

func VisitLabeledStmt(lexer *lex.Lexer) *LabeledStmt {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	clone := lexer.Clone()

	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // identifier

	colon := lexer.LA()
	if colon.Type_() != lex.GoLexerCOLON {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // colon

	// todo 冒号后面有一个statement是什么语法？
	statement := VisitStatement(lexer)

	return &LabeledStmt{identifier: identifier, colon: colon, statement: statement}
}

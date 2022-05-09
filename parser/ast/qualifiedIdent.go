package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
)

type QualifiedIdent struct {
	// qualifiedIdent: IDENTIFIER DOT IDENTIFIER;
	identifier1 *lex.Token
	dot         *lex.Token
	identifier2 *lex.Token
}

func (a *QualifiedIdent) CodeBuilder() *util.CodeBuilder {
	return util.NewCB().AppendToken(a.identifier1).AppendToken(a.dot).AppendToken(a.identifier2)
}

func (a *QualifiedIdent) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*QualifiedIdent)(nil)

func VisitQualifiedIdent(lexer *lex.Lexer) *QualifiedIdent {
	clone := lexer.Clone()

	identifier1 := lexer.LA()
	if identifier1.Type_() != lex.GoLexerIDENTIFIER {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // identifier1

	dot := lexer.LA()
	if dot.Type_() != lex.GoLexerDOT {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // dot

	identifier2 := lexer.LA()
	if identifier2.Type_() != lex.GoLexerIDENTIFIER {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // identifier2

	return &QualifiedIdent{identifier1: identifier1, dot: dot, identifier2: identifier2}
}

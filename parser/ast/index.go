package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
)

type Index struct {
	// index: L_BRACKET expression R_BRACKET;
	lBracket   *lex.Token
	expression *Expression
	rBracket   *lex.Token
}

func (a *Index) CodeBuilder() *util.CodeBuilder {
	return util.NewCB().AppendToken(a.lBracket).AppendTreeNode(a.expression).AppendToken(a.rBracket)
}

func (a *Index) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*Index)(nil)

func VisitIndex(lexer *lex.Lexer) *Index {
	clone := lexer.Clone()

	lBracket := lexer.LA()
	if lBracket.Type_() != lex.GoLexerL_BRACKET {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lBracket

	expression := VisitExpression(lexer)
	if expression == nil {
		lexer.Recover(clone)
		return nil
	}

	rBracket := lexer.LA()
	if rBracket.Type_() != lex.GoLexerR_BRACKET {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rBracket

	return &Index{lBracket: lBracket, expression: expression, rBracket: rBracket}
}

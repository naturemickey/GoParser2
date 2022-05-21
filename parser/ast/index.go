package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type Index struct {
	// index: L_BRACKET expression R_BRACKET;
	lBracket   *lex.Token
	expression *Expression
	rBracket   *lex.Token
}

func (a *Index) LBracket() *lex.Token {
	return a.lBracket
}

func (a *Index) SetLBracket(lBracket *lex.Token) {
	a.lBracket = lBracket
}

func (a *Index) Expression() *Expression {
	return a.expression
}

func (a *Index) SetExpression(expression *Expression) {
	a.expression = expression
}

func (a *Index) RBracket() *lex.Token {
	return a.rBracket
}

func (a *Index) SetRBracket(rBracket *lex.Token) {
	a.rBracket = rBracket
}

func (a *Index) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.lBracket).AppendTreeNode(a.expression).AppendToken(a.rBracket)
}

func (a *Index) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*Index)(nil)

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

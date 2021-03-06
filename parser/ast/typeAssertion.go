package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type TypeAssertion struct {
	// typeAssertion: DOT L_PAREN type_ R_PAREN;
	dot    *lex.Token
	lParen *lex.Token
	type_  *Type_
	rParen *lex.Token
}

func (a *TypeAssertion) Dot() *lex.Token {
	return a.dot
}

func (a *TypeAssertion) SetDot(dot *lex.Token) {
	a.dot = dot
}

func (a *TypeAssertion) LParen() *lex.Token {
	return a.lParen
}

func (a *TypeAssertion) SetLParen(lParen *lex.Token) {
	a.lParen = lParen
}

func (a *TypeAssertion) Type_() *Type_ {
	return a.type_
}

func (a *TypeAssertion) SetType_(type_ *Type_) {
	a.type_ = type_
}

func (a *TypeAssertion) RParen() *lex.Token {
	return a.rParen
}

func (a *TypeAssertion) SetRParen(rParen *lex.Token) {
	a.rParen = rParen
}

func (a *TypeAssertion) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.dot).AppendToken(a.lParen).AppendTreeNode(a.type_).AppendToken(a.rParen)
}

func (a *TypeAssertion) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*TypeAssertion)(nil)

func VisitTypeAssertion(lexer *lex.Lexer) *TypeAssertion {
	clone := lexer.Clone()

	dot := lexer.LA()
	if dot.Type_() != lex.GoLexerDOT {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // dot

	lParen := lexer.LA()
	if lParen.Type_() != lex.GoLexerL_PAREN {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lParen

	type_ := VisitType_(lexer)
	if type_ == nil {
		lexer.Recover(clone)
		return nil
	}

	rParen := lexer.LA()
	if rParen.Type_() != lex.GoLexerR_PAREN {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rParen

	return &TypeAssertion{dot: dot, lParen: lParen, type_: type_, rParen: rParen}
}

package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type NonNamedType struct {
	// nonNamedType: typeLit | L_PAREN nonNamedType R_PAREN;
	typeLit TypeLit

	lParen       *lex.Token
	nonNamedType *NonNamedType
	rParen       *lex.Token
}

func (a *NonNamedType) TypeLit() TypeLit {
	return a.typeLit
}

func (a *NonNamedType) SetTypeLit(typeLit TypeLit) {
	a.typeLit = typeLit
}

func (a *NonNamedType) LParen() *lex.Token {
	return a.lParen
}

func (a *NonNamedType) SetLParen(lParen *lex.Token) {
	a.lParen = lParen
}

func (a *NonNamedType) NonNamedType() *NonNamedType {
	return a.nonNamedType
}

func (a *NonNamedType) SetNonNamedType(nonNamedType *NonNamedType) {
	a.nonNamedType = nonNamedType
}

func (a *NonNamedType) RParen() *lex.Token {
	return a.rParen
}

func (a *NonNamedType) SetRParen(rParen *lex.Token) {
	a.rParen = rParen
}

func (a *NonNamedType) CodeBuilder() *CodeBuilder {
	return NewCB().AppendTreeNode(a.typeLit).AppendToken(a.lParen).AppendTreeNode(a.nonNamedType).AppendToken(a.rParen)
}

func (a *NonNamedType) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*NonNamedType)(nil)

func VisitNonNamedType(lexer *lex.Lexer) *NonNamedType {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	clone := lexer.Clone()

	lParen := lexer.LA()
	if lParen.Type_() == lex.GoLexerL_PAREN {
		lexer.Pop() // lParen

		nonNamedType := VisitNonNamedType(lexer)
		if nonNamedType == nil {
			lexer.Recover(clone)
			return nil
		}

		rParen := lexer.LA()
		if rParen.Type_() != lex.GoLexerR_PAREN {
			fmt.Printf("nonNamedType,此处应该是一个')'。%s\n", rParen.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		lexer.Pop()

		return &NonNamedType{lParen: lParen, nonNamedType: nonNamedType, rParen: rParen}
	}

	typeLit := VisitTypeLit(lexer)
	if typeLit == nil {
		return nil
	}
	return &NonNamedType{typeLit: typeLit}
}

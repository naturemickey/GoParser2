package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"fmt"
)

type NonNamedType struct {
	// nonNamedType: typeLit | L_PAREN nonNamedType R_PAREN;
	typeLit TypeLit

	lParen       *lex.Token
	nonNamedType *NonNamedType
	rParen       *lex.Token
}

func (a *NonNamedType) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*NonNamedType)(nil)

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
			fmt.Printf("此处应该是一个')'。%s\n", rParen.ErrorMsg())
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

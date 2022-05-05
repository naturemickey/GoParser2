package ast

import (
	"GoParser2/lex"
	"fmt"
)

type Parameters struct {
	// parameters:
	//	L_PAREN (parameterDecl (COMMA parameterDecl)* COMMA?)? R_PAREN;
	lParen         *lex.Token
	parameterDecls []*ParameterDecl
	rParen         *lex.Token
}

func VisitParameters(lexer *lex.Lexer) *Parameters {
	clone := lexer.Clone()

	lParen := lexer.LA()
	if lParen.Type_() != lex.GoLexerL_PAREN {
		return nil
	}
	lexer.Pop() // lParen

	var parameterDecls []*ParameterDecl
	for {
		parameterDecl := VisitParameterDecl(lexer)
		if parameterDecl == nil {
			break
		}
		parameterDecls = append(parameterDecls, parameterDecl)

		comma := lexer.LA()
		if comma.Type_() != lex.GoLexerCOMMA {
			break
		}
		lexer.Pop() // comma
	}

	rParen := lexer.LA()
	if rParen.Type_() != lex.GoLexerR_PAREN {
		fmt.Printf("此处应该是一个')'。 %s\n", rParen.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &Parameters{lParen: lParen, parameterDecls: parameterDecls, rParen: rParen}
}

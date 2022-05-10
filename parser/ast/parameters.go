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

func (a *Parameters) CodeBuilder() *CodeBuilder {
	cb := NewCB().AppendToken(a.lParen)
	for i, decl := range a.parameterDecls {
		if i == 0 {
			cb.AppendTreeNode(decl)
		} else {
			cb.AppendString(",").AppendTreeNode(decl)
		}
	}
	cb.AppendToken(a.rParen)
	return cb
}

func (a *Parameters) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*Parameters)(nil)

func (p Parameters) __Result__() {
	panic("imposible")
}

var _ Result = (*Parameters)(nil)

func VisitParameters(lexer *lex.Lexer) *Parameters {
	clone := lexer.Clone()

	lParen := lexer.LA()
	if lParen == nil || lParen.Type_() != lex.GoLexerL_PAREN {
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
		fmt.Printf("parameters,此处应该是一个')'。 %s\n", rParen.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rParen

	return &Parameters{lParen: lParen, parameterDecls: parameterDecls, rParen: rParen}
}

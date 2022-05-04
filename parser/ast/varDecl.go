package ast

import (
	"GoParser2/lex"
	"fmt"
)

type VarDecl struct {
	// varDecl: VAR (varSpec | L_PAREN (varSpec eos)* R_PAREN);
	var_     *lex.Token
	lParen   *lex.Token
	varSpecs []*VarSpec
	rParen   *lex.Token
}

func (v VarDecl) __IFunctionMethodDeclaration__() {
	//TODO implement me
	panic("implement me")
}

func (v VarDecl) __Declaration__() {
	//TODO implement me
	panic("implement me")
}

var _ Declaration = (*VarDecl)(nil)

func VisitVarDecl(lexer *lex.Lexer) *VarDecl {
	clone := lexer.Clone()

	var_ := lexer.LA()
	if var_.Type_() != lex.GoLexerVAR {
		return nil
	}
	lexer.Pop()

	lParen := lexer.LA()
	if lParen.Type_() != lex.GoLexerL_PAREN {
		lParen = nil
	} else {
		lexer.Pop()
	}

	var varSpecs []*VarSpec

	if lParen == nil {
		varSpec := VisitVarSpec(lexer)
		if varSpec == nil {
			fmt.Printf("没找到变量的定义。%s\n", var_.ErrorMsg())
			lexer.Recover(clone)
			return nil
		} else {
			varSpecs = append(varSpecs, varSpec)
		}
		return &VarDecl{var_: var_, varSpecs: varSpecs}
	} else {
		for true {
			varSpec := VisitVarSpec(lexer)
			if varSpec != nil {
				varSpecs = append(varSpecs, varSpec)
			}
		}
		rParen := lexer.LA()
		if rParen.Type_() != lex.GoLexerR_PAREN {
			fmt.Printf("没找到右括号。%s\n", rParen.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return &VarDecl{var_: var_, lParen: lParen, varSpecs: varSpecs, rParen: rParen}
	}
}

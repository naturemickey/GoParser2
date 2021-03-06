package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type VarDecl struct {
	// varDecl: VAR (varSpec | L_PAREN (varSpec eos)* R_PAREN);
	var_     *lex.Token
	lParen   *lex.Token
	varSpecs []*VarSpec
	rParen   *lex.Token
}

func (a *VarDecl) Var_() *lex.Token {
	return a.var_
}

func (a *VarDecl) SetVar_(var_ *lex.Token) {
	a.var_ = var_
}

func (a *VarDecl) LParen() *lex.Token {
	return a.lParen
}

func (a *VarDecl) SetLParen(lParen *lex.Token) {
	a.lParen = lParen
}

func (a *VarDecl) VarSpecs() []*VarSpec {
	return a.varSpecs
}

func (a *VarDecl) SetVarSpecs(varSpecs []*VarSpec) {
	a.varSpecs = varSpecs
}

func (a *VarDecl) RParen() *lex.Token {
	return a.rParen
}

func (a *VarDecl) SetRParen(rParen *lex.Token) {
	a.rParen = rParen
}

func (a *VarDecl) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.var_)
	if a.lParen != nil {
		cb.AppendToken(a.lParen).Newline()
		for _, spec := range a.varSpecs {
			cb.AppendTreeNode(spec).Newline()
		}
		cb.AppendToken(a.rParen)
	} else {
		cb.AppendTreeNode(a.varSpecs[0])
	}
	return cb
}

func (a *VarDecl) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*VarDecl)(nil)

func (v VarDecl) __Statement__() {
	panic("imposible")
}

func (v VarDecl) __IFunctionMethodDeclaration__() {
	panic("imposible")
}

func (v VarDecl) __Declaration__() {
	panic("imposible")
}

var _ Declaration = (*VarDecl)(nil)

func VisitVarDecl(lexer *lex.Lexer) *VarDecl {
	clone := lexer.Clone()

	var_ := lexer.LA()
	if var_.Type_() != lex.GoLexerVAR {
		return nil
	}
	lexer.Pop() // var_

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
			fmt.Printf("varDecl,没找到变量的定义。%s\n", var_.ErrorMsg())
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
			} else {
				break
			}
		}
		rParen := lexer.LA()
		if rParen.Type_() != lex.GoLexerR_PAREN {
			fmt.Printf("varDecl,没找到右括号。%s\n", rParen.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		lexer.Pop()
		return &VarDecl{var_: var_, lParen: lParen, varSpecs: varSpecs, rParen: rParen}
	}
}

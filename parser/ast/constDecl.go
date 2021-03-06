package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type ConstDecl struct {
	// constDecl: CONST (constSpec | L_PAREN (constSpec eos)* R_PAREN);
	const_     *lex.Token
	lParen     *lex.Token
	constSpecs []*ConstSpec
	rParen     *lex.Token
}

func (a *ConstDecl) Const_() *lex.Token {
	return a.const_
}

func (a *ConstDecl) SetConst_(const_ *lex.Token) {
	a.const_ = const_
}

func (a *ConstDecl) LParen() *lex.Token {
	return a.lParen
}

func (a *ConstDecl) SetLParen(lParen *lex.Token) {
	a.lParen = lParen
}

func (a *ConstDecl) ConstSpecs() []*ConstSpec {
	return a.constSpecs
}

func (a *ConstDecl) SetConstSpecs(constSpecs []*ConstSpec) {
	a.constSpecs = constSpecs
}

func (a *ConstDecl) RParen() *lex.Token {
	return a.rParen
}

func (a *ConstDecl) SetRParen(rParen *lex.Token) {
	a.rParen = rParen
}

func (a *ConstDecl) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.const_)
	if a.lParen != nil {
		cb.AppendToken(a.lParen).Newline()
		for _, spec := range a.constSpecs {
			cb.AppendTreeNode(spec).Newline()
		}
		cb.AppendToken(a.rParen)
	} else {
		for _, spec := range a.constSpecs {
			cb.AppendTreeNode(spec)
		}
	}
	return cb
}

func (a *ConstDecl) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ConstDecl)(nil)

func (c ConstDecl) __Statement__() {
	panic("imposible")
}

func (c ConstDecl) __IFunctionMethodDeclaration__() {
	panic("imposible")
}

func (c ConstDecl) __Declaration__() {
	panic("imposible")
}

var _ Declaration = (*ConstDecl)(nil)

func VisitConstDecl(lexer *lex.Lexer) *ConstDecl {
	const_ := lexer.LA()
	if const_.Type_() != lex.GoLexerCONST {
		return nil
	}
	lexer.Pop()

	lParen := lexer.LA()
	if lParen.Type_() == lex.GoLexerL_PAREN {
		lexer.Pop()
	} else {
		lParen = nil
	}

	var constSpecs []*ConstSpec

	if lParen == nil {
		constSpec := VisitConstSpec(lexer)
		if constSpec == nil {
			fmt.Printf("constDecl,const后面跟的东西不对。%s\n")
			return nil
		}
		constSpecs = append(constSpecs, constSpec)

		return &ConstDecl{const_: const_, constSpecs: constSpecs}
	} else {
		for true {
			constSpec := VisitConstSpec(lexer)
			if constSpec != nil {
				constSpecs = append(constSpecs, constSpec)
				VisitEos(lexer)
			} else {
				break
			}
		}

		rParen := lexer.LA()
		if rParen.Type_() != lex.GoLexerR_PAREN {
			fmt.Printf("constDecl,有左括号，但没找到右括号。%s\n", rParen.ErrorMsg())
			return nil
		}
		lexer.Pop()

		return &ConstDecl{const_: const_, lParen: lParen, constSpecs: constSpecs, rParen: rParen}
	}
}

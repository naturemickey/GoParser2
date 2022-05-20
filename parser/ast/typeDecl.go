package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type TypeDecl struct {
	// typeDecl: TYPE (typeSpec | L_PAREN (typeSpec eos)* R_PAREN);
	type_     *lex.Token
	lParen    *lex.Token
	typeSpecs []*TypeSpec
	rParen    *lex.Token
}

func (a *TypeDecl) TypeSpecs() []*TypeSpec {
	return a.typeSpecs
}

func (a *TypeDecl) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.type_)
	if a.lParen != nil {
		cb.AppendToken(a.lParen).Newline()
		for _, spec := range a.typeSpecs {
			cb.AppendTreeNode(spec).Newline()
		}
		cb.AppendToken(a.rParen)
	} else {
		cb.AppendTreeNode(a.typeSpecs[0])
	}
	return cb
}

func (a *TypeDecl) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*TypeDecl)(nil)

func (t TypeDecl) __Statement__() {
	panic("imposible")
}

func (t TypeDecl) __IFunctionMethodDeclaration__() {
	panic("imposible")
}

func (t TypeDecl) __Declaration__() {
	panic("imposible")
}

var _ Declaration = (*TypeDecl)(nil)

func VisitTypeDecl(lexer *lex.Lexer) *TypeDecl {
	clone := lexer.Clone()

	type_ := lexer.LA()
	if type_.Type_() != lex.GoLexerTYPE {
		return nil
	}
	lexer.Pop() // type_

	lParen := lexer.LA()
	if lParen.Type_() == lex.GoLexerL_PAREN {
		lexer.Pop()
	} else {
		lParen = nil
	}

	var typeSpecs []*TypeSpec
	if lParen == nil {
		typeSpec := VisitTypeSpec(lexer)
		if typeSpec == nil {
			fmt.Printf("typeDecl,type后面要跟着类型定义。%s\n", type_.ErrorMsg())
			lexer.Recover(clone)
			return nil
		} else {
			typeSpecs = append(typeSpecs, typeSpec)
		}
		return &TypeDecl{type_: type_, typeSpecs: typeSpecs}
	} else {
		for {
			typeSpec := VisitTypeSpec(lexer)
			if typeSpec != nil {
				typeSpecs = append(typeSpecs, typeSpec)
			} else {
				break
			}
		}
		rParen := lexer.LA()
		if rParen.Type_() != lex.GoLexerR_PAREN {
			fmt.Printf("typeDecl,此处应该有一个')'。%s\n", rParen.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		lexer.Pop()
		return &TypeDecl{type_: type_, lParen: lParen, typeSpecs: typeSpecs, rParen: rParen}
	}
}

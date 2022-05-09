package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"fmt"
)

type TypeDecl struct {
	// typeDecl: TYPE (typeSpec | L_PAREN (typeSpec eos)* R_PAREN);
	type_     *lex.Token
	lParen    *lex.Token
	typeSpecs []*TypeSpec
	rParen    *lex.Token
}

func (a *TypeDecl) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*TypeDecl)(nil)

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
			fmt.Printf("type后面要跟着类型定义。%s\n", type_.ErrorMsg())
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
			fmt.Printf("此处应该有一个')'。%s\n", rParen.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return &TypeDecl{type_: type_, lParen: lParen, typeSpecs: typeSpecs, rParen: rParen}
	}
}

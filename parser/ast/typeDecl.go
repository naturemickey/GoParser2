package ast

import (
	"GoParser2/lex"
	"fmt"
)

type TypeDecl struct {
	// typeDecl: TYPE (typeSpec | L_PAREN (typeSpec eos)* R_PAREN);
	typeToken *lex.Token
	typeSpecs []*TypeSpec
}

func (t TypeDecl) __IFunctionMethodDeclaration__() {
	//TODO implement me
	panic("implement me")
}

func (t TypeDecl) __Declaration__() {
	//TODO implement me
	panic("implement me")
}

var _ Declaration = (*TypeDecl)(nil)

func VisitTypeDecl(lexer *lex.Lexer) *TypeDecl {
	clone := lexer.Clone()

	typeToken := lexer.LA()
	if typeToken.Type_() != lex.GoLexerTYPE {
		return nil
	}
	lexer.Pop()

	la := lexer.LA()
	hasParen := false
	if la.Type_() == lex.GoLexerL_PAREN {
		hasParen = true
		lexer.Pop()
	}

	var typeSpecs []*TypeSpec
	if !hasParen {
		typeSpec := VisitTypeSpec(lexer)
		if typeSpec == nil {
			fmt.Printf("type后面要跟着类型定义。%s\n", la)
			lexer.Recover(clone)
			return nil
		} else {
			typeSpecs = append(typeSpecs, typeSpec)
		}
	} else {
		for true {
			typeSpec := VisitTypeSpec(lexer)
			if typeSpec != nil {
				typeSpecs = append(typeSpecs, typeSpec)
			}
			break
		}
	}
	return &TypeDecl{typeToken: typeToken, typeSpecs: typeSpecs}
}

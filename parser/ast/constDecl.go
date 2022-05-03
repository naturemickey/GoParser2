package ast

import (
	"GoParser2/lex"
	"fmt"
)

type ConstDecl struct {
	// constDecl: CONST (constSpec | L_PAREN (constSpec eos)* R_PAREN);
	const_     *lex.Token
	constSpecs []*ConstSpec
}

func (c ConstDecl) __IFunctionMethodDeclaration__() {
	//TODO implement me
	panic("implement me")
}

func (c ConstDecl) __Declaration__() {
	//TODO implement me
	panic("implement me")
}

var _ Declaration = (*ConstDecl)(nil)

func VisitConstDecl(lexer *lex.Lexer) *ConstDecl {
	const_ := lexer.LA()
	if const_.Type_() != lex.GoLexerCONST {
		return nil
	}
	lexer.Pop()

	la := lexer.LA()
	hasParen := false
	if la.Type_() == lex.GoLexerL_PAREN {
		hasParen = true
		lexer.Pop()
	}

	var constSpecs []*ConstSpec

	if !hasParen {
		constSpec := VisitConstSpec(lexer)
		if constSpec == nil {
			fmt.Printf("const后面跟的东西不对。%s\n")
			return nil
		}
		constSpecs = append(constSpecs, constSpec)
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

		la := lexer.LA()
		if la.Type_() != lex.GoLexerR_PAREN {
			fmt.Printf("有左括号，但没找到右括号。%s\n", la.ErrorMsg())
			return nil
		}
	}

	return &ConstDecl{const_: const_, constSpecs: constSpecs}
}

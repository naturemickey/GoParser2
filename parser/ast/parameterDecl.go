package ast

import (
	"GoParser2/lex"
)

type ParameterDecl struct {
	// parameterDecl: identifierList? ELLIPSIS? type_;
	identifierList *IdentifierList
	ellipsis       *lex.Token
	type_          *Type_
}

func VisitParameterDecl(lexer *lex.Lexer) *ParameterDecl {
	clone := lexer.Clone()
	identifierList := VisitIdentifierList(lexer)
	ellipsis := lexer.LA()
	if ellipsis.Type_() == lex.GoLexerELLIPSIS {
		lexer.Pop() // ellipsis
	} else {
		ellipsis = nil
	}
	type_ := VisitType_(lexer)
	if type_ == nil {
		// int, int 这样的会被前面的identifierList抢到
		lexer.Recover(clone)
		type_ := VisitType_(lexer)
		if type_ == nil {
			return nil
		}
		return &ParameterDecl{type_: type_}
	}
	return &ParameterDecl{identifierList: identifierList, ellipsis: ellipsis, type_: type_}
}

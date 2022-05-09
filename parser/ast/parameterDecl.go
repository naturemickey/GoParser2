package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
)

type ParameterDecl struct {
	// parameterDecl: identifierList? ELLIPSIS? type_;
	identifierList *IdentifierList
	ellipsis       *lex.Token
	type_          *Type_
}

func (a *ParameterDecl) CodeBuilder() *util.CodeBuilder {
	return util.NewCB().AppendTreeNode(a.identifierList).AppendToken(a.ellipsis).AppendTreeNode(a.type_)
}

func (a *ParameterDecl) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*ParameterDecl)(nil)

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

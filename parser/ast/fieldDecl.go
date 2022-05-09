package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
)

type FieldDecl struct {
	// fieldDecl: (
	//		identifierList type_
	//		| embeddedField
	//	) tag = string_?;
	// string_: RAW_STRING_LIT | INTERPRETED_STRING_LIT;

	identifierList *IdentifierList
	type_          *Type_
	embeddedField  *EmbeddedField
	tag            *lex.Token
}

func (a *FieldDecl) CodeBuilder() *util.CodeBuilder {
	cb := util.NewCB()
	cb.AppendTreeNode(a.identifierList).AppendTreeNode(a.type_)
	cb.AppendTreeNode(a.embeddedField)
	cb.AppendToken(a.tag)
	return cb
}

func (a *FieldDecl) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*FieldDecl)(nil)

func VisitFieldDecl(lexer *lex.Lexer) *FieldDecl {
	clone := lexer.Clone()

	var (
		identifierList *IdentifierList
		type_          *Type_
		embeddedField  *EmbeddedField
		tag            *lex.Token
	)

	identifierList = VisitIdentifierList(lexer)
	if identifierList != nil {
		type_ = VisitType_(lexer)
		if type_ == nil {
			identifierList = nil
			lexer.Recover(clone)
		}
	}
	if identifierList == nil {
		embeddedField = VisitEmbeddedField(lexer)
		if embeddedField == nil {
			lexer.Recover(clone)
			return nil
		}
	}

	tag = lexer.LA()
	if tag.Type_() != lex.GoLexerRAW_STRING_LIT && tag.Type_() != lex.GoLexerINTERPRETED_STRING_LIT {
		tag = nil
	} else {
		lexer.Pop() // tag
	}

	return &FieldDecl{identifierList: identifierList, type_: type_, embeddedField: embeddedField, tag: tag}
}

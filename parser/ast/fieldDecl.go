package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type FieldDecl struct {
	// fieldDecl: annotationList (
	//		identifierList type_
	//		| embeddedField
	//	) tag = string_?;
	// string_: RAW_STRING_LIT | INTERPRETED_STRING_LIT;

	annotationList *AnnotationList
	identifierList *IdentifierList
	type_          *Type_
	embeddedField  *EmbeddedField
	tag            *lex.Token
}

func (a *FieldDecl) SetAnnotationList(annotationList *AnnotationList) {
	a.annotationList = annotationList
}

func (a *FieldDecl) IdentifierList() *IdentifierList {
	return a.identifierList
}

func (a *FieldDecl) SetIdentifierList(identifierList *IdentifierList) {
	a.identifierList = identifierList
}

func (a *FieldDecl) Type_() *Type_ {
	return a.type_
}

func (a *FieldDecl) SetType_(type_ *Type_) {
	a.type_ = type_
}

func (a *FieldDecl) EmbeddedField() *EmbeddedField {
	return a.embeddedField
}

func (a *FieldDecl) SetEmbeddedField(embeddedField *EmbeddedField) {
	a.embeddedField = embeddedField
}

func (a *FieldDecl) Tag() *lex.Token {
	return a.tag
}

func (a *FieldDecl) SetTag(tag *lex.Token) {
	a.tag = tag
}

func (a *FieldDecl) AnnotationList() *AnnotationList {
	return a.annotationList
}

func (a *FieldDecl) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendTreeNode(a.identifierList).AppendTreeNode(a.type_)
	cb.AppendTreeNode(a.embeddedField)
	cb.AppendToken(a.tag)
	return cb
}

func (a *FieldDecl) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*FieldDecl)(nil)

func VisitFieldDecl(lexer *lex.Lexer) *FieldDecl {
	clone := lexer.Clone()

	var (
		annotationList *AnnotationList
		identifierList *IdentifierList
		type_          *Type_
		embeddedField  *EmbeddedField
		tag            *lex.Token
	)

	annotationList = VisitAnnotationList(lexer)

	identifierList = VisitIdentifierList(lexer)
	if identifierList != nil {
		last_id := identifierList.identifiers[len(identifierList.identifiers)-1]
		la := lexer.LA()

		if last_id.Line() < la.Line() { // 后面的type不能换行
			identifierList = nil
			lexer.Recover(clone)
		} else {
			type_ = VisitType_(lexer)
			if type_ == nil {
				identifierList = nil
				lexer.Recover(clone)
			}
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

	return &FieldDecl{annotationList: annotationList, identifierList: identifierList, type_: type_, embeddedField: embeddedField, tag: tag}
}

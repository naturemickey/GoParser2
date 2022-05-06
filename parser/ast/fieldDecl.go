package ast

import "GoParser2/lex"

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

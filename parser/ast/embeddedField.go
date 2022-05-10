package ast

import (
	"GoParser2/lex"
)

type EmbeddedField struct {
	// embeddedField: STAR? typeName;
	star     *lex.Token
	typeName *TypeName
}

func (a *EmbeddedField) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.star)
	cb.AppendTreeNode(a.typeName)
	return cb
}

func (a *EmbeddedField) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*EmbeddedField)(nil)

func VisitEmbeddedField(lexer *lex.Lexer) *EmbeddedField {
	clone := lexer.Clone()

	star := lexer.LA()
	if star.Type_() != lex.GoLexerSTAR {
		star = nil
	} else {
		lexer.Pop() // star
	}

	typeName := VisitTypeName(lexer)
	if typeName == nil {
		lexer.Recover(clone)
		return nil
	}

	return &EmbeddedField{star: star, typeName: typeName}
}

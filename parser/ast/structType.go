package ast

import (
	"GoParser2/lex"
)

type StructType struct {
	// structType: STRUCT L_CURLY (fieldDecl eos)* R_CURLY;
	struct_    *lex.Token
	lCurly     *lex.Token
	fieldDecls []*FieldDecl
	rCurly     *lex.Token
}

func (a *StructType) CodeBuilder() *CodeBuilder {
	cb := NewCB().AppendToken(a.struct_).AppendToken(a.lCurly)
	if len(a.fieldDecls) > 0 {
		cb.Newline()
		for _, decl := range a.fieldDecls {
			cb.AppendTreeNode(decl).Newline()
		}
	}
	cb.AppendToken(a.rCurly)
	return cb
}

func (a *StructType) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*StructType)(nil)

func (s StructType) __TypeLit__() {
	panic("imposible")
}

var _ TypeLit = (*StructType)(nil)

func VisitStructType(lexer *lex.Lexer) *StructType {
	clone := lexer.Clone()

	struct_ := lexer.LA()
	if struct_.Type_() != lex.GoLexerSTRUCT {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // struct_

	lCurly := lexer.LA()
	if lCurly.Type_() != lex.GoLexerL_CURLY {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lCurly

	var fieldDecls []*FieldDecl
	for {
		fieldDecl := VisitFieldDecl(lexer)
		if fieldDecl != nil {
			fieldDecls = append(fieldDecls, fieldDecl)
			VisitEos(lexer)
		} else {
			break
		}
	}

	rCurly := lexer.LA()
	if rCurly.Type_() != lex.GoLexerR_CURLY {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rCurly

	return &StructType{struct_: struct_, lCurly: lCurly, fieldDecls: fieldDecls, rCurly: rCurly}
}

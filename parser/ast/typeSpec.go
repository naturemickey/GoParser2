package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type TypeSpec struct {
	// typeSpec: annotationList IDENTIFIER ASSIGN? type_;
	annotationList *AnnotationList
	identifier     *lex.Token
	assign         *lex.Token
	type_          *Type_
}

func (this *TypeSpec) AnnotationList() *AnnotationList {
	return this.annotationList
}

func (this *TypeSpec) Name() string {
	return this.identifier.Literal()
}

func (a *TypeSpec) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.identifier).AppendToken(a.assign).AppendTreeNode(a.type_)
}

func (a *TypeSpec) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*TypeSpec)(nil)

func VisitTypeSpec(lexer *lex.Lexer) *TypeSpec {
	clone := lexer.Clone()

	annotationList := VisitAnnotationList(lexer)

	identifier := lexer.LA()
	if identifier == nil || identifier.Type_() != lex.GoLexerIDENTIFIER {
		return nil
	}
	lexer.Pop()

	assign := lexer.LA()
	if assign.Type_() != lex.GoLexerASSIGN {
		assign = nil
	} else {
		lexer.Pop() // assign
	}

	type_ := VisitType_(lexer)

	if type_ == nil {
		fmt.Printf("typeSpec,后面没看到类型的描述。%s\n", identifier.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &TypeSpec{annotationList: annotationList, identifier: identifier, assign: assign, type_: type_}
}

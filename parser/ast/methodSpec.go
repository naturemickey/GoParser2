package ast

import (
	"GoParser2/lex"
)

type MethodSpec struct {
	// methodSpec:
	//	IDENTIFIER annotationList parameters result
	//	| IDENTIFIER annotationList parameters;
	identifier     *lex.Token
	annotationList *AnnotationList
	parameters     *Parameters
	result         Result
}

func (a *MethodSpec) AnnotationList() *AnnotationList {
	return a.annotationList
}

func (a *MethodSpec) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.identifier).AppendTreeNode(a.parameters).AppendTreeNode(a.result)
}

func (a *MethodSpec) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*MethodSpec)(nil)

func (m MethodSpec) __IMethodspecOrTypename__() {
	panic("imposible")
}

var _ IMethodspecOrTypename = (*MethodSpec)(nil)

func VisitMethodSpec(lexer *lex.Lexer) *MethodSpec {
	clone := lexer.Clone()

	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		return nil
	}
	lexer.Pop() // identifier

	annotationList := VisitAnnotationList(lexer)

	parameters := VisitParameters(lexer)
	if parameters == nil {
		//fmt.Printf("methodSpec,方法名后面找不到参数列表。%s\n", identifier.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	paren := parameters.rParen
	la := lexer.LA()
	if la != nil && la.Line() > paren.Line() {
		// result如果和parameters换行了，那么result很可能是一个接口继承
		return &MethodSpec{identifier: identifier, parameters: parameters}
	} else {
		result := VisitResult(lexer)
		return &MethodSpec{identifier: identifier, annotationList: annotationList, parameters: parameters, result: result}
	}
}

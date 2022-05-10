package ast

import (
	"GoParser2/lex"
)

type Signature struct {
	// signature:
	//	parameters result
	//	| parameters;
	parameters *Parameters
	result     Result
}

func (a *Signature) CodeBuilder() *CodeBuilder {
	return NewCB().AppendTreeNode(a.parameters).AppendTreeNode(a.result)
}

func (a *Signature) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*Signature)(nil)

func VisitSignature(lexer *lex.Lexer) *Signature {
	parameters := VisitParameters(lexer)
	if parameters == nil {
		return nil
	}

	result := VisitResult(lexer)

	return &Signature{parameters: parameters, result: result}
}

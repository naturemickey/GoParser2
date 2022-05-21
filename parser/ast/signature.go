package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type Signature struct {
	// signature:
	//	parameters result
	//	| parameters;
	parameters *Parameters
	result     Result
}

func (a *Signature) Parameters() *Parameters {
	return a.parameters
}

func (a *Signature) SetParameters(parameters *Parameters) {
	a.parameters = parameters
}

func (a *Signature) Result() Result {
	return a.result
}

func (a *Signature) SetResult(result Result) {
	a.result = result
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
	var result Result

	if parameters == nil {
		return nil
	}

	rParen := parameters.rParen
	la := lexer.LA()
	if la != nil && la.Line() == rParen.Line() {
		result = VisitResult(lexer)
	}

	return &Signature{parameters: parameters, result: result}
}

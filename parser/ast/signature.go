package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
)

type Signature struct {
	// signature:
	//	parameters result
	//	| parameters;
	parameters *Parameters
	result     Result
}

func (a *Signature) CodeBuilder() *util.CodeBuilder {
	return util.NewCB().AppendTreeNode(a.parameters).AppendTreeNode(a.result)
}

func (a *Signature) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*Signature)(nil)

func VisitSignature(lexer *lex.Lexer) *Signature {
	parameters := VisitParameters(lexer)
	if parameters == nil {
		return nil
	}

	result := VisitResult(lexer)

	return &Signature{parameters: parameters, result: result}
}

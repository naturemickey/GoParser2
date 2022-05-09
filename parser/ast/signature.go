package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type Signature struct {
	// signature:
	//	parameters result
	//	| parameters;
	parameters *Parameters
	result     Result
}

func (a *Signature) String() string {
	//TODO implement me
	panic("implement me")
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

package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type Result interface {
	parser.ITreeNode
	// result: parameters | type_;

	__Result__()
}

func VisitResult(lexer *lex.Lexer) Result {
	parameters := VisitParameters(lexer)
	if parameters != nil {
		return parameters
	}

	type_ := VisitType_(lexer)
	if type_ != nil {
		return type_
	}

	return nil
}

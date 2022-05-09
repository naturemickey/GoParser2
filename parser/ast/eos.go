package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type Eos struct {
	semi *lex.Token
}

func (a *Eos) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*Eos)(nil)

func VisitEos(lexer *lex.Lexer) *Eos {
	semi := lexer.LA()
	if semi != nil && semi.Type_() == lex.GoLexerSEMI {
		lexer.Pop() // semi
		return &Eos{semi: semi}
	}
	return nil
}

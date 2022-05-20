package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type Eos struct {
	semi *lex.Token
}

func (a *Eos) CodeBuilder() *CodeBuilder {
	return NewCB().AppendString(";")
}

func (a *Eos) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*Eos)(nil)

func VisitEos(lexer *lex.Lexer) *Eos {
	semi := lexer.LA()
	if semi != nil && semi.Type_() == lex.GoLexerSEMI {
		lexer.Pop() // semi
		return &Eos{semi: semi}
	}
	return nil
}

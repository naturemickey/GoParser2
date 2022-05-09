package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
)

type Eos struct {
	semi *lex.Token
}

func (a *Eos) CodeBuilder() *util.CodeBuilder {
	return util.NewCB().AppendString(";")
}

func (a *Eos) String() string {
	return a.CodeBuilder().String()
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

package ast

import "GoParser2/lex"

type Eos struct {
	semi *lex.Token
}

func VisitEos(lexer *lex.Lexer) *Eos {
	semi := lexer.LA()
	if semi.Type_() == lex.GoLexerSEMI {
		lexer.Pop() // semi
		return &Eos{semi: semi}
	}
	return nil
}

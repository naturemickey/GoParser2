package ast

import "GoParser2/lex"

type Index struct {
	// index: L_BRACKET expression R_BRACKET;
	lBracket   *lex.Token
	expression *Expression
	rBracket   *lex.Token
}

func VisitIndex(lexer *lex.Lexer) *Index {
	clone := lexer.Clone()

	lBracket := lexer.LA()
	if lBracket.Type_() != lex.GoLexerL_BRACKET {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lBracket

	expression := VisitExpression(lexer)
	if expression == nil {
		lexer.Recover(clone)
		return nil
	}

	rBracket := lexer.LA()
	if rBracket.Type_() != lex.GoLexerR_BRACKET {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rBracket

	return &Index{lBracket: lBracket, expression: expression, rBracket: rBracket}
}

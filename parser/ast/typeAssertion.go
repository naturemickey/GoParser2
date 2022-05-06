package ast

import "GoParser2/lex"

type TypeAssertion struct {
	// typeAssertion: DOT L_PAREN type_ R_PAREN;
	dot    *lex.Token
	lParen *lex.Token
	type_  *Type_
	rParen *lex.Token
}

func VisitTypeAssertion(lexer *lex.Lexer) *TypeAssertion {
	clone := lexer.Clone()

	dot := lexer.LA()
	if dot.Type_() != lex.GoLexerDOT {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // dot

	lParen := lexer.LA()
	if lParen.Type_() != lex.GoLexerR_PAREN {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lParen

	type_ := VisitType_(lexer)
	if type_ == nil {
		lexer.Recover(clone)
		return nil
	}

	rParen := lexer.LA()
	if rParen.Type_() != lex.GoLexerR_PAREN {
		lexer.Recover(clone)
		return nil
	}

	return &TypeAssertion{dot: dot, lParen: lParen, type_: type_, rParen: rParen}
}

package ast

import "GoParser2/lex"

type Slice struct {
	// slice:
	//	L_BRACKET (
	//		expression? COLON expression?
	//		| expression? COLON expression COLON expression
	//	) R_BRACKET;
	lBracket *lex.Token
	rBracket *lex.Token

	expression1 *Expression
	expression2 *Expression
	expression3 *Expression

	colon1 *lex.Token
	colon2 *lex.Token
}

func VisitSlice(lexer *lex.Lexer) *Slice {
	clone := lexer.Clone()

	lBracket := lexer.LA()
	if lBracket.Type_() != lex.GoLexerL_BRACKET {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lBracket

	expression1 := VisitExpression(lexer)

	colon1 := lexer.LA()
	if colon1.Type_() != lex.GoLexerCOLON {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // colon1

	expression2 := VisitExpression(lexer)
	var expression3 *Expression

	colon2 := lexer.LA()
	if colon2.Type_() != lex.GoLexerCOLON {
		lexer.Recover(clone)
		colon2 = nil
	} else {
		lexer.Pop() // colon2
		expression3 = VisitExpression(lexer)
	}

	if colon2 != nil {
		if expression2 == nil || expression3 == nil {
			lexer.Recover(clone)
			return nil
		}
	}

	rBracket := lexer.LA()
	if rBracket.Type_() != lex.GoLexerR_BRACKET {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rBracket

	return &Slice{lBracket: lBracket, rBracket: rBracket, colon1: colon1, colon2: colon2,
		expression1: expression1, expression2: expression2, expression3: expression3}
}

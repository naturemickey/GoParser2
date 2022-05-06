package ast

import "GoParser2/lex"

type LiteralValue struct {
	// literalValue: L_CURLY (elementList COMMA?)? R_CURLY;
	lCurly      *lex.Token
	elementList *ElementList
	comma       *lex.Token
	rCurly      *lex.Token
}

func (l LiteralValue) __Key__() {
	//TODO implement me
	panic("implement me")
}

func (l LiteralValue) __Element__() {
	//TODO implement me
	panic("implement me")
}

var _ Element = (*LiteralValue)(nil)
var _ Key = (*LiteralValue)(nil)

func VisitLiteralValue(lexer *lex.Lexer) *LiteralValue {
	clone := lexer.Clone()

	lCurly := lexer.LA()
	if lCurly.Type_() != lex.GoLexerL_CURLY {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lCurly

	elementList := VisitElementList(lexer)
	var comma *lex.Token
	if elementList != nil {
		comma = lexer.LA()
		if comma.Type_() != lex.GoLexerCOMMA {
			comma = nil
		}
	}

	rCurly := lexer.LA()
	if rCurly.Type_() != lex.GoLexerR_CURLY {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rCurly

	return &LiteralValue{lCurly: lCurly, elementList: elementList, comma: comma, rCurly: rCurly}

}

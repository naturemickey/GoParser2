package ast

import (
	"GoParser2/lex"
)

type LiteralValue struct {
	// literalValue: L_CURLY (elementList COMMA?)? R_CURLY;
	lCurly      *lex.Token
	elementList *ElementList
	comma       *lex.Token
	rCurly      *lex.Token
}

func (a *LiteralValue) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.lCurly).AppendTreeNode(a.elementList).AppendToken(a.comma).AppendToken(a.rCurly)
}

func (a *LiteralValue) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*LiteralValue)(nil)

func (l LiteralValue) __Key__() {
	panic("imposible")
}

func (l LiteralValue) __Element__() {
	panic("imposible")
}

var _ Element = (*LiteralValue)(nil)
var _ Key = (*LiteralValue)(nil)

func VisitLiteralValue(lexer *lex.Lexer) *LiteralValue {
	clone := lexer.Clone()

	lCurly := lexer.LA()
	if lCurly == nil || lCurly.Type_() != lex.GoLexerL_CURLY {
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
		} else {
			lexer.Pop() // comma
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

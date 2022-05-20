package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type BasicLit struct {
	// basicLit:
	//	NIL_LIT
	//	| integer
	//	| string_
	//	| FLOAT_LIT;

	// integer:
	//	DECIMAL_LIT
	//	| BINARY_LIT
	//	| OCTAL_LIT
	//	| HEX_LIT
	//	| IMAGINARY_LIT
	//	| RUNE_LIT;

	// string_: RAW_STRING_LIT | INTERPRETED_STRING_LIT;

	nil_lit   *lex.Token
	integer   *lex.Token
	string_   *lex.Token
	float_lit *lex.Token
}

func (a *BasicLit) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.nil_lit)
	cb.AppendToken(a.integer)
	cb.AppendToken(a.string_)
	cb.AppendToken(a.float_lit)
	return cb
}

func (a *BasicLit) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*BasicLit)(nil)

func (b BasicLit) __Literal__() {
	panic("imposible")
}

var _ Literal = (*BasicLit)(nil)

func VisitBasicLit(lexer *lex.Lexer) *BasicLit {
	la := lexer.LA()
	switch la.Type_() {
	case lex.GoLexerNIL_LIT:
		lexer.Pop() // nil_lit
		return &BasicLit{nil_lit: la}
	case lex.GoLexerDECIMAL_LIT, lex.GoLexerBINARY_LIT, lex.GoLexerOCTAL_LIT,
		lex.GoLexerHEX_LIT, lex.GoLexerIMAGINARY_LIT, lex.GoLexerRUNE_LIT:
		lexer.Pop() // integer
		return &BasicLit{integer: la}
	case lex.GoLexerRAW_STRING_LIT, lex.GoLexerINTERPRETED_STRING_LIT:
		lexer.Pop() // string
		return &BasicLit{string_: la}
	case lex.GoLexerFLOAT_LIT:
		lexer.Pop() // float_lit
		return &BasicLit{float_lit: la}
	default:
		return nil
	}
}

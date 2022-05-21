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

func (a *BasicLit) Nil_lit() *lex.Token {
	return a.nil_lit
}

func (a *BasicLit) SetNil_lit(nil_lit *lex.Token) {
	a.nil_lit = nil_lit
}

func (a *BasicLit) Integer() *lex.Token {
	return a.integer
}

func (a *BasicLit) SetInteger(integer *lex.Token) {
	a.integer = integer
}

func (a *BasicLit) String_() *lex.Token {
	return a.string_
}

func (a *BasicLit) SetString_(string_ *lex.Token) {
	a.string_ = string_
}

func (a *BasicLit) Float_lit() *lex.Token {
	return a.float_lit
}

func (a *BasicLit) SetFloat_lit(float_lit *lex.Token) {
	a.float_lit = float_lit
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

package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type Assign_op struct {
	// assign_op: prefix=(
	//		PLUS
	//		| MINUS
	//		| OR
	//		| CARET
	//		| STAR
	//		| DIV
	//		| MOD
	//		| LSHIFT
	//		| RSHIFT
	//		| AMPERSAND
	//		| BIT_CLEAR
	//	)? ASSIGN;

	prefix *lex.Token
	assign *lex.Token
}

func (a *Assign_op) Prefix() *lex.Token {
	return a.prefix
}

func (a *Assign_op) SetPrefix(prefix *lex.Token) {
	a.prefix = prefix
}

func (a *Assign_op) Assign() *lex.Token {
	return a.assign
}

func (a *Assign_op) SetAssign(assign *lex.Token) {
	a.assign = assign
}

func (a *Assign_op) CodeBuilder() *CodeBuilder {
	if a.prefix != nil {
		return NewCB().AppendString(a.prefix.Literal() + a.assign.Literal())
	} else {
		return NewCB().AppendToken(a.assign)

	}
}

func (a *Assign_op) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*Assign_op)(nil)

func VisitAssign_op(lexer *lex.Lexer) *Assign_op {
	clone := lexer.Clone()
	la := lexer.LA()
	if la == nil {
		return nil
	}
	if la.Type_() == lex.GoLexerASSIGN {
		lexer.Pop()
		return &Assign_op{assign: la}
	} else {
		var prefix *lex.Token
		switch la.Type_() {
		case lex.GoLexerPLUS, lex.GoLexerMINUS, lex.GoLexerOR, lex.GoLexerCARET, lex.GoLexerSTAR,
			lex.GoLexerDIV, lex.GoLexerMOD, lex.GoLexerLSHIFT, lex.GoLexerRSHIFT, lex.GoLexerAMPERSAND, lex.GoLexerBIT_CLEAR:
			prefix = la
			lexer.Pop() // prefix
		default:
			lexer.Recover(clone)
			return nil
		}

		la := lexer.LA()
		if la.Type_() != lex.GoLexerASSIGN {
			//fmt.Printf("assign_op,等号在哪里？%s\n", la)
			lexer.Recover(clone)
			return nil
		}
		lexer.Pop()

		return &Assign_op{prefix: prefix, assign: la}
	}
}

package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

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

func (a *Slice) LBracket() *lex.Token {
	return a.lBracket
}

func (a *Slice) SetLBracket(lBracket *lex.Token) {
	a.lBracket = lBracket
}

func (a *Slice) RBracket() *lex.Token {
	return a.rBracket
}

func (a *Slice) SetRBracket(rBracket *lex.Token) {
	a.rBracket = rBracket
}

func (a *Slice) Expression1() *Expression {
	return a.expression1
}

func (a *Slice) SetExpression1(expression1 *Expression) {
	a.expression1 = expression1
}

func (a *Slice) Expression2() *Expression {
	return a.expression2
}

func (a *Slice) SetExpression2(expression2 *Expression) {
	a.expression2 = expression2
}

func (a *Slice) Expression3() *Expression {
	return a.expression3
}

func (a *Slice) SetExpression3(expression3 *Expression) {
	a.expression3 = expression3
}

func (a *Slice) Colon1() *lex.Token {
	return a.colon1
}

func (a *Slice) SetColon1(colon1 *lex.Token) {
	a.colon1 = colon1
}

func (a *Slice) Colon2() *lex.Token {
	return a.colon2
}

func (a *Slice) SetColon2(colon2 *lex.Token) {
	a.colon2 = colon2
}

func (a *Slice) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.lBracket).AppendTreeNode(a.expression1).AppendToken(a.colon1).
		AppendTreeNode(a.expression2).AppendToken(a.colon2).AppendTreeNode(a.expression3).AppendToken(a.rBracket)
}

func (a *Slice) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*Slice)(nil)

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
		// lexer.Recover(clone)
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

package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type IncDecStmt struct {
	// incDecStmt: expression (PLUS_PLUS | MINUS_MINUS);
	expression *Expression
	plusplus   *lex.Token
	minusminus *lex.Token
}

func (a *IncDecStmt) Expression() *Expression {
	return a.expression
}

func (a *IncDecStmt) SetExpression(expression *Expression) {
	a.expression = expression
}

func (a *IncDecStmt) Plusplus() *lex.Token {
	return a.plusplus
}

func (a *IncDecStmt) SetPlusplus(plusplus *lex.Token) {
	a.plusplus = plusplus
}

func (a *IncDecStmt) Minusminus() *lex.Token {
	return a.minusminus
}

func (a *IncDecStmt) SetMinusminus(minusminus *lex.Token) {
	a.minusminus = minusminus
}

func (a *IncDecStmt) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	return cb.AppendTreeNode(a.expression).AppendToken(a.plusplus).AppendToken(a.minusminus)
}

func (a *IncDecStmt) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*IncDecStmt)(nil)

func (i IncDecStmt) __Statement__() {
	panic("imposible")
}

func (i IncDecStmt) __SimpleStmt__() {
	panic("imposible")
}

var _ SimpleStmt = (*IncDecStmt)(nil)

func VisitIncDecStmt(lexer *lex.Lexer) *IncDecStmt {
	clone := lexer.Clone()

	expression := VisitExpression(lexer)
	if expression == nil {
		lexer.Recover(clone)
		return nil
	}

	la := lexer.LA()
	if la == nil {
		lexer.Recover(clone)
		return nil
	}
	if la.Type_() == lex.GoLexerPLUS_PLUS {
		lexer.Pop()
		return &IncDecStmt{expression: expression, plusplus: la}
	} else if la.Type_() == lex.GoLexerMINUS_MINUS {
		lexer.Pop()
		return &IncDecStmt{expression: expression, minusminus: la}
	} else {
		lexer.Recover(clone)
		return nil
	}
}

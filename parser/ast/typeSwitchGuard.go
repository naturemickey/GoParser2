package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type TypeSwitchGuard struct {
	// typeSwitchGuard: (IDENTIFIER DECLARE_ASSIGN)? primaryExpr DOT L_PAREN TYPE R_PAREN;
	identifier     *lex.Token
	declare_assign *lex.Token
	primaryExpr    *PrimaryExpr
	dot            *lex.Token
	lParen         *lex.Token
	type_          *lex.Token
	rParen         *lex.Token
}

func (a *TypeSwitchGuard) Identifier() *lex.Token {
	return a.identifier
}

func (a *TypeSwitchGuard) SetIdentifier(identifier *lex.Token) {
	a.identifier = identifier
}

func (a *TypeSwitchGuard) Declare_assign() *lex.Token {
	return a.declare_assign
}

func (a *TypeSwitchGuard) SetDeclare_assign(declare_assign *lex.Token) {
	a.declare_assign = declare_assign
}

func (a *TypeSwitchGuard) PrimaryExpr() *PrimaryExpr {
	return a.primaryExpr
}

func (a *TypeSwitchGuard) SetPrimaryExpr(primaryExpr *PrimaryExpr) {
	a.primaryExpr = primaryExpr
}

func (a *TypeSwitchGuard) Dot() *lex.Token {
	return a.dot
}

func (a *TypeSwitchGuard) SetDot(dot *lex.Token) {
	a.dot = dot
}

func (a *TypeSwitchGuard) LParen() *lex.Token {
	return a.lParen
}

func (a *TypeSwitchGuard) SetLParen(lParen *lex.Token) {
	a.lParen = lParen
}

func (a *TypeSwitchGuard) Type_() *lex.Token {
	return a.type_
}

func (a *TypeSwitchGuard) SetType_(type_ *lex.Token) {
	a.type_ = type_
}

func (a *TypeSwitchGuard) RParen() *lex.Token {
	return a.rParen
}

func (a *TypeSwitchGuard) SetRParen(rParen *lex.Token) {
	a.rParen = rParen
}

func (a *TypeSwitchGuard) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.identifier).AppendToken(a.declare_assign).AppendTreeNode(a.primaryExpr).
		AppendToken(a.dot).AppendToken(a.lParen).AppendToken(a.type_).AppendToken(a.rParen)
}

func (a *TypeSwitchGuard) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*TypeSwitchGuard)(nil)

func VisitTypeSwitchGuard(lexer *lex.Lexer) *TypeSwitchGuard {
	clone := lexer.Clone()

	declare_assign := lexer.LA1()
	var identifier *lex.Token

	if declare_assign.Type_() == lex.GoLexerDECLARE_ASSIGN {
		identifier = lexer.LA()
		if identifier.Type_() != lex.GoLexerIDENTIFIER {
			lexer.Recover(clone)
			return nil
		}
		lexer.Pop() // identifier
		lexer.Pop() // declare_assign
	} else {
		declare_assign = nil
	}

	primaryExpr := VisitPrimaryExpr(lexer)
	if primaryExpr == nil {
		lexer.Recover(clone)
		return nil
	}

	dot := lexer.LA()
	if dot.Type_() != lex.GoLexerDOT {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // dot
	lParen := lexer.LA()
	if lParen.Type_() != lex.GoLexerL_PAREN {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lParen
	type_ := lexer.LA()
	if type_.Type_() != lex.GoLexerTYPE {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // type_
	rParen := lexer.LA()
	if rParen.Type_() != lex.GoLexerR_PAREN {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rParen

	return &TypeSwitchGuard{identifier: identifier, declare_assign: declare_assign, primaryExpr: primaryExpr,
		dot: dot, lParen: lParen, type_: type_, rParen: rParen}
}

package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type MethodExpr struct {
	// methodExpr: nonNamedType DOT IDENTIFIER;
	nonNamedType *NonNamedType
	dot          *lex.Token
	identifier   *lex.Token
}

func (a *MethodExpr) NonNamedType() *NonNamedType {
	return a.nonNamedType
}

func (a *MethodExpr) SetNonNamedType(nonNamedType *NonNamedType) {
	a.nonNamedType = nonNamedType
}

func (a *MethodExpr) Dot() *lex.Token {
	return a.dot
}

func (a *MethodExpr) SetDot(dot *lex.Token) {
	a.dot = dot
}

func (a *MethodExpr) Identifier() *lex.Token {
	return a.identifier
}

func (a *MethodExpr) SetIdentifier(identifier *lex.Token) {
	a.identifier = identifier
}

func (a *MethodExpr) CodeBuilder() *CodeBuilder {
	return NewCB().AppendTreeNode(a.nonNamedType).AppendToken(a.dot).AppendToken(a.identifier)
}

func (a *MethodExpr) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*MethodExpr)(nil)

func VisitMethodExpr(lexer *lex.Lexer) *MethodExpr {
	clone := lexer.Clone()

	nonNamedType := VisitNonNamedType(lexer)
	if nonNamedType == nil {
		lexer.Recover(clone)
		return nil
	}

	dot := lexer.LA()
	if dot.Type_() != lex.GoLexerDOT {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // dot

	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // identifier

	return &MethodExpr{nonNamedType: nonNamedType, dot: dot, identifier: identifier}
}

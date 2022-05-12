package ast

import (
	"GoParser2/lex"
)

type RangeClause struct {
	// rangeClause: (
	//		expressionList ASSIGN
	//		| identifierList DECLARE_ASSIGN
	//	)? RANGE expression;
	expressionList *ExpressionList
	assign         *lex.Token

	identifierList *IdentifierList
	declare_assign *lex.Token

	range_     *lex.Token
	expression *Expression
}

func (a *RangeClause) CodeBuilder() *CodeBuilder {
	return NewCB().AppendTreeNode(a.expressionList).AppendToken(a.assign).
		AppendTreeNode(a.identifierList).AppendToken(a.declare_assign).AppendToken(a.range_).
		AppendTreeNode(a.expression)
}

func (a *RangeClause) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*RangeClause)(nil)

func VisitRangeClause(lexer *lex.Lexer) *RangeClause {
	clone := lexer.Clone()

	var expressionList *ExpressionList
	var identifierList *IdentifierList
	var assign *lex.Token
	var declare_assign *lex.Token

	expressionList, assign = _expressionList_ASSIGN(lexer)
	if expressionList == nil {
		identifierList, declare_assign = _identifierList_DECLARE_ASSIGN(lexer)
	}

	range_ := lexer.LA()
	if range_.Type_() != lex.GoLexerRANGE {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop()

	expression := VisitExpression(lexer)

	return &RangeClause{expressionList: expressionList, assign: assign, identifierList: identifierList, declare_assign: declare_assign,
		range_: range_, expression: expression}
}

func _expressionList_ASSIGN(lexer *lex.Lexer) (*ExpressionList, *lex.Token) {
	clone := lexer.Clone()

	var expressionList = VisitExpressionList(lexer)
	var assign *lex.Token

	if expressionList == nil {
		lexer.Recover(clone)
		return nil, nil
	}

	assign = lexer.LA()
	if assign.Type_() != lex.GoLexerASSIGN {
		lexer.Recover(clone)
		return nil, nil
	}
	lexer.Pop() // assign
	return expressionList, assign
}

func _identifierList_DECLARE_ASSIGN(lexer *lex.Lexer) (*IdentifierList, *lex.Token) {
	clone := lexer.Clone()

	identifierList := VisitIdentifierList(lexer)
	if identifierList == nil {
		lexer.Recover(clone)
		return nil, nil
	}
	declare_assign := lexer.LA()
	if declare_assign.Type_() != lex.GoLexerDECLARE_ASSIGN {
		lexer.Recover(clone)
		return nil, nil
	}
	lexer.Pop() // declare_assign

	return identifierList, declare_assign
}

package ast

import "GoParser2/lex"

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

func VisitRangeClause(lexer *lex.Lexer) *RangeClause {
	clone := lexer.Clone()

	var expressionList = VisitExpressionList(lexer)
	var identifierList *IdentifierList
	var assign *lex.Token
	var declare_assign *lex.Token

	if expressionList == nil {
		identifierList = VisitIdentifierList(lexer)
		if identifierList != nil {
			declare_assign = lexer.LA()
			if declare_assign.Type_() != lex.GoLexerDECLARE_ASSIGN {
				identifierList = nil
				declare_assign = nil
				lexer.Recover(clone)
			} else {
				lexer.Pop() // declare_assign
			}
		}
	} else {
		assign = lexer.LA()
		if assign.Type_() != lex.GoLexerASSIGN {
			assign = nil
			expressionList = nil
			lexer.Recover(clone)
		} else {
			lexer.Pop() // assign
		}
	}

	range_ := lexer.LA()
	if range_.Type_() != lex.GoLexerRANGE {
		return nil
	}
	lexer.Pop()

	expression := VisitExpression(lexer)

	return &RangeClause{expressionList: expressionList, assign: assign, identifierList: identifierList, declare_assign: declare_assign,
		range_: range_, expression: expression}
}

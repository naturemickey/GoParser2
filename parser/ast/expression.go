package ast

import (
	"GoParser2/lex"
	"fmt"
)

type Expression struct {
	// expression:
	//	primaryExpr
	//	| unary_op = (
	//		PLUS
	//		| MINUS
	//		| EXCLAMATION
	//		| CARET
	//		| STAR
	//		| AMPERSAND
	//		| RECEIVE
	//	) expression
	//	| expression mul_op = (
	//		STAR
	//		| DIV
	//		| MOD
	//		| LSHIFT
	//		| RSHIFT
	//		| AMPERSAND
	//		| BIT_CLEAR
	//	) expression
	//	| expression add_op = (PLUS | MINUS | OR | CARET) expression
	//	| expression rel_op = (
	//		EQUALS
	//		| NOT_EQUALS
	//		| LESS
	//		| LESS_OR_EQUALS
	//		| GREATER
	//		| GREATER_OR_EQUALS
	//	) expression
	//	| expression LOGICAL_AND expression
	//	| expression LOGICAL_OR expression;

	primaryExpr *PrimaryExpr

	unary_op   *lex.Token
	expression *Expression

	mul_op      *lex.Token
	expression2 *Expression

	add_op      *lex.Token
	rel_op      *lex.Token
	logical_and *lex.Token
	logical_or  *lex.Token
}

func (e Expression) __Key__() {
	//TODO implement me
	panic("implement me")
}

func (e Expression) __Element__() {
	//TODO implement me
	panic("implement me")
}

func (e Expression) __Statement__() {
	//TODO implement me
	panic("implement me")
}

func (e Expression) __SimpleStmt__() {
	//TODO implement me
	panic("implement me")
}

func (e Expression) __ExpressionStmt__() {
	//TODO implement me
	panic("implement me")
}

var _ ExpressionStmt = (*Expression)(nil)
var _ Element = (*Expression)(nil)
var _ Key = (*Expression)(nil)

func VisitExpression(lexer *lex.Lexer) *Expression {
	clone := lexer.Clone()
	//	unary_op = (
	//		PLUS
	//		| MINUS
	//		| EXCLAMATION
	//		| CARET
	//		| STAR
	//		| AMPERSAND
	//		| RECEIVE
	//	) expression
	la := lexer.LA()
	switch la.Type_() {
	case lex.GoLexerPLUS,
		lex.GoLexerMINUS,
		lex.GoLexerEXCLAMATION,
		lex.GoLexerCARET,
		lex.GoLexerSTAR,
		lex.GoLexerAMPERSAND,
		lex.GoLexerRECEIVE:
		unary_op := la
		lexer.Pop() // unary_op
		expression := VisitExpression(lexer)
		if expression == nil {
			fmt.Printf("这个一元符号后面应该是一个表达式。%s\n", unary_op.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return &Expression{unary_op: unary_op, expression: expression}
	}
	//	primaryExpr
	primaryExpr := VisitPrimaryExpr(lexer)
	if primaryExpr != nil {
		return &Expression{primaryExpr: primaryExpr}
	}

	// 以下所有分支都要先识别到expression
	expression := VisitExpression(lexer)
	if expression == nil {
		lexer.Recover(clone)
		return nil
	}

	//	expression mul_op = (
	//		STAR
	//		| DIV
	//		| MOD
	//		| LSHIFT
	//		| RSHIFT
	//		| AMPERSAND
	//		| BIT_CLEAR
	//	) expression
	mul_op := lexer.LA()
	switch mul_op.Type_() {
	case lex.GoLexerSTAR,
		lex.GoLexerDIV,
		lex.GoLexerMOD,
		lex.GoLexerLSHIFT,
		lex.GoLexerRSHIFT,
		lex.GoLexerAMPERSAND,
		lex.GoLexerBIT_CLEAR:
		lexer.Pop() // mul_op
		expression2 := VisitExpression(lexer)
		if expression2 == nil {
			fmt.Printf("此符号后面需要一个表达式。%s\n", mul_op.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return &Expression{expression: expression, mul_op: mul_op, expression2: expression2}
	}

	//	expression add_op = (PLUS | MINUS | OR | CARET) expression
	add_op := lexer.LA()
	switch add_op.Type_() {
	case lex.GoLexerPLUS,
		lex.GoLexerMINUS,
		lex.GoLexerOR,
		lex.GoLexerCARET:
		lexer.Pop() // add_op
		expression2 := VisitExpression(lexer)
		if expression2 == nil {
			fmt.Printf("此符号后面需要一个表达式。%s\n", add_op.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return &Expression{expression: expression, add_op: add_op, expression2: expression2}
	}

	//	expression rel_op = (
	//		EQUALS
	//		| NOT_EQUALS
	//		| LESS
	//		| LESS_OR_EQUALS
	//		| GREATER
	//		| GREATER_OR_EQUALS
	//	) expression
	rel_op := lexer.LA()
	switch rel_op.Type_() {
	case lex.GoLexerEQUALS,
		lex.GoLexerNOT_EQUALS,
		lex.GoLexerLESS,
		lex.GoLexerLESS_OR_EQUALS,
		lex.GoLexerGREATER,
		lex.GoLexerGREATER_OR_EQUALS:
		lexer.Pop() // rel_op
		expression2 := VisitExpression(lexer)
		if expression2 == nil {
			fmt.Printf("此符号后面需要一个表达式。%s\n", rel_op.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return &Expression{expression: expression, rel_op: rel_op, expression2: expression2}
	}

	//	expression LOGICAL_AND expression
	logical_and := lexer.LA()
	if logical_and.Type_() == lex.GoLexerLOGICAL_AND {
		lexer.Pop() // logical_and
		expression2 := VisitExpression(lexer)
		if expression2 == nil {
			fmt.Printf("此符号后面需要一个表达式。%s\n", logical_and.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return &Expression{expression: expression, logical_and: logical_and, expression2: expression2}
	}

	//	expression LOGICAL_OR expression
	logical_or := lexer.LA()
	if logical_or.Type_() == lex.GoLexerLOGICAL_AND {
		lexer.Pop() // logical_or
		expression2 := VisitExpression(lexer)
		if expression2 == nil {
			fmt.Printf("此符号后面需要一个表达式。%s\n", logical_or.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return &Expression{expression: expression, logical_or: logical_or, expression2: expression2}
	}

	fmt.Printf("这里应该是某个运算符吧？%s\n", logical_or.ErrorMsg())
	lexer.Recover(clone)
	return nil
}

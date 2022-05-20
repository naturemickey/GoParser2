package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type Expression struct {
	// 注意：这个定义里面有左递归

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

	// 调一下格式：
	// expression:
	//	(primaryExpr | unary_op = (PLUS | MINUS | EXCLAMATION | CARET | STAR | AMPERSAND | RECEIVE) expression)
	//	| expression ( mul_op   = (STAR | DIV | MOD | LSHIFT | RSHIFT | AMPERSAND | BIT_CLEAR)                expression
	//	             | add_op   = (PLUS | MINUS | OR | CARET)                                                 expression
	//	             | rel_op   = (EQUALS | NOT_EQUALS | LESS | LESS_OR_EQUALS | GREATER | GREATER_OR_EQUALS) expression
	//	             | LOGICAL_AND                                                                            expression
	//	             | LOGICAL_OR                                                                             expression
	//	             ) ;
	//
	// 等价于：
	// exp1 : (primaryExpr | unary_op = (PLUS | MINUS | EXCLAMATION | CARET | STAR | AMPERSAND | RECEIVE) expression)
	// exp2 :        ( mul_op   = (STAR | DIV | MOD | LSHIFT | RSHIFT | AMPERSAND | BIT_CLEAR)                expression
	//	             | add_op   = (PLUS | MINUS | OR | CARET)                                                 expression
	//	             | rel_op   = (EQUALS | NOT_EQUALS | LESS | LESS_OR_EQUALS | GREATER | GREATER_OR_EQUALS) expression
	//	             | LOGICAL_AND                                                                            expression
	//	             | LOGICAL_OR                                                                             expression
	//	             ) ;
	// expression : exp1 | expression exp2 ;
	// 左递归改造（公式）：
	// expression : exp1 exp_p
	// exp_p      : exp2 exp_p | ε
	// 简化：
	// expression : exp1 exp2*
	//

	primaryExpr *PrimaryExpr

	unary_op   *lex.Token
	expression *Expression

	exp2s []*exp2
}

func (a *Expression) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendTreeNode(a.primaryExpr)
	cb.AppendToken(a.unary_op).AppendTreeNode(a.expression)
	for _, e := range a.exp2s {
		cb.AppendToken(e.mul_op)
		cb.AppendToken(e.add_op)
		cb.AppendToken(e.rel_op)
		cb.AppendToken(e.logical_and)
		cb.AppendToken(e.logical_or)
		cb.AppendTreeNode(e.expression2)
	}
	return cb
}

func (a *Expression) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*Expression)(nil)

type exp2 struct {
	mul_op      *lex.Token
	add_op      *lex.Token
	rel_op      *lex.Token
	logical_and *lex.Token
	logical_or  *lex.Token
	expression2 *Expression
}

func (e Expression) __Key__() {
	panic("imposible")
}

func (e Expression) __Element__() {
	panic("imposible")
}

func (e Expression) __Statement__() {
	panic("imposible")
}

func (e Expression) __SimpleStmt__() {
	panic("imposible")
}

func (e Expression) __ExpressionStmt__() {
	panic("imposible")
}

var _ ExpressionStmt = (*Expression)(nil)
var _ Element = (*Expression)(nil)
var _ Key = (*Expression)(nil)

func VisitExpression(lexer *lex.Lexer) *Expression {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	clone := lexer.Clone()

	exp1 := _visitExp1(lexer)
	if exp1 == nil {
		lexer.Recover(clone)
		return nil
	}

	var exp2s []*exp2
	for {
		e2 := _visitExp2(lexer)
		if e2 == nil {
			break
		}
		exp2s = append(exp2s, e2)
	}

	return &Expression{primaryExpr: exp1.primaryExpr, unary_op: exp1.unary_op, expression: exp1.expression, exp2s: exp2s}
}

func _visitExp1(lexer *lex.Lexer) *struct {
	primaryExpr *PrimaryExpr
	unary_op    *lex.Token
	expression  *Expression
} {
	exp1 := &struct {
		primaryExpr *PrimaryExpr
		unary_op    *lex.Token
		expression  *Expression
	}{}
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
			fmt.Printf("expression,这个一元符号后面应该是一个表达式。%s\n", unary_op.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		exp1.unary_op = unary_op
		exp1.expression = expression
		return exp1
	}
	//	primaryExpr
	primaryExpr := VisitPrimaryExpr(lexer)
	if primaryExpr != nil {
		exp1.primaryExpr = primaryExpr
		return exp1
	}
	return nil
}

func _visitExp2(lexer *lex.Lexer) *exp2 {
	// expression 在符号前不换行
	lb := lexer.LB()
	la := lexer.LA()
	if lb != nil && la != nil && lb.Line() < la.Line() {
		return nil
	}

	exp2 := &exp2{}

	clone := lexer.Clone()

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

	if mul_op == nil {
		return nil
	}

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
			//fmt.Printf("expression,此符号后面需要一个表达式。%s\n", mul_op.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		exp2.mul_op = mul_op
		exp2.expression2 = expression2
		return exp2
	}
	mul_op = nil

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
			//fmt.Printf("expression,此符号后面需要一个表达式。%s\n", add_op.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		exp2.add_op = add_op
		exp2.expression2 = expression2
		return exp2
	}
	add_op = nil

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
			//fmt.Printf("expression,此符号后面需要一个表达式。%s\n", rel_op.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		exp2.rel_op = rel_op
		exp2.expression2 = expression2
		return exp2
	}
	rel_op = nil

	//	expression LOGICAL_AND expression
	logical_and := lexer.LA()
	if logical_and.Type_() == lex.GoLexerLOGICAL_AND {
		lexer.Pop() // logical_and
		expression2 := VisitExpression(lexer)
		if expression2 == nil {
			//fmt.Printf("expression,此符号后面需要一个表达式。%s\n", logical_and.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		exp2.logical_and = logical_and
		exp2.expression2 = expression2
		return exp2
	}
	logical_and = nil

	//	expression LOGICAL_OR expression
	logical_or := lexer.LA()
	if logical_or.Type_() == lex.GoLexerLOGICAL_OR {
		lexer.Pop() // logical_or
		expression2 := VisitExpression(lexer)
		if expression2 == nil {
			//fmt.Printf("expression,此符号后面需要一个表达式。%s\n", logical_or.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		exp2.logical_or = logical_or
		exp2.expression2 = expression2
		return exp2
	}
	logical_or = nil

	lexer.Recover(clone)
	return nil
}

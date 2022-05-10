package ast

import (
	"GoParser2/lex"
)

type PrimaryExpr struct {
	// primaryExpr:
	//	operand
	//	| conversion
	//	| methodExpr
	//	| primaryExpr (
	//		(DOT IDENTIFIER)
	//		| index
	//		| slice
	//		| typeAssertion
	//		| arguments
	//	);

	operand    *Operand
	conversion *Conversion
	methodExpr *MethodExpr

	primaryExpr *PrimaryExpr // 这里有一个首递归
	suffix      *PrimaryExprSuffix
}

func (a *PrimaryExpr) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendTreeNode(a.operand)
	cb.AppendTreeNode(a.conversion)
	cb.AppendTreeNode(a.methodExpr)
	cb.AppendTreeNode(a.primaryExpr)

	if a.suffix != nil {
		cb.AppendToken(a.suffix.dot).AppendToken(a.suffix.identifier)
	}

	if a.suffix != nil {
		cb.AppendTreeNode(a.suffix.index)
		cb.AppendTreeNode(a.suffix.slice)
		cb.AppendTreeNode(a.suffix.typeAssertion)
		cb.AppendTreeNode(a.suffix.arguments)
	}

	return cb
}

func (a *PrimaryExpr) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*PrimaryExpr)(nil)

type PrimaryExprSuffix struct {
	dot, identifier *lex.Token
	index           *Index
	slice           *Slice
	typeAssertion   *TypeAssertion
	arguments       *Arguments
}

func VisitPrimaryExpr(lexer *lex.Lexer) *PrimaryExpr {
	clone := lexer.Clone()

	var primaryExpr *PrimaryExpr
	operand := VisitOperand(lexer)
	if operand != nil {
		primaryExpr = &PrimaryExpr{operand: operand}
	} else {
		conversion := VisitConversion(lexer)
		if conversion != nil {
			primaryExpr = &PrimaryExpr{conversion: conversion}
		} else {
			methodExpr := VisitMethodExpr(lexer)
			if conversion != nil {
				primaryExpr = &PrimaryExpr{methodExpr: methodExpr}
			}
		}
	}
	if primaryExpr == nil {
		lexer.Recover(clone)
		return nil
	}

	for { // 处理首递归
		suffix := _tryVisitSuffix(lexer)
		if suffix == nil {
			break
		} else {
			primaryExpr = &PrimaryExpr{primaryExpr: primaryExpr, suffix: suffix}
		}
	}
	return primaryExpr
}

func _tryVisitSuffix(lexer *lex.Lexer) *PrimaryExprSuffix {
	//	(DOT IDENTIFIER)
	//	| index
	//	| slice
	//	| typeAssertion
	//	| arguments
	clone := lexer.Clone()

	dot := lexer.LA()
	if dot == nil {
		return nil
	}
	if dot.Type_() == lex.GoLexerDOT {
		lexer.Pop() // dot

		identifier := lexer.LA()
		if identifier.Type_() == lex.GoLexerIDENTIFIER {
			lexer.Pop() // identifier
			return &PrimaryExprSuffix{dot: dot, identifier: identifier}
		} else {
			lexer.Recover(clone)
			dot = nil
		}
	}

	index := VisitIndex(lexer)
	if index != nil {
		return &PrimaryExprSuffix{index: index}
	}

	slice := VisitSlice(lexer)
	if slice != nil {
		return &PrimaryExprSuffix{slice: slice}
	}

	typeAssertion := VisitTypeAssertion(lexer)
	if typeAssertion != nil {
		return &PrimaryExprSuffix{typeAssertion: typeAssertion}
	}

	arguments := VisitArguments(lexer)
	if arguments != nil {
		return &PrimaryExprSuffix{arguments: arguments}
	}

	lexer.Recover(clone)
	return nil
}

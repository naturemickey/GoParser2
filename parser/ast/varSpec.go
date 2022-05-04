package ast

import (
	"GoParser2/lex"
	"fmt"
)

type VarSpec struct {
	// varSpec:
	//	identifierList (
	//		type_ (ASSIGN expressionList)?
	//		| ASSIGN expressionList
	//	);
	identifierList *IdentifierList
	type_          *Type_
	assign         *lex.Token
	expressionList *ExpressionList
}

func VisitVarSpec(lexer *lex.Lexer) *VarSpec {
	clone := lexer.Clone()

	identifierList := VisitIdentifierList(lexer)
	if identifierList == nil {
		return nil
	}

	type_ := VisitType_(lexer)

	if type_ != nil {
		assign := lexer.LA()
		if assign.Type_() == lex.GoLexerASSIGN {
			lexer.Pop()
			expressionList := VisitExpressionList(lexer)
			if expressionList == nil {
				lexer.Recover(clone)
				return nil
			}
			return &VarSpec{identifierList: identifierList, type_: type_, assign: assign, expressionList: expressionList}
		} else {
			return &VarSpec{identifierList: identifierList, type_: type_}
		}
	} else {
		assign := lexer.LA()
		if assign.Type_() != lex.GoLexerASSIGN {
			fmt.Printf("此处应该有一个等号。%s\n", assign.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		lexer.Pop()
		expressionList := VisitExpressionList(lexer)
		if expressionList == nil {
			lexer.Recover(clone)
			return nil
		}
		return &VarSpec{identifierList: identifierList, assign: assign, expressionList: expressionList}
	}
}

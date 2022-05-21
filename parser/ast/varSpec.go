package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
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

func (a *VarSpec) IdentifierList() *IdentifierList {
	return a.identifierList
}

func (a *VarSpec) SetIdentifierList(identifierList *IdentifierList) {
	a.identifierList = identifierList
}

func (a *VarSpec) Type_() *Type_ {
	return a.type_
}

func (a *VarSpec) SetType_(type_ *Type_) {
	a.type_ = type_
}

func (a *VarSpec) Assign() *lex.Token {
	return a.assign
}

func (a *VarSpec) SetAssign(assign *lex.Token) {
	a.assign = assign
}

func (a *VarSpec) ExpressionList() *ExpressionList {
	return a.expressionList
}

func (a *VarSpec) SetExpressionList(expressionList *ExpressionList) {
	a.expressionList = expressionList
}

func (a *VarSpec) CodeBuilder() *CodeBuilder {
	return NewCB().AppendTreeNode(a.identifierList).AppendTreeNode(a.type_).AppendToken(a.assign).AppendTreeNode(a.expressionList)
}

func (a *VarSpec) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*VarSpec)(nil)

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
			lexer.Pop() // assign
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
			fmt.Printf("varSpec,此处应该有一个等号。%s\n", assign.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		lexer.Pop() // assign
		expressionList := VisitExpressionList(lexer)
		if expressionList == nil {
			lexer.Recover(clone)
			return nil
		}
		return &VarSpec{identifierList: identifierList, assign: assign, expressionList: expressionList}
	}
}

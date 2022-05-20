package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type ConstSpec struct {
	// constSpec: identifierList (type_? ASSIGN expressionList)?;

	identifierList *IdentifierList
	type_          *Type_
	assign         *lex.Token
	expressionList *ExpressionList
}

func (a *ConstSpec) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendTreeNode(a.identifierList)
	cb.AppendTreeNode(a.type_)
	cb.AppendToken(a.assign)
	cb.AppendTreeNode(a.expressionList)
	return cb
}

func (a *ConstSpec) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ConstSpec)(nil)

func VisitConstSpec(lexer *lex.Lexer) *ConstSpec {
	// 识别失败不是真失败，需要恢复lexer，因为外面可能会继续使用
	clone := lexer.Clone()

	identifierList := VisitIdentifierList(lexer)
	if identifierList == nil {
		lexer.Recover(clone)
		return nil
	}

	clone2 := lexer.Clone()
	type_ := VisitType_(lexer)

	assign := lexer.LA()
	if assign.Type_() != lex.GoLexerASSIGN {
		// 如果assign不在，那么这一整段就只有identifierList了
		assign = nil
		lexer.Recover(clone2)
		return &ConstSpec{identifierList: identifierList}
	}
	lexer.Pop() // assign

	expressionList := VisitExpressionList(lexer)
	if expressionList == nil {
		fmt.Printf("constSpec,'='后面跟的东西不对。%s\n", assign.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &ConstSpec{identifierList: identifierList, type_: type_, assign: assign, expressionList: expressionList}
}

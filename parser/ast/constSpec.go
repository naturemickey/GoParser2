package ast

import (
	"GoParser2/lex"
	"fmt"
)

type ConstSpec struct {
	// constSpec: identifierList (type_? ASSIGN expressionList)?;

	identifierList *IdentifierList
	type_          *Type_
	assign         *lex.Token
	expressionList *ExpressionList
}

func VisitConstSpec(lexer *lex.Lexer) *ConstSpec {
	// 识别失败不是真失败，需要恢复lexer，因为外面可能会继续使用
	clone := lexer.Clone()

	identifierList := VisitIdentifierList(lexer)
	if identifierList == nil {
		lexer.Recover(clone)
		return nil
	}

	type_ := VisitType_(lexer)

	assign := lexer.LA()
	if assign.Type_() != lex.GoLexerASSIGN {
		// 如果assign不在，那么这一整段就只有identifierList了
		assign = nil
		return &ConstSpec{identifierList: identifierList}
	}

	expressionList := VisitExpressionList(lexer)
	if expressionList == nil {
		fmt.Printf("'='后面跟的东西不对。%s\n", assign.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &ConstSpec{identifierList: identifierList, type_: type_, assign: assign, expressionList: expressionList}
}

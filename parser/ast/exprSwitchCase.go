package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type ExprSwitchCase struct {
	// exprSwitchCase: CASE expressionList | DEFAULT;
	case_          *lex.Token
	expressionList *ExpressionList
	default_       *lex.Token
}

func (a *ExprSwitchCase) Case_() *lex.Token {
	return a.case_
}

func (a *ExprSwitchCase) SetCase_(case_ *lex.Token) {
	a.case_ = case_
}

func (a *ExprSwitchCase) ExpressionList() *ExpressionList {
	return a.expressionList
}

func (a *ExprSwitchCase) SetExpressionList(expressionList *ExpressionList) {
	a.expressionList = expressionList
}

func (a *ExprSwitchCase) Default_() *lex.Token {
	return a.default_
}

func (a *ExprSwitchCase) SetDefault_(default_ *lex.Token) {
	a.default_ = default_
}

func (a *ExprSwitchCase) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.case_)
	cb.AppendTreeNode(a.expressionList)
	cb.AppendToken(a.default_)
	return cb
}

func (a *ExprSwitchCase) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ExprSwitchCase)(nil)

func VisitExprSwitchCase(lexer *lex.Lexer) *ExprSwitchCase {
	clone := lexer.Clone()

	case_ := lexer.LA()

	if case_.Type_() == lex.GoLexerCASE {
		lexer.Pop() // case_
		expressionList := VisitExpressionList(lexer)
		if expressionList == nil {
			fmt.Printf("exprSwitchCase,case后面要有表达式。%s\n", case_.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return &ExprSwitchCase{case_: case_, expressionList: expressionList}
	} else if case_.Type_() == lex.GoLexerDEFAULT {
		lexer.Pop() // default
		return &ExprSwitchCase{default_: case_}
	} else {
		return nil
	}
}

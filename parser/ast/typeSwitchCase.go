package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type TypeSwitchCase struct {
	// typeSwitchCase: CASE typeList | DEFAULT;
	case_    *lex.Token
	typeList *TypeList
	default_ *lex.Token
}

func (a *TypeSwitchCase) Case_() *lex.Token {
	return a.case_
}

func (a *TypeSwitchCase) SetCase_(case_ *lex.Token) {
	a.case_ = case_
}

func (a *TypeSwitchCase) TypeList() *TypeList {
	return a.typeList
}

func (a *TypeSwitchCase) SetTypeList(typeList *TypeList) {
	a.typeList = typeList
}

func (a *TypeSwitchCase) Default_() *lex.Token {
	return a.default_
}

func (a *TypeSwitchCase) SetDefault_(default_ *lex.Token) {
	a.default_ = default_
}

func (a *TypeSwitchCase) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.case_).AppendTreeNode(a.typeList).AppendToken(a.default_)
}

func (a *TypeSwitchCase) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*TypeSwitchCase)(nil)

func VisitTypeSwitchCase(lexer *lex.Lexer) *TypeSwitchCase {
	clone := lexer.Clone()

	la := lexer.LA()

	if la.Type_() == lex.GoLexerCASE {
		lexer.Pop() // la
		typeList := VisitTypeList(lexer)
		if typeList == nil {
			fmt.Printf("typeSwitchCase,case后面要有表达式。%s\n", la.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return &TypeSwitchCase{case_: la, typeList: typeList}
	} else if la.Type_() == lex.GoLexerDEFAULT {
		lexer.Pop() // la
		return &TypeSwitchCase{default_: la}
	} else {
		return nil
	}
}

package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
	"fmt"
)

type TypeSwitchCase struct {
	// typeSwitchCase: CASE typeList | DEFAULT;
	case_    *lex.Token
	typeList *TypeList
	default_ *lex.Token
}

func (a *TypeSwitchCase) CodeBuilder() *util.CodeBuilder {
	return util.NewCB().AppendToken(a.case_).AppendTreeNode(a.typeList).AppendToken(a.default_)
}

func (a *TypeSwitchCase) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*TypeSwitchCase)(nil)

func VisitTypeSwitchCase(lexer *lex.Lexer) *TypeSwitchCase {
	clone := lexer.Clone()

	la := lexer.LA()

	if la.Type_() == lex.GoLexerCASE {
		lexer.Pop() // la
		typeList := VisitTypeList(lexer)
		if typeList == nil {
			fmt.Printf("case后面要有表达式。%s\n", la.ErrorMsg())
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

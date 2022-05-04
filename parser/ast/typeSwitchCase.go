package ast

import (
	"GoParser2/lex"
	"fmt"
)

type TypeSwitchCase struct {
	// typeSwitchCase: CASE typeList | DEFAULT;
	case_    *lex.Token
	typeList *TypeList
	default_ *lex.Token
}

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

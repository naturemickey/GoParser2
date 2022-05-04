package ast

import (
	"GoParser2/lex"
	"fmt"
)

type ExprSwitchCase struct {
	// exprSwitchCase: CASE expressionList | DEFAULT;
	case_          *lex.Token
	expressionList *ExpressionList
	default_       *lex.Token
}

func VisitExprSwitchCase(lexer *lex.Lexer) *ExprSwitchCase {
	clone := lexer.Clone()

	la := lexer.LA()

	if la.Type_() == lex.GoLexerCASE {
		lexer.Pop() // la
		expressionList := VisitExpressionList(lexer)
		if expressionList == nil {
			fmt.Printf("case后面要有表达式。%s\n", la.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return &ExprSwitchCase{case_: la, expressionList: expressionList}
	} else if la.Type_() == lex.GoLexerDEFAULT {
		lexer.Pop() // la
		return &ExprSwitchCase{default_: la}
	} else {
		return nil
	}
}

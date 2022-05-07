package ast

import (
	"GoParser2/lex"
	"fmt"
)

type IdentifierList struct {
	// identifierList: IDENTIFIER (COMMA IDENTIFIER)*;
	identifiers []*lex.Token
}

func VisitIdentifierList(lexer *lex.Lexer) *IdentifierList {
	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		return nil
	}
	lexer.Pop() // identifier

	var identifiers []*lex.Token
	identifiers = append(identifiers, identifier)

	for true {
		comma := lexer.LA()
		if comma.Type_() == lex.GoLexerCOMMA {
			lexer.Pop()

			identifier := lexer.LA()
			if identifier.Type_() != lex.GoLexerIDENTIFIER {
				fmt.Printf("逗号后面要跟着一个标识符才对。%s\n", comma.ErrorMsg())
				return nil
			}
			lexer.Pop()
			identifiers = append(identifiers, identifier)
		} else {
			break
		}
	}
	return &IdentifierList{identifiers: identifiers}
}

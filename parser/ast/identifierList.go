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
	var identifiers []*lex.Token
	identifiers = append(identifiers, identifier)

	for true {
		la := lexer.LA()
		if la.Type_() == lex.GoLexerCOMMA {
			identifier := lexer.LA()
			if identifier.Type_() != lex.GoLexerIDENTIFIER {
				fmt.Printf("逗号后面要跟着一个标识符才对。%s\n", la.ErrorMsg())
				return nil
			}
			identifiers = append(identifiers, identifier)
		} else {
			break
		}
	}
	return &IdentifierList{identifiers: identifiers}
}

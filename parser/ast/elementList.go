package ast

import "GoParser2/lex"

type ElementList struct {
	// elementList: keyedElement (COMMA keyedElement)*;
	keyedElements []*KeyedElement
}

func VisitElementList(lexer *lex.Lexer) *ElementList {
	clone := lexer.Clone()

	var keyedElements []*KeyedElement

	keyedElement := VisitKeyedElement(lexer)
	if keyedElement == nil {
		lexer.Recover(clone)
		return nil
	}
	keyedElements = append(keyedElements, keyedElement)

	for {
		comma := lexer.LA()
		if comma.Type_() != lex.GoLexerCOMMA {
			break
		}
		lexer.Pop() // comma

		keyedElement := VisitKeyedElement(lexer)
		if keyedElement == nil {
			lexer.Recover(clone)
			return nil
		}
		keyedElements = append(keyedElements, keyedElement)
	}

	return &ElementList{keyedElements: keyedElements}
}

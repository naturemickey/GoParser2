package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type ElementList struct {
	// elementList: keyedElement (COMMA keyedElement)*;
	keyedElements []*KeyedElement
}

func (a *ElementList) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*ElementList)(nil)

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

package ast

import (
	"GoParser2/lex"
)

type ElementList struct {
	// elementList: keyedElement (COMMA keyedElement)*;
	keyedElements []*KeyedElement
}

func (a *ElementList) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	for i, element := range a.keyedElements {
		if i == 0 {
			cb.AppendTreeNode(element)
		} else {
			cb.AppendString(",").AppendTreeNode(element)
		}
	}
	return cb
}

func (a *ElementList) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ElementList)(nil)

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

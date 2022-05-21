package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type KeyedElement struct {
	// keyedElement: (key COLON)? element;
	key     Key
	colon   *lex.Token
	element Element
}

func (a *KeyedElement) Key() Key {
	return a.key
}

func (a *KeyedElement) SetKey(key Key) {
	a.key = key
}

func (a *KeyedElement) Colon() *lex.Token {
	return a.colon
}

func (a *KeyedElement) SetColon(colon *lex.Token) {
	a.colon = colon
}

func (a *KeyedElement) Element() Element {
	return a.element
}

func (a *KeyedElement) SetElement(element Element) {
	a.element = element
}

func (a *KeyedElement) CodeBuilder() *CodeBuilder {
	return NewCB().AppendTreeNode(a.key).AppendToken(a.colon).AppendTreeNode(a.element)
}

func (a *KeyedElement) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*KeyedElement)(nil)

func VisitKeyedElement(lexer *lex.Lexer) *KeyedElement {
	clone := lexer.Clone()

	key := VisitKey(lexer)
	var colon *lex.Token
	if key != nil {
		colon = lexer.LA()
		if colon.Type_() != lex.GoLexerCOLON {
			key = nil
			colon = nil
			lexer.Recover(clone)
		} else {
			lexer.Pop() // key
		}
	}

	element := VisitElement(lexer)
	if element == nil {
		lexer.Recover(clone)
		return nil
	}
	return &KeyedElement{key: key, colon: colon, element: element}
}

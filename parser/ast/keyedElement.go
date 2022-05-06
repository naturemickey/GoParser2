package ast

import "GoParser2/lex"

type KeyedElement struct {
	// keyedElement: (key COLON)? element;
	key     Key
	colon   *lex.Token
	element Element
}

func VisitKeyedElement(lexer *lex.Lexer) *KeyedElement {
	clone := lexer.Clone()

	key := VisitKey(lexer)
	var colon *lex.Token
	if key != nil {
		colon = lexer.LA()
		if colon.Type_() != lex.GoLexerCOLON {
			lexer.Recover(clone)
			return nil
		}
		lexer.Pop() // key
	}

	element := VisitElement(lexer)
	if element == nil {
		lexer.Recover(clone)
		return nil
	}
	return &KeyedElement{key: key, colon: colon, element: element}
}

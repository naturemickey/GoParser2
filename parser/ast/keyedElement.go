package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
)

type KeyedElement struct {
	// keyedElement: (key COLON)? element;
	key     Key
	colon   *lex.Token
	element Element
}

func (a *KeyedElement) CodeBuilder() *util.CodeBuilder {
	return util.NewCB().AppendTreeNode(a.key).AppendToken(a.colon).AppendTreeNode(a.element)
}

func (a *KeyedElement) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*KeyedElement)(nil)

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

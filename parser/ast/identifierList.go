package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type IdentifierList struct {
	// identifierList: IDENTIFIER (COMMA IDENTIFIER)*;
	identifiers []*lex.Token
}

func (a *IdentifierList) Identifiers() []*lex.Token {
	return a.identifiers
}

func (a *IdentifierList) SetIdentifiers(identifiers []*lex.Token) {
	a.identifiers = identifiers
}

func (a *IdentifierList) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	for i, identifier := range a.identifiers {
		if i == 0 {
			cb.AppendToken(identifier)
		} else {
			cb.AppendString(",").AppendToken(identifier)
		}
	}
	return cb
}

func (a *IdentifierList) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*IdentifierList)(nil)

func VisitIdentifierList(lexer *lex.Lexer) *IdentifierList {
	if lexer.LA() == nil { // 文件结束
		return nil
	}
	clone := lexer.Clone()

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
				//fmt.Printf("identifierList,逗号后面要跟着一个标识符才对。%s\n", comma.ErrorMsg())
				lexer.Recover(clone)
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

package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
	"fmt"
)

type IdentifierList struct {
	// identifierList: IDENTIFIER (COMMA IDENTIFIER)*;
	identifiers []*lex.Token
}

func (a *IdentifierList) CodeBuilder() *util.CodeBuilder {
	cb := util.NewCB()
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

var _ parser.ITreeNode = (*IdentifierList)(nil)

func VisitIdentifierList(lexer *lex.Lexer) *IdentifierList {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

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

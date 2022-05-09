package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
	"fmt"
)

type SliceType struct {
	// sliceType  : L_BRACKET R_BRACKET elementType;
	// elementType: type_;
	lBracket    *lex.Token
	rBracket    *lex.Token
	elementType *Type_
}

func (a *SliceType) CodeBuilder() *util.CodeBuilder {
	return util.NewCB().AppendToken(a.lBracket).AppendToken(a.rBracket).AppendTreeNode(a.elementType)
}

func (a *SliceType) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*SliceType)(nil)

func (s SliceType) __TypeLit__() {
	panic("imposible")
}

var _ TypeLit = (*SliceType)(nil)

func VisitSliceType(lexer *lex.Lexer) *SliceType {
	clone := lexer.Clone()

	lBracket := lexer.LA()
	if lBracket.Type_() != lex.GoLexerL_BRACKET {
		return nil
	}
	lexer.Pop() // lBracket

	rBracket := lexer.LA()
	if rBracket.Type_() != lex.GoLexerR_BRACKET {
		fmt.Printf("此处应该是一个']'。%s\n", rBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rBracket

	elementType := VisitType_(lexer)
	if elementType == nil {
		fmt.Printf("']'后面应该是类型描述。%s\n", rBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &SliceType{lBracket: lBracket, rBracket: rBracket, elementType: elementType}
}

package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"fmt"
)

type MapType struct {
	// mapType    : MAP L_BRACKET type_ R_BRACKET elementType;
	// elementType: type_;
	map_        *lex.Token
	lBracket    *lex.Token
	type_       *Type_
	rBracket    *lex.Token
	elementType *Type_
}

func (a *MapType) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*MapType)(nil)

func (m MapType) __TypeLit__() {
	panic("imposible")
}

var _ TypeLit = (*MapType)(nil)

func VisitMapType(lexer *lex.Lexer) *MapType {
	clone := lexer.Clone()

	map_ := lexer.LA()
	if map_.Type_() != lex.GoLexerMAP {
		return nil
	}
	lexer.Pop() // map_

	lBracket := lexer.LA()
	if lBracket.Type_() != lex.GoLexerL_BRACKET {
		fmt.Printf("此处需要一个'['。%s\n", lBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lBracket

	type_ := VisitType_(lexer)
	if type_ == nil {
		fmt.Printf("'['后面需要跟着一个类型定义。%s\n", lBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	rBracket := lexer.LA()
	if rBracket.Type_() != lex.GoLexerR_BRACKET {
		fmt.Printf("此处应该是一个']'。%s\n", rBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rBracket

	elementType := VisitType_(lexer)
	if elementType == nil {
		fmt.Printf("']'右边应该是一个类型描述。%s\n", rBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &MapType{map_: map_, lBracket: lBracket, type_: type_, rBracket: rBracket, elementType: elementType}
}

package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
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

func (a *MapType) Map_() *lex.Token {
	return a.map_
}

func (a *MapType) SetMap_(map_ *lex.Token) {
	a.map_ = map_
}

func (a *MapType) LBracket() *lex.Token {
	return a.lBracket
}

func (a *MapType) SetLBracket(lBracket *lex.Token) {
	a.lBracket = lBracket
}

func (a *MapType) Type_() *Type_ {
	return a.type_
}

func (a *MapType) SetType_(type_ *Type_) {
	a.type_ = type_
}

func (a *MapType) RBracket() *lex.Token {
	return a.rBracket
}

func (a *MapType) SetRBracket(rBracket *lex.Token) {
	a.rBracket = rBracket
}

func (a *MapType) ElementType() *Type_ {
	return a.elementType
}

func (a *MapType) SetElementType(elementType *Type_) {
	a.elementType = elementType
}

func (a *MapType) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.map_).AppendToken(a.lBracket).AppendTreeNode(a.type_).AppendToken(a.rBracket).AppendTreeNode(a.elementType)
}

func (a *MapType) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*MapType)(nil)

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
		fmt.Printf("mapType,此处需要一个'['。%s\n", lBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lBracket

	type_ := VisitType_(lexer)
	if type_ == nil {
		fmt.Printf("mapType,'['后面需要跟着一个类型定义。%s\n", lBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	rBracket := lexer.LA()
	if rBracket.Type_() != lex.GoLexerR_BRACKET {
		fmt.Printf("mapType,此处应该是一个']'。%s\n", rBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rBracket

	elementType := VisitType_(lexer)
	if elementType == nil {
		fmt.Printf("mapType,']'右边应该是一个类型描述。%s\n", rBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &MapType{map_: map_, lBracket: lBracket, type_: type_, rBracket: rBracket, elementType: elementType}
}

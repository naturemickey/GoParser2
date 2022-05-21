package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type PointerType struct {
	// pointerType: STAR type_;
	star  *lex.Token
	type_ *Type_
}

func (a *PointerType) Star() *lex.Token {
	return a.star
}

func (a *PointerType) SetStar(star *lex.Token) {
	a.star = star
}

func (a *PointerType) Type_() *Type_ {
	return a.type_
}

func (a *PointerType) SetType_(type_ *Type_) {
	a.type_ = type_
}

func (a *PointerType) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.star).AppendTreeNode(a.type_)
}

func (a *PointerType) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*PointerType)(nil)

func (p PointerType) __TypeLit__() {
	panic("imposible")
}

var _ TypeLit = (*PointerType)(nil)

func VisitPointerType(lexer *lex.Lexer) *PointerType {
	clone := lexer.Clone()

	star := lexer.LA()
	if star.Type_() != lex.GoLexerSTAR {
		return nil
	}
	lexer.Pop() // star

	type_ := VisitType_(lexer)
	if type_ == nil {
		fmt.Printf("pointerType,'*'后面需要一个类型描述。%s\n", star.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &PointerType{star: star, type_: type_}
}

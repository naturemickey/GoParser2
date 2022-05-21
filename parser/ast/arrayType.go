package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type ArrayType struct {
	// arrayType  : L_BRACKET arrayLength R_BRACKET elementType;
	// arrayLength: expression;
	// elementType: type_;

	lBracket    *lex.Token
	arrayLength *Expression
	rBracket    *lex.Token
	elementType *Type_
}

func (a *ArrayType) LBracket() *lex.Token {
	return a.lBracket
}

func (a *ArrayType) SetLBracket(lBracket *lex.Token) {
	a.lBracket = lBracket
}

func (a *ArrayType) ArrayLength() *Expression {
	return a.arrayLength
}

func (a *ArrayType) SetArrayLength(arrayLength *Expression) {
	a.arrayLength = arrayLength
}

func (a *ArrayType) RBracket() *lex.Token {
	return a.rBracket
}

func (a *ArrayType) SetRBracket(rBracket *lex.Token) {
	a.rBracket = rBracket
}

func (a *ArrayType) ElementType() *Type_ {
	return a.elementType
}

func (a *ArrayType) SetElementType(elementType *Type_) {
	a.elementType = elementType
}

func (a *ArrayType) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.lBracket)
	cb.AppendTreeNode(a.arrayLength)
	cb.AppendToken(a.rBracket)
	cb.AppendTreeNode(a.elementType)
	return cb
}

func (a *ArrayType) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ArrayType)(nil)

func (a ArrayType) __TypeLit__() {
	panic("imposible")
}

var _ TypeLit = (*ArrayType)(nil)

func VisitArrayType(lexer *lex.Lexer) *ArrayType {
	clone := lexer.Clone()

	lBracket := lexer.LA()
	if lBracket.Type_() != lex.GoLexerL_BRACKET {
		return nil
	}
	lexer.Pop() // lBracket

	arrayLength := VisitExpression(lexer)
	if arrayLength == nil {
		fmt.Printf("arrayType,方括号中间应该是一个值为数字的表达式。%s\n", lBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	rBracket := lexer.LA()
	if rBracket.Type_() != lex.GoLexerR_BRACKET {
		fmt.Printf("arrayType,此处应该有一个右方括号。%s\n", rBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rBracket

	elementType := VisitType_(lexer)
	if elementType == nil {
		fmt.Printf("arrayType,这里（']'右边）需要一个类型描述。%s\n", rBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &ArrayType{lBracket: lBracket, arrayLength: arrayLength, rBracket: rBracket, elementType: elementType}
}

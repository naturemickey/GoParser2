package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"fmt"
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

func (a *ArrayType) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*ArrayType)(nil)

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
		fmt.Printf("方括号中间应该是一个值为数字的表达式。%s\n", lBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	rBracket := lexer.LA()
	if rBracket.Type_() != lex.GoLexerR_BRACKET {
		fmt.Printf("此处应该有一个右方括号。%s\n", rBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rBracket

	elementType := VisitType_(lexer)
	if elementType == nil {
		fmt.Printf("这里（']'右边）需要一个类型描述。%s\n", rBracket.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &ArrayType{lBracket: lBracket, arrayLength: arrayLength, rBracket: rBracket, elementType: elementType}
}

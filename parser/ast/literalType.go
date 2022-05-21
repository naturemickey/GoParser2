package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type LiteralType struct {
	// literalType:
	//	structType
	//	| arrayType
	//	| L_BRACKET ELLIPSIS R_BRACKET elementType
	//	| sliceType
	//	| mapType
	//	| typeName;

	// elementType: type_;

	structType *StructType
	arrayType  *ArrayType

	lBracket    *lex.Token
	ellipsis    *lex.Token
	rBracket    *lex.Token
	elementType *Type_

	sliceType *SliceType
	mapType   *MapType
	typeName  *TypeName
}

func (a *LiteralType) StructType() *StructType {
	return a.structType
}

func (a *LiteralType) SetStructType(structType *StructType) {
	a.structType = structType
}

func (a *LiteralType) ArrayType() *ArrayType {
	return a.arrayType
}

func (a *LiteralType) SetArrayType(arrayType *ArrayType) {
	a.arrayType = arrayType
}

func (a *LiteralType) LBracket() *lex.Token {
	return a.lBracket
}

func (a *LiteralType) SetLBracket(lBracket *lex.Token) {
	a.lBracket = lBracket
}

func (a *LiteralType) Ellipsis() *lex.Token {
	return a.ellipsis
}

func (a *LiteralType) SetEllipsis(ellipsis *lex.Token) {
	a.ellipsis = ellipsis
}

func (a *LiteralType) RBracket() *lex.Token {
	return a.rBracket
}

func (a *LiteralType) SetRBracket(rBracket *lex.Token) {
	a.rBracket = rBracket
}

func (a *LiteralType) ElementType() *Type_ {
	return a.elementType
}

func (a *LiteralType) SetElementType(elementType *Type_) {
	a.elementType = elementType
}

func (a *LiteralType) SliceType() *SliceType {
	return a.sliceType
}

func (a *LiteralType) SetSliceType(sliceType *SliceType) {
	a.sliceType = sliceType
}

func (a *LiteralType) MapType() *MapType {
	return a.mapType
}

func (a *LiteralType) SetMapType(mapType *MapType) {
	a.mapType = mapType
}

func (a *LiteralType) TypeName() *TypeName {
	return a.typeName
}

func (a *LiteralType) SetTypeName(typeName *TypeName) {
	a.typeName = typeName
}

func (a *LiteralType) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendTreeNode(a.structType)
	cb.AppendTreeNode(a.arrayType)
	cb.AppendToken(a.lBracket).AppendToken(a.ellipsis).AppendToken(a.rBracket).AppendTreeNode(a.elementType)
	cb.AppendTreeNode(a.sliceType)
	cb.AppendTreeNode(a.mapType)
	cb.AppendTreeNode(a.typeName)
	return cb
}

func (a *LiteralType) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*LiteralType)(nil)

func VisitLiteralType(lexer *lex.Lexer) *LiteralType {
	clone := lexer.Clone()

	la := lexer.LA()
	if la.Type_() == lex.GoLexerSTRUCT {
		//	structType
		structType := VisitStructType(lexer)
		if structType == nil {
			return nil
		}
		return &LiteralType{structType: structType}
	} else if la.Type_() == lex.GoLexerL_BRACKET {
		//	| arrayType
		//	| L_BRACKET ELLIPSIS R_BRACKET elementType
		//	| sliceType
		la1 := lexer.LA1()
		if la1.Type_() == lex.GoLexerR_BRACKET {
			//	| sliceType
			sliceType := VisitSliceType(lexer)
			if sliceType == nil {
				lexer.Recover(clone)
				return nil
			}
			return &LiteralType{sliceType: sliceType}
		} else if la1.Type_() == lex.GoLexerELLIPSIS {
			//	| L_BRACKET ELLIPSIS R_BRACKET elementType
			lBracket := la
			ellipsis := la1
			lexer.Pop() // lBracket
			lexer.Pop() // ellipsis
			rBracket := lexer.LA()
			if rBracket.Type_() != lex.GoLexerR_BRACKET {
				lexer.Recover(clone)
				return nil
			}
			lexer.Pop() // rBracket
			elementType := VisitType_(lexer)
			if elementType == nil {
				lexer.Recover(clone)
				return nil
			}
			return &LiteralType{lBracket: lBracket, ellipsis: ellipsis, rBracket: rBracket, elementType: elementType}
		} else {
			//	| arrayType
			arrayType := VisitArrayType(lexer)
			if arrayType == nil {
				lexer.Recover(clone)
				return nil
			}
			return &LiteralType{arrayType: arrayType}
		}
	} else if la.Type_() == lex.GoLexerMAP {
		//	| mapType
		mapType := VisitMapType(lexer)
		if mapType == nil {
			lexer.Recover(clone)
			return nil
		}
		return &LiteralType{mapType: mapType}
	} else {
		//	| typeName;
		typeName := VisitTypeName(lexer)
		if typeName == nil {
			lexer.Recover(clone)
			return nil
		}
		return &LiteralType{typeName: typeName}
	}
}

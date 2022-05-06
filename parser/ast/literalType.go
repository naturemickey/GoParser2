package ast

import "GoParser2/lex"

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

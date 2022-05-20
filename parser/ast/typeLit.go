package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type TypeLit interface {
	ITreeNode
	// typeLit:
	//	arrayType
	//	| structType
	//	| pointerType
	//	| functionType
	//	| interfaceType
	//	| sliceType
	//	| mapType
	//	| channelType;

	__TypeLit__()
}

func VisitTypeLit(lexer *lex.Lexer) TypeLit {
	clone := lexer.Clone()

	la := lexer.LA()

	switch la.Type_() {
	case lex.GoLexerL_BRACKET: // arrayType | sliceType
		la1 := lexer.LA1()
		if la1.Type_() == lex.GoLexerR_BRACKET { // sliceType
			sliceType := VisitSliceType(lexer)
			if sliceType == nil {
				fmt.Printf("typeLit,slice类型的描述不完整。%s\n", la.ErrorMsg())
				lexer.Recover(clone)
				return nil
			}
			return sliceType
		} else {
			arrayType := VisitArrayType(lexer)
			if arrayType == nil {
				fmt.Printf("typeLit,array类型的描述不完整。%s\n", la.ErrorMsg())
				lexer.Recover(clone)
				return nil
			}
			return arrayType
		}
	case lex.GoLexerSTRUCT: // structType
		structType := VisitStructType(lexer)
		if structType == nil {
			fmt.Printf("typeLit,struct类型的描述不完整。%s\n", la.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return structType
	case lex.GoLexerSTAR: // pointerType
		pointerType := VisitPointerType(lexer)
		if pointerType == nil {
			fmt.Printf("typeLit,指针类型的描述不完整。%s\n", la.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return pointerType
	case lex.GoLexerFUNC: // functionType
		functionType := VisitFunctionType(lexer)
		if functionType == nil {
			fmt.Printf("typeLit,func类型的描述不完整。%s\n", la.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return functionType
	case lex.GoLexerINTERFACE: // interfaceType
		interfaceType := VisitInterfaceType(lexer)
		if interfaceType == nil {
			fmt.Printf("typeLit,interface类型的描述不完整。%s\n", la.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return interfaceType
	case lex.GoLexerMAP: // mapType
		mapType := VisitMapType(lexer)
		if mapType == nil {
			fmt.Printf("typeLit,map类型的描述不完整。%s\n", la.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		return mapType
	default: // channelType
		channelType := VisitChannelType(lexer)
		if channelType == nil {
			lexer.Recover(clone)
			return nil
		}
		return channelType
	}
}

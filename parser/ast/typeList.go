package ast

import (
	"GoParser2/lex"
	"fmt"
)

type TypeList struct {
	// typeList: (type_ | NIL_LIT) (COMMA (type_ | NIL_LIT))*;
	type_s []*Type_ // 为nil就表示NIL_LIT
}

func VisitTypeList(lexer *lex.Lexer) *TypeList {
	clone := lexer.Clone()

	var type_s []*Type_

	type_, ok := _visitTypeOrNil(lexer)
	if !ok {
		lexer.Recover(clone)
		return nil
	}
	type_s = append(type_s, type_)

	for true {
		comma := lexer.LA()
		if comma.Type_() != lex.GoLexerCOMMA {
			break
		}
		type_, ok := _visitTypeOrNil(lexer)
		if !ok {
			fmt.Printf("逗号后面需要跟着另外一个类型（或nil）。%s\n", comma.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		type_s = append(type_s, type_)
	}
	return &TypeList{type_s: type_s}
}

func _visitTypeOrNil(lexer *lex.Lexer) (*Type_, bool) {
	la := lexer.LA()
	if la.Type_() == lex.GoLexerNIL_LIT {
		lexer.Pop()
		return nil, true
	}
	type_ := VisitType_(lexer)
	if type_ != nil {
		return type_, true
	}
	return nil, false
}

package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type TypeList struct {
	// typeList: (type_ | NIL_LIT) (COMMA (type_ | NIL_LIT))*;
	type_s []*Type_ // 为nil就表示NIL_LIT
}

func (a *TypeList) Type_s() []*Type_ {
	return a.type_s
}

func (a *TypeList) SetType_s(type_s []*Type_) {
	a.type_s = type_s
}

func (a *TypeList) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	for i, type_ := range a.type_s {
		if i == 0 {
			if type_ == nil {
				cb.AppendString("nil")
			} else {
				cb.AppendTreeNode(type_)
			}
		} else {
			if type_ == nil {
				cb.AppendString(",").AppendString("nil")
			} else {
				cb.AppendString(",").AppendTreeNode(type_)
			}
		}
	}
	return cb
}

func (a *TypeList) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*TypeList)(nil)

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
		lexer.Pop() // comma

		type_, ok := _visitTypeOrNil(lexer)
		if !ok {
			fmt.Printf("typeLit,逗号后面需要跟着另外一个类型（或nil）。%s\n", comma.ErrorMsg())
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

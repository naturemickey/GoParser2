package ast

import (
	"GoParser2/lex"
	"fmt"
)

type Type_ struct {
	// type_: typeName | typeLit | L_PAREN type_ R_PAREN;
	typeName *TypeName
	typeLit  TypeLit

	lParen *lex.Token
	type_  *Type_
	rParen *lex.Token
}

func (a *Type_) CodeBuilder() *CodeBuilder {
	return NewCB().AppendTreeNode(a.typeName).AppendTreeNode(a.typeLit).AppendToken(a.lParen).AppendTreeNode(a.type_).AppendToken(a.rParen)
}

func (a *Type_) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*Type_)(nil)

func (t Type_) __Result__() {
	panic("imposible")
}

var _ Result = (*Type_)(nil)

func VisitType_(lexer *lex.Lexer) *Type_ {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	clone := lexer.Clone()

	lParen := lexer.LA()
	if lParen.Type_() == lex.GoLexerL_PAREN {
		lexer.Pop() // lParen

		type_ := VisitType_(lexer)
		if type_ == nil {
			fmt.Printf("此处括号里面需要是一个类型。%s\n", lParen.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}

		rParen := lexer.LA()
		if rParen.Type_() != lex.GoLexerR_PAREN {
			fmt.Printf("此处应该有一个')'。%s\n", rParen.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		lexer.Pop() // rParen

		return &Type_{lParen: lParen, type_: type_, rParen: rParen}
	}

	typeLit := VisitTypeLit(lexer)
	if typeLit != nil {
		return &Type_{typeLit: typeLit}
	}

	typeName := VisitTypeName(lexer)
	if typeName != nil {
		return &Type_{typeName: typeName}
	}

	lexer.Recover(clone)
	return nil
}

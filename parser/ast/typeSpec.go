package ast

import (
	"GoParser2/lex"
	"fmt"
)

type TypeSpec struct {
	// typeSpec: IDENTIFIER ASSIGN? type_;
	identifier *lex.Token
	assign     *lex.Token
	type_      *Type_
}

func (a *TypeSpec) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.identifier).AppendToken(a.assign).AppendTreeNode(a.type_)
}

func (a *TypeSpec) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*TypeSpec)(nil)

func VisitTypeSpec(lexer *lex.Lexer) *TypeSpec {
	clone := lexer.Clone()

	identifier := lexer.LA()
	if identifier == nil {
		return nil
	}
	lexer.Pop()

	assign := lexer.LA()
	if assign.Type_() != lex.GoLexerASSIGN {
		assign = nil
	} else {
		lexer.Pop() // assign
	}

	type_ := VisitType_(lexer)

	if type_ == nil {
		fmt.Println("后面没看到类型的描述。%s\n", identifier.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	return &TypeSpec{identifier: identifier, assign: assign, type_: type_}
}

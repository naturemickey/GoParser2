package ast

import (
	"GoParser2/lex"
	"fmt"
)

type Block struct {
	// block: L_CURLY statementList? R_CURLY;
	lCurly        *lex.Token
	statementList *StatementList
	rCurly        *lex.Token
}

func (b Block) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*Block)(nil)

func VisitBlock(lexer *lex.Lexer) *Block {
	clone := lexer.Clone()
	lCurly := lexer.LA()
	if lCurly.Type_() != lex.GoLexerL_CURLY {
		return nil
	}
	lexer.Pop() // lCurly

	statementList := VisitStatementList(lexer)
	// todo 判断statementList的语句个数大于等于1

	rCurly := lexer.LA()
	if rCurly.Type_() != lex.GoLexerL_CURLY {
		fmt.Printf("此处没有看到右花括号。%s\n", rCurly.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rCurly

	return &Block{lCurly: lCurly, statementList: statementList, rCurly: lCurly}
}

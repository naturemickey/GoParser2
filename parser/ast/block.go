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

func (this *Block) AddStatement(statement Statement) {
	if this.statementList == nil {
		this.statementList = NewStatementList()
	}
	this.statementList.statements = append(this.statementList.statements, statement)
}

func (a *Block) StatementList() *StatementList {
	return a.statementList
}

func (a *Block) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.lCurly)
	if a.statementList != nil && len(a.statementList.statements) > 0 {
		cb.Newline()
		cb.AppendTreeNode(a.statementList)
	}
	cb.AppendToken(a.rCurly)
	return cb
}

func (a *Block) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*Block)(nil)

func (b Block) __Statement__() {
	panic("imposible")
}

var _ Statement = (*Block)(nil)

func VisitBlock(lexer *lex.Lexer) *Block {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	clone := lexer.Clone()
	lCurly := lexer.LA()
	if lCurly.Type_() != lex.GoLexerL_CURLY {
		return nil
	}
	lexer.Pop() // lCurly

	statementList := VisitStatementList(lexer)
	// todo 判断statementList的语句个数大于等于1

	rCurly := lexer.LA()
	if rCurly.Type_() != lex.GoLexerR_CURLY {
		fmt.Printf("block,此处没有看到右花括号。%s\n", rCurly.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rCurly

	return &Block{lCurly: lCurly, statementList: statementList, rCurly: rCurly}
}

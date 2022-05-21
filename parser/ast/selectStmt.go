package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type SelectStmt struct {
	// selectStmt: SELECT L_CURLY commClause* R_CURLY;
	select_     *lex.Token
	lCurly      *lex.Token
	commClauses []*CommClause
	rCurly      *lex.Token
}

func (a *SelectStmt) Select_() *lex.Token {
	return a.select_
}

func (a *SelectStmt) SetSelect_(select_ *lex.Token) {
	a.select_ = select_
}

func (a *SelectStmt) LCurly() *lex.Token {
	return a.lCurly
}

func (a *SelectStmt) SetLCurly(lCurly *lex.Token) {
	a.lCurly = lCurly
}

func (a *SelectStmt) CommClauses() []*CommClause {
	return a.commClauses
}

func (a *SelectStmt) SetCommClauses(commClauses []*CommClause) {
	a.commClauses = commClauses
}

func (a *SelectStmt) RCurly() *lex.Token {
	return a.rCurly
}

func (a *SelectStmt) SetRCurly(rCurly *lex.Token) {
	a.rCurly = rCurly
}

func (a *SelectStmt) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.select_).AppendToken(a.lCurly)

	if len(a.commClauses) > 0 {
		cb.Newline()
		for _, clause := range a.commClauses {
			cb.AppendTreeNode(clause).Newline()
		}
	}
	cb.AppendToken(a.rCurly)
	return cb
}

func (a *SelectStmt) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*SelectStmt)(nil)

func (s SelectStmt) __Statement__() {
	panic("imposible")
}

var _ Statement = (*SelectStmt)(nil)

func VisitSelectStmt(lexer *lex.Lexer) *SelectStmt {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	clone := lexer.Clone()

	select_ := lexer.LA()
	if select_.Type_() != lex.GoLexerSELECT {
		return nil
	}
	lexer.Pop() // select_

	lCurly := lexer.LA()
	if lCurly.Type_() != lex.GoLexerL_CURLY {
		fmt.Printf("selectStmt,select后面必须是一个{。%s\n", lCurly.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lCurly

	var commClauses []*CommClause
	for true {
		commClause := VisitCommClause(lexer)
		if commClause != nil {
			commClauses = append(commClauses, commClause)
		} else {
			break
		}
	}
	rCurly := lexer.LA()
	if rCurly.Type_() != lex.GoLexerR_CURLY {
		fmt.Printf("selectStmt,select需要一个'}'，在这里。%s\n", rCurly.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rCurly

	return &SelectStmt{select_: select_, lCurly: lCurly, commClauses: commClauses, rCurly: rCurly}
}

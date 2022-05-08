package ast

import (
	"GoParser2/lex"
	"fmt"
)

type SelectStmt struct {
	// selectStmt: SELECT L_CURLY commClause* R_CURLY;
	select_     *lex.Token
	lCurly      *lex.Token
	commClauses []*CommClause
	rCurly      *lex.Token
}

func (s SelectStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
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
		fmt.Printf("select后面必须是一个{。%s\n", lCurly.ErrorMsg())
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
		fmt.Printf("select需要一个'}'，在这里。%s\n", rCurly.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rCurly

	return &SelectStmt{select_: select_, lCurly: lCurly, commClauses: commClauses, rCurly: rCurly}
}

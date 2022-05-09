package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"fmt"
)

type TypeSwitchStmt struct {
	// typeSwitchStmt:
	//	SWITCH ( typeSwitchGuard
	//					| eos typeSwitchGuard
	//					| simpleStmt eos typeSwitchGuard)
	//					 L_CURLY typeCaseClause* R_CURLY;
	switch_ *lex.Token

	typeSwitchGuard *TypeSwitchGuard
	simpleStmt      SimpleStmt

	lCurly          *lex.Token
	typeCaseClauses []*TypeCaseClause
	rCurly          *lex.Token
}

func (a *TypeSwitchStmt) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*TypeSwitchStmt)(nil)

func (t TypeSwitchStmt) __Statement__() {
	panic("imposible")
}

func (t TypeSwitchStmt) __SwitchStmt__() {
	panic("imposible")
}

var _ SwitchStmt = (*TypeSwitchStmt)(nil)

func VisitTypeSwitchStmt(lexer *lex.Lexer) *TypeSwitchStmt {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	clone := lexer.Clone()

	switch_ := lexer.LA()
	if switch_.Type_() != lex.GoLexerSWITCH {
		return nil
	}
	lexer.Pop() // switch_

	VisitEos(lexer)

	typeSwitchGuard := VisitTypeSwitchGuard(lexer)
	var simpleStmt SimpleStmt
	if typeSwitchGuard == nil {
		simpleStmt = VisitSimpleStmt(lexer)
		if simpleStmt == nil {
			fmt.Printf("switch后面的表达式语法不正确。%s\n", switch_.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		VisitEos(lexer)
		typeSwitchGuard = VisitTypeSwitchGuard(lexer)
	}
	lCurly := lexer.LA()
	if lCurly.Type_() != lex.GoLexerL_CURLY {
		fmt.Println("switch语句的左花括号没有找到。%s\n", switch_.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lCurly

	var typeCaseClauses []*TypeCaseClause
	for true {
		typeCaseClause := VisitTypeCaseClause(lexer)
		if typeCaseClause != nil {
			typeCaseClauses = append(typeCaseClauses, typeCaseClause)
		} else {
			break
		}
	}
	rCurly := lexer.LA()
	if rCurly.Type_() != lex.GoLexerR_CURLY {
		fmt.Println("switch语句的右花括号没有找到。%s\n", switch_.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rCurly
	return &TypeSwitchStmt{switch_: switch_, typeSwitchGuard: typeSwitchGuard, simpleStmt: simpleStmt, lCurly: lCurly, typeCaseClauses: typeCaseClauses, rCurly: rCurly}
}

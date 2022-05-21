package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type TypeSwitchStmt struct {
	// typeSwitchStmt:
	//	SWITCH ( typeSwitchGuard
	//					| eos typeSwitchGuard
	//					| simpleStmt eos typeSwitchGuard)
	//					 L_CURLY typeCaseClause* R_CURLY;
	switch_ *lex.Token

	eos             *Eos
	typeSwitchGuard *TypeSwitchGuard
	simpleStmt      SimpleStmt

	lCurly          *lex.Token
	typeCaseClauses []*TypeCaseClause
	rCurly          *lex.Token
}

func (a *TypeSwitchStmt) Switch_() *lex.Token {
	return a.switch_
}

func (a *TypeSwitchStmt) SetSwitch_(switch_ *lex.Token) {
	a.switch_ = switch_
}

func (a *TypeSwitchStmt) Eos() *Eos {
	return a.eos
}

func (a *TypeSwitchStmt) SetEos(eos *Eos) {
	a.eos = eos
}

func (a *TypeSwitchStmt) TypeSwitchGuard() *TypeSwitchGuard {
	return a.typeSwitchGuard
}

func (a *TypeSwitchStmt) SetTypeSwitchGuard(typeSwitchGuard *TypeSwitchGuard) {
	a.typeSwitchGuard = typeSwitchGuard
}

func (a *TypeSwitchStmt) SimpleStmt() SimpleStmt {
	return a.simpleStmt
}

func (a *TypeSwitchStmt) SetSimpleStmt(simpleStmt SimpleStmt) {
	a.simpleStmt = simpleStmt
}

func (a *TypeSwitchStmt) LCurly() *lex.Token {
	return a.lCurly
}

func (a *TypeSwitchStmt) SetLCurly(lCurly *lex.Token) {
	a.lCurly = lCurly
}

func (a *TypeSwitchStmt) TypeCaseClauses() []*TypeCaseClause {
	return a.typeCaseClauses
}

func (a *TypeSwitchStmt) SetTypeCaseClauses(typeCaseClauses []*TypeCaseClause) {
	a.typeCaseClauses = typeCaseClauses
}

func (a *TypeSwitchStmt) RCurly() *lex.Token {
	return a.rCurly
}

func (a *TypeSwitchStmt) SetRCurly(rCurly *lex.Token) {
	a.rCurly = rCurly
}

func (a *TypeSwitchStmt) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.switch_)
	if a.eos != nil {
		cb.AppendString(";").AppendTreeNode(a.typeSwitchGuard)
	} else if a.simpleStmt != nil {
		cb.AppendTreeNode(a.simpleStmt).AppendString(";").AppendTreeNode(a.typeSwitchGuard)
	} else {
		cb.AppendTreeNode(a.typeSwitchGuard)
	}
	cb.AppendToken(a.lCurly).Newline()
	for _, clause := range a.typeCaseClauses {
		cb.AppendTreeNode(clause).Newline()
	}
	cb.AppendToken(a.rCurly)
	return cb
}

func (a *TypeSwitchStmt) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*TypeSwitchStmt)(nil)

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

	typeSwitchGuard := VisitTypeSwitchGuard(lexer)
	var eos *Eos
	var simpleStmt SimpleStmt
	if typeSwitchGuard == nil {
		eos = VisitEos(lexer)
		if eos != nil {
			typeSwitchGuard = VisitTypeSwitchGuard(lexer)
			if typeSwitchGuard == nil {
				fmt.Printf("TypeSwitchStmt,分号后面没有类型条件。%s\n", eos.semi.ErrorMsg())
				lexer.Recover(clone)
				return nil
			}
		} else {
			simpleStmt = VisitSimpleStmt(lexer)
			if simpleStmt == nil {
				fmt.Printf("TypeSwitchStmt,此处应该是一个简单表达式。%s\n", lexer.LA().ErrorMsg())
				lexer.Recover(clone)
				return nil
			}
			eos = VisitEos(lexer)
			if eos == nil {
				fmt.Printf("TypeSwitchStmt,此处应该是一个分号。%s\n", lexer.LA().ErrorMsg())
				lexer.Recover(clone)
				return nil
			}
			typeSwitchGuard = VisitTypeSwitchGuard(lexer)
			if typeSwitchGuard == nil {
				fmt.Printf("TypeSwitchStmt,此处应该是一个类型条件。%s\n", lexer.LA().ErrorMsg())
				lexer.Recover(clone)
				return nil
			}
		}
	}

	lCurly := lexer.LA()
	if lCurly.Type_() != lex.GoLexerL_CURLY {
		fmt.Printf("typeSwitchStmt,switch语句的左花括号没有找到。%s\n", switch_.ErrorMsg())
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
		fmt.Printf("typeSwitchStmt,switch语句的右花括号没有找到。%s\n", switch_.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rCurly
	return &TypeSwitchStmt{switch_: switch_, eos: eos, typeSwitchGuard: typeSwitchGuard, simpleStmt: simpleStmt, lCurly: lCurly, typeCaseClauses: typeCaseClauses, rCurly: rCurly}
}

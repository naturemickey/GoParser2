package ast

import (
	"GoParser2/lex"
	"fmt"
)

type IfStmt struct {
	// ifStmt:
	//	IF ( expression
	//			| eos expression
	//			| simpleStmt eos expression
	//		) block (
	//		ELSE (ifStmt | block)
	//	)?;

	if_ *lex.Token

	expression *Expression
	simpleStmt SimpleStmt

	block *Block

	else_ *lex.Token

	ifStmt    *IfStmt
	elseBlock *Block
}

func (i IfStmt) __Statement__() {
	//TODO implement me
	panic("implement me")
}

var _ Statement = (*IfStmt)(nil)

func VisitIfStmt(lexer *lex.Lexer) *IfStmt {
	if lexer.LA() == nil { // 文件结束
		return nil
	}

	clone := lexer.Clone()

	if_ := lexer.LA()
	if if_.Type_() != lex.GoLexerIF {
		return nil
	}
	lexer.Pop() // if_

	VisitEos(lexer)

	expression := VisitExpression(lexer)
	var simpleStmt SimpleStmt
	if expression == nil {
		simpleStmt = VisitSimpleStmt(lexer)
		if simpleStmt == nil {
			lexer.Recover(clone)
			return nil
		}
		VisitEos(lexer)
		expression = VisitExpression(lexer)
		if expression == nil {
			lexer.Recover(clone)
			return nil
		}
	}

	block := VisitBlock(lexer)
	if block == nil {
		lexer.Recover(clone)
		return nil
	}

	else_ := lexer.LA()
	if else_ != nil {
		lexer.Pop() // else_

		ifStmt := VisitIfStmt(lexer)
		if ifStmt != nil {
			return &IfStmt{if_: if_, expression: expression, simpleStmt: simpleStmt, block: block, else_: else_, ifStmt: ifStmt}
		} else {
			elseBlock := VisitBlock(lexer)
			if elseBlock == nil {
				fmt.Printf("else后面的语句块不对。%s\n", else_.ErrorMsg())
				lexer.Recover(clone)
				return nil
			}
			return &IfStmt{if_: if_, expression: expression, simpleStmt: simpleStmt, block: block, else_: else_, elseBlock: elseBlock}
		}
	} else {
		return &IfStmt{if_: if_, expression: expression, simpleStmt: simpleStmt, block: block}
	}
}

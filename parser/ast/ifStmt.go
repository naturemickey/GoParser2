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

	semi       *lex.Token // eos
	expression *Expression
	simpleStmt SimpleStmt

	block *Block

	else_ *lex.Token

	ifStmt    *IfStmt
	elseBlock *Block
}

func (a *IfStmt) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.if_)
	if a.simpleStmt != nil {
		cb.AppendTreeNode(a.simpleStmt).AppendString(";").AppendTreeNode(a.expression)
	} else if a.semi != nil {
		cb.AppendString(";").AppendTreeNode(a.expression)
	} else {
		cb.AppendTreeNode(a.expression)
	}
	cb.AppendTreeNode(a.block)
	cb.AppendToken(a.else_)
	cb.AppendTreeNode(a.ifStmt)
	cb.AppendTreeNode(a.block)
	return cb
}

func (a *IfStmt) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*IfStmt)(nil)

func (i IfStmt) __Statement__() {
	panic("imposible")
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

	//var expression *Expression
	//var simpleStmt SimpleStmt
	//var semi = lexer.LA()
	expression, simpleStmt, semi, success := _visitIfCondition(lexer)
	if !success {
		lexer.Recover(clone)
		return nil
	}
	//if semi.Type_() == lex.GoLexerSEMI { // eos expression
	//	lexer.Pop() // semi
	//	expression = VisitExpression(lexer)
	//	if expression == nil {
	//		lexer.Recover(clone)
	//		return nil
	//	}
	//} else {
	//	semi = nil
	//	// 先识别 simpleStmt eos expression
	//	simpleStmt = VisitSimpleStmt(lexer)
	//	if simpleStmt != nil {
	//		semi = lexer.LA()
	//		if semi.Type_() != lex.GoLexerSEMI {
	//			lexer.Recover(clone)
	//			return nil
	//		}
	//		lexer.Pop() // semi
	//		expression = VisitExpression(lexer)
	//		if expression == nil {
	//			lexer.Recover(clone)
	//			return nil
	//		}
	//	} else {
	//		// 再识别 expression
	//		expression = VisitExpression(lexer)
	//		if expression == nil {
	//			lexer.Recover(clone)
	//			return nil
	//		}
	//	}
	//}

	block := VisitBlock(lexer)
	if block == nil {
		// 	todo 修复 LiteralValue 与block 模式相同的问题
		pe := expression.primaryExpr
		if pe != nil {
			opd := pe.operand
			if opd != nil {
				cmpl := opd.literal
				if cmpl != nil {
					if c, ok := cmpl.(*CompositeLit); ok {
						literalValue := c.literalValue
						if literalValue != nil { // 几乎可以确定literalValue不会为空
							c.literalValue = nil
							blockLexer := lex.NewLexerWithCode(literalValue.String())
							expressionLexer := lex.NewLexerWithCode(expression.String())
							block = VisitBlock(blockLexer)
							expression = VisitExpression(expressionLexer)
							if expression == nil || block == nil {
								lexer.Recover(clone)
								return nil
							}
							goto L
						}
					}
				}
			}
		}
		lexer.Recover(clone)
		return nil
	}
L:

	else_ := lexer.LA()
	if else_.Type_() == lex.GoLexerELSE {
		lexer.Pop() // else_

		ifStmt := VisitIfStmt(lexer)
		if ifStmt != nil {
			return &IfStmt{if_: if_, semi: semi, expression: expression, simpleStmt: simpleStmt, block: block, else_: else_, ifStmt: ifStmt}
		} else {
			elseBlock := VisitBlock(lexer)
			if elseBlock == nil {
				fmt.Printf("else后面的语句块不对。%s\n", else_.ErrorMsg())
				lexer.Recover(clone)
				return nil
			}
			return &IfStmt{if_: if_, semi: semi, expression: expression, simpleStmt: simpleStmt, block: block, else_: else_, elseBlock: elseBlock}
		}
	} else {
		return &IfStmt{if_: if_, semi: semi, expression: expression, simpleStmt: simpleStmt, block: block}
	}
}

func _visitIfCondition(lexer *lex.Lexer) (*Expression, SimpleStmt, *lex.Token, bool) {
	clone := lexer.Clone()

	var expression *Expression
	var simpleStmt SimpleStmt
	var semi = lexer.LA()

	if semi.Type_() == lex.GoLexerSEMI { // eos expression
		lexer.Pop() // semi
		expression = VisitExpression(lexer)
		if expression == nil {
			lexer.Recover(clone)
			return nil, nil, nil, false
		}
	} else {
		semi = nil
		// 先识别 simpleStmt eos expression
		simpleStmt = VisitSimpleStmt(lexer)
		if simpleStmt != nil {
			semi = lexer.LA()
			if semi.Type_() != lex.GoLexerSEMI {
				lexer.Recover(clone)
				return nil, nil, nil, false
			}
			lexer.Pop() // semi
			expression = VisitExpression(lexer)
			if expression == nil {
				lexer.Recover(clone)
				return nil, nil, nil, false
			}
		} else {
			// 再识别 expression
			expression = VisitExpression(lexer)
			if expression == nil {
				lexer.Recover(clone)
				return nil, nil, nil, false
			}
		}
	}

	return expression, simpleStmt, semi, true
}

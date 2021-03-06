package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
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

func (a *IfStmt) If_() *lex.Token {
	return a.if_
}

func (a *IfStmt) SetIf_(if_ *lex.Token) {
	a.if_ = if_
}

func (a *IfStmt) Semi() *lex.Token {
	return a.semi
}

func (a *IfStmt) SetSemi(semi *lex.Token) {
	a.semi = semi
}

func (a *IfStmt) Expression() *Expression {
	return a.expression
}

func (a *IfStmt) SetExpression(expression *Expression) {
	a.expression = expression
}

func (a *IfStmt) SimpleStmt() SimpleStmt {
	return a.simpleStmt
}

func (a *IfStmt) SetSimpleStmt(simpleStmt SimpleStmt) {
	a.simpleStmt = simpleStmt
}

func (a *IfStmt) Block() *Block {
	return a.block
}

func (a *IfStmt) SetBlock(block *Block) {
	a.block = block
}

func (a *IfStmt) Else_() *lex.Token {
	return a.else_
}

func (a *IfStmt) SetElse_(else_ *lex.Token) {
	a.else_ = else_
}

func (a *IfStmt) IfStmt() *IfStmt {
	return a.ifStmt
}

func (a *IfStmt) SetIfStmt(ifStmt *IfStmt) {
	a.ifStmt = ifStmt
}

func (a *IfStmt) ElseBlock() *Block {
	return a.elseBlock
}

func (a *IfStmt) SetElseBlock(elseBlock *Block) {
	a.elseBlock = elseBlock
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
	//if a.else_ != nil {
	cb.AppendToken(a.else_)
	cb.AppendTreeNode(a.ifStmt)
	cb.AppendTreeNode(a.elseBlock)
	//}
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
	if lexer.LA() == nil { // ????????????
		return nil
	}

	clone := lexer.Clone()

	if_ := lexer.LA()
	if if_.Type_() != lex.GoLexerIF {
		return nil
	}
	lexer.Pop() // if_

	expression, simpleStmt, semi, success := _visitIfCondition(lexer)
	if !success {
		lexer.Recover(clone)
		return nil
	}

	block := VisitBlock(lexer)
	if block == nil {
		//  ?????? LiteralValue ???block ?????????????????????
		var pe = _getPrimaryFromExpression(expression)

		if pe != nil {
			opd := pe.operand
			if opd != nil {
				cmpl := opd.literal
				if cmpl != nil {
					if c, ok := cmpl.(*CompositeLit); ok {
						literalValue := c.literalValue
						if literalValue != nil { // ??????????????????literalValue????????????
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
				fmt.Printf("ifStmt,else???????????????????????????%s\n", else_.ErrorMsg())
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
		clone2 := lexer.Clone()
		// ????????? simpleStmt eos expression
		simpleStmt = VisitSimpleStmt(lexer)
		if simpleStmt != nil {
			semi = lexer.LA()
			if semi.Type_() != lex.GoLexerSEMI { // simpleStmt???????????????????????????expression?????????
				simpleStmt = nil
				semi = nil
				lexer.Recover(clone2)
				expression = VisitExpression(lexer)
				if expression == nil {
					lexer.Recover(clone)
					return nil, nil, nil, false
				}
			} else {
				lexer.Pop() // semi
				expression = VisitExpression(lexer)
				if expression == nil {
					lexer.Recover(clone)
					return nil, nil, nil, false
				}
			}
		} else {
			// ????????? expression
			expression = VisitExpression(lexer)
			if expression == nil {
				lexer.Recover(clone)
				return nil, nil, nil, false
			}
		}
	}

	return expression, simpleStmt, semi, true
}

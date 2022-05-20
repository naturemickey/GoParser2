package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type ForStmt struct {
	// forStmt: FOR (expression | forClause | rangeClause)? block;
	for_        *lex.Token
	expression  *Expression
	forClause   *ForClause
	rangeClause *RangeClause
	block       *Block
}

func (a *ForStmt) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.for_)
	cb.AppendTreeNode(a.expression)
	cb.AppendTreeNode(a.forClause)
	cb.AppendTreeNode(a.rangeClause)
	cb.AppendTreeNode(a.block)
	return cb
}

func (a *ForStmt) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ForStmt)(nil)

func (f ForStmt) __Statement__() {
	panic("imposible")
}

var _ Statement = (*ForStmt)(nil)

func VisitForStmt(lexer *lex.Lexer) *ForStmt {
	clone := lexer.Clone()

	if lexer.LA() == nil { // 文件结束
		return nil
	}

	for_ := lexer.LA()
	if for_.Type_() != lex.GoLexerFOR {
		return nil
	}
	lexer.Pop() // for_

	var expression *Expression
	var forClause *ForClause
	var rangeClause *RangeClause

	rangeClause = VisitRangeClause(lexer)
	if rangeClause == nil {
		forClause = VisitForClause(lexer)
		if forClause == nil {
			expression = VisitExpression(lexer)
		}
	}

	block := VisitBlock(lexer)
	if block == nil {
		// 修复 LiteralValue 与block 模式相同的问题
		// todo 下面有多个 if pe != nil { ... }，重复代码，研究如何去掉。
		if expression != nil {
			pe := _getPrimaryFromExpression(expression)
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
		}
		if rangeClause != nil {
			pe := _getPrimaryFromExpression(rangeClause.expression) // rangeClause.expression不可能为nil
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
								rangeClauseLexer := lex.NewLexerWithCode(rangeClause.String())
								block = VisitBlock(blockLexer)
								rangeClause = VisitRangeClause(rangeClauseLexer)
								if rangeClause == nil || block == nil {
									lexer.Recover(clone)
									return nil
								}
								goto L
							}
						}
					}
				}
			}
		}
		if forClause != nil {
			var pe *PrimaryExpr
			if forClause.postStmt != nil {
				switch simpleStmt := forClause.postStmt.(type) {
				case *SendStmt:
					pe = _getPrimaryFromExpression(simpleStmt.expression)
				case *IncDecStmt: // 不可能
				case *Assignment:
					pe = _getPrimaryFromExpressionList(simpleStmt.rExpressionList)
				case *ShortVarDecl:
					pe = _getPrimaryFromExpressionList(simpleStmt.expressionList)
				case *Expression:
					pe = _getPrimaryFromExpression(simpleStmt)
				}
			}
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
								forClauseLexer := lex.NewLexerWithCode(forClause.String())
								block = VisitBlock(blockLexer)
								forClause = VisitForClause(forClauseLexer)
								if forClause == nil || block == nil {
									lexer.Recover(clone)
									return nil
								}
								goto L
							}
						}
					}
				}
			}
		}
		lexer.Recover(clone)
		return nil
	}
L:

	return &ForStmt{for_: for_, expression: expression, forClause: forClause, rangeClause: rangeClause, block: block}
}

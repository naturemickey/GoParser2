package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type Operand struct {
	// operand    : literal | operandName | L_PAREN expression R_PAREN;
	// operandName: IDENTIFIER;

	literal     Literal
	operandName *lex.Token
	lParen      *lex.Token
	expression  *Expression
	rParen      *lex.Token
}

func (a *Operand) Literal() Literal {
	return a.literal
}

func (a *Operand) SetLiteral(literal Literal) {
	a.literal = literal
}

func (a *Operand) OperandName() *lex.Token {
	return a.operandName
}

func (a *Operand) SetOperandName(operandName *lex.Token) {
	a.operandName = operandName
}

func (a *Operand) LParen() *lex.Token {
	return a.lParen
}

func (a *Operand) SetLParen(lParen *lex.Token) {
	a.lParen = lParen
}

func (a *Operand) Expression() *Expression {
	return a.expression
}

func (a *Operand) SetExpression(expression *Expression) {
	a.expression = expression
}

func (a *Operand) RParen() *lex.Token {
	return a.rParen
}

func (a *Operand) SetRParen(rParen *lex.Token) {
	a.rParen = rParen
}

func (a *Operand) CodeBuilder() *CodeBuilder {
	return NewCB().AppendTreeNode(a.literal).AppendToken(a.operandName).
		AppendToken(a.lParen).AppendTreeNode(a.expression).AppendToken(a.rParen)
}

func (a *Operand) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*Operand)(nil)

func VisitOperand(lexer *lex.Lexer) *Operand {
	clone := lexer.Clone()

	lParen := lexer.LA()
	if lParen.Type_() == lex.GoLexerL_PAREN {
		lexer.Pop() // lParen

		expression := VisitExpression(lexer)
		if expression == nil {
			//fmt.Printf("operand,'('后面应该是一个表达式才对。%s\n", lParen.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}

		rParen := lexer.LA()
		if rParen.Type_() != lex.GoLexerR_PAREN {
			//fmt.Printf("operand,此处应该有一个')'。%s\n", rParen.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		lexer.Pop() // rParen

		return &Operand{lParen: lParen, expression: expression, rParen: rParen}
	}

	literal := VisitLiteral(lexer)
	if literal != nil {
		return &Operand{literal: literal}
	}

	operandName := lexer.LA()
	if operandName.Type_() == lex.GoLexerIDENTIFIER {
		lexer.Pop() // operandName
		return &Operand{operandName: operandName}
	}

	lexer.Recover(clone)
	return nil
}

package ast

import (
	"GoParser2/lex"
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

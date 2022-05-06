package ast

import "GoParser2/lex"

type Operand struct {
	// operand    : literal | operandName | L_PAREN expression R_PAREN;
	// operandName: IDENTIFIER;

	literal     *Literal
	operandName *lex.Token
	lParen      *lex.Token
	expression  *Expression
	rParen      *lex.Token
}

func VisitOperand(lexer *lex.Lexer) *Operand {
	panic("todo")
}

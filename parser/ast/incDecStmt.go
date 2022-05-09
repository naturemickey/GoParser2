package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
)

type IncDecStmt struct {
	// incDecStmt: expression (PLUS_PLUS | MINUS_MINUS);
	expression *Expression
	plusplus   *lex.Token
	minusminus *lex.Token
}

func (a *IncDecStmt) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*IncDecStmt)(nil)

func (i IncDecStmt) __Statement__() {
	panic("imposible")
}

func (i IncDecStmt) __SimpleStmt__() {
	panic("imposible")
}

var _ SimpleStmt = (*IncDecStmt)(nil)

func VisitIncDecStmt(lexer *lex.Lexer) *IncDecStmt {
	clone := lexer.Clone()

	expression := VisitExpression(lexer)
	if expression == nil {
		lexer.Recover(clone)
		return nil
	}

	la := lexer.LA()
	if la.Type_() == lex.GoLexerPLUS_PLUS {
		lexer.Pop()
		return &IncDecStmt{expression: expression, plusplus: la}
	} else if la.Type_() == lex.GoLexerMINUS_MINUS {
		lexer.Pop()
		return &IncDecStmt{expression: expression, minusminus: la}
	} else {
		lexer.Recover(clone)
		return nil
	}
}

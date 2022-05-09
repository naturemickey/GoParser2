package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
)

type Conversion struct {
	// conversion: nonNamedType L_PAREN expression COMMA? R_PAREN;
	nonNamedType *NonNamedType
	lParen       *lex.Token
	expression   *Expression
	comma        *lex.Token
	rParen       *lex.Token
}

func (a *Conversion) CodeBuilder() *util.CodeBuilder {
	cb := util.NewCB()
	cb.AppendTreeNode(a.nonNamedType)
	cb.AppendToken(a.lParen)
	cb.AppendTreeNode(a.expression)
	cb.AppendToken(a.comma)
	cb.AppendToken(a.rParen)
	return cb
}

func (a *Conversion) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*Conversion)(nil)

func VisitConversion(lexer *lex.Lexer) *Conversion {
	clone := lexer.Clone()

	nonNamedType := VisitNonNamedType(lexer)
	if nonNamedType == nil {
		return nil
	}

	lParen := lexer.LA()
	if lParen.Type_() != lex.GoLexerL_PAREN {
		// fmt.Printf("此处应该是一个'('。%s\n", lParen.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lParen

	expression := VisitExpression(lexer)
	if expression == nil {
		// fmt.Printf("这个'('后面应该有一个表达式。%s\n", lParen.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}

	comma := lexer.LA()
	if comma.Type_() == lex.GoLexerCOMMA {
		lexer.Pop() // comma
	}

	rParen := lexer.LA()
	if rParen.Type_() != lex.GoLexerR_PAREN {
		// fmt.Printf("此处应该是一个')'%s\n", rParen.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rParen

	return &Conversion{nonNamedType: nonNamedType, lParen: lParen, expression: expression, comma: comma, rParen: rParen}
}

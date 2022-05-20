package ast

import (
	"github.com/naturemickey/GoParser2/lex"
)

type Arguments struct {
	// arguments:
	//	L_PAREN (
	//		(expressionList | nonNamedType (COMMA expressionList)?) ELLIPSIS? COMMA?
	//	)? R_PAREN;

	lParen *lex.Token

	expressionList *ExpressionList
	nonNamedType   *NonNamedType
	comma          *lex.Token

	ellipsis *lex.Token
	comma2   *lex.Token

	rParen *lex.Token
}

func (a *Arguments) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.lParen)
	if a.nonNamedType != nil {
		cb.AppendTreeNode(a.nonNamedType)
		cb.AppendToken(a.comma)
		cb.AppendTreeNode(a.expressionList)
	} else {
		cb.AppendTreeNode(a.expressionList)
	}
	cb.AppendToken(a.ellipsis)
	cb.AppendToken(a.comma2)
	cb.AppendToken(a.rParen)
	return cb
}

func (a *Arguments) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*Arguments)(nil)

func VisitArguments(lexer *lex.Lexer) *Arguments {
	clone := lexer.Clone()

	lParen := lexer.LA()
	if lParen.Type_() != lex.GoLexerL_PAREN {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // lParen

	var (
		expressionList *ExpressionList
		nonNamedType   *NonNamedType
		comma          *lex.Token
	)

	expressionList = VisitExpressionList(lexer)
	if expressionList == nil {
		nonNamedType = VisitNonNamedType(lexer)
		comma = lexer.LA()
		if comma.Type_() == lex.GoLexerCOMMA {
			lexer.Pop() // comma
			expressionList = VisitExpressionList(lexer)
			if expressionList == nil {
				lexer.Recover(clone)
				return nil
			}
		} else {
			comma = nil
		}
	}

	ellipsis := lexer.LA()
	if ellipsis.Type_() != lex.GoLexerELLIPSIS {
		ellipsis = nil
	} else {
		lexer.Pop() // ellipsis
	}

	comma2 := lexer.LA()
	if comma2.Type_() != lex.GoLexerCOMMA {
		comma2 = nil
	} else {
		lexer.Pop() // comma2
	}

	rParen := lexer.LA()
	if rParen.Type_() != lex.GoLexerR_PAREN {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // rParen

	return &Arguments{lParen: lParen, expressionList: expressionList, nonNamedType: nonNamedType, comma: comma, ellipsis: ellipsis,
		comma2: comma2, rParen: rParen}
}

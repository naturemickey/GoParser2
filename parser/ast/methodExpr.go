package ast

import "GoParser2/lex"

type MethodExpr struct {
	// methodExpr: nonNamedType DOT IDENTIFIER;
	nonNamedType *NonNamedType
	dot          *lex.Token
	identifier   *lex.Token
}

func VisitMethodExpr(lexer *lex.Lexer) *MethodExpr {
	clone := lexer.Clone()

	nonNamedType := VisitNonNamedType(lexer)
	if nonNamedType == nil {
		lexer.Recover(clone)
		return nil
	}

	dot := lexer.LA()
	if dot.Type_() != lex.GoLexerDOT {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // dot

	identifier := lexer.LA()
	if identifier.Type_() != lex.GoLexerIDENTIFIER {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // identifier

	return &MethodExpr{nonNamedType: nonNamedType, dot: dot, identifier: identifier}
}

package ast

import "GoParser2/lex"

type TypeName struct {
	// typeName: qualifiedIdent | IDENTIFIER;
	qualifiedIdent *QualifiedIdent
	identifier     *lex.Token
}

func (t TypeName) __IMethodspecOrTypename__() {
	panic("imposible")
}

var _ IMethodspecOrTypename = (*TypeName)(nil)

func VisitTypeName(lexer *lex.Lexer) *TypeName {
	qualifiedIdent := VisitQualifiedIdent(lexer)
	if qualifiedIdent == nil {
		identifier := lexer.LA()
		if identifier.Type_() != lex.GoLexerIDENTIFIER {
			return nil
		}
		lexer.Pop() // identifier
		return &TypeName{identifier: identifier}
	}
	return &TypeName{qualifiedIdent: qualifiedIdent}
}

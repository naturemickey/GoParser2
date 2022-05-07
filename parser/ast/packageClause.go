package ast

import (
	"GoParser2/lex"
	"fmt"
)

type PackageClause struct {
	// packageClause: PACKAGE packageName = IDENTIFIER;
	package_    *lex.Token
	packageName *lex.Token
}

func VisitPackageClause(lexer *lex.Lexer) *PackageClause {
	package_ := lexer.LA()
	if package_.Type_() != lex.GoLexerPACKAGE {
		return nil
	}
	lexer.Pop() // 把package扔掉

	packageName := lexer.LA()
	if packageName.Type_() != lex.GoLexerIDENTIFIER {
		fmt.Printf("package后面需要跟着一个id做为包名， %s", packageName.ErrorMsg())
		return nil
	}
	lexer.Pop()

	return &PackageClause{package_: package_, packageName: packageName}
}

package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type PackageClause struct {
	// packageClause: PACKAGE packageName = IDENTIFIER;
	package_    *lex.Token
	packageName *lex.Token
}

func (this *PackageClause) Package_() *lex.Token {
	return this.package_
}

func (this *PackageClause) SetPackage_(package_ *lex.Token) {
	this.package_ = package_
}

func (this *PackageClause) SetPackageName(packageName *lex.Token) {
	this.packageName = packageName
}

func (this *PackageClause) PackageName() string {
	return this.packageName.Literal()
}

func (a *PackageClause) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.package_).AppendToken(a.packageName).Newline()
}

func (a *PackageClause) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*PackageClause)(nil)

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

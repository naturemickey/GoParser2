package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
	"github.com/naturemickey/GoParser2/parser/util"
	"reflect"
)

type SourceFile struct {
	// sourceFile:
	//	packageClause eos (importDecl eos)* (
	//		(functionDecl | methodDecl | declaration) eos
	//	)* EOF;
	packageClause *PackageClause
	importDecls   []*ImportDecl
	fmds          []IFunctionMethodDeclaration
}

func (a *SourceFile) SetPackageClause(packageClause *PackageClause) {
	a.packageClause = packageClause
}

func (a *SourceFile) SetImportDecls(importDecls []*ImportDecl) {
	a.importDecls = importDecls
}

func (a *SourceFile) SetFmds(fmds []IFunctionMethodDeclaration) {
	a.fmds = fmds
}

func (a *SourceFile) PackageClause() *PackageClause {
	return a.packageClause
}

func (a *SourceFile) ImportDecls() []*ImportDecl {
	return a.importDecls
}

func (a *SourceFile) Fmds() []IFunctionMethodDeclaration {
	return a.fmds
}

func NewSourceFile(packageClause *PackageClause, importDecls []*ImportDecl, fmds []IFunctionMethodDeclaration) *SourceFile {
	return &SourceFile{packageClause: packageClause, importDecls: importDecls, fmds: fmds}
}

func (a *SourceFile) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendTreeNode(a.packageClause).Newline()
	for _, decl := range a.importDecls {
		cb.AppendTreeNode(decl).Newline()
	}
	for _, fmd := range a.fmds {
		cb.AppendTreeNode(fmd).Newline().Newline()
	}
	return cb
}

func (a *SourceFile) String() string {
	code := a.CodeBuilder().String()
	return util.GoFmt(code)
	//return code
}

var _ ITreeNode = (*SourceFile)(nil)

func VisitSourceFile(lexer *lex.Lexer) *SourceFile {
	var packageClause = VisitPackageClause(lexer)

	if packageClause == nil {
		fmt.Println("sourceFile,文件第一个部分必须是package语句")
		return nil
	}

	VisitEos(lexer)

	var importDecls []*ImportDecl
	var fmds []IFunctionMethodDeclaration

	for la := lexer.LA(); la != nil && la.Type_() == lex.GoLexerIMPORT; {
		importDecl := VisitImportDecl(lexer)
		if importDecl != nil {
			importDecls = append(importDecls, importDecl)
		} else {
			break
		}
		VisitEos(lexer)
	}
	for {
		la := lexer.LA()
		if la != nil && (la.Type_() == lex.GoLexerFUNC || la.Type_() == lex.GoLexerVAR ||
			la.Type_() == lex.GoLexerCONST || la.Type_() == lex.GoLexerTYPE) {

			if la.Type_() == lex.GoLexerFUNC {
				functionDecl := VisitFunctionDecl(lexer)
				if functionDecl != nil {
					fmds = append(fmds, functionDecl)
					continue
				}
				methodDecl := VisitMethodDecl(lexer)
				if methodDecl != nil {
					fmds = append(fmds, methodDecl)
					continue
				}
				fmt.Printf("sourceFile,func关键字后面的语法不对 %s\n", la.ErrorMsg())
				return nil
			} else if la.Type_() == lex.GoLexerVAR || la.Type_() == lex.GoLexerCONST || la.Type_() == lex.GoLexerTYPE {
				declaration := VisitDeclaration(lexer)
				if declaration != nil && !reflect.ValueOf(declaration).IsNil() {
					fmds = append(fmds, declaration)
				} else {
					fmt.Printf("sourceFile,看到了%s关键字，但识别不到后面的定义。 %s\n", la.Literal(), la.ErrorMsg())
					return nil
				}
			}
			VisitEos(lexer)
		} else {
			break
		}
	}

	la := lexer.LA()
	if la != nil {
		fmt.Printf("sourceFile,无法识别的语法元素，%s\n", la.ErrorMsg())
	}

	return NewSourceFile(packageClause, importDecls, fmds)
}

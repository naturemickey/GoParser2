package ast

import (
	"GoParser2/lex"
	"fmt"
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

func NewSourceFile(packageClause *PackageClause, importDecls []*ImportDecl, fmds []IFunctionMethodDeclaration) *SourceFile {
	return &SourceFile{packageClause: packageClause, importDecls: importDecls, fmds: fmds}
}

func VisitSourceFile(lexer *lex.Lexer) *SourceFile {
	var packageClause = VisitPackageClause(lexer)
	var importDecls []*ImportDecl
	var fmds []IFunctionMethodDeclaration

	if packageClause == nil {
		fmt.Println("文件第一个部分必须是package语句")
		return nil
	}

	for la := lexer.LA(); la.Type_() == lex.GoLexerIMPORT; {
		importDecl := VisitImportDecl(*lexer)
		if importDecl != nil {
			importDecls = append(importDecls, importDecl)
		} else {
			break
		}
	}

	for la := lexer.LA(); la.Type_() == lex.GoLexerFUNC || la.Type_() == lex.GoLexerVAR ||
		la.Type_() == lex.GoLexerCONST || la.Type_() == lex.GoLexerTYPE; {

		if la.Type_() == lex.GoLexerFUNC {
			functionDecl := VisitFunctionDecl(lexer)
			if functionDecl != nil {
				fmds = append(fmds, functionDecl)
			} else {
				methodDecl := VisitMethodDecl(lexer)
				if methodDecl != nil {
					fmds = append(fmds, methodDecl)
				} else {
					fmt.Printf("看到了func关键字，但识别不到function或method的定义。 %s\n", la.ErrorMsg())
					return nil
				}
			}
		} else if la.Type_() == lex.GoLexerVAR || la.Type_() == lex.GoLexerCONST || la.Type_() == lex.GoLexerTYPE {
			declaration := VisitDeclaration(lexer)
			if declaration != nil {
				fmds = append(fmds, declaration)
			} else {
				fmt.Printf("看到了%s关键字，但识别不到后面的定义。 %s\n", la.Literal(), la.ErrorMsg())
				return nil
			}
		}
	}
	return NewSourceFile(packageClause, importDecls, fmds)
}

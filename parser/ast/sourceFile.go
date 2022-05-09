package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"fmt"
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

func (a *SourceFile) String() string {
	//TODO implement me
	panic("implement me")
}

var _ parser.ITreeNode = (*SourceFile)(nil)

func NewSourceFile(packageClause *PackageClause, importDecls []*ImportDecl, fmds []IFunctionMethodDeclaration) *SourceFile {
	return &SourceFile{packageClause: packageClause, importDecls: importDecls, fmds: fmds}
}

func VisitSourceFile(lexer *lex.Lexer) *SourceFile {
	var packageClause = VisitPackageClause(lexer)

	if packageClause == nil {
		fmt.Println("文件第一个部分必须是package语句")
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
				la1 := lexer.LA1()
				if la1 != nil && la1.Type_() == lex.GoLexerL_PAREN { // method
					methodDecl := VisitMethodDecl(lexer)
					if methodDecl != nil {
						fmds = append(fmds, methodDecl)
					} else {
						fmt.Printf("看到了func关键字，但识别不到method的定义。 %s\n", la.ErrorMsg())
						return nil
					}
				} else if la1 != nil && la1.Type_() == lex.GoLexerIDENTIFIER { // function
					functionDecl := VisitFunctionDecl(lexer)
					if functionDecl != nil {
						fmds = append(fmds, functionDecl)
					} else {
						fmt.Printf("看到了func关键字，但识别不到function的定义。 %s\n", la.ErrorMsg())
						return nil
					}
				} else { // unknown, error
					fmt.Printf("func关键字后面必须跟着函数名或者receiver。 %s\n", la.ErrorMsg())
					return nil
				}

			} else if la.Type_() == lex.GoLexerVAR || la.Type_() == lex.GoLexerCONST || la.Type_() == lex.GoLexerTYPE {
				declaration := VisitDeclaration(lexer)
				if declaration != nil && !reflect.ValueOf(declaration).IsNil() {
					fmds = append(fmds, declaration)
				} else {
					fmt.Printf("看到了%s关键字，但识别不到后面的定义。 %s\n", la.Literal(), la.ErrorMsg())
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
		fmt.Printf("无法识别的语法元素，%s\n", la.ErrorMsg())
	}

	return NewSourceFile(packageClause, importDecls, fmds)
}

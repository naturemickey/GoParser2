package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type ImportDecl struct {
	// importDecl: IMPORT (importSpec | L_PAREN (importSpec eos)* R_PAREN);
	import_     *lex.Token
	lParen      *lex.Token
	importSpecs []*ImportSpec
	rParen      *lex.Token
}

func (a *ImportDecl) Import_() *lex.Token {
	return a.import_
}

func (a *ImportDecl) SetImport_(import_ *lex.Token) {
	a.import_ = import_
}

func (a *ImportDecl) LParen() *lex.Token {
	return a.lParen
}

func (a *ImportDecl) SetLParen(lParen *lex.Token) {
	a.lParen = lParen
}

func (a *ImportDecl) ImportSpecs() []*ImportSpec {
	return a.importSpecs
}

func (a *ImportDecl) SetImportSpecs(importSpecs []*ImportSpec) {
	a.importSpecs = importSpecs
}

func (a *ImportDecl) RParen() *lex.Token {
	return a.rParen
}

func (a *ImportDecl) SetRParen(rParen *lex.Token) {
	a.rParen = rParen
}

func NewImportDecl(import_ *lex.Token, lParen *lex.Token, importSpecs []*ImportSpec, rParen *lex.Token) *ImportDecl {
	return &ImportDecl{import_: import_, lParen: lParen, importSpecs: importSpecs, rParen: rParen}
}

func (a *ImportDecl) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	cb.AppendToken(a.import_)
	if a.lParen != nil {
		cb.AppendToken(a.lParen).Newline()
		for _, spec := range a.importSpecs {
			cb.AppendTreeNode(spec).Newline()
		}
		cb.AppendToken(a.rParen)
	} else {
		cb.AppendTreeNode(a.importSpecs[0])
	}
	return cb
}

func (a *ImportDecl) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ImportDecl)(nil)

func VisitImportDecl(lexer *lex.Lexer) *ImportDecl {
	clone := lexer.Clone()

	import_ := lexer.LA()
	if import_.Type_() != lex.GoLexerIMPORT {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // import_

	var importSpecs []*ImportSpec

	lParen := lexer.LA()
	if lParen.Type_() == lex.GoLexerL_PAREN {
		lexer.Pop() // 丢弃左括号

		for true {
			importSpec := VisitImportSpec(lexer)
			if importSpec != nil {
				importSpecs = append(importSpecs, importSpec)
				VisitEos(lexer)
			} else {
				break
			}
		}

		rParen := lexer.LA()
		if rParen.Type_() != lex.GoLexerR_PAREN {
			fmt.Printf("importDecl,此处应该是一个')'才对。%s\n", rParen.ErrorMsg())
			lexer.Recover(clone)
			return nil
		}
		lexer.Pop() // rParen

		return &ImportDecl{import_: import_, lParen: lParen, importSpecs: importSpecs, rParen: rParen}
	} else {
		importSpec := VisitImportSpec(lexer)
		if importSpec == nil {
			fmt.Printf("importDecl,import后面没看到路径描述，%s\n", import_.ErrorMsg())
			lexer.Recover(clone)
			return nil
		} else {
			importSpecs = append(importSpecs, importSpec)
		}
		return &ImportDecl{import_: import_, importSpecs: importSpecs}
	}

}

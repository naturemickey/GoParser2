package ast

import (
	"GoParser2/lex"
	"fmt"
)

type ImportDecl struct {
	// importDecl: IMPORT (importSpec | L_PAREN (importSpec eos)* R_PAREN);
	import_     *lex.Token
	importSpecs []*ImportSpec
}

func VisitImportDecl(lexer *lex.Lexer) *ImportDecl {
	import_ := lexer.LA()
	if import_.Type_() != lex.GoLexerIMPORT {
		return nil
	}
	lexer.Pop() // 丢弃import关键字

	la := lexer.LA()
	hasParen := false
	if la.Type_() == lex.GoLexerL_PAREN {
		hasParen = true
		lexer.Pop() // 丢弃左括号
	}

	var importSpecs []*ImportSpec

	if !hasParen {
		// 如果没有括号，则后面一定有一个importSpec
		importSpec := VisitImportSpec(lexer)
		if importSpec == nil {
			fmt.Printf("import后面没看到路径描述，%s\n", la.ErrorMsg())
			return nil
		} else {
			importSpecs = append(importSpecs, importSpec)
		}
	} else {
		for true {
			importSpec := VisitImportSpec(lexer)
			if importSpec != nil {
				importSpecs = append(importSpecs, importSpec)
				VisitEos(lexer)
			} else {
				break
			}
		}

		la := lexer.LA()
		if la.Type_() != lex.GoLexerR_PAREN {
			fmt.Printf("有左括号，但没找到右括号。%s\n", la.ErrorMsg())
			return nil
		}
	}

	return &ImportDecl{import_: import_, importSpecs: importSpecs}
}

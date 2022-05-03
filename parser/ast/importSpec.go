package ast

import "GoParser2/lex"

type ImportSpec struct {
	// importSpec: alias = (DOT | IDENTIFIER)? importPath;
	alias      *lex.Token
	importPath *String_
}

func VisitImportSpec(lexer *lex.Lexer) *ImportSpec {
	// 识别失败不是真失败，需要恢复lexer，因为外面可能会继续使用
	clone := lexer.Clone()

	alias := lexer.LA()
	if alias.Type_() != lex.GoLexerDOT && alias.Type_() != lex.GoLexerIDENTIFIER {
		alias = nil
	} else {
		lexer.Pop()
	}

	importPath := VisitString(lexer)
	if importPath == nil {
		lexer.Recover(clone)
		return nil
	}
	return &ImportSpec{alias: alias, importPath: importPath}
}

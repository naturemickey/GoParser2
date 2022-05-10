package ast

import (
	"GoParser2/lex"
)

type ImportSpec struct {
	// importSpec: alias = (DOT | IDENTIFIER)? importPath;
	// importPath: string_;
	// string_   : RAW_STRING_LIT | INTERPRETED_STRING_LIT;
	alias      *lex.Token
	importPath *lex.Token
}

func (a *ImportSpec) CodeBuilder() *CodeBuilder {
	return NewCB().AppendToken(a.alias).AppendToken(a.importPath)
}

func (a *ImportSpec) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*ImportSpec)(nil)

func VisitImportSpec(lexer *lex.Lexer) *ImportSpec {
	// 识别失败不是真失败，需要恢复lexer，因为外面可能会继续使用
	clone := lexer.Clone()

	alias := lexer.LA()
	if alias.Type_() != lex.GoLexerDOT && alias.Type_() != lex.GoLexerIDENTIFIER {
		alias = nil
	} else {
		lexer.Pop() // alias
	}

	importPath := lexer.LA()
	if importPath.Type_() != lex.GoLexerRAW_STRING_LIT &&
		importPath.Type_() != lex.GoLexerINTERPRETED_STRING_LIT {
		lexer.Recover(clone)
		return nil
	}
	lexer.Pop() // importPath

	return &ImportSpec{alias: alias, importPath: importPath}
}

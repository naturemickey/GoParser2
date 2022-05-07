package parser

import (
	"GoParser2/lex"
	"GoParser2/parser/ast"
)

func Parse(filepath string) *ast.SourceFile {
	lexer := lex.NewLexerWithFile(filepath)
	return ast.VisitSourceFile(lexer)
}

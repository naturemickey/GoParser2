package parser

import (
	"github.com/naturemickey/GoParser2/lex"
	"github.com/naturemickey/GoParser2/parser/ast"
)

func Parse(filepath string) *ast.SourceFile {
	lexer := lex.NewLexerWithFile(filepath)
	return ast.VisitSourceFile(lexer)
}

func ParseCode(code string) *ast.SourceFile {
	lexer := lex.NewLexerWithCode(code)
	return ast.VisitSourceFile(lexer)
}

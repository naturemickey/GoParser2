package parser

import (
	"GoParser2/lex"
	"GoParser2/parser/ast"
)

type parser struct {
}

func (this *parser) parse(filepath string) *ast.SourceFile {
	lexer := lex.NewLexer2WithFile(filepath)
	return ast.VisitSourceFile(lexer)
}

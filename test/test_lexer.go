package main

import (
	"GoParser2/lex"
)

func main() {
	filepath := "test/test_example/unicodeIdentifier_go"
	lexer := lex.NewLexerInnerWithFileInner(filepath, lex.NFA)
	for {
		token := lexer.NextToken()
		if token == nil {
			break
		}
		println(token.String())
	}
}

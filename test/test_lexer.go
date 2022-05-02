package main

import "GoParser2/lex"

func main() {
	lexer := lex.NewLexerWithFile("./test/test_example/arrayEllipsisDecls_go")

	for true {
		token := lexer.NextToken()
		if token == nil {
			break
		}
		println(token.String())
	}
}

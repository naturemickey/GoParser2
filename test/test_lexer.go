package main

import "GoParser2/lex"

func main() {
	lexer := lex.NewLexerWithFile("./test/test_example/arrayEllipsisDecls_go")

	for token := lexer.Pop(); token != nil; {
		println(token.String())
		token = lexer.Pop()
	}
}

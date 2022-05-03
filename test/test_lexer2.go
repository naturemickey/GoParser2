package main

import "GoParser2/lex"

func main() {
	lexer := lex.NewLexerWithCode(".")

	for token := lexer.Pop(); token != nil; {
		println(token.String())
		token = lexer.Pop()
	}
}

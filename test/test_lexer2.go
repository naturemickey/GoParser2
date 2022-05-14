package main

import (
	"GoParser2/lex"
)

func main() {
	lexer := lex.NewLexerWithCode("/*@Bean(name=abc,cached=true)*/")

	for {
		token := lexer.LA()
		if token != nil {
			println(token.String())
		} else {
			break
		}
		lexer.Pop()
	}
}

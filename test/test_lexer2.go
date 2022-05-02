package main

import "GoParser2/lex"

func main() {

	lexer := lex.NewLexerWithCode(".")

	for true {
		token := lexer.NextToken()
		if token == nil {
			break
		}
		println(token.String())
	}
}

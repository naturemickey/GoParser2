package main

import "GoParser2/lex"

func main() {
	nfa := lex.Or(
		//lex.NewNfaWithString("func", lex.GoLexerFUNC),
		//lex.New_WS_nfa(),
		//lex.New_TERMINATOR_nfa()
		lex.NewNfaWithString(".", lex.GoLexerDOT),
	)
	lexer := lex.NewLexerWithCodeInner(".", nfa)

	for true {
		token := lexer.NextToken()
		if token == nil {
			break
		}
		println(token.String())
	}
}

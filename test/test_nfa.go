package main

import "GoParser2/lex"

func main() {
	nfa := lex.Or(
		//lex.NewNfaWithString("func", lex.GoLexerFUNC),
		//lex.New_WS_nfa(),
		//lex.New_TERMINATOR_nfa()
		lex.NewNfaWithString(".", lex.GoLexerDOT),
	)
	lexer := lex.NewLexerInnerWithCodeInner(".", nfa)

	for token := lexer.Pop(); token != nil; {
		println(token.String())
		token = lexer.Pop()
	}
}

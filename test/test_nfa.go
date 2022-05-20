package main

import (
	"github.com/naturemickey/GoParser2/lex"
)

func main() {
	//nfa := lex.Or(
	//	//lex.NewNfaWithString("func", lex.GoLexerFUNC),
	//	//lex.New_WS_nfa(),
	//	//lex.New_TERMINATOR_nfa()
	//	lex.NewNfaWithString(".", lex.GoLexerDOT),
	//)
	nfa := lex.New_LINE_COMMENT_nfa()
	lexer := lex.NewLexerInnerWithCodeInner("//\tlambd := func(s string) { Sleep(10); fmt.Println(s) }\n//\tlambd(\"From lambda!\")\n//\tfunc() { fmt.Println(\"Create and invoke!\")}()",
		nfa)

	for {
		token := lexer.LA()
		if token == nil {
			break
		}

		println(token.String())
		lexer.Pop()
	}
}

package main

import (
	"GoParser2/lex"
	"GoParser2/parser/ast"
)

func main() {
	//{
	//	lexer := lex.NewLexerWithCode("fields := map[string]interface{}{\n\t\t\"extra\": `{\"test_update_fields\":\"update_fields_by_test\"}`,\n\t}")
	//	a := ast.VisitStatement(lexer)
	//	println(a, a.String())
	//}

	{
		lexer := lex.NewLexerWithCode("{\nbreak\n\tLabel:\n\t}")
		a := ast.VisitBlock(lexer)
		println(a, a.String())
	}
}

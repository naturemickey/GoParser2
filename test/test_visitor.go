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
		lexer := lex.NewLexerWithCode("{\nif string(decrypt) != theMsg {\n\t\tt.Fatal(\"test fail! the msg is not equal\")\n\t}\n}")
		a := ast.VisitBlock(lexer)
		println(a.String())
	}
}

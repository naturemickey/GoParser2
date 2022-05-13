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
		lexer := lex.NewLexerWithCode("for a < b {\n\t\tfmt.Println(\"From condition-only ForStmt\")\n\t\tbreak\n\t}")
		a := ast.VisitForStmt(lexer)
		println(a.String())
	}
}

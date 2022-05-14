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
		lexer := lex.NewLexerWithCode("{\nif len(*partnerIdList) > 0 {\n\t\t\t\tfilter.AndEntityIdIn(*partnerIdList)\n\t\t\t\t*partnerIdList = []int64{}\n\t\t\t}\n}")
		a := ast.VisitBlock(lexer)
		println(a.String())
	}
}

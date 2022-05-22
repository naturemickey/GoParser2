package main

import (
	"github.com/naturemickey/GoParser2/lex"
	"github.com/naturemickey/GoParser2/parser/ast"
)

func main() {
	//{
	//	lexer := lex.NewLexerWithCode("fields := map[string]interface{}{\n\t\t\"extra\": `{\"test_update_fields\":\"update_fields_by_test\"}`,\n\t}")
	//	a := ast.VisitStatement(lexer)
	//	println(a, a.String())
	//}

	{
		lexer := lex.NewLexerWithCode("interface {\n\tMethodExample /*@MethodSpecAnnoName*/ ()\n}")
		a := ast.VisitInterfaceType(lexer)
		println(a.String())
	}
}

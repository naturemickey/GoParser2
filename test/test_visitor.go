package main

import (
	"GoParser2/lex"
	"GoParser2/parser/ast"
)

func main() {
	lexer := lex.NewLexerWithCode("for range kvs {\n\t\tfmt.Printf(\"empty range\\n\")\n\t}")
	a := ast.VisitForStmt(lexer)
	println(a, a.String())

}

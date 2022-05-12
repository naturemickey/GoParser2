package main

import (
	"GoParser2/lex"
	"GoParser2/parser/ast"
)

func main() {
	lexer := lex.NewLexerWithCode("sum += num")
	a := ast.VisitExpression(lexer)
	println(a)

}

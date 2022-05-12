package main

import "GoParser2/parser"

func main() {
	sourceFile := parser.Parse("test/test_example/foreachStmts_go")

	println(sourceFile)
}

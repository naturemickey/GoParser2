package main

import "GoParser2/parser"

func main() {
	sourceFile := parser.Parse("test/test_example/sliceDecls_go")

	println(sourceFile)
}

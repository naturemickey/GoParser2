package main

import "GoParser2/parser"

func main() {
	sourceFile := parser.Parse("test/test_example/empty_init_go")

	println(sourceFile)
}

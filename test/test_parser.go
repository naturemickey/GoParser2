package main

import "GoParser2/parser"

func main() {
	sourceFile := parser.Parse("test/test_example/interface_inheritance_go")

	println(sourceFile)
}

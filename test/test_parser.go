package main

import "GoParser2/parser"

func main() {
	sourceFile := parser.Parse("test/test_example/anonymousMethods_go")

	println(sourceFile)
}

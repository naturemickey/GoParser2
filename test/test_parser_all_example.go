package main

import (
	"GoParser2/parser"
	utils "GoParser2/test/util"
)

func main() {
	utils.WalkDir("test/test_example/", func(path string) {
		println(path)
		parser.Parse(path)
		//sourceFile := parser.Parse(path)

		// println(sourceFile)
	})
}

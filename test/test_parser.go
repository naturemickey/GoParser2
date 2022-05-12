package main

import (
	"GoParser2/parser"
	"go/format"
)

func main() {
	sourceFile := parser.Parse("test/test_example/select_go")

	println(sourceFile)
	//println("============================")
	//println(sourceFile.String())
	//println("============================")
	//println(string(GoFmt(sourceFile.String())))
}

func GoFmt(content string) string {
	bs, err := format.Source([]byte(content))
	if err != nil {
		panic(err)
	}
	return string(bs)
}

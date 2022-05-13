package main

import (
	"GoParser2/parser"
	"go/format"
)

func main() {
	sourceFile := parser.Parse("test/test_example/typeAsserts_go")

	println(sourceFile)
	println(sourceFile.String())
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

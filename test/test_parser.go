package main

import (
	"github.com/naturemickey/GoParser2/parser"
	"github.com/naturemickey/GoParser2/parser/util"
)

func main() {
	sourceFile := parser.Parse("test/test_example/switchesExpr_go")

	//println(sourceFile)
	//println(sourceFile.String())
	println(util.GoFmt(sourceFile.String()))
	//println("============================")
	//println(sourceFile.String())
	//println("============================")
	//println(string(GoFmt(sourceFile.String())))
}

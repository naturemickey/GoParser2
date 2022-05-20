package main

import (
	"github.com/naturemickey/GoParser2/parser"
	"github.com/naturemickey/GoParser2/parser/util"
	utils "github.com/naturemickey/GoParser2/test/util"
	"io/ioutil"
)

func main() {
	utils.WalkDir("test/test_example/", func(path string) {
		println("====================================")
		println(path)
		println("------------------------------------")
		c, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		code := string(c)

		//println(code)
		//println("------------------------------------")
		sourceFile := parser.ParseCode(code)

		code = sourceFile.String()

		//println(code)
		//println("------------------------------------")

		println(util.GoFmt(code))
		println("------------------------------------")
	})
}

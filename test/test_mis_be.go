package main

import (
	"GoParser2/parser"
	"GoParser2/parser/util"
	utils "GoParser2/test/util"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	utils.WalkDir("/Users/mickey/git/mis-backend/", parse_go_file, "/Users/mickey/git/mis-backend/vendor/")
	println(time.Since(start).Seconds())
}

func parse_go_file(filePath string) {
	if strings.HasSuffix(filePath, ".go") {
		//println("=======================================================================================")
		println("开始处理文件：", filePath)

		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err.Error())
		}
		originContent := string(file)

		//println("原始文件：\n", originContent)

		sourceFile := parser.ParseCode(originContent)

		//input := antlr.NewInputStream(originContent)
		//lexer := antlr4.NewGoLexer(input)
		//stream := antlr.NewCommonTokenStream(lexer, 0)
		//p := antlr4.NewGoParser(stream)
		//p.SetErrorHandler(antlr.NewBailErrorStrategy())
		//// p.SetErrorHandler(antlr.NewDefaultErrorStrategy())
		//// p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
		//// p.BuildParseTrees = true
		//tree := p.SourceFile()
		//
		//visitor := &parser.GoParserVisitorImpl{}
		//
		//accept := tree.Accept(visitor).(ast.INode)

		content := sourceFile.String()

		//println("------------------------")
		//println("解析后：\n", content)

		//bs := ast.GoFmt(content)
		_ = util.GoFmt(content)

		//println("------------------------")
		//println("gofmt后：\n", string(bs))
		//println("=======================================================================================")
	}
}

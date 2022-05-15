package main

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/ast"
	"io/ioutil"
	"strings"
)

func main() {
	caseFile := parser.Parse("test/test_annotation/test_case/case.go")

	importDecl := ast.VisitImportDecl(lex.NewLexerWithCode("import (\n\"GoParser2/test/test_annotation/framework\"\n\"GoParser2/test/test_annotation/test_case\")"))
	functionDecl := ast.VisitFunctionDecl(lex.NewLexerWithCode("func init(){}"))

	packageName := caseFile.PackageClause().PackageName()

	genSourceFile := ast.NewSourceFile(caseFile.PackageClause(), []*ast.ImportDecl{importDecl}, []ast.IFunctionMethodDeclaration{functionDecl})

	for _, _fmd := range caseFile.Fmds() {
		switch fmd := _fmd.(type) {
		case *ast.FunctionDecl:
			annotation := fmd.Annotation()
			if annotation == nil {
				continue
			}
			kv := ""
			for k, v := range annotation.Value() {
				kv += "\"" + k + "\":\"" + v + "\","
			}
			statement := "framework.RegisterFunction(" + packageName + "." + fmd.Name() + "," +
				"framework.NewAnnotation(\"" + annotation.Name() + "\", " + " map[string]string{" + kv + "}" + "))"

			//println(statement)

			functionDecl.Block().AddStatement(ast.VisitStatement(lex.NewLexerWithCode(statement)))
		case *ast.MethodDecl:
			annotation := fmd.Annotation()
			if annotation == nil {
				continue
			}
			kv := ""
			for k, v := range annotation.Value() {
				kv += "\"" + k + "\":\"" + v + "\","
			}

			//varType := packageName + "." + fmd.Receiver().VarType()
			varType := fmd.Receiver().VarType()
			if strings.HasPrefix(varType, "*") {
				varType = "(*" + packageName + "." + strings.TrimPrefix(varType, "*") + ")"
			} else {
				varType = packageName + "." + varType
			}

			statement := "framework.RegisterMethod(" + varType + "." + fmd.Name() + "," +
				"framework.NewAnnotation(\"" + annotation.Name() + "\", " + " map[string]string{" + kv + "}" + "))"

			//println(statement)

			functionDecl.Block().AddStatement(ast.VisitStatement(lex.NewLexerWithCode(statement)))
		case *ast.TypeDecl:
			for _, ts := range fmd.TypeSpecs() {
				annotation := ts.Annotation()
				if annotation == nil {
					continue
				}
				kv := ""
				for k, v := range annotation.Value() {
					kv += "\"" + k + "\":\"" + v + "\","
				}
				statement := "framework.RegisterType(\"" + packageName + "." + ts.Name() + "\"," +
					"framework.NewAnnotation(\"" + annotation.Name() + "\", " + " map[string]string{" + kv + "}" + "))"

				//println(statement)

				functionDecl.Block().AddStatement(ast.VisitStatement(lex.NewLexerWithCode(statement)))
			}
		}
	}

	ioutil.WriteFile("test/test_annotation/test_case/gen/case.go", []byte(genSourceFile.String()), 0644)
}

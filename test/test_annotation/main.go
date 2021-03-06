package main

import (
	"github.com/naturemickey/GoParser2/test/test_annotation/framework"
	"github.com/naturemickey/GoParser2/test/test_annotation/test_case"
	_ "github.com/naturemickey/GoParser2/test/test_annotation/test_case/gen"
	"reflect"
)

func main() {
	a := &test_case.StructA{}

	fa := framework.GetFunctionAnnotation(test_case.FunctionA)
	println(fa.Name())
	println(fa.Value("k1"), fa.Value("k2"))

	ma := framework.GetMethodAnnotation((*test_case.StructA).MethodSG)
	println(ma.Name())

	ta := framework.GetTypeAnnotation(reflect.TypeOf(a))
	println(ta.Name())
}

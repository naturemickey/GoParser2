package test_case

import (
	"github.com/naturemickey/GoParser2/test/test_annotation/framework"
	"github.com/naturemickey/GoParser2/test/test_annotation/test_case"
)

func init() {
	framework.RegisterType("test_case.StructA", framework.NewAnnotation("IamAannotationName", map[string]string{}))
	framework.RegisterFunction(test_case.FunctionA, framework.NewAnnotation("Abc", map[string]string{"k1": "v1", "k2": "v2"}))
	framework.RegisterMethod((*test_case.StructA).MethodSG, framework.NewAnnotation("SG", map[string]string{}))
}

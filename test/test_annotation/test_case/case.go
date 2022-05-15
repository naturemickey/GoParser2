package test_case

type /*@IamAannotationName*/ StructA struct {
}

func /*@Abc(k1=v1,k2=v2)*/ FunctionA(a int) *StructA {
	return nil
}

func /*@SG*/ (this *StructA) MethodSG() int {
	return 0
}

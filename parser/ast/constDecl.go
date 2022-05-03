package ast

type ConstDecl struct {
}

func (c ConstDecl) __IFunctionMethodDeclaration__() {
	//TODO implement me
	panic("implement me")
}

func (c ConstDecl) __Declaration__() {
	//TODO implement me
	panic("implement me")
}

var _ Declaration = (*ConstDecl)(nil)

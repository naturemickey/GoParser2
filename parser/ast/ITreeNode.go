package ast

type ITreeNode interface {
	String() string
	CodeBuilder() *CodeBuilder
}

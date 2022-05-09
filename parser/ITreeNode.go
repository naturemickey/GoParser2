package parser

import "GoParser2/parser/util"

type ITreeNode interface {
	String() string
	CodeBuilder() *util.CodeBuilder
}

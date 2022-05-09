package ast

import "GoParser2/parser"

type IMethodspecOrTypename interface {
	parser.ITreeNode
	__IMethodspecOrTypename__()
}

package ast

import (
	"GoParser2/lex"
	"GoParser2/parser"
	"GoParser2/parser/util"
	"fmt"
)

type Receiver struct {
	// receiver: parameters;
	parameters *Parameters
}

func (a *Receiver) CodeBuilder() *util.CodeBuilder {
	return util.NewCB().AppendTreeNode(a.parameters)
}

func (a *Receiver) String() string {
	return a.CodeBuilder().String()
}

var _ parser.ITreeNode = (*Receiver)(nil)

func VisitReceiver(lexer *lex.Lexer) *Receiver {
	clone := lexer.Clone()

	parameters := VisitParameters(lexer)
	if parameters == nil {
		lexer.Recover(clone)
		return nil
	}
	onlyOneParams := true // todo 需要判断receiver中是否只有一个参数
	if !onlyOneParams {
		fmt.Printf("receiver只可以有1个。%s\n", lexer.LA().ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	return &Receiver{parameters: parameters}
}

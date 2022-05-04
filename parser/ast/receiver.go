package ast

import (
	"GoParser2/lex"
	"fmt"
)

type Receiver struct {
	// receiver: parameters;
	parameters *Parameters
}

func VisitReceiver(lexer *lex.Lexer) *Receiver {
	clone := lexer.Clone()
	la := lexer.LA()
	parameters := VisitParameters(lexer)
	if parameters == nil {
		return nil
	}
	onlyOneParams := true // todo 需要判断receiver中是否只有一个参数
	if !onlyOneParams {
		fmt.Printf("receiver只可以有1个。%s\n", la.ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	return &Receiver{parameters: parameters}
}

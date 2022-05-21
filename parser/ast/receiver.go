package ast

import (
	"fmt"
	"github.com/naturemickey/GoParser2/lex"
)

type Receiver struct {
	// receiver: parameters;
	parameters *Parameters
}

func (this *Receiver) SetParameters(parameters *Parameters) {
	this.parameters = parameters
}

func (this *Receiver) VarName() string {
	return this.parameters.parameterDecls[0].identifierList.identifiers[0].Literal()
}

func (this *Receiver) VarType() string {
	t := this.parameters.parameterDecls[0].type_.String()
	//return strings.TrimPrefix(t, "*")
	return t
}

func (a *Receiver) Parameters() *Parameters {
	return a.parameters
}

func (a *Receiver) CodeBuilder() *CodeBuilder {
	return NewCB().AppendTreeNode(a.parameters)
}

func (a *Receiver) String() string {
	return a.CodeBuilder().String()
}

var _ ITreeNode = (*Receiver)(nil)

func VisitReceiver(lexer *lex.Lexer) *Receiver {
	clone := lexer.Clone()

	parameters := VisitParameters(lexer)
	if parameters == nil {
		lexer.Recover(clone)
		return nil
	}
	onlyOneParams := true // todo 需要判断receiver中是否只有一个参数
	if !onlyOneParams {
		fmt.Printf("receiver,receiver只可以有1个。%s\n", lexer.LA().ErrorMsg())
		lexer.Recover(clone)
		return nil
	}
	return &Receiver{parameters: parameters}
}

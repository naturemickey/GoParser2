package ast

import (
	"github.com/naturemickey/GoParser2/lex"
	"reflect"
	"strings"
)

type CodeBuilder struct {
	code []string
}

func NewCB() *CodeBuilder {
	return new(CodeBuilder)
}

func (this *CodeBuilder) AppendString(str string) *CodeBuilder {
	this.code = append(this.code, str)
	return this
}

func (this *CodeBuilder) AppendToken(token *lex.Token) *CodeBuilder {
	if token == nil || reflect.ValueOf(token).IsNil() {
		return this
	}
	this.code = append(this.code, token.Literal())
	return this
}

func (this *CodeBuilder) AppendTreeNode(node ITreeNode) *CodeBuilder {
	if node == nil || reflect.ValueOf(node).IsNil() {
		return this
	}
	this.code = append(this.code, node.CodeBuilder().code...)
	return this
}

func (this *CodeBuilder) Newline() *CodeBuilder {
	this.code = append(this.code, "\n")
	return this
}

func (this *CodeBuilder) Tab() *CodeBuilder {
	this.code = append(this.code, "\t")
	return this
}

func (this *CodeBuilder) String() string {
	return strings.Join(this.code, " ")
}

func (this *CodeBuilder) StringCompact() string {
	return strings.Join(this.code, "")
}

func (this *CodeBuilder) popLast() *CodeBuilder {
	len := len(this.code)
	if len > 0 {
		this.code = this.code[:len-1]
	}
	return this
}

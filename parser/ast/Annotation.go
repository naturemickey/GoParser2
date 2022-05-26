package ast

import (
	"github.com/naturemickey/GoParser2/lex"
	"strings"
)

type Annotation struct {
	name  string
	value map[string]string
}

func (a *Annotation) SetName(name string) {
	a.name = name
}

func (a *Annotation) SetValue(value map[string]string) {
	a.value = value
}

func (a Annotation) Name() string {
	return a.name
}

func (a Annotation) Value() map[string]string {
	return a.value
}

func (a Annotation) String() string {
	return a.CodeBuilder().StringCompact()
}

func (a Annotation) CodeBuilder() *CodeBuilder {
	cb := &CodeBuilder{}
	cb.AppendString("/*@")
	cb.AppendString(a.name)
	if len(a.value) > 0 {
		cb.AppendString("(")
		for k, v := range a.value {
			cb.AppendString(k).AppendString("=").AppendString(v).AppendString(",")
		}
		cb.popLast()
		cb.AppendString(")")
	}
	cb.AppendString("*/")
	return cb
}

var _ ITreeNode = (*Annotation)(nil)

func VisitAnnotation(lexer *lex.Lexer) *Annotation {
	// 类似这样：/*@Bean(name=abc,cached=true)*/
	annotation := lexer.LA()
	if annotation.Type_() != lex.GoLexerANNOTATION_COMMENT {
		return nil
	}
	lexer.Pop() // annotation

	var name string
	var value = map[string]string{}
	var paramstr string

	literal := annotation.Literal()
	literal = strings.TrimPrefix(literal, "/*@")
	literal = strings.TrimSuffix(literal, "*/")
	lParenIdx := strings.Index(literal, "(")
	if lParenIdx < 0 {
		name = literal
	} else {
		name = literal[:lParenIdx]
		paramstr = literal[lParenIdx+1 : len(literal)-1]
	}
	if paramstr != "" {
		for _, kv := range strings.Split(paramstr, ",") {
			k_v := strings.Split(kv, "=")
			k := k_v[0]
			v := k_v[1]
			value[k] = v
		}
	}

	return &Annotation{name: name, value: value}
}

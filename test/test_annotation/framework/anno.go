package framework

import (
	"reflect"
	"strings"
)

var _func_map = map[string]*Annotation{}
var _method_map = map[string]*Annotation{}
var _type_map = map[string]*Annotation{}

func RegisterFunction(f interface{}, annotation *Annotation) {
	_func_map[reflect.TypeOf(f).String()] = annotation
}

func RegisterMethod(m interface{}, annotation *Annotation) {
	_method_map[reflect.TypeOf(m).String()] = annotation
}

func RegisterType(t string, annotation *Annotation) {
	_type_map[t] = annotation
}

func GetFunctionAnnotation(f interface{}) *Annotation {
	return _func_map[reflect.TypeOf(f).String()]
}

func GetTypeAnnotation(t reflect.Type) *Annotation {
	return _type_map[strings.TrimPrefix(t.String(), "*")]
}

func GetMethodAnnotation(m interface{}) *Annotation {
	return _method_map[reflect.TypeOf(m).String()]
}

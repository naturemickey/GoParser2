package ast

import "GoParser2/lex"

type AnnotationList struct {
	annotations []*Annotation
}

func (this *AnnotationList) String() string {
	return this.CodeBuilder().String()
}

func (this *AnnotationList) CodeBuilder() *CodeBuilder {
	cb := NewCB()
	for _, annotation := range this.annotations {
		cb.AppendTreeNode(annotation)
	}
	return cb
}

var _ ITreeNode = (*AnnotationList)(nil)

func VisitAnnotationList(lexer *lex.Lexer) *AnnotationList {
	var annotations []*Annotation
	for {
		annotation := VisitAnnotation(lexer)
		if annotation != nil {
			annotations = append(annotations, annotation)
		} else {
			break
		}
	}
	return &AnnotationList{annotations: annotations}
}

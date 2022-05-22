package ast

import "github.com/naturemickey/GoParser2/lex"

type AnnotationList struct {
	annotations []*Annotation
}

func (this *AnnotationList) Annotations() []*Annotation {
	return this.annotations
}

func (this *AnnotationList) SetAnnotations(annotations []*Annotation) {
	this.annotations = annotations
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
	if len(annotations) == 0 {
		return nil
	}
	return &AnnotationList{annotations: annotations}
}

package framework

type Annotation struct {
	name  string
	value map[string]string
}

func NewAnnotation(name string, value map[string]string) *Annotation {
	return &Annotation{name: name, value: value}
}

func (this *Annotation) Name() string {
	return this.name
}

func (this *Annotation) Value(key string) string {
	return this.value[key]
}

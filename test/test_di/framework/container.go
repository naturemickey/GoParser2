package framework

import "sync"

type container struct {
	beanMap sync.Map
}

func NewContainer() *container {
	return &container{beanMap: sync.Map{}}
}

func (this *container) RegisterBean(beanName string, bean interface{}) {
	this.beanMap.Store(beanName, bean)
}

func (this *container) GetBean(beanName string) (bean interface{}, ok bool) {
	return this.beanMap.Load(beanName)
}

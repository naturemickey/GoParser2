package service

import (
	"GoParser2/test/test_di/biz"
	"GoParser2/test/test_di/domain"
	"context"
)

type /*@Bean*/ XxxService struct {
	/*Inject*/ XxxBiz biz.IXxxBiz
}

func (this *XxxService) Save(ctx context.Context, xxx *domain.Xxx) {
	this.XxxBiz.Save(ctx, xxx)
}

func (this *XxxService) Query(ctx context.Context) {

}

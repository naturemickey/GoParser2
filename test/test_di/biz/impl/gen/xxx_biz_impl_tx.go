package gen

import (
	"GoParser2/test/test_di/biz"
	"GoParser2/test/test_di/domain"
	"GoParser2/test/test_di/framework"
	"context"
)

type XxxBizImplTx struct {
	/*Inject*/ XxxBiz biz.IXxxBiz
}

func (x *XxxBizImplTx) Save(ctx context.Context, xxx *domain.Xxx) {
	framework.DoTransaction(ctx, func(ctx context.Context) {
		x.XxxBiz.Save(ctx, xxx)
	})
}

func (x *XxxBizImplTx) Query(ctx context.Context) *domain.Xxx {
	return x.XxxBiz.Query(ctx)
}

var _ biz.IXxxBiz = (*XxxBizImplTx)(nil)

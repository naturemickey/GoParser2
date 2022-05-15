package impl

import (
	"GoParser2/test/test_di/biz"
	"GoParser2/test/test_di/dao"
	"GoParser2/test/test_di/domain"
	"context"
)

type /*@Bean*/ XxxBizImpl struct {
	/*@Inject*/ XxxDao dao.IXxxDao
}

func (x *XxxBizImpl) Save(ctx context.Context, xxx *domain.Xxx) {
	x.XxxDao.Save(ctx, xxx)
}

func (x *XxxBizImpl) Query(ctx context.Context) *domain.Xxx {
	return x.XxxDao.Query(ctx)
}

var _ biz.IXxxBiz = (*XxxBizImpl)(nil)

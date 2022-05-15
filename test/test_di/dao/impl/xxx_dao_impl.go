package impl

import (
	"GoParser2/test/test_di/dao"
	"GoParser2/test/test_di/domain"
	"context"
)

type /*@Bean*/ XxxDaoImpl struct {
}

func (x *XxxDaoImpl) Query(ctx context.Context) *domain.Xxx {
	// do something
	return nil
}

func (x *XxxDaoImpl) Save(ctx context.Context, xxx *domain.Xxx) {
	// do something
}

var _ dao.IXxxDao = (*XxxDaoImpl)(nil)

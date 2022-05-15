package biz

import (
	"GoParser2/test/test_di/domain"
	"context"
)

type IXxxBiz interface {
	Save /*@Transaction*/ (ctx context.Context, xxx *domain.Xxx)
	Query(ctx context.Context) *domain.Xxx
}

package dao

import (
	"GoParser2/test/test_di/domain"
	"context"
)

type IXxxDao interface {
	Save(ctx context.Context, xxx *domain.Xxx)
	Query(ctx context.Context) *domain.Xxx
}

package framework

import "context"

func DoTransaction(ctx context.Context, f func(context.Context)) {
	// do something
}

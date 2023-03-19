package pool

import (
	"context"
	"io"
	"time"
)

type Closer interface {
	io.Closer
}

type Pool[T Closer] interface {
	Put(ctx context.Context, t T) error
	Get(ctx context.Context) (T, error)
}

type Conn[T Closer] struct {
	t              T
	lastActiveTime time.Time // 记录当前连接最后一次放回连接池的时间，用来判断是否空闲连接
}

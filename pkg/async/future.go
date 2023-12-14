package async

import (
	"context"
)

type Future[T any] struct {
	l   TaskCancellation
	v   T
	err error
}

func (fut *Future[T]) Done() <-chan struct{} {
	return fut.l.Done()
}

func (fut *Future[T]) Value() (T, error) {
	Await(fut)
	return fut.v, fut.err
}

func (fut *Future[T]) ValueCtx(ctx context.Context) (T, error) {
	if err := AwaitCtx(ctx, fut); err != nil {
		var zero T
		return zero, err
	}
	return fut.v, fut.err
}

func ResolveFuture[T any](fut *Future[T], v T, err error) {
	Resolve(&fut.l, func() {
		fut.v = v
		fut.err = err
	})
}

// Package async provides asynchronous primitives and utilities.
package async

import (
	"context"
	"errors"
)

var ErrAlreadyResolved = errors.New("async: already resolved")

type Awaiter interface {
	Done() <-chan struct{}
}

func Await(a Awaiter) {
	<-a.Done()
}

func AwaitCtx(ctx context.Context, a Awaiter) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-a.Done():
		return nil
	}
}

type Resolver interface {
	resolve(func())
}

func Resolve(r Resolver, fn func()) {
	r.resolve(fn)
}

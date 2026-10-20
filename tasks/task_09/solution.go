package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](ctx context.Context, workers int, in []T, fn func(context.Context, T) (R, error)) ([]Result[R], error) {
	if workers < 1 {
		return nil, ErrInvalidWorkers
	}

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	if len(in) == 0 {
		return []Result[R]{}, nil
	}

	ch := make(chan int)
	answer := make([]Result[R], len(in))
	var wg sync.WaitGroup
	for range min(workers, len(in)) {
		wg.Go(func() {
			for index := range ch {
				if ctx.Err() != nil {
					return
				}
				res, err := fn(ctx, in[index])
				if err != nil {
					answer[index].Err = err
				} else {
					answer[index].Value = res
				}
			}
		})
	}

outer:
	for m := range len(in) {
		select {
		case ch <- m:
		case <-ctx.Done():
			break outer
		}
	}

	close(ch)
	wg.Wait()

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	return answer, nil
}

package ugo

import (
	"context"
	"time"
)

func RunWithTimeout(dur time.Duration, f func(ctx context.Context) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), dur)
	defer cancel()
	return f(ctx)
}

func Run(ctx context.Context, f func()) {
	quit := make(chan struct{})
	go func() {
		f()
		quit <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		return
	case <-quit:
		return
	}
}

func ForceTimeout(timeout time.Duration, f func() error) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan error)
	go func() {
		defer close(done)
		done <- f()
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

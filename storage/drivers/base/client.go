package base

import (
	"context"

	"github.com/gabesullice/phargo/storage"
	"github.com/pkg/errors"
)

type Doer struct{}

func (c *Doer) Do(ctx context.Context, r storage.Runner) (storage.Fetcher, error) {
	if err := ctx.Err(); err != nil {
		return nil, errors.Wrap(err, "Context already closed")
	}
	res := make(chan storage.Fetcher, 1)
	ch := run(res, r)
	select {
	case <-ctx.Done():
		return nil, errors.Wrap(ctx.Err(), "Context closed")
	case result := <-res:
		return result, nil
	case err := <-ch:
		return nil, errors.Wrap(err, "Transaction failed")
	}
}

func run(res chan storage.Fetcher, r storage.Runner) chan error {
	ch := make(chan error, 1)
	go func() {
		ch <- func() error {
			result, err := r.Run()
			if err == nil {
				res <- result
			}
			return err
		}()
	}()
	return ch
}

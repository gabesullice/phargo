package boltdb

import (
	"context"
	"errors"
	"testing"

	"github.com/gabesullice/phargo/storage"
)

func TestDo(t *testing.T) {
	t.Parallel()
	c := Client{}
	t.Run("Should call Run on the Runner", func(t *testing.T) {
		var called bool
		runner := storage.MockRunner(func() (storage.Fetcher, error) {
			called = true
			return result{}, nil
		})
		if c.Do(context.Background(), runner); !called {
			t.Errorf("Expected call")
		}
	})
	t.Run("Should not run a query if context.Done() channel is closed", func(t *testing.T) {
		var called bool
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		runner := storage.MockRunner(func() (storage.Fetcher, error) {
			called = true
			return result{}, nil
		})
		if c.Do(ctx, runner); called {
			t.Errorf("Expected nothing to be run")
		}
	})
	t.Run("Should return an error if the Runner fails", func(t *testing.T) {
		runner := storage.MockRunner(func() (storage.Fetcher, error) {
			return nil, errors.New("Expected")
		})
		if _, err := c.Do(context.Background(), runner); err == nil {
			t.Errorf("Expected error")
		}
	})
	t.Run("Should return a result on success", func(t *testing.T) {
		runner := storage.MockRunner(func() (storage.Fetcher, error) {
			return result{}, nil
		})
		if res, err := c.Do(context.Background(), runner); err != nil {
			t.Fail()
		} else if _, ok := res.(result); !ok {
			t.Errorf("Expected result. Got: %T", res)
		}
	})
}

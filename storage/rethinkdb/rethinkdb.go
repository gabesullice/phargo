package rethinkdb

import (
	"context"

	"github.com/gabesullice/phargo/storage"
)

type Client struct{}

type result struct{}

func (c Client) Do(ctx context.Context, q storage.QueryCompiler) (storage.Fetcher, error) {
	return result{}, nil
}

func (r result) Fetch(target interface{}) error      { return nil }
func (r result) FetchOne(target interface{}) error   { return nil }
func (r result) FetchAll(target []interface{}) error { return nil }

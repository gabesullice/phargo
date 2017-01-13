package boltdb

import (
	"context"

	b "github.com/boltdb/bolt"
	"github.com/gabesullice/phargo/storage"
	"github.com/pkg/errors"
)

// A Client can connect to and execute queries against BoltDB
type Client struct {
	db *b.DB
}

// A query that is prepared and ready to be run. Implements storage.Runner.
type preparedQuery struct {
	db       *b.DB
	dataType []byte
	term     term
}

// Wraps the BoltDB DB connection struct.
type db b.DB

// Internal representation of a storage.Query
type term struct{}

// Internal representation of a query result. Implements storage.Fetcher.
type result struct {
	cursor *b.Cursor
	close  chan struct{}
}

// Do takes a runnable query and executes it. A Client may have timeouts and
// other custom configuration defined, e.g. retry logic.
func (c Client) Do(ctx context.Context, r storage.Runner) (storage.Fetcher, error) {
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

// Prepare takes a standard query and turns it into a query that can be run
// against Client's connection to a BoltDB instance.
func (c Client) Prepare(q storage.Query) (storage.Runner, error) {
	return c.compile(q)
}

// Execute attempt to prepare a standard Query and then immediately executes it,
// returning a Fetchable response.
func (c Client) Execute(q storage.Query) (storage.Fetcher, error) {
	p, err := c.Prepare(q)
	if err != nil {
		return nil, err
	}
	return p.Run()
}

func (c Client) compile(q storage.Query) (storage.Runner, error) {
	return preparedQuery{db: c.db}, nil
}

func (q preparedQuery) Run() (storage.Fetcher, error) {
	var r result
	fn := func() func(tx *b.Tx) error {
		return q.buildTransaction(&r)
	}
	err := q.db.View(fn())
	return r, err
}

func (q preparedQuery) buildTransaction(r *result) func(tx *b.Tx) error {
	return func(tx *b.Tx) error {
		bkt := tx.Bucket([]byte("datatypes")).Bucket(q.dataType).Cursor()
		if bkt == nil {
			return errors.Errorf("Bucket %s.%s does not exist", "datatypes", q.dataType)
		}
		r.close = make(chan struct{})
		go func(tx *b.Tx) {
			<-r.close
		}(tx)
		return nil
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

func (r result) Fetch(target interface{}) (more bool, err error) { return }
func (r result) FetchOne(target interface{}) (err error)         { return }
func (r result) FetchAll(target []interface{}) (err error)       { return }
func (r result) Empty() bool                                     { return false }

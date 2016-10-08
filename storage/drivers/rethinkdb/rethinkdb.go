package rethinkdb

import (
	"context"

	"github.com/dancannon/gorethink"
	"github.com/gabesullice/phargo/storage"
)

// A Client can connect to and execute queries against a RethinkDB instance.
type Client struct {
	session *session
}

type preparedQuery struct {
	query   term
	session *session
}

type session gorethink.Session

type term gorethink.Term

type result struct{}

// Do takes a runnable query and executes it. A Client may have timeouts and
// other custom configuration defined, e.g. retry logic.
func (c Client) Do(ctx context.Context, q storage.Runner) (storage.Fetcher, error) {
	return q.Run()
}

// Prepare takes a standard query and turns it into a query that can be run
// against Client's connection to RethinkDB.
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
	return preparedQuery{term{}, c.session}, nil
}

func (q preparedQuery) Run() (storage.Fetcher, error) {
	return result{}, nil
}

func (r result) Fetch(target interface{}) (more bool, err error) { return }
func (r result) FetchOne(target interface{}) (err error)         { return }
func (r result) FetchAll(target []interface{}) (err error)       { return }
func (r result) Empty() bool                                     { return false }

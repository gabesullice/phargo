package storage

import "context"

// A Client is able to create and execute queries against a datastore.
type Client interface {
	Doer
	Preparer
	Executor
}

// A Doer reveives a context and a Query. A Doer should know how to run the
// query that it reveives. That is, it should know how to set any session
// and/or connection details prior to running the query.
type Doer interface {
	Do(ctx context.Context, r Runner) (Fetcher, error)
}

// A Preparer creates a runnable query.
type Preparer interface {
	Prepare(q Query) (Runner, error)
}

// An Executor prepares and does a query in one step.
type Executor interface {
	Execute(q Query) (Fetcher, error)
}

// Runner runs a compiled query.
type Runner interface {
	Run() (Fetcher, error)
}

// Fetcher can marshal one or many query results into an interface.
type Fetcher interface {
	Fetch(interface{}) (more bool, err error)
	FetchOne(interface{}) error
	FetchAll([]interface{}) error
	Empty() bool
}

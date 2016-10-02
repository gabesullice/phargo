package storage

import "context"

// A Client is able to create and execute queries against a datastore.
type Client interface {
	//NewQuery(datatype string) Query
	Doer
}

// Doer reveives a context and a Query. A Doer should know how to run the
// query that it reveives. That is, it should know how to set any session
// and/or connection details prior to running the query.
type Doer interface {
	Do(ctx context.Context, q QueryCompiler) (Fetcher, error)
}

// Runner runs a compiled query.
type Runner interface {
	Run() (Fetcher, error)
}

// Fetcher can marshal one or many query results into an interface.
type Fetcher interface {
	Fetch(interface{}) error
	FetchOne(interface{}) error
	FetchAll([]interface{}) error
}

// Query defines a set of conditions and pagination options used to fetch data
// from a datastore.
type Query interface {
	Condition(prop string, value interface{}, operator string) Query
	AddCondition(Condition) Query
	Exists(prop string, negate bool) Query
	ConditionGroup(conjunction string) Condition
	Conditions() []Condition
	Count() Query
	Range(offset, count int) Query
	QueryCompiler
}

// QueryCompiler can return a compiled query, ready to run.
type QueryCompiler interface {
	Compile() Runner
}

// Condition records an assertion to which fetched data must adhere. A
// condition may contain child conditions and a conjunction to create complex
// queries.
type Condition interface {
	Property() string
	Value() interface{}
	IsGroup() bool
	Conjunction() string
	Conditions() []Condition
}

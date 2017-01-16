package sqlite

import "database/sql"

type fetcher struct {
	rows *sql.Rows
}

func NewFetcher(rows *sql.Rows) *fetcher {
	return &fetcher{rows}
}

func (r *fetcher) Fetch(interface{}) (more bool, err error) {
	panic("not implemented")
}

func (r *fetcher) FetchOne(interface{}) error {
	panic("not implemented")
}

func (r *fetcher) FetchAll([]interface{}) error {
	panic("not implemented")
}

func (r *fetcher) Empty() bool {
	panic("not implemented")
}

package sqlite

import (
	"database/sql"

	"github.com/gabesullice/phargo/storage"
)

type runner struct {
	stmt *sql.Stmt
}

func NewRunner(stmt *sql.Stmt) *runner {
	r := new(runner)
	return r
}

func (r *runner) Run() (storage.Fetcher, error) {
	rows, err := r.stmt.Query()
	return NewFetcher(rows), err
}

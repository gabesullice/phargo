package sqlite

import (
	"database/sql"

	"github.com/gabesullice/phargo/storage"
	"github.com/gabesullice/phargo/storage/drivers/base"
	"github.com/pkg/errors"

	"gopkg.in/doug-martin/goqu.v3"

	_ "github.com/mattn/go-sqlite3"
	_ "gopkg.in/doug-martin/goqu.v3/adapters/sqlite3"
)

type DB interface {
	QueryPreparer
}

type QueryPreparer interface {
	Prepare(query string) (*sql.Stmt, error)
}

// A Client can connect to and execute queries against a sqlite database.
type Client struct {
	base.Doer
	conn DB
}

func (c *Client) Prepare(q storage.Query) (storage.Runner, error) {
	term := goqu.From(q.Type())

	if q.Count() {
		term = term.Select(goqu.COUNT("uuid"))
	} else {
		term = term.Select("uuid")
	}

	term, err := conditionize(term, q.Conditions())
	if err != nil {
		return nil, errors.Wrap(err, "Unable to apply condition to SQL query")
	}

	offset, limit := q.Range()
	term = term.Offset(offset)
	term = term.Limit(limit)

	query, _, err := term.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to compile SQL query string")
	}

	stmt, err := c.conn.Prepare(query)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to prepare SQL statement")
	}

	return NewRunner(stmt), nil
}

func (c *Client) Execute(q storage.Query) (storage.Fetcher, error) {
	panic("not implemented")
}

func conditionize(term *goqu.Dataset, cnds []storage.Condition) (*goqu.Dataset, error) {
	return term, nil
}

package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"testing"

	"github.com/gabesullice/phargo/storage/mocks"
	"github.com/stretchr/testify/mock"
)

func TestClient_Do(t *testing.T) {
	t.Parallel()
	c := Client{}
	t.Run("Should call Run on the Runner", func(t *testing.T) {
		r := new(mocks.Runner)
		r.On("Run").Return(nil, nil)
		c.Do(context.Background(), r)
		r.AssertCalled(t, "Run")
	})
	t.Run("Should not run a query if context.Done() channel is closed", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		r := new(mocks.Runner)
		r.On("Run").Return(nil, nil)
		c.Do(ctx, r)
		r.AssertNotCalled(t, "Run")
	})
	t.Run("Should return an error if the Runner fails", func(t *testing.T) {
		r := new(mocks.Runner)
		r.On("Run").Return(nil, errors.New("Expected"))
		if _, err := c.Do(context.Background(), r); err == nil {
			t.Errorf("Expected error")
		}
	})
	t.Run("Should return a result on success", func(t *testing.T) {
		r := new(mocks.Runner)
		r.On("Run").Return(&fetcher{}, nil)
		if res, err := c.Do(context.Background(), r); err != nil {
			t.Fail()
		} else if _, ok := res.(*fetcher); !ok {
			t.Errorf("Expected result. Got: %T", res)
		}
	})
}

func TestClient_Prepare(t *testing.T) {
	t.Parallel()
	t.Run("Should call Conditions, Range, and Count on a Query", func(t *testing.T) {
		c := Client{conn: NewDBMock()}
		q := NewQueryMock()
		_, _ = c.Prepare(q)

		q.AssertCalled(t, "Conditions")
		q.AssertCalled(t, "Range")
		q.AssertCalled(t, "Count")
	})
	t.Run("Should call Prepare on its DB connection", func(t *testing.T) {
		db := NewDBMock()
		c := Client{conn: db}
		q := NewQueryMock()
		_, _ = c.Prepare(q)

		db.AssertExpectations(t)
	})
	t.Run("Should build a COUNT query", func(t *testing.T) {
		q := new(mocks.Query)
		q.On("Conditions").Return(nil)
		q.On("Range").Return(uint(0), uint(10))
		q.On("Type").Return("user")

		db := NewDBMock()
		db.On("Prepare", mock.MatchedBy(func(s string) bool {
			return strings.Contains(s, "COUNT")
		}))
		c := Client{conn: db}
		q.On("Count").Return(true)

		_, _ = c.Prepare(q)

		db.AssertExpectations(t)
	})
	t.Run("Should build an OFFSET/LIMIT query", func(t *testing.T) {
		q := new(mocks.Query)
		q.On("Conditions").Return(nil)
		q.On("Type").Return("user")
		q.On("Count").Return(false)

		db := NewDBMock()
		db.On("Prepare", mock.MatchedBy(func(s string) bool {
			return strings.Contains(s, "OFFSET 10") && strings.Contains(s, "LIMIT 20")
		}))
		c := Client{conn: db}
		q.On("Range").Return(uint(10), uint(20))

		_, _ = c.Prepare(q)

		db.AssertExpectations(t)
	})
}

func NewQueryMock() *mocks.Query {
	q := new(mocks.Query)
	q.On("Conditions").Return(nil)
	q.On("Range").Return(uint(0), uint(10))
	q.On("Count").Return(true)
	q.On("Type").Return("user")
	return q
}

func NewDBMock() *MockDB {
	db := new(MockDB)
	db.On(
		"Prepare",
		mock.MatchedBy(func(s string) bool { return true }),
	).Return(&sql.Stmt{}, nil)
	return db
}

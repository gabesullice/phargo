package rethinkdb

import (
	"context"
	"testing"

	"github.com/gabesullice/phargo/storage"
)

func TestDo(t *testing.T) {
	t.Run("Do should compile and run a query.", func(t *testing.T) {
		var c storage.Client = Client{}
		var compiled, run bool
		q := MockQuery{compilerFunc(&compiled), runnerFunc(&run, false)}
		_, err := c.Do(context.Background(), q)
		if err != nil {
			t.Errorf("Expected success, got error: %v", err)
		}
		if !q.ran() {
			t.Error("Expected query to be run.")
		}
	})
}

type MockQuery struct {
	compileFunc func()
	runFunc     func() (shouldFail bool)
}

func (q MockQuery) Compile() storage.Runner {
	q.compileFunc()
	return q
}

func (q MockQuery) Run() (res storage.Fetcher, err error) {
	if q.runFunc() {
		err = errors.New("MockQuery.Run() failed.")
	}
	return result{}, err
}

func compilerFunc(compiled *bool) {
	return func() {
		compiled = true
	}
}

func runnerFunc(ran *bool, shouldFail bool) {
	return func() bool {
		ran = true
		return shouldFail
	}
}

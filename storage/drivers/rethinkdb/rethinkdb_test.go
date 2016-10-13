package rethinkdb

import (
	"context"
	"testing"

	"github.com/gabesullice/phargo/storage"
)

func TestDo(t *testing.T) {
	t.Run("Do should call run on a Runner", func(t *testing.T) {
		var c storage.Client = Client{}
		var run bool
		q := storage.MockRunner{func() bool {
			run = true
			return false
		}}
		_, err := c.Do(context.Background(), q)
		if err != nil {
			t.Errorf("Expected success, got error: %v", err)
		}
		if !run {
			t.Error("Expected query to be run.")
		}
	})
	// @todo test that the Do honors the passed in context
}

func TestPrepare(t *testing.T) {
	t.Run("Should compile a query to a gorethink.Term", func(t *testing.T) {
		var _ storage.Client = Client{}
		var _ storage.Query = storage.MockQuery{}
		// @todo implement a test that should see of preparedQuery is the
		// appropriate gorethink.Term
	})
}

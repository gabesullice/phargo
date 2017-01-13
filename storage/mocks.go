package storage

type MockRunner func() (Fetcher, error)

func (f MockRunner) Run() (Fetcher, error) {
	return f()
}

type MockQuery struct {
	ConditionsFunc func() []Condition
	RangeFunc      func() (offset, limit int)
	CountFunc      func() bool
}

func (m MockQuery) Conditions() []Condition {
	return m.ConditionsFunc()
}

func (m MockQuery) Range() (offset, limit int) {
	return m.RangeFunc()
}

func (m MockQuery) Count() bool {
	return m.CountFunc()
}

type MockFetcher struct{}

func (r MockFetcher) Fetch(target interface{}) (more bool, err error) { return }
func (r MockFetcher) FetchOne(target interface{}) (err error)         { return }
func (r MockFetcher) FetchAll(target []interface{}) (err error)       { return }
func (r MockFetcher) Empty() bool                                     { return false }

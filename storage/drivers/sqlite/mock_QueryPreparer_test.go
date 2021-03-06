package sqlite

import mock "github.com/stretchr/testify/mock"
import sql "database/sql"

// MockQueryPreparer is an autogenerated mock type for the QueryPreparer type
type MockQueryPreparer struct {
	mock.Mock
}

// Prepare provides a mock function with given fields: query
func (_m *MockQueryPreparer) Prepare(query string) (*sql.Stmt, error) {
	ret := _m.Called(query)

	var r0 *sql.Stmt
	if rf, ok := ret.Get(0).(func(string) *sql.Stmt); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Stmt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ QueryPreparer = (*MockQueryPreparer)(nil)

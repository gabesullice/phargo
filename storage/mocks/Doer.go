package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import storage "github.com/gabesullice/phargo/storage"

// Doer is an autogenerated mock type for the Doer type
type Doer struct {
	mock.Mock
}

// Do provides a mock function with given fields: ctx, r
func (_m *Doer) Do(ctx context.Context, r storage.Runner) (storage.Fetcher, error) {
	ret := _m.Called(ctx, r)

	var r0 storage.Fetcher
	if rf, ok := ret.Get(0).(func(context.Context, storage.Runner) storage.Fetcher); ok {
		r0 = rf(ctx, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.Fetcher)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, storage.Runner) error); ok {
		r1 = rf(ctx, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ storage.Doer = (*Doer)(nil)
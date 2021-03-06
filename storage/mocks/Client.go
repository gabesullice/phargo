package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import storage "github.com/gabesullice/phargo/storage"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Do provides a mock function with given fields: ctx, r
func (_m *Client) Do(ctx context.Context, r storage.Runner) (storage.Fetcher, error) {
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

// Execute provides a mock function with given fields: q
func (_m *Client) Execute(q storage.Query) (storage.Fetcher, error) {
	ret := _m.Called(q)

	var r0 storage.Fetcher
	if rf, ok := ret.Get(0).(func(storage.Query) storage.Fetcher); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.Fetcher)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(storage.Query) error); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Prepare provides a mock function with given fields: q
func (_m *Client) Prepare(q storage.Query) (storage.Runner, error) {
	ret := _m.Called(q)

	var r0 storage.Runner
	if rf, ok := ret.Get(0).(func(storage.Query) storage.Runner); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.Runner)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(storage.Query) error); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ storage.Client = (*Client)(nil)

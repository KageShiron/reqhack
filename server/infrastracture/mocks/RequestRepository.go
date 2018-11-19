// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "github.com/KageShiron/reqhack/server/domain"

import mock "github.com/stretchr/testify/mock"

// RequestRepository is an autogenerated mock type for the RequestRepository type
type RequestRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: r
func (_m *RequestRepository) Add(r *domain.Request) error {
	ret := _m.Called(r)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Request) error); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: binID, id
func (_m *RequestRepository) Get(binID int64, id int64) (*domain.Request, error) {
	ret := _m.Called(binID, id)

	var r0 *domain.Request
	if rf, ok := ret.Get(0).(func(int64, int64) *domain.Request); ok {
		r0 = rf(binID, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Request)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(binID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRange provides a mock function with given fields: binID, start, length
func (_m *RequestRepository) GetRange(binID int64, start int64, length int64) ([]*domain.Request, error) {
	ret := _m.Called(binID, start, length)

	var r0 []*domain.Request
	if rf, ok := ret.Get(0).(func(int64, int64, int64) []*domain.Request); ok {
		r0 = rf(binID, start, length)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Request)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64, int64) error); ok {
		r1 = rf(binID, start, length)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Length provides a mock function with given fields: binID
func (_m *RequestRepository) Length(binID int64) (int64, error) {
	ret := _m.Called(binID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(binID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(binID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
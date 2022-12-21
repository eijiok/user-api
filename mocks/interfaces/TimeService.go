// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// TimeService is an autogenerated mock type for the TimeService type
type TimeService struct {
	mock.Mock
}

// Now provides a mock function with given fields:
func (_m *TimeService) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

type mockConstructorTestingTNewTimeService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTimeService creates a new instance of TimeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTimeService(t mockConstructorTestingTNewTimeService) *TimeService {
	mock := &TimeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

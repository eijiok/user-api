// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// RequestResponseFunc is an autogenerated mock type for the RequestResponseFunc type
type RequestResponseFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: writer, request
func (_m *RequestResponseFunc) Execute(writer http.ResponseWriter, request *http.Request) {
	_m.Called(writer, request)
}

type mockConstructorTestingTNewRequestResponseFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequestResponseFunc creates a new instance of RequestResponseFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequestResponseFunc(t mockConstructorTestingTNewRequestResponseFunc) *RequestResponseFunc {
	mock := &RequestResponseFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

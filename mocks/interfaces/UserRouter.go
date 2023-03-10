// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	mux "github.com/gorilla/mux"
	mock "github.com/stretchr/testify/mock"
)

// UserRouter is an autogenerated mock type for the UserRouter type
type UserRouter struct {
	mock.Mock
}

// ConfigRoutes provides a mock function with given fields: apiRouter, pathPrefix, pathUserApi
func (_m *UserRouter) ConfigRoutes(apiRouter *mux.Router, pathPrefix string, pathUserApi string) {
	_m.Called(apiRouter, pathPrefix, pathUserApi)
}

type mockConstructorTestingTNewUserRouter interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRouter creates a new instance of UserRouter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRouter(t mockConstructorTestingTNewUserRouter) *UserRouter {
	mock := &UserRouter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

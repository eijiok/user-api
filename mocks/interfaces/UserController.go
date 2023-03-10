// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	errors "github.com/eijiok/user-api/errors"

	mock "github.com/stretchr/testify/mock"
)

// UserController is an autogenerated mock type for the UserController type
type UserController struct {
	mock.Mock
}

// Create provides a mock function with given fields: writer, request
func (_m *UserController) Create(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	ret := _m.Called(writer, request)

	var r0 *errors.HttpError
	if rf, ok := ret.Get(0).(func(http.ResponseWriter, *http.Request) *errors.HttpError); ok {
		r0 = rf(writer, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.HttpError)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: writer, request
func (_m *UserController) Delete(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	ret := _m.Called(writer, request)

	var r0 *errors.HttpError
	if rf, ok := ret.Get(0).(func(http.ResponseWriter, *http.Request) *errors.HttpError); ok {
		r0 = rf(writer, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.HttpError)
		}
	}

	return r0
}

// List provides a mock function with given fields: writer, request
func (_m *UserController) List(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	ret := _m.Called(writer, request)

	var r0 *errors.HttpError
	if rf, ok := ret.Get(0).(func(http.ResponseWriter, *http.Request) *errors.HttpError); ok {
		r0 = rf(writer, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.HttpError)
		}
	}

	return r0
}

// Read provides a mock function with given fields: writer, request
func (_m *UserController) Read(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	ret := _m.Called(writer, request)

	var r0 *errors.HttpError
	if rf, ok := ret.Get(0).(func(http.ResponseWriter, *http.Request) *errors.HttpError); ok {
		r0 = rf(writer, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.HttpError)
		}
	}

	return r0
}

// Update provides a mock function with given fields: writer, request
func (_m *UserController) Update(writer http.ResponseWriter, request *http.Request) *errors.HttpError {
	ret := _m.Called(writer, request)

	var r0 *errors.HttpError
	if rf, ok := ret.Get(0).(func(http.ResponseWriter, *http.Request) *errors.HttpError); ok {
		r0 = rf(writer, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.HttpError)
		}
	}

	return r0
}

type mockConstructorTestingTNewUserController interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserController creates a new instance of UserController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserController(t mockConstructorTestingTNewUserController) *UserController {
	mock := &UserController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

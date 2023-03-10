// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/eijiok/user-api/common"

	mock "github.com/stretchr/testify/mock"

	model "github.com/eijiok/user-api/model"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, filter
func (_m *UserRepository) Count(ctx context.Context, filter *common.UserFilter) (int64, error) {
	ret := _m.Called(ctx, filter)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *common.UserFilter) int64); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *common.UserFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *UserRepository) Delete(ctx context.Context, id *primitive.ObjectID) (int64, error) {
	ret := _m.Called(ctx, id)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *primitive.ObjectID) int64); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *primitive.ObjectID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *UserRepository) GetById(ctx context.Context, id *primitive.ObjectID) (*model.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, *primitive.ObjectID) *model.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *primitive.ObjectID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0
func (_m *UserRepository) List(_a0 context.Context) ([]model.User, error) {
	ret := _m.Called(_a0)

	var r0 []model.User
	if rf, ok := ret.Get(0).(func(context.Context) []model.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, user
func (_m *UserRepository) Save(ctx context.Context, user *model.User) (*primitive.ObjectID, error) {
	ret := _m.Called(ctx, user)

	var r0 *primitive.ObjectID
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) *primitive.ObjectID); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*primitive.ObjectID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, user
func (_m *UserRepository) Update(ctx context.Context, user *model.User) (int64, error) {
	ret := _m.Called(ctx, user)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) int64); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	twitter "github.com/kaungmyathan22/golang-graphql-social-network"
	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *UserService) GetByID(ctx context.Context, id string) (twitter.User, error) {
	ret := _m.Called(ctx, id)

	var r0 twitter.User
	if rf, ok := ret.Get(0).(func(context.Context, string) twitter.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(twitter.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
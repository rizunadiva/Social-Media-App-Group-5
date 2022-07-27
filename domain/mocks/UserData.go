// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "socialmedia-app/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserData is an autogenerated mock type for the UserData type
type UserData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: userID
func (_m *UserData) Delete(userID int) (int, error) {
	ret := _m.Called(userID)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSpecific provides a mock function with given fields: userID
func (_m *UserData) GetSpecific(userID int) (domain.User, error) {
	ret := _m.Called(userID)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(int) domain.User); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newUser
func (_m *UserData) Insert(newUser domain.User) (domain.User, error) {
	ret := _m.Called(newUser)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: userLogin
func (_m *UserData) Login(userLogin domain.User) (int, domain.User, error) {
	ret := _m.Called(userLogin)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.User) int); ok {
		r0 = rf(userLogin)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 domain.User
	if rf, ok := ret.Get(1).(func(domain.User) domain.User); ok {
		r1 = rf(userLogin)
	} else {
		r1 = ret.Get(1).(domain.User)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(domain.User) error); ok {
		r2 = rf(userLogin)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: userID, updatedData
func (_m *UserData) Update(userID int, updatedData domain.User) domain.User {
	ret := _m.Called(userID, updatedData)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(int, domain.User) domain.User); ok {
		r0 = rf(userID, updatedData)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	return r0
}

type mockConstructorTestingTNewUserData interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserData creates a new instance of UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserData(t mockConstructorTestingTNewUserData) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

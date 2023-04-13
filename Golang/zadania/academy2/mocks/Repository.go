// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	academy "github.com/grupawp/akademia-programowania/Golang/zadania/academy2"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Get provides a mock function with given fields: name
func (_m *Repository) Get(name string) (academy.Student, error) {
	ret := _m.Called(name)

	var r0 academy.Student
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (academy.Student, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) academy.Student); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(academy.Student)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Graduate provides a mock function with given fields: name
func (_m *Repository) Graduate(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: year
func (_m *Repository) List(year uint8) ([]string, error) {
	ret := _m.Called(year)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(uint8) ([]string, error)); ok {
		return rf(year)
	}
	if rf, ok := ret.Get(0).(func(uint8) []string); ok {
		r0 = rf(year)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(uint8) error); ok {
		r1 = rf(year)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: name, year
func (_m *Repository) Save(name string, year uint8) error {
	ret := _m.Called(name, year)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint8) error); ok {
		r0 = rf(name, year)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
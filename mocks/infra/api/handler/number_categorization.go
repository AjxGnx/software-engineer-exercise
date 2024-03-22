// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// NumberCategorization is an autogenerated mock type for the NumberCategorization type
type NumberCategorization struct {
	mock.Mock
}

// Categorize provides a mock function with given fields: context
func (_m *NumberCategorization) Categorize(context echo.Context) error {
	ret := _m.Called(context)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(context)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: context
func (_m *NumberCategorization) Get(context echo.Context) error {
	ret := _m.Called(context)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(context)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByNumber provides a mock function with given fields: context
func (_m *NumberCategorization) GetByNumber(context echo.Context) error {
	ret := _m.Called(context)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(context)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewNumberCategorization interface {
	mock.TestingT
	Cleanup(func())
}

// NewNumberCategorization creates a new instance of NumberCategorization. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNumberCategorization(t mockConstructorTestingTNewNumberCategorization) *NumberCategorization {
	mock := &NumberCategorization{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

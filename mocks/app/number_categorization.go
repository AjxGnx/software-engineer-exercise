// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	dto "github.com/AjxGnx/software-engineer-exercise/internal/domain/dto"
	entity "github.com/AjxGnx/software-engineer-exercise/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// NumberCategorization is an autogenerated mock type for the NumberCategorization type
type NumberCategorization struct {
	mock.Mock
}

// Categorize provides a mock function with given fields: number
func (_m *NumberCategorization) Categorize(number int64) (entity.Categorization, error) {
	ret := _m.Called(number)

	var r0 entity.Categorization
	if rf, ok := ret.Get(0).(func(int64) entity.Categorization); ok {
		r0 = rf(number)
	} else {
		r0 = ret.Get(0).(entity.Categorization)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: paginate
func (_m *NumberCategorization) Get(paginate dto.Paginate) (*entity.Paginator, error) {
	ret := _m.Called(paginate)

	var r0 *entity.Paginator
	if rf, ok := ret.Get(0).(func(dto.Paginate) *entity.Paginator); ok {
		r0 = rf(paginate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Paginator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.Paginate) error); ok {
		r1 = rf(paginate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByNumber provides a mock function with given fields: number
func (_m *NumberCategorization) GetByNumber(number int64) (entity.Categorization, error) {
	ret := _m.Called(number)

	var r0 entity.Categorization
	if rf, ok := ret.Get(0).(func(int64) entity.Categorization); ok {
		r0 = rf(number)
	} else {
		r0 = ret.Get(0).(entity.Categorization)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
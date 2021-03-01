// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/odpf/guardian/domain"
	mock "github.com/stretchr/testify/mock"
)

// PolicyRepository is an autogenerated mock type for the PolicyRepository type
type PolicyRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *PolicyRepository) Create(_a0 *domain.Policy) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Policy) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields:
func (_m *PolicyRepository) Find() ([]*domain.Policy, error) {
	ret := _m.Called()

	var r0 []*domain.Policy
	if rf, ok := ret.Get(0).(func() []*domain.Policy); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Policy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOne provides a mock function with given fields: id, version
func (_m *PolicyRepository) GetOne(id string, version int) (*domain.Policy, error) {
	ret := _m.Called(id, version)

	var r0 *domain.Policy
	if rf, ok := ret.Get(0).(func(string, int) *domain.Policy); ok {
		r0 = rf(id, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Policy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(id, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

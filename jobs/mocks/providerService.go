// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/goto/guardian/domain"

	mock "github.com/stretchr/testify/mock"
)

// ProviderService is an autogenerated mock type for the providerService type
type ProviderService struct {
	mock.Mock
}

// FetchResources provides a mock function with given fields: _a0
func (_m *ProviderService) FetchResources(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: _a0
func (_m *ProviderService) Find(_a0 context.Context) ([]*domain.Provider, error) {
	ret := _m.Called(_a0)

	var r0 []*domain.Provider
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*domain.Provider, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Provider); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Provider)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProviderService interface {
	mock.TestingT
	Cleanup(func())
}

// NewProviderService creates a new instance of ProviderService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProviderService(t mockConstructorTestingTNewProviderService) *ProviderService {
	mock := &ProviderService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

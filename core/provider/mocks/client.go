// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/odpf/guardian/domain"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateConfig provides a mock function with given fields: _a0
func (_m *Client) CreateConfig(_a0 *domain.ProviderConfig) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.ProviderConfig) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAccountTypes provides a mock function with given fields:
func (_m *Client) GetAccountTypes() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetPermissions provides a mock function with given fields: p, resourceType, role
func (_m *Client) GetPermissions(p *domain.ProviderConfig, resourceType string, role string) ([]interface{}, error) {
	ret := _m.Called(p, resourceType, role)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(*domain.ProviderConfig, string, string) []interface{}); ok {
		r0 = rf(p, resourceType, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.ProviderConfig, string, string) error); ok {
		r1 = rf(p, resourceType, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResources provides a mock function with given fields: pc
func (_m *Client) GetResources(pc *domain.ProviderConfig) ([]*domain.Resource, error) {
	ret := _m.Called(pc)

	var r0 []*domain.Resource
	if rf, ok := ret.Get(0).(func(*domain.ProviderConfig) []*domain.Resource); ok {
		r0 = rf(pc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.ProviderConfig) error); ok {
		r1 = rf(pc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoles provides a mock function with given fields: pc, resourceType
func (_m *Client) GetRoles(pc *domain.ProviderConfig, resourceType string) ([]*domain.Role, error) {
	ret := _m.Called(pc, resourceType)

	var r0 []*domain.Role
	if rf, ok := ret.Get(0).(func(*domain.ProviderConfig, string) []*domain.Role); ok {
		r0 = rf(pc, resourceType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.ProviderConfig, string) error); ok {
		r1 = rf(pc, resourceType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetType provides a mock function with given fields:
func (_m *Client) GetType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GrantAccess provides a mock function with given fields: _a0, _a1
func (_m *Client) GrantAccess(_a0 *domain.ProviderConfig, _a1 domain.Access) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.ProviderConfig, domain.Access) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RevokeAccess provides a mock function with given fields: _a0, _a1
func (_m *Client) RevokeAccess(_a0 *domain.ProviderConfig, _a1 domain.Access) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.ProviderConfig, domain.Access) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

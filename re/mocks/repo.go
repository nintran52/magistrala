// Code generated by mockery v2.53.3. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	re "github.com/absmach/magistrala/re"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddRule provides a mock function with given fields: ctx, r
func (_m *Repository) AddRule(ctx context.Context, r re.Rule) (re.Rule, error) {
	ret := _m.Called(ctx, r)

	if len(ret) == 0 {
		panic("no return value specified for AddRule")
	}

	var r0 re.Rule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, re.Rule) (re.Rule, error)); ok {
		return rf(ctx, r)
	}
	if rf, ok := ret.Get(0).(func(context.Context, re.Rule) re.Rule); ok {
		r0 = rf(ctx, r)
	} else {
		r0 = ret.Get(0).(re.Rule)
	}

	if rf, ok := ret.Get(1).(func(context.Context, re.Rule) error); ok {
		r1 = rf(ctx, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRules provides a mock function with given fields: ctx, pm
func (_m *Repository) ListRules(ctx context.Context, pm re.PageMeta) (re.Page, error) {
	ret := _m.Called(ctx, pm)

	if len(ret) == 0 {
		panic("no return value specified for ListRules")
	}

	var r0 re.Page
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, re.PageMeta) (re.Page, error)); ok {
		return rf(ctx, pm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, re.PageMeta) re.Page); ok {
		r0 = rf(ctx, pm)
	} else {
		r0 = ret.Get(0).(re.Page)
	}

	if rf, ok := ret.Get(1).(func(context.Context, re.PageMeta) error); ok {
		r1 = rf(ctx, pm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveRule provides a mock function with given fields: ctx, id
func (_m *Repository) RemoveRule(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for RemoveRule")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRule provides a mock function with given fields: ctx, r
func (_m *Repository) UpdateRule(ctx context.Context, r re.Rule) (re.Rule, error) {
	ret := _m.Called(ctx, r)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRule")
	}

	var r0 re.Rule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, re.Rule) (re.Rule, error)); ok {
		return rf(ctx, r)
	}
	if rf, ok := ret.Get(0).(func(context.Context, re.Rule) re.Rule); ok {
		r0 = rf(ctx, r)
	} else {
		r0 = ret.Get(0).(re.Rule)
	}

	if rf, ok := ret.Get(1).(func(context.Context, re.Rule) error); ok {
		r1 = rf(ctx, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRuleStatus provides a mock function with given fields: ctx, id, status
func (_m *Repository) UpdateRuleStatus(ctx context.Context, id string, status re.Status) (re.Rule, error) {
	ret := _m.Called(ctx, id, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRuleStatus")
	}

	var r0 re.Rule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, re.Status) (re.Rule, error)); ok {
		return rf(ctx, id, status)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, re.Status) re.Rule); ok {
		r0 = rf(ctx, id, status)
	} else {
		r0 = ret.Get(0).(re.Rule)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, re.Status) error); ok {
		r1 = rf(ctx, id, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ViewRule provides a mock function with given fields: ctx, id
func (_m *Repository) ViewRule(ctx context.Context, id string) (re.Rule, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for ViewRule")
	}

	var r0 re.Rule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (re.Rule, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) re.Rule); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(re.Rule)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

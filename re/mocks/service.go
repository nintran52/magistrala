// Code generated by mockery v2.53.3. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	authn "github.com/absmach/supermq/pkg/authn"

	mock "github.com/stretchr/testify/mock"

	re "github.com/absmach/magistrala/re"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddRule provides a mock function with given fields: ctx, session, r
func (_m *Service) AddRule(ctx context.Context, session authn.Session, r re.Rule) (re.Rule, error) {
	ret := _m.Called(ctx, session, r)

	if len(ret) == 0 {
		panic("no return value specified for AddRule")
	}

	var r0 re.Rule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, re.Rule) (re.Rule, error)); ok {
		return rf(ctx, session, r)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, re.Rule) re.Rule); ok {
		r0 = rf(ctx, session, r)
	} else {
		r0 = ret.Get(0).(re.Rule)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, re.Rule) error); ok {
		r1 = rf(ctx, session, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConsumeAsync provides a mock function with given fields: ctx, messages
func (_m *Service) ConsumeAsync(ctx context.Context, messages interface{}) {
	_m.Called(ctx, messages)
}

// DisableRule provides a mock function with given fields: ctx, session, id
func (_m *Service) DisableRule(ctx context.Context, session authn.Session, id string) (re.Rule, error) {
	ret := _m.Called(ctx, session, id)

	if len(ret) == 0 {
		panic("no return value specified for DisableRule")
	}

	var r0 re.Rule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) (re.Rule, error)); ok {
		return rf(ctx, session, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) re.Rule); ok {
		r0 = rf(ctx, session, id)
	} else {
		r0 = ret.Get(0).(re.Rule)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string) error); ok {
		r1 = rf(ctx, session, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableRule provides a mock function with given fields: ctx, session, id
func (_m *Service) EnableRule(ctx context.Context, session authn.Session, id string) (re.Rule, error) {
	ret := _m.Called(ctx, session, id)

	if len(ret) == 0 {
		panic("no return value specified for EnableRule")
	}

	var r0 re.Rule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) (re.Rule, error)); ok {
		return rf(ctx, session, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) re.Rule); ok {
		r0 = rf(ctx, session, id)
	} else {
		r0 = ret.Get(0).(re.Rule)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string) error); ok {
		r1 = rf(ctx, session, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Errors provides a mock function with no fields
func (_m *Service) Errors() <-chan error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Errors")
	}

	var r0 <-chan error
	if rf, ok := ret.Get(0).(func() <-chan error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan error)
		}
	}

	return r0
}

// ListRules provides a mock function with given fields: ctx, session, pm
func (_m *Service) ListRules(ctx context.Context, session authn.Session, pm re.PageMeta) (re.Page, error) {
	ret := _m.Called(ctx, session, pm)

	if len(ret) == 0 {
		panic("no return value specified for ListRules")
	}

	var r0 re.Page
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, re.PageMeta) (re.Page, error)); ok {
		return rf(ctx, session, pm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, re.PageMeta) re.Page); ok {
		r0 = rf(ctx, session, pm)
	} else {
		r0 = ret.Get(0).(re.Page)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, re.PageMeta) error); ok {
		r1 = rf(ctx, session, pm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveRule provides a mock function with given fields: ctx, session, id
func (_m *Service) RemoveRule(ctx context.Context, session authn.Session, id string) error {
	ret := _m.Called(ctx, session, id)

	if len(ret) == 0 {
		panic("no return value specified for RemoveRule")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) error); ok {
		r0 = rf(ctx, session, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartScheduler provides a mock function with given fields: ctx
func (_m *Service) StartScheduler(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for StartScheduler")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRule provides a mock function with given fields: ctx, session, r
func (_m *Service) UpdateRule(ctx context.Context, session authn.Session, r re.Rule) (re.Rule, error) {
	ret := _m.Called(ctx, session, r)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRule")
	}

	var r0 re.Rule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, re.Rule) (re.Rule, error)); ok {
		return rf(ctx, session, r)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, re.Rule) re.Rule); ok {
		r0 = rf(ctx, session, r)
	} else {
		r0 = ret.Get(0).(re.Rule)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, re.Rule) error); ok {
		r1 = rf(ctx, session, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ViewRule provides a mock function with given fields: ctx, session, id
func (_m *Service) ViewRule(ctx context.Context, session authn.Session, id string) (re.Rule, error) {
	ret := _m.Called(ctx, session, id)

	if len(ret) == 0 {
		panic("no return value specified for ViewRule")
	}

	var r0 re.Rule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) (re.Rule, error)); ok {
		return rf(ctx, session, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authn.Session, string) re.Rule); ok {
		r0 = rf(ctx, session, id)
	} else {
		r0 = ret.Get(0).(re.Rule)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authn.Session, string) error); ok {
		r1 = rf(ctx, session, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

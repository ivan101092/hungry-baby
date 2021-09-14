// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	user "hungry-baby/businesses/user"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, userDomain
func (_m *Repository) Delete(ctx context.Context, userDomain *user.Domain) (user.Domain, error) {
	ret := _m.Called(ctx, userDomain)

	var r0 user.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *user.Domain) user.Domain); ok {
		r0 = rf(ctx, userDomain)
	} else {
		r0 = ret.Get(0).(user.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.Domain) error); ok {
		r1 = rf(ctx, userDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, search, status, page, perpage
func (_m *Repository) Find(ctx context.Context, search string, status string, page int, perpage int) ([]user.Domain, int, error) {
	ret := _m.Called(ctx, search, status, page, perpage)

	var r0 []user.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, int) []user.Domain); ok {
		r0 = rf(ctx, search, status, page, perpage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.Domain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, int) int); ok {
		r1 = rf(ctx, search, status, page, perpage)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, int, int) error); ok {
		r2 = rf(ctx, search, status, page, perpage)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindAll provides a mock function with given fields: ctx, search, status
func (_m *Repository) FindAll(ctx context.Context, search string, status string) ([]user.Domain, error) {
	ret := _m.Called(ctx, search, status)

	var r0 []user.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []user.Domain); ok {
		r0 = rf(ctx, search, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, search, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByCode provides a mock function with given fields: ctx, code, status
func (_m *Repository) FindByCode(ctx context.Context, code string, status string) (user.Domain, error) {
	ret := _m.Called(ctx, code, status)

	var r0 user.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, string) user.Domain); ok {
		r0 = rf(ctx, code, status)
	} else {
		r0 = ret.Get(0).(user.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, code, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEmail provides a mock function with given fields: ctx, email, status
func (_m *Repository) FindByEmail(ctx context.Context, email string, status string) (user.Domain, error) {
	ret := _m.Called(ctx, email, status)

	var r0 user.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, string) user.Domain); ok {
		r0 = rf(ctx, email, status)
	} else {
		r0 = ret.Get(0).(user.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, id, status
func (_m *Repository) FindByID(ctx context.Context, id int, status string) (user.Domain, error) {
	ret := _m.Called(ctx, id, status)

	var r0 user.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int, string) user.Domain); ok {
		r0 = rf(ctx, id, status)
	} else {
		r0 = ret.Get(0).(user.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, string) error); ok {
		r1 = rf(ctx, id, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, userDomain
func (_m *Repository) Store(ctx context.Context, userDomain *user.Domain) (user.Domain, error) {
	ret := _m.Called(ctx, userDomain)

	var r0 user.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *user.Domain) user.Domain); ok {
		r0 = rf(ctx, userDomain)
	} else {
		r0 = ret.Get(0).(user.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.Domain) error); ok {
		r1 = rf(ctx, userDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, userDomain
func (_m *Repository) Update(ctx context.Context, userDomain *user.Domain) (user.Domain, error) {
	ret := _m.Called(ctx, userDomain)

	var r0 user.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *user.Domain) user.Domain); ok {
		r0 = rf(ctx, userDomain)
	} else {
		r0 = ret.Get(0).(user.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.Domain) error); ok {
		r1 = rf(ctx, userDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

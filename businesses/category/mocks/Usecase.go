// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"
	category "hungry-baby/businesses/category"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx
func (_m *Usecase) GetAll(ctx context.Context) ([]category.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []category.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []category.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]category.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByActive provides a mock function with given fields: ctx, active
func (_m *Usecase) GetByActive(ctx context.Context, active bool) ([]category.Domain, error) {
	ret := _m.Called(ctx, active)

	var r0 []category.Domain
	if rf, ok := ret.Get(0).(func(context.Context, bool) []category.Domain); ok {
		r0 = rf(ctx, active)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]category.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, bool) error); ok {
		r1 = rf(ctx, active)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Usecase) GetByID(ctx context.Context, id int) (category.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 category.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) category.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(category.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
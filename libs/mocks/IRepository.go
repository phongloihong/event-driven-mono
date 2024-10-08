// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	mongo "go.mongodb.org/mongo-driver/mongo"

	options "go.mongodb.org/mongo-driver/mongo/options"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// IRepository is an autogenerated mock type for the IRepository type
type IRepository[T interface{}] struct {
	mock.Mock
}

// DeleteOne provides a mock function with given fields: ctx, filter, opts
func (_m *IRepository[T]) DeleteOne(ctx context.Context, filter primitive.M, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.DeleteResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.M, ...*options.DeleteOptions) (*mongo.DeleteResult, error)); ok {
		return rf(ctx, filter, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.M, ...*options.DeleteOptions) *mongo.DeleteResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.DeleteResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.M, ...*options.DeleteOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, filter, opts
func (_m *IRepository[T]) Find(ctx context.Context, filter primitive.M, opts ...*options.FindOptions) ([]T, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []T
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.M, ...*options.FindOptions) ([]T, error)); ok {
		return rf(ctx, filter, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.M, ...*options.FindOptions) []T); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]T)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.M, ...*options.FindOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: ctx, filter, opts
func (_m *IRepository[T]) FindOne(ctx context.Context, filter primitive.M, opts ...*options.FindOneOptions) (*T, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *T
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.M, ...*options.FindOneOptions) (*T, error)); ok {
		return rf(ctx, filter, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.M, ...*options.FindOneOptions) *T); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.M, ...*options.FindOneOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertMany provides a mock function with given fields: ctx, data, opts
func (_m *IRepository[T]) InsertMany(ctx context.Context, data []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, data)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.InsertManyResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []interface{}, ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)); ok {
		return rf(ctx, data, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []interface{}, ...*options.InsertManyOptions) *mongo.InsertManyResult); ok {
		r0 = rf(ctx, data, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.InsertManyResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []interface{}, ...*options.InsertManyOptions) error); ok {
		r1 = rf(ctx, data, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertOne provides a mock function with given fields: ctx, data, opts
func (_m *IRepository[T]) InsertOne(ctx context.Context, data interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, data)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.InsertOneResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)); ok {
		return rf(ctx, data, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.InsertOneOptions) *mongo.InsertOneResult); ok {
		r0 = rf(ctx, data, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.InsertOneResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.InsertOneOptions) error); ok {
		r1 = rf(ctx, data, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMany provides a mock function with given fields: ctx, filter, data, opts
func (_m *IRepository[T]) UpdateMany(ctx context.Context, filter primitive.M, data interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, data)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.UpdateResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.M, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)); ok {
		return rf(ctx, filter, data, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.M, interface{}, ...*options.UpdateOptions) *mongo.UpdateResult); ok {
		r0 = rf(ctx, filter, data, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.UpdateResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.M, interface{}, ...*options.UpdateOptions) error); ok {
		r1 = rf(ctx, filter, data, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOne provides a mock function with given fields: ctx, filter, data, opts
func (_m *IRepository[T]) UpdateOne(ctx context.Context, filter primitive.M, data interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, data)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo.UpdateResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.M, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)); ok {
		return rf(ctx, filter, data, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.M, interface{}, ...*options.UpdateOptions) *mongo.UpdateResult); ok {
		r0 = rf(ctx, filter, data, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo.UpdateResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.M, interface{}, ...*options.UpdateOptions) error); ok {
		r1 = rf(ctx, filter, data, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIRepository creates a new instance of IRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRepository[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *IRepository[T] {
	mock := &IRepository[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

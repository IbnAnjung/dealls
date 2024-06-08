// Code generated by mockery v2.43.2. DO NOT EDIT.

package cache

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockCacheService is an autogenerated mock type for the CacheService type
type MockCacheService struct {
	mock.Mock
}

type MockCacheService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCacheService) EXPECT() *MockCacheService_Expecter {
	return &MockCacheService_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, key, val
func (_m *MockCacheService) Get(ctx context.Context, key string, val interface{}) error {
	ret := _m.Called(ctx, key, val)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, key, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCacheService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockCacheService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - val interface{}
func (_e *MockCacheService_Expecter) Get(ctx interface{}, key interface{}, val interface{}) *MockCacheService_Get_Call {
	return &MockCacheService_Get_Call{Call: _e.mock.On("Get", ctx, key, val)}
}

func (_c *MockCacheService_Get_Call) Run(run func(ctx context.Context, key string, val interface{})) *MockCacheService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *MockCacheService_Get_Call) Return(err error) *MockCacheService_Get_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockCacheService_Get_Call) RunAndReturn(run func(context.Context, string, interface{}) error) *MockCacheService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAndLock provides a mock function with given fields: ctx, key, val
func (_m *MockCacheService) GetAndLock(ctx context.Context, key string, val interface{}) error {
	ret := _m.Called(ctx, key, val)

	if len(ret) == 0 {
		panic("no return value specified for GetAndLock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, key, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCacheService_GetAndLock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAndLock'
type MockCacheService_GetAndLock_Call struct {
	*mock.Call
}

// GetAndLock is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - val interface{}
func (_e *MockCacheService_Expecter) GetAndLock(ctx interface{}, key interface{}, val interface{}) *MockCacheService_GetAndLock_Call {
	return &MockCacheService_GetAndLock_Call{Call: _e.mock.On("GetAndLock", ctx, key, val)}
}

func (_c *MockCacheService_GetAndLock_Call) Run(run func(ctx context.Context, key string, val interface{})) *MockCacheService_GetAndLock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *MockCacheService_GetAndLock_Call) Return(err error) *MockCacheService_GetAndLock_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockCacheService_GetAndLock_Call) RunAndReturn(run func(context.Context, string, interface{}) error) *MockCacheService_GetAndLock_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, key, val, expTime
func (_m *MockCacheService) Set(ctx context.Context, key string, val interface{}, expTime time.Duration) error {
	ret := _m.Called(ctx, key, val, expTime)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = rf(ctx, key, val, expTime)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCacheService_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockCacheService_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - val interface{}
//   - expTime time.Duration
func (_e *MockCacheService_Expecter) Set(ctx interface{}, key interface{}, val interface{}, expTime interface{}) *MockCacheService_Set_Call {
	return &MockCacheService_Set_Call{Call: _e.mock.On("Set", ctx, key, val, expTime)}
}

func (_c *MockCacheService_Set_Call) Run(run func(ctx context.Context, key string, val interface{}, expTime time.Duration)) *MockCacheService_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}), args[3].(time.Duration))
	})
	return _c
}

func (_c *MockCacheService_Set_Call) Return(_a0 error) *MockCacheService_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCacheService_Set_Call) RunAndReturn(run func(context.Context, string, interface{}, time.Duration) error) *MockCacheService_Set_Call {
	_c.Call.Return(run)
	return _c
}

// Unlock provides a mock function with given fields: ctx, key
func (_m *MockCacheService) Unlock(ctx context.Context, key string) (interface{}, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Unlock")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (interface{}, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) interface{}); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCacheService_Unlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unlock'
type MockCacheService_Unlock_Call struct {
	*mock.Call
}

// Unlock is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *MockCacheService_Expecter) Unlock(ctx interface{}, key interface{}) *MockCacheService_Unlock_Call {
	return &MockCacheService_Unlock_Call{Call: _e.mock.On("Unlock", ctx, key)}
}

func (_c *MockCacheService_Unlock_Call) Run(run func(ctx context.Context, key string)) *MockCacheService_Unlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockCacheService_Unlock_Call) Return(val interface{}, err error) *MockCacheService_Unlock_Call {
	_c.Call.Return(val, err)
	return _c
}

func (_c *MockCacheService_Unlock_Call) RunAndReturn(run func(context.Context, string) (interface{}, error)) *MockCacheService_Unlock_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCacheService creates a new instance of MockCacheService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCacheService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCacheService {
	mock := &MockCacheService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

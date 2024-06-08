// Code generated by mockery v2.43.2. DO NOT EDIT.

package string

import mock "github.com/stretchr/testify/mock"

// MockRandomString is an autogenerated mock type for the RandomString type
type MockRandomString struct {
	mock.Mock
}

type MockRandomString_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRandomString) EXPECT() *MockRandomString_Expecter {
	return &MockRandomString_Expecter{mock: &_m.Mock}
}

// GenerateID provides a mock function with given fields:
func (_m *MockRandomString) GenerateID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GenerateID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockRandomString_GenerateID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateID'
type MockRandomString_GenerateID_Call struct {
	*mock.Call
}

// GenerateID is a helper method to define mock.On call
func (_e *MockRandomString_Expecter) GenerateID() *MockRandomString_GenerateID_Call {
	return &MockRandomString_GenerateID_Call{Call: _e.mock.On("GenerateID")}
}

func (_c *MockRandomString_GenerateID_Call) Run(run func()) *MockRandomString_GenerateID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRandomString_GenerateID_Call) Return(_a0 string) *MockRandomString_GenerateID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRandomString_GenerateID_Call) RunAndReturn(run func() string) *MockRandomString_GenerateID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRandomString creates a new instance of MockRandomString. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRandomString(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRandomString {
	mock := &MockRandomString{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
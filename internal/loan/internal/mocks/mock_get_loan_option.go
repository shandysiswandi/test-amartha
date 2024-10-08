// Code generated by mockery. DO NOT EDIT.

package loanmocks

import (
	goqu "github.com/doug-martin/goqu/v9"
	mock "github.com/stretchr/testify/mock"
)

// MockGetLoanOption is an autogenerated mock type for the GetLoanOption type
type MockGetLoanOption struct {
	mock.Mock
}

type MockGetLoanOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGetLoanOption) EXPECT() *MockGetLoanOption_Expecter {
	return &MockGetLoanOption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MockGetLoanOption) Execute(_a0 *goqu.SelectDataset) *goqu.SelectDataset {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *goqu.SelectDataset
	if rf, ok := ret.Get(0).(func(*goqu.SelectDataset) *goqu.SelectDataset); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*goqu.SelectDataset)
		}
	}

	return r0
}

// MockGetLoanOption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockGetLoanOption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *goqu.SelectDataset
func (_e *MockGetLoanOption_Expecter) Execute(_a0 interface{}) *MockGetLoanOption_Execute_Call {
	return &MockGetLoanOption_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MockGetLoanOption_Execute_Call) Run(run func(_a0 *goqu.SelectDataset)) *MockGetLoanOption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*goqu.SelectDataset))
	})
	return _c
}

func (_c *MockGetLoanOption_Execute_Call) Return(_a0 *goqu.SelectDataset) *MockGetLoanOption_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGetLoanOption_Execute_Call) RunAndReturn(run func(*goqu.SelectDataset) *goqu.SelectDataset) *MockGetLoanOption_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGetLoanOption creates a new instance of MockGetLoanOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGetLoanOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGetLoanOption {
	mock := &MockGetLoanOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

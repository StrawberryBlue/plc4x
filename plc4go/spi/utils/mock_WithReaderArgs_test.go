/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.42.2. DO NOT EDIT.

package utils

import mock "github.com/stretchr/testify/mock"

// MockWithReaderArgs is an autogenerated mock type for the WithReaderArgs type
type MockWithReaderArgs struct {
	mock.Mock
}

type MockWithReaderArgs_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWithReaderArgs) EXPECT() *MockWithReaderArgs_Expecter {
	return &MockWithReaderArgs_Expecter{mock: &_m.Mock}
}

// isReaderArgs provides a mock function with given fields:
func (_m *MockWithReaderArgs) isReaderArgs() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for isReaderArgs")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockWithReaderArgs_isReaderArgs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isReaderArgs'
type MockWithReaderArgs_isReaderArgs_Call struct {
	*mock.Call
}

// isReaderArgs is a helper method to define mock.On call
func (_e *MockWithReaderArgs_Expecter) isReaderArgs() *MockWithReaderArgs_isReaderArgs_Call {
	return &MockWithReaderArgs_isReaderArgs_Call{Call: _e.mock.On("isReaderArgs")}
}

func (_c *MockWithReaderArgs_isReaderArgs_Call) Run(run func()) *MockWithReaderArgs_isReaderArgs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWithReaderArgs_isReaderArgs_Call) Return(_a0 bool) *MockWithReaderArgs_isReaderArgs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWithReaderArgs_isReaderArgs_Call) RunAndReturn(run func() bool) *MockWithReaderArgs_isReaderArgs_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWithReaderArgs creates a new instance of MockWithReaderArgs. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWithReaderArgs(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWithReaderArgs {
	mock := &MockWithReaderArgs{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

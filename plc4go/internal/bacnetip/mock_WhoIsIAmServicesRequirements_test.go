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

package bacnetip

import mock "github.com/stretchr/testify/mock"

// MockWhoIsIAmServicesRequirements is an autogenerated mock type for the WhoIsIAmServicesRequirements type
type MockWhoIsIAmServicesRequirements struct {
	mock.Mock
}

type MockWhoIsIAmServicesRequirements_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWhoIsIAmServicesRequirements) EXPECT() *MockWhoIsIAmServicesRequirements_Expecter {
	return &MockWhoIsIAmServicesRequirements_Expecter{mock: &_m.Mock}
}

// Request provides a mock function with given fields: args, kwargs
func (_m *MockWhoIsIAmServicesRequirements) Request(args _args, kwargs _kwargs) error {
	ret := _m.Called(args, kwargs)

	if len(ret) == 0 {
		panic("no return value specified for Request")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(_args, _kwargs) error); ok {
		r0 = rf(args, kwargs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockWhoIsIAmServicesRequirements_Request_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Request'
type MockWhoIsIAmServicesRequirements_Request_Call struct {
	*mock.Call
}

// Request is a helper method to define mock.On call
//   - args _args
//   - kwargs _kwargs
func (_e *MockWhoIsIAmServicesRequirements_Expecter) Request(args interface{}, kwargs interface{}) *MockWhoIsIAmServicesRequirements_Request_Call {
	return &MockWhoIsIAmServicesRequirements_Request_Call{Call: _e.mock.On("Request", args, kwargs)}
}

func (_c *MockWhoIsIAmServicesRequirements_Request_Call) Run(run func(args _args, kwargs _kwargs)) *MockWhoIsIAmServicesRequirements_Request_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_args), args[1].(_kwargs))
	})
	return _c
}

func (_c *MockWhoIsIAmServicesRequirements_Request_Call) Return(_a0 error) *MockWhoIsIAmServicesRequirements_Request_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWhoIsIAmServicesRequirements_Request_Call) RunAndReturn(run func(_args, _kwargs) error) *MockWhoIsIAmServicesRequirements_Request_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWhoIsIAmServicesRequirements creates a new instance of MockWhoIsIAmServicesRequirements. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWhoIsIAmServicesRequirements(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWhoIsIAmServicesRequirements {
	mock := &MockWhoIsIAmServicesRequirements{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

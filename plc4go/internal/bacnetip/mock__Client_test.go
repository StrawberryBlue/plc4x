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

// mock_Client is an autogenerated mock type for the _Client type
type mock_Client struct {
	mock.Mock
}

type mock_Client_Expecter struct {
	mock *mock.Mock
}

func (_m *mock_Client) EXPECT() *mock_Client_Expecter {
	return &mock_Client_Expecter{mock: &_m.Mock}
}

// Confirmation provides a mock function with given fields: args, kwargs
func (_m *mock_Client) Confirmation(args Args, kwargs KWArgs) error {
	ret := _m.Called(args, kwargs)

	if len(ret) == 0 {
		panic("no return value specified for Confirmation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Args, KWArgs) error); ok {
		r0 = rf(args, kwargs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_Client_Confirmation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Confirmation'
type mock_Client_Confirmation_Call struct {
	*mock.Call
}

// Confirmation is a helper method to define mock.On call
//   - args Args
//   - kwargs KWArgs
func (_e *mock_Client_Expecter) Confirmation(args interface{}, kwargs interface{}) *mock_Client_Confirmation_Call {
	return &mock_Client_Confirmation_Call{Call: _e.mock.On("Confirmation", args, kwargs)}
}

func (_c *mock_Client_Confirmation_Call) Run(run func(args Args, kwargs KWArgs)) *mock_Client_Confirmation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Args), args[1].(KWArgs))
	})
	return _c
}

func (_c *mock_Client_Confirmation_Call) Return(_a0 error) *mock_Client_Confirmation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_Client_Confirmation_Call) RunAndReturn(run func(Args, KWArgs) error) *mock_Client_Confirmation_Call {
	_c.Call.Return(run)
	return _c
}

// Request provides a mock function with given fields: args, kwargs
func (_m *mock_Client) Request(args Args, kwargs KWArgs) error {
	ret := _m.Called(args, kwargs)

	if len(ret) == 0 {
		panic("no return value specified for Request")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Args, KWArgs) error); ok {
		r0 = rf(args, kwargs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_Client_Request_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Request'
type mock_Client_Request_Call struct {
	*mock.Call
}

// Request is a helper method to define mock.On call
//   - args Args
//   - kwargs KWArgs
func (_e *mock_Client_Expecter) Request(args interface{}, kwargs interface{}) *mock_Client_Request_Call {
	return &mock_Client_Request_Call{Call: _e.mock.On("Request", args, kwargs)}
}

func (_c *mock_Client_Request_Call) Run(run func(args Args, kwargs KWArgs)) *mock_Client_Request_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Args), args[1].(KWArgs))
	})
	return _c
}

func (_c *mock_Client_Request_Call) Return(_a0 error) *mock_Client_Request_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_Client_Request_Call) RunAndReturn(run func(Args, KWArgs) error) *mock_Client_Request_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *mock_Client) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// mock_Client_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type mock_Client_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *mock_Client_Expecter) String() *mock_Client_String_Call {
	return &mock_Client_String_Call{Call: _e.mock.On("String")}
}

func (_c *mock_Client_String_Call) Run(run func()) *mock_Client_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_Client_String_Call) Return(_a0 string) *mock_Client_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_Client_String_Call) RunAndReturn(run func() string) *mock_Client_String_Call {
	_c.Call.Return(run)
	return _c
}

// _setClientPeer provides a mock function with given fields: server
func (_m *mock_Client) _setClientPeer(server _Server) {
	_m.Called(server)
}

// mock_Client__setClientPeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method '_setClientPeer'
type mock_Client__setClientPeer_Call struct {
	*mock.Call
}

// _setClientPeer is a helper method to define mock.On call
//   - server _Server
func (_e *mock_Client_Expecter) _setClientPeer(server interface{}) *mock_Client__setClientPeer_Call {
	return &mock_Client__setClientPeer_Call{Call: _e.mock.On("_setClientPeer", server)}
}

func (_c *mock_Client__setClientPeer_Call) Run(run func(server _Server)) *mock_Client__setClientPeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_Server))
	})
	return _c
}

func (_c *mock_Client__setClientPeer_Call) Return() *mock_Client__setClientPeer_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_Client__setClientPeer_Call) RunAndReturn(run func(_Server)) *mock_Client__setClientPeer_Call {
	_c.Call.Return(run)
	return _c
}

// getClientId provides a mock function with given fields:
func (_m *mock_Client) getClientId() *int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for getClientId")
	}

	var r0 *int
	if rf, ok := ret.Get(0).(func() *int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	return r0
}

// mock_Client_getClientId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getClientId'
type mock_Client_getClientId_Call struct {
	*mock.Call
}

// getClientId is a helper method to define mock.On call
func (_e *mock_Client_Expecter) getClientId() *mock_Client_getClientId_Call {
	return &mock_Client_getClientId_Call{Call: _e.mock.On("getClientId")}
}

func (_c *mock_Client_getClientId_Call) Run(run func()) *mock_Client_getClientId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mock_Client_getClientId_Call) Return(_a0 *int) *mock_Client_getClientId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_Client_getClientId_Call) RunAndReturn(run func() *int) *mock_Client_getClientId_Call {
	_c.Call.Return(run)
	return _c
}

// newMock_Client creates a new instance of mock_Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMock_Client(t interface {
	mock.TestingT
	Cleanup(func())
}) *mock_Client {
	mock := &mock_Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

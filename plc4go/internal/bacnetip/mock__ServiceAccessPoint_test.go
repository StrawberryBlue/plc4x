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

// mock_ServiceAccessPoint is an autogenerated mock type for the _ServiceAccessPoint type
type mock_ServiceAccessPoint struct {
	mock.Mock
}

type mock_ServiceAccessPoint_Expecter struct {
	mock *mock.Mock
}

func (_m *mock_ServiceAccessPoint) EXPECT() *mock_ServiceAccessPoint_Expecter {
	return &mock_ServiceAccessPoint_Expecter{mock: &_m.Mock}
}

// SapConfirmation provides a mock function with given fields: _a0, _a1
func (_m *mock_ServiceAccessPoint) SapConfirmation(_a0 Args, _a1 KWArgs) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SapConfirmation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Args, KWArgs) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_ServiceAccessPoint_SapConfirmation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapConfirmation'
type mock_ServiceAccessPoint_SapConfirmation_Call struct {
	*mock.Call
}

// SapConfirmation is a helper method to define mock.On call
//   - _a0 Args
//   - _a1 KWArgs
func (_e *mock_ServiceAccessPoint_Expecter) SapConfirmation(_a0 interface{}, _a1 interface{}) *mock_ServiceAccessPoint_SapConfirmation_Call {
	return &mock_ServiceAccessPoint_SapConfirmation_Call{Call: _e.mock.On("SapConfirmation", _a0, _a1)}
}

func (_c *mock_ServiceAccessPoint_SapConfirmation_Call) Run(run func(_a0 Args, _a1 KWArgs)) *mock_ServiceAccessPoint_SapConfirmation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Args), args[1].(KWArgs))
	})
	return _c
}

func (_c *mock_ServiceAccessPoint_SapConfirmation_Call) Return(_a0 error) *mock_ServiceAccessPoint_SapConfirmation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_ServiceAccessPoint_SapConfirmation_Call) RunAndReturn(run func(Args, KWArgs) error) *mock_ServiceAccessPoint_SapConfirmation_Call {
	_c.Call.Return(run)
	return _c
}

// SapIndication provides a mock function with given fields: _a0, _a1
func (_m *mock_ServiceAccessPoint) SapIndication(_a0 Args, _a1 KWArgs) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SapIndication")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Args, KWArgs) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_ServiceAccessPoint_SapIndication_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapIndication'
type mock_ServiceAccessPoint_SapIndication_Call struct {
	*mock.Call
}

// SapIndication is a helper method to define mock.On call
//   - _a0 Args
//   - _a1 KWArgs
func (_e *mock_ServiceAccessPoint_Expecter) SapIndication(_a0 interface{}, _a1 interface{}) *mock_ServiceAccessPoint_SapIndication_Call {
	return &mock_ServiceAccessPoint_SapIndication_Call{Call: _e.mock.On("SapIndication", _a0, _a1)}
}

func (_c *mock_ServiceAccessPoint_SapIndication_Call) Run(run func(_a0 Args, _a1 KWArgs)) *mock_ServiceAccessPoint_SapIndication_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Args), args[1].(KWArgs))
	})
	return _c
}

func (_c *mock_ServiceAccessPoint_SapIndication_Call) Return(_a0 error) *mock_ServiceAccessPoint_SapIndication_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_ServiceAccessPoint_SapIndication_Call) RunAndReturn(run func(Args, KWArgs) error) *mock_ServiceAccessPoint_SapIndication_Call {
	_c.Call.Return(run)
	return _c
}

// SapRequest provides a mock function with given fields: _a0, _a1
func (_m *mock_ServiceAccessPoint) SapRequest(_a0 Args, _a1 KWArgs) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SapRequest")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Args, KWArgs) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_ServiceAccessPoint_SapRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapRequest'
type mock_ServiceAccessPoint_SapRequest_Call struct {
	*mock.Call
}

// SapRequest is a helper method to define mock.On call
//   - _a0 Args
//   - _a1 KWArgs
func (_e *mock_ServiceAccessPoint_Expecter) SapRequest(_a0 interface{}, _a1 interface{}) *mock_ServiceAccessPoint_SapRequest_Call {
	return &mock_ServiceAccessPoint_SapRequest_Call{Call: _e.mock.On("SapRequest", _a0, _a1)}
}

func (_c *mock_ServiceAccessPoint_SapRequest_Call) Run(run func(_a0 Args, _a1 KWArgs)) *mock_ServiceAccessPoint_SapRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Args), args[1].(KWArgs))
	})
	return _c
}

func (_c *mock_ServiceAccessPoint_SapRequest_Call) Return(_a0 error) *mock_ServiceAccessPoint_SapRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_ServiceAccessPoint_SapRequest_Call) RunAndReturn(run func(Args, KWArgs) error) *mock_ServiceAccessPoint_SapRequest_Call {
	_c.Call.Return(run)
	return _c
}

// SapResponse provides a mock function with given fields: _a0, _a1
func (_m *mock_ServiceAccessPoint) SapResponse(_a0 Args, _a1 KWArgs) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SapResponse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Args, KWArgs) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_ServiceAccessPoint_SapResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SapResponse'
type mock_ServiceAccessPoint_SapResponse_Call struct {
	*mock.Call
}

// SapResponse is a helper method to define mock.On call
//   - _a0 Args
//   - _a1 KWArgs
func (_e *mock_ServiceAccessPoint_Expecter) SapResponse(_a0 interface{}, _a1 interface{}) *mock_ServiceAccessPoint_SapResponse_Call {
	return &mock_ServiceAccessPoint_SapResponse_Call{Call: _e.mock.On("SapResponse", _a0, _a1)}
}

func (_c *mock_ServiceAccessPoint_SapResponse_Call) Run(run func(_a0 Args, _a1 KWArgs)) *mock_ServiceAccessPoint_SapResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Args), args[1].(KWArgs))
	})
	return _c
}

func (_c *mock_ServiceAccessPoint_SapResponse_Call) Return(_a0 error) *mock_ServiceAccessPoint_SapResponse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_ServiceAccessPoint_SapResponse_Call) RunAndReturn(run func(Args, KWArgs) error) *mock_ServiceAccessPoint_SapResponse_Call {
	_c.Call.Return(run)
	return _c
}

// _setServiceElement provides a mock function with given fields: serviceElement
func (_m *mock_ServiceAccessPoint) _setServiceElement(serviceElement _ApplicationServiceElement) {
	_m.Called(serviceElement)
}

// mock_ServiceAccessPoint__setServiceElement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method '_setServiceElement'
type mock_ServiceAccessPoint__setServiceElement_Call struct {
	*mock.Call
}

// _setServiceElement is a helper method to define mock.On call
//   - serviceElement _ApplicationServiceElement
func (_e *mock_ServiceAccessPoint_Expecter) _setServiceElement(serviceElement interface{}) *mock_ServiceAccessPoint__setServiceElement_Call {
	return &mock_ServiceAccessPoint__setServiceElement_Call{Call: _e.mock.On("_setServiceElement", serviceElement)}
}

func (_c *mock_ServiceAccessPoint__setServiceElement_Call) Run(run func(serviceElement _ApplicationServiceElement)) *mock_ServiceAccessPoint__setServiceElement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_ApplicationServiceElement))
	})
	return _c
}

func (_c *mock_ServiceAccessPoint__setServiceElement_Call) Return() *mock_ServiceAccessPoint__setServiceElement_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_ServiceAccessPoint__setServiceElement_Call) RunAndReturn(run func(_ApplicationServiceElement)) *mock_ServiceAccessPoint__setServiceElement_Call {
	_c.Call.Return(run)
	return _c
}

// newMock_ServiceAccessPoint creates a new instance of mock_ServiceAccessPoint. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMock_ServiceAccessPoint(t interface {
	mock.TestingT
	Cleanup(func())
}) *mock_ServiceAccessPoint {
	mock := &mock_ServiceAccessPoint{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

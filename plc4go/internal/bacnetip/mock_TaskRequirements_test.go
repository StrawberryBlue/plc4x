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

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockTaskRequirements is an autogenerated mock type for the TaskRequirements type
type MockTaskRequirements struct {
	mock.Mock
}

type MockTaskRequirements_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTaskRequirements) EXPECT() *MockTaskRequirements_Expecter {
	return &MockTaskRequirements_Expecter{mock: &_m.Mock}
}

// GetIsScheduled provides a mock function with given fields:
func (_m *MockTaskRequirements) GetIsScheduled() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetIsScheduled")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockTaskRequirements_GetIsScheduled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIsScheduled'
type MockTaskRequirements_GetIsScheduled_Call struct {
	*mock.Call
}

// GetIsScheduled is a helper method to define mock.On call
func (_e *MockTaskRequirements_Expecter) GetIsScheduled() *MockTaskRequirements_GetIsScheduled_Call {
	return &MockTaskRequirements_GetIsScheduled_Call{Call: _e.mock.On("GetIsScheduled")}
}

func (_c *MockTaskRequirements_GetIsScheduled_Call) Run(run func()) *MockTaskRequirements_GetIsScheduled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTaskRequirements_GetIsScheduled_Call) Return(_a0 bool) *MockTaskRequirements_GetIsScheduled_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTaskRequirements_GetIsScheduled_Call) RunAndReturn(run func() bool) *MockTaskRequirements_GetIsScheduled_Call {
	_c.Call.Return(run)
	return _c
}

// GetTaskTime provides a mock function with given fields:
func (_m *MockTaskRequirements) GetTaskTime() *time.Time {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTaskTime")
	}

	var r0 *time.Time
	if rf, ok := ret.Get(0).(func() *time.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Time)
		}
	}

	return r0
}

// MockTaskRequirements_GetTaskTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTaskTime'
type MockTaskRequirements_GetTaskTime_Call struct {
	*mock.Call
}

// GetTaskTime is a helper method to define mock.On call
func (_e *MockTaskRequirements_Expecter) GetTaskTime() *MockTaskRequirements_GetTaskTime_Call {
	return &MockTaskRequirements_GetTaskTime_Call{Call: _e.mock.On("GetTaskTime")}
}

func (_c *MockTaskRequirements_GetTaskTime_Call) Run(run func()) *MockTaskRequirements_GetTaskTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTaskRequirements_GetTaskTime_Call) Return(_a0 *time.Time) *MockTaskRequirements_GetTaskTime_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTaskRequirements_GetTaskTime_Call) RunAndReturn(run func() *time.Time) *MockTaskRequirements_GetTaskTime_Call {
	_c.Call.Return(run)
	return _c
}

// InstallTask provides a mock function with given fields: when, delta
func (_m *MockTaskRequirements) InstallTask(when *time.Time, delta *time.Duration) {
	_m.Called(when, delta)
}

// MockTaskRequirements_InstallTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InstallTask'
type MockTaskRequirements_InstallTask_Call struct {
	*mock.Call
}

// InstallTask is a helper method to define mock.On call
//   - when *time.Time
//   - delta *time.Duration
func (_e *MockTaskRequirements_Expecter) InstallTask(when interface{}, delta interface{}) *MockTaskRequirements_InstallTask_Call {
	return &MockTaskRequirements_InstallTask_Call{Call: _e.mock.On("InstallTask", when, delta)}
}

func (_c *MockTaskRequirements_InstallTask_Call) Run(run func(when *time.Time, delta *time.Duration)) *MockTaskRequirements_InstallTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*time.Time), args[1].(*time.Duration))
	})
	return _c
}

func (_c *MockTaskRequirements_InstallTask_Call) Return() *MockTaskRequirements_InstallTask_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTaskRequirements_InstallTask_Call) RunAndReturn(run func(*time.Time, *time.Duration)) *MockTaskRequirements_InstallTask_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessTask provides a mock function with given fields:
func (_m *MockTaskRequirements) ProcessTask() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ProcessTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTaskRequirements_ProcessTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessTask'
type MockTaskRequirements_ProcessTask_Call struct {
	*mock.Call
}

// ProcessTask is a helper method to define mock.On call
func (_e *MockTaskRequirements_Expecter) ProcessTask() *MockTaskRequirements_ProcessTask_Call {
	return &MockTaskRequirements_ProcessTask_Call{Call: _e.mock.On("ProcessTask")}
}

func (_c *MockTaskRequirements_ProcessTask_Call) Run(run func()) *MockTaskRequirements_ProcessTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTaskRequirements_ProcessTask_Call) Return(_a0 error) *MockTaskRequirements_ProcessTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTaskRequirements_ProcessTask_Call) RunAndReturn(run func() error) *MockTaskRequirements_ProcessTask_Call {
	_c.Call.Return(run)
	return _c
}

// SetIsScheduled provides a mock function with given fields: isScheduled
func (_m *MockTaskRequirements) SetIsScheduled(isScheduled bool) {
	_m.Called(isScheduled)
}

// MockTaskRequirements_SetIsScheduled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetIsScheduled'
type MockTaskRequirements_SetIsScheduled_Call struct {
	*mock.Call
}

// SetIsScheduled is a helper method to define mock.On call
//   - isScheduled bool
func (_e *MockTaskRequirements_Expecter) SetIsScheduled(isScheduled interface{}) *MockTaskRequirements_SetIsScheduled_Call {
	return &MockTaskRequirements_SetIsScheduled_Call{Call: _e.mock.On("SetIsScheduled", isScheduled)}
}

func (_c *MockTaskRequirements_SetIsScheduled_Call) Run(run func(isScheduled bool)) *MockTaskRequirements_SetIsScheduled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *MockTaskRequirements_SetIsScheduled_Call) Return() *MockTaskRequirements_SetIsScheduled_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTaskRequirements_SetIsScheduled_Call) RunAndReturn(run func(bool)) *MockTaskRequirements_SetIsScheduled_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTaskRequirements creates a new instance of MockTaskRequirements. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTaskRequirements(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTaskRequirements {
	mock := &MockTaskRequirements{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

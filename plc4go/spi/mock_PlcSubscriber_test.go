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

package spi

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"
)

// MockPlcSubscriber is an autogenerated mock type for the PlcSubscriber type
type MockPlcSubscriber struct {
	mock.Mock
}

type MockPlcSubscriber_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcSubscriber) EXPECT() *MockPlcSubscriber_Expecter {
	return &MockPlcSubscriber_Expecter{mock: &_m.Mock}
}

// Register provides a mock function with given fields: consumer, handles
func (_m *MockPlcSubscriber) Register(consumer model.PlcSubscriptionEventConsumer, handles []model.PlcSubscriptionHandle) model.PlcConsumerRegistration {
	ret := _m.Called(consumer, handles)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 model.PlcConsumerRegistration
	if rf, ok := ret.Get(0).(func(model.PlcSubscriptionEventConsumer, []model.PlcSubscriptionHandle) model.PlcConsumerRegistration); ok {
		r0 = rf(consumer, handles)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcConsumerRegistration)
		}
	}

	return r0
}

// MockPlcSubscriber_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type MockPlcSubscriber_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - consumer model.PlcSubscriptionEventConsumer
//   - handles []model.PlcSubscriptionHandle
func (_e *MockPlcSubscriber_Expecter) Register(consumer interface{}, handles interface{}) *MockPlcSubscriber_Register_Call {
	return &MockPlcSubscriber_Register_Call{Call: _e.mock.On("Register", consumer, handles)}
}

func (_c *MockPlcSubscriber_Register_Call) Run(run func(consumer model.PlcSubscriptionEventConsumer, handles []model.PlcSubscriptionHandle)) *MockPlcSubscriber_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.PlcSubscriptionEventConsumer), args[1].([]model.PlcSubscriptionHandle))
	})
	return _c
}

func (_c *MockPlcSubscriber_Register_Call) Return(_a0 model.PlcConsumerRegistration) *MockPlcSubscriber_Register_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriber_Register_Call) RunAndReturn(run func(model.PlcSubscriptionEventConsumer, []model.PlcSubscriptionHandle) model.PlcConsumerRegistration) *MockPlcSubscriber_Register_Call {
	_c.Call.Return(run)
	return _c
}

// Subscribe provides a mock function with given fields: ctx, subscriptionRequest
func (_m *MockPlcSubscriber) Subscribe(ctx context.Context, subscriptionRequest model.PlcSubscriptionRequest) <-chan model.PlcSubscriptionRequestResult {
	ret := _m.Called(ctx, subscriptionRequest)

	if len(ret) == 0 {
		panic("no return value specified for Subscribe")
	}

	var r0 <-chan model.PlcSubscriptionRequestResult
	if rf, ok := ret.Get(0).(func(context.Context, model.PlcSubscriptionRequest) <-chan model.PlcSubscriptionRequestResult); ok {
		r0 = rf(ctx, subscriptionRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan model.PlcSubscriptionRequestResult)
		}
	}

	return r0
}

// MockPlcSubscriber_Subscribe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Subscribe'
type MockPlcSubscriber_Subscribe_Call struct {
	*mock.Call
}

// Subscribe is a helper method to define mock.On call
//   - ctx context.Context
//   - subscriptionRequest model.PlcSubscriptionRequest
func (_e *MockPlcSubscriber_Expecter) Subscribe(ctx interface{}, subscriptionRequest interface{}) *MockPlcSubscriber_Subscribe_Call {
	return &MockPlcSubscriber_Subscribe_Call{Call: _e.mock.On("Subscribe", ctx, subscriptionRequest)}
}

func (_c *MockPlcSubscriber_Subscribe_Call) Run(run func(ctx context.Context, subscriptionRequest model.PlcSubscriptionRequest)) *MockPlcSubscriber_Subscribe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.PlcSubscriptionRequest))
	})
	return _c
}

func (_c *MockPlcSubscriber_Subscribe_Call) Return(_a0 <-chan model.PlcSubscriptionRequestResult) *MockPlcSubscriber_Subscribe_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriber_Subscribe_Call) RunAndReturn(run func(context.Context, model.PlcSubscriptionRequest) <-chan model.PlcSubscriptionRequestResult) *MockPlcSubscriber_Subscribe_Call {
	_c.Call.Return(run)
	return _c
}

// Unregister provides a mock function with given fields: registration
func (_m *MockPlcSubscriber) Unregister(registration model.PlcConsumerRegistration) {
	_m.Called(registration)
}

// MockPlcSubscriber_Unregister_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unregister'
type MockPlcSubscriber_Unregister_Call struct {
	*mock.Call
}

// Unregister is a helper method to define mock.On call
//   - registration model.PlcConsumerRegistration
func (_e *MockPlcSubscriber_Expecter) Unregister(registration interface{}) *MockPlcSubscriber_Unregister_Call {
	return &MockPlcSubscriber_Unregister_Call{Call: _e.mock.On("Unregister", registration)}
}

func (_c *MockPlcSubscriber_Unregister_Call) Run(run func(registration model.PlcConsumerRegistration)) *MockPlcSubscriber_Unregister_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.PlcConsumerRegistration))
	})
	return _c
}

func (_c *MockPlcSubscriber_Unregister_Call) Return() *MockPlcSubscriber_Unregister_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPlcSubscriber_Unregister_Call) RunAndReturn(run func(model.PlcConsumerRegistration)) *MockPlcSubscriber_Unregister_Call {
	_c.Call.Return(run)
	return _c
}

// Unsubscribe provides a mock function with given fields: ctx, unsubscriptionRequest
func (_m *MockPlcSubscriber) Unsubscribe(ctx context.Context, unsubscriptionRequest model.PlcUnsubscriptionRequest) <-chan model.PlcUnsubscriptionRequestResult {
	ret := _m.Called(ctx, unsubscriptionRequest)

	if len(ret) == 0 {
		panic("no return value specified for Unsubscribe")
	}

	var r0 <-chan model.PlcUnsubscriptionRequestResult
	if rf, ok := ret.Get(0).(func(context.Context, model.PlcUnsubscriptionRequest) <-chan model.PlcUnsubscriptionRequestResult); ok {
		r0 = rf(ctx, unsubscriptionRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan model.PlcUnsubscriptionRequestResult)
		}
	}

	return r0
}

// MockPlcSubscriber_Unsubscribe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unsubscribe'
type MockPlcSubscriber_Unsubscribe_Call struct {
	*mock.Call
}

// Unsubscribe is a helper method to define mock.On call
//   - ctx context.Context
//   - unsubscriptionRequest model.PlcUnsubscriptionRequest
func (_e *MockPlcSubscriber_Expecter) Unsubscribe(ctx interface{}, unsubscriptionRequest interface{}) *MockPlcSubscriber_Unsubscribe_Call {
	return &MockPlcSubscriber_Unsubscribe_Call{Call: _e.mock.On("Unsubscribe", ctx, unsubscriptionRequest)}
}

func (_c *MockPlcSubscriber_Unsubscribe_Call) Run(run func(ctx context.Context, unsubscriptionRequest model.PlcUnsubscriptionRequest)) *MockPlcSubscriber_Unsubscribe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.PlcUnsubscriptionRequest))
	})
	return _c
}

func (_c *MockPlcSubscriber_Unsubscribe_Call) Return(_a0 <-chan model.PlcUnsubscriptionRequestResult) *MockPlcSubscriber_Unsubscribe_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcSubscriber_Unsubscribe_Call) RunAndReturn(run func(context.Context, model.PlcUnsubscriptionRequest) <-chan model.PlcUnsubscriptionRequestResult) *MockPlcSubscriber_Unsubscribe_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcSubscriber creates a new instance of MockPlcSubscriber. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcSubscriber(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcSubscriber {
	mock := &MockPlcSubscriber{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

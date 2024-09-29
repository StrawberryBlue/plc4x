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

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// CBusMessage is the corresponding interface of CBusMessage
type CBusMessage interface {
	CBusMessageContract
	CBusMessageRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsCBusMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusMessage()
	// CreateBuilder creates a CBusMessageBuilder
	CreateCBusMessageBuilder() CBusMessageBuilder
}

// CBusMessageContract provides a set of functions which can be overwritten by a sub struct
type CBusMessageContract interface {
	// GetRequestContext() returns a parser argument
	GetRequestContext() RequestContext
	// GetCBusOptions() returns a parser argument
	GetCBusOptions() CBusOptions
	// IsCBusMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusMessage()
	// CreateBuilder creates a CBusMessageBuilder
	CreateCBusMessageBuilder() CBusMessageBuilder
}

// CBusMessageRequirements provides a set of functions which need to be implemented by a sub struct
type CBusMessageRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetIsResponse returns IsResponse (discriminator field)
	GetIsResponse() bool
}

// _CBusMessage is the data-structure of this message
type _CBusMessage struct {
	_SubType CBusMessage

	// Arguments.
	RequestContext RequestContext
	CBusOptions    CBusOptions
}

var _ CBusMessageContract = (*_CBusMessage)(nil)

// NewCBusMessage factory function for _CBusMessage
func NewCBusMessage(requestContext RequestContext, cBusOptions CBusOptions) *_CBusMessage {
	return &_CBusMessage{RequestContext: requestContext, CBusOptions: cBusOptions}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CBusMessageBuilder is a builder for CBusMessage
type CBusMessageBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() CBusMessageBuilder
	// AsCBusMessageToServer converts this build to a subType of CBusMessage. It is always possible to return to current builder using Done()
	AsCBusMessageToServer() interface {
		CBusMessageToServerBuilder
		Done() CBusMessageBuilder
	}
	// AsCBusMessageToClient converts this build to a subType of CBusMessage. It is always possible to return to current builder using Done()
	AsCBusMessageToClient() interface {
		CBusMessageToClientBuilder
		Done() CBusMessageBuilder
	}
	// Build builds the CBusMessage or returns an error if something is wrong
	PartialBuild() (CBusMessageContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() CBusMessageContract
	// Build builds the CBusMessage or returns an error if something is wrong
	Build() (CBusMessage, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CBusMessage
}

// NewCBusMessageBuilder() creates a CBusMessageBuilder
func NewCBusMessageBuilder() CBusMessageBuilder {
	return &_CBusMessageBuilder{_CBusMessage: new(_CBusMessage)}
}

type _CBusMessageChildBuilder interface {
	utils.Copyable
	setParent(CBusMessageContract)
	buildForCBusMessage() (CBusMessage, error)
}

type _CBusMessageBuilder struct {
	*_CBusMessage

	childBuilder _CBusMessageChildBuilder

	err *utils.MultiError
}

var _ (CBusMessageBuilder) = (*_CBusMessageBuilder)(nil)

func (b *_CBusMessageBuilder) WithMandatoryFields() CBusMessageBuilder {
	return b
}

func (b *_CBusMessageBuilder) PartialBuild() (CBusMessageContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CBusMessage.deepCopy(), nil
}

func (b *_CBusMessageBuilder) PartialMustBuild() CBusMessageContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CBusMessageBuilder) AsCBusMessageToServer() interface {
	CBusMessageToServerBuilder
	Done() CBusMessageBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CBusMessageToServerBuilder
		Done() CBusMessageBuilder
	}); ok {
		return cb
	}
	cb := NewCBusMessageToServerBuilder().(*_CBusMessageToServerBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CBusMessageBuilder) AsCBusMessageToClient() interface {
	CBusMessageToClientBuilder
	Done() CBusMessageBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		CBusMessageToClientBuilder
		Done() CBusMessageBuilder
	}); ok {
		return cb
	}
	cb := NewCBusMessageToClientBuilder().(*_CBusMessageToClientBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_CBusMessageBuilder) Build() (CBusMessage, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForCBusMessage()
}

func (b *_CBusMessageBuilder) MustBuild() CBusMessage {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CBusMessageBuilder) DeepCopy() any {
	_copy := b.CreateCBusMessageBuilder().(*_CBusMessageBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_CBusMessageChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCBusMessageBuilder creates a CBusMessageBuilder
func (b *_CBusMessage) CreateCBusMessageBuilder() CBusMessageBuilder {
	if b == nil {
		return NewCBusMessageBuilder()
	}
	return &_CBusMessageBuilder{_CBusMessage: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCBusMessage(structType any) CBusMessage {
	if casted, ok := structType.(CBusMessage); ok {
		return casted
	}
	if casted, ok := structType.(*CBusMessage); ok {
		return *casted
	}
	return nil
}

func (m *_CBusMessage) GetTypeName() string {
	return "CBusMessage"
}

func (m *_CBusMessage) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_CBusMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func CBusMessageParse[T CBusMessage](ctx context.Context, theBytes []byte, isResponse bool, requestContext RequestContext, cBusOptions CBusOptions) (T, error) {
	return CBusMessageParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), isResponse, requestContext, cBusOptions)
}

func CBusMessageParseWithBufferProducer[T CBusMessage](isResponse bool, requestContext RequestContext, cBusOptions CBusOptions) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CBusMessageParseWithBuffer[T](ctx, readBuffer, isResponse, requestContext, cBusOptions)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func CBusMessageParseWithBuffer[T CBusMessage](ctx context.Context, readBuffer utils.ReadBuffer, isResponse bool, requestContext RequestContext, cBusOptions CBusOptions) (T, error) {
	v, err := (&_CBusMessage{RequestContext: requestContext, CBusOptions: cBusOptions}).parse(ctx, readBuffer, isResponse, requestContext, cBusOptions)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_CBusMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer, isResponse bool, requestContext RequestContext, cBusOptions CBusOptions) (__cBusMessage CBusMessage, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((requestContext) != (nil))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "requestContext required"})
	}

	// Validation
	if !(bool((cBusOptions) != (nil))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "cBusOptions required"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CBusMessage
	switch {
	case isResponse == bool(false): // CBusMessageToServer
		if _child, err = new(_CBusMessageToServer).parse(ctx, readBuffer, m, isResponse, requestContext, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CBusMessageToServer for type-switch of CBusMessage")
		}
	case isResponse == bool(true): // CBusMessageToClient
		if _child, err = new(_CBusMessageToClient).parse(ctx, readBuffer, m, isResponse, requestContext, cBusOptions); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CBusMessageToClient for type-switch of CBusMessage")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [isResponse=%v]", isResponse)
	}

	if closeErr := readBuffer.CloseContext("CBusMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusMessage")
	}

	return _child, nil
}

func (pm *_CBusMessage) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CBusMessage, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CBusMessage"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CBusMessage")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CBusMessage"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CBusMessage")
	}
	return nil
}

////
// Arguments Getter

func (m *_CBusMessage) GetRequestContext() RequestContext {
	return m.RequestContext
}
func (m *_CBusMessage) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}

//
////

func (m *_CBusMessage) IsCBusMessage() {}

func (m *_CBusMessage) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CBusMessage) deepCopy() *_CBusMessage {
	if m == nil {
		return nil
	}
	_CBusMessageCopy := &_CBusMessage{
		nil, // will be set by child
		m.RequestContext,
		m.CBusOptions,
	}
	return _CBusMessageCopy
}

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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// CreateSubscriptionRequest is the corresponding interface of CreateSubscriptionRequest
type CreateSubscriptionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetRequestedPublishingInterval returns RequestedPublishingInterval (property field)
	GetRequestedPublishingInterval() float64
	// GetRequestedLifetimeCount returns RequestedLifetimeCount (property field)
	GetRequestedLifetimeCount() uint32
	// GetRequestedMaxKeepAliveCount returns RequestedMaxKeepAliveCount (property field)
	GetRequestedMaxKeepAliveCount() uint32
	// GetMaxNotificationsPerPublish returns MaxNotificationsPerPublish (property field)
	GetMaxNotificationsPerPublish() uint32
	// GetPublishingEnabled returns PublishingEnabled (property field)
	GetPublishingEnabled() bool
	// GetPriority returns Priority (property field)
	GetPriority() uint8
	// IsCreateSubscriptionRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCreateSubscriptionRequest()
	// CreateBuilder creates a CreateSubscriptionRequestBuilder
	CreateCreateSubscriptionRequestBuilder() CreateSubscriptionRequestBuilder
}

// _CreateSubscriptionRequest is the data-structure of this message
type _CreateSubscriptionRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader               ExtensionObjectDefinition
	RequestedPublishingInterval float64
	RequestedLifetimeCount      uint32
	RequestedMaxKeepAliveCount  uint32
	MaxNotificationsPerPublish  uint32
	PublishingEnabled           bool
	Priority                    uint8
	// Reserved Fields
	reservedField0 *uint8
}

var _ CreateSubscriptionRequest = (*_CreateSubscriptionRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CreateSubscriptionRequest)(nil)

// NewCreateSubscriptionRequest factory function for _CreateSubscriptionRequest
func NewCreateSubscriptionRequest(requestHeader ExtensionObjectDefinition, requestedPublishingInterval float64, requestedLifetimeCount uint32, requestedMaxKeepAliveCount uint32, maxNotificationsPerPublish uint32, publishingEnabled bool, priority uint8) *_CreateSubscriptionRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for CreateSubscriptionRequest must not be nil")
	}
	_result := &_CreateSubscriptionRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		RequestedPublishingInterval:       requestedPublishingInterval,
		RequestedLifetimeCount:            requestedLifetimeCount,
		RequestedMaxKeepAliveCount:        requestedMaxKeepAliveCount,
		MaxNotificationsPerPublish:        maxNotificationsPerPublish,
		PublishingEnabled:                 publishingEnabled,
		Priority:                          priority,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CreateSubscriptionRequestBuilder is a builder for CreateSubscriptionRequest
type CreateSubscriptionRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader ExtensionObjectDefinition, requestedPublishingInterval float64, requestedLifetimeCount uint32, requestedMaxKeepAliveCount uint32, maxNotificationsPerPublish uint32, publishingEnabled bool, priority uint8) CreateSubscriptionRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(ExtensionObjectDefinition) CreateSubscriptionRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) CreateSubscriptionRequestBuilder
	// WithRequestedPublishingInterval adds RequestedPublishingInterval (property field)
	WithRequestedPublishingInterval(float64) CreateSubscriptionRequestBuilder
	// WithRequestedLifetimeCount adds RequestedLifetimeCount (property field)
	WithRequestedLifetimeCount(uint32) CreateSubscriptionRequestBuilder
	// WithRequestedMaxKeepAliveCount adds RequestedMaxKeepAliveCount (property field)
	WithRequestedMaxKeepAliveCount(uint32) CreateSubscriptionRequestBuilder
	// WithMaxNotificationsPerPublish adds MaxNotificationsPerPublish (property field)
	WithMaxNotificationsPerPublish(uint32) CreateSubscriptionRequestBuilder
	// WithPublishingEnabled adds PublishingEnabled (property field)
	WithPublishingEnabled(bool) CreateSubscriptionRequestBuilder
	// WithPriority adds Priority (property field)
	WithPriority(uint8) CreateSubscriptionRequestBuilder
	// Build builds the CreateSubscriptionRequest or returns an error if something is wrong
	Build() (CreateSubscriptionRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CreateSubscriptionRequest
}

// NewCreateSubscriptionRequestBuilder() creates a CreateSubscriptionRequestBuilder
func NewCreateSubscriptionRequestBuilder() CreateSubscriptionRequestBuilder {
	return &_CreateSubscriptionRequestBuilder{_CreateSubscriptionRequest: new(_CreateSubscriptionRequest)}
}

type _CreateSubscriptionRequestBuilder struct {
	*_CreateSubscriptionRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (CreateSubscriptionRequestBuilder) = (*_CreateSubscriptionRequestBuilder)(nil)

func (b *_CreateSubscriptionRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_CreateSubscriptionRequestBuilder) WithMandatoryFields(requestHeader ExtensionObjectDefinition, requestedPublishingInterval float64, requestedLifetimeCount uint32, requestedMaxKeepAliveCount uint32, maxNotificationsPerPublish uint32, publishingEnabled bool, priority uint8) CreateSubscriptionRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithRequestedPublishingInterval(requestedPublishingInterval).WithRequestedLifetimeCount(requestedLifetimeCount).WithRequestedMaxKeepAliveCount(requestedMaxKeepAliveCount).WithMaxNotificationsPerPublish(maxNotificationsPerPublish).WithPublishingEnabled(publishingEnabled).WithPriority(priority)
}

func (b *_CreateSubscriptionRequestBuilder) WithRequestHeader(requestHeader ExtensionObjectDefinition) CreateSubscriptionRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_CreateSubscriptionRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) CreateSubscriptionRequestBuilder {
	builder := builderSupplier(b.RequestHeader.CreateExtensionObjectDefinitionBuilder())
	var err error
	b.RequestHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return b
}

func (b *_CreateSubscriptionRequestBuilder) WithRequestedPublishingInterval(requestedPublishingInterval float64) CreateSubscriptionRequestBuilder {
	b.RequestedPublishingInterval = requestedPublishingInterval
	return b
}

func (b *_CreateSubscriptionRequestBuilder) WithRequestedLifetimeCount(requestedLifetimeCount uint32) CreateSubscriptionRequestBuilder {
	b.RequestedLifetimeCount = requestedLifetimeCount
	return b
}

func (b *_CreateSubscriptionRequestBuilder) WithRequestedMaxKeepAliveCount(requestedMaxKeepAliveCount uint32) CreateSubscriptionRequestBuilder {
	b.RequestedMaxKeepAliveCount = requestedMaxKeepAliveCount
	return b
}

func (b *_CreateSubscriptionRequestBuilder) WithMaxNotificationsPerPublish(maxNotificationsPerPublish uint32) CreateSubscriptionRequestBuilder {
	b.MaxNotificationsPerPublish = maxNotificationsPerPublish
	return b
}

func (b *_CreateSubscriptionRequestBuilder) WithPublishingEnabled(publishingEnabled bool) CreateSubscriptionRequestBuilder {
	b.PublishingEnabled = publishingEnabled
	return b
}

func (b *_CreateSubscriptionRequestBuilder) WithPriority(priority uint8) CreateSubscriptionRequestBuilder {
	b.Priority = priority
	return b
}

func (b *_CreateSubscriptionRequestBuilder) Build() (CreateSubscriptionRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CreateSubscriptionRequest.deepCopy(), nil
}

func (b *_CreateSubscriptionRequestBuilder) MustBuild() CreateSubscriptionRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_CreateSubscriptionRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_CreateSubscriptionRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_CreateSubscriptionRequestBuilder) DeepCopy() any {
	_copy := b.CreateCreateSubscriptionRequestBuilder().(*_CreateSubscriptionRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCreateSubscriptionRequestBuilder creates a CreateSubscriptionRequestBuilder
func (b *_CreateSubscriptionRequest) CreateCreateSubscriptionRequestBuilder() CreateSubscriptionRequestBuilder {
	if b == nil {
		return NewCreateSubscriptionRequestBuilder()
	}
	return &_CreateSubscriptionRequestBuilder{_CreateSubscriptionRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CreateSubscriptionRequest) GetIdentifier() string {
	return "787"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CreateSubscriptionRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CreateSubscriptionRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_CreateSubscriptionRequest) GetRequestedPublishingInterval() float64 {
	return m.RequestedPublishingInterval
}

func (m *_CreateSubscriptionRequest) GetRequestedLifetimeCount() uint32 {
	return m.RequestedLifetimeCount
}

func (m *_CreateSubscriptionRequest) GetRequestedMaxKeepAliveCount() uint32 {
	return m.RequestedMaxKeepAliveCount
}

func (m *_CreateSubscriptionRequest) GetMaxNotificationsPerPublish() uint32 {
	return m.MaxNotificationsPerPublish
}

func (m *_CreateSubscriptionRequest) GetPublishingEnabled() bool {
	return m.PublishingEnabled
}

func (m *_CreateSubscriptionRequest) GetPriority() uint8 {
	return m.Priority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCreateSubscriptionRequest(structType any) CreateSubscriptionRequest {
	if casted, ok := structType.(CreateSubscriptionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CreateSubscriptionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CreateSubscriptionRequest) GetTypeName() string {
	return "CreateSubscriptionRequest"
}

func (m *_CreateSubscriptionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (requestedPublishingInterval)
	lengthInBits += 64

	// Simple field (requestedLifetimeCount)
	lengthInBits += 32

	// Simple field (requestedMaxKeepAliveCount)
	lengthInBits += 32

	// Simple field (maxNotificationsPerPublish)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (publishingEnabled)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CreateSubscriptionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CreateSubscriptionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__createSubscriptionRequest CreateSubscriptionRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CreateSubscriptionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CreateSubscriptionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	requestedPublishingInterval, err := ReadSimpleField(ctx, "requestedPublishingInterval", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedPublishingInterval' field"))
	}
	m.RequestedPublishingInterval = requestedPublishingInterval

	requestedLifetimeCount, err := ReadSimpleField(ctx, "requestedLifetimeCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedLifetimeCount' field"))
	}
	m.RequestedLifetimeCount = requestedLifetimeCount

	requestedMaxKeepAliveCount, err := ReadSimpleField(ctx, "requestedMaxKeepAliveCount", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedMaxKeepAliveCount' field"))
	}
	m.RequestedMaxKeepAliveCount = requestedMaxKeepAliveCount

	maxNotificationsPerPublish, err := ReadSimpleField(ctx, "maxNotificationsPerPublish", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxNotificationsPerPublish' field"))
	}
	m.MaxNotificationsPerPublish = maxNotificationsPerPublish

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	publishingEnabled, err := ReadSimpleField(ctx, "publishingEnabled", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publishingEnabled' field"))
	}
	m.PublishingEnabled = publishingEnabled

	priority, err := ReadSimpleField(ctx, "priority", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	m.Priority = priority

	if closeErr := readBuffer.CloseContext("CreateSubscriptionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CreateSubscriptionRequest")
	}

	return m, nil
}

func (m *_CreateSubscriptionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CreateSubscriptionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateSubscriptionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CreateSubscriptionRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[float64](ctx, "requestedPublishingInterval", m.GetRequestedPublishingInterval(), WriteDouble(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedPublishingInterval' field")
		}

		if err := WriteSimpleField[uint32](ctx, "requestedLifetimeCount", m.GetRequestedLifetimeCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedLifetimeCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "requestedMaxKeepAliveCount", m.GetRequestedMaxKeepAliveCount(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedMaxKeepAliveCount' field")
		}

		if err := WriteSimpleField[uint32](ctx, "maxNotificationsPerPublish", m.GetMaxNotificationsPerPublish(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxNotificationsPerPublish' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "publishingEnabled", m.GetPublishingEnabled(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'publishingEnabled' field")
		}

		if err := WriteSimpleField[uint8](ctx, "priority", m.GetPriority(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'priority' field")
		}

		if popErr := writeBuffer.PopContext("CreateSubscriptionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CreateSubscriptionRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CreateSubscriptionRequest) IsCreateSubscriptionRequest() {}

func (m *_CreateSubscriptionRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CreateSubscriptionRequest) deepCopy() *_CreateSubscriptionRequest {
	if m == nil {
		return nil
	}
	_CreateSubscriptionRequestCopy := &_CreateSubscriptionRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.RequestedPublishingInterval,
		m.RequestedLifetimeCount,
		m.RequestedMaxKeepAliveCount,
		m.MaxNotificationsPerPublish,
		m.PublishingEnabled,
		m.Priority,
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _CreateSubscriptionRequestCopy
}

func (m *_CreateSubscriptionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

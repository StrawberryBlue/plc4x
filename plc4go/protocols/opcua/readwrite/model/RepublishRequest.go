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

// RepublishRequest is the corresponding interface of RepublishRequest
type RepublishRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetRetransmitSequenceNumber returns RetransmitSequenceNumber (property field)
	GetRetransmitSequenceNumber() uint32
	// IsRepublishRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRepublishRequest()
	// CreateBuilder creates a RepublishRequestBuilder
	CreateRepublishRequestBuilder() RepublishRequestBuilder
}

// _RepublishRequest is the data-structure of this message
type _RepublishRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader            ExtensionObjectDefinition
	SubscriptionId           uint32
	RetransmitSequenceNumber uint32
}

var _ RepublishRequest = (*_RepublishRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RepublishRequest)(nil)

// NewRepublishRequest factory function for _RepublishRequest
func NewRepublishRequest(requestHeader ExtensionObjectDefinition, subscriptionId uint32, retransmitSequenceNumber uint32) *_RepublishRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for RepublishRequest must not be nil")
	}
	_result := &_RepublishRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		SubscriptionId:                    subscriptionId,
		RetransmitSequenceNumber:          retransmitSequenceNumber,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RepublishRequestBuilder is a builder for RepublishRequest
type RepublishRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader ExtensionObjectDefinition, subscriptionId uint32, retransmitSequenceNumber uint32) RepublishRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(ExtensionObjectDefinition) RepublishRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) RepublishRequestBuilder
	// WithSubscriptionId adds SubscriptionId (property field)
	WithSubscriptionId(uint32) RepublishRequestBuilder
	// WithRetransmitSequenceNumber adds RetransmitSequenceNumber (property field)
	WithRetransmitSequenceNumber(uint32) RepublishRequestBuilder
	// Build builds the RepublishRequest or returns an error if something is wrong
	Build() (RepublishRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() RepublishRequest
}

// NewRepublishRequestBuilder() creates a RepublishRequestBuilder
func NewRepublishRequestBuilder() RepublishRequestBuilder {
	return &_RepublishRequestBuilder{_RepublishRequest: new(_RepublishRequest)}
}

type _RepublishRequestBuilder struct {
	*_RepublishRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (RepublishRequestBuilder) = (*_RepublishRequestBuilder)(nil)

func (b *_RepublishRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_RepublishRequestBuilder) WithMandatoryFields(requestHeader ExtensionObjectDefinition, subscriptionId uint32, retransmitSequenceNumber uint32) RepublishRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithSubscriptionId(subscriptionId).WithRetransmitSequenceNumber(retransmitSequenceNumber)
}

func (b *_RepublishRequestBuilder) WithRequestHeader(requestHeader ExtensionObjectDefinition) RepublishRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_RepublishRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) RepublishRequestBuilder {
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

func (b *_RepublishRequestBuilder) WithSubscriptionId(subscriptionId uint32) RepublishRequestBuilder {
	b.SubscriptionId = subscriptionId
	return b
}

func (b *_RepublishRequestBuilder) WithRetransmitSequenceNumber(retransmitSequenceNumber uint32) RepublishRequestBuilder {
	b.RetransmitSequenceNumber = retransmitSequenceNumber
	return b
}

func (b *_RepublishRequestBuilder) Build() (RepublishRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._RepublishRequest.deepCopy(), nil
}

func (b *_RepublishRequestBuilder) MustBuild() RepublishRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_RepublishRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_RepublishRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_RepublishRequestBuilder) DeepCopy() any {
	_copy := b.CreateRepublishRequestBuilder().(*_RepublishRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateRepublishRequestBuilder creates a RepublishRequestBuilder
func (b *_RepublishRequest) CreateRepublishRequestBuilder() RepublishRequestBuilder {
	if b == nil {
		return NewRepublishRequestBuilder()
	}
	return &_RepublishRequestBuilder{_RepublishRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RepublishRequest) GetIdentifier() string {
	return "832"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RepublishRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RepublishRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_RepublishRequest) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_RepublishRequest) GetRetransmitSequenceNumber() uint32 {
	return m.RetransmitSequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRepublishRequest(structType any) RepublishRequest {
	if casted, ok := structType.(RepublishRequest); ok {
		return casted
	}
	if casted, ok := structType.(*RepublishRequest); ok {
		return *casted
	}
	return nil
}

func (m *_RepublishRequest) GetTypeName() string {
	return "RepublishRequest"
}

func (m *_RepublishRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (retransmitSequenceNumber)
	lengthInBits += 32

	return lengthInBits
}

func (m *_RepublishRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RepublishRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__republishRequest RepublishRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RepublishRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RepublishRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	subscriptionId, err := ReadSimpleField(ctx, "subscriptionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionId' field"))
	}
	m.SubscriptionId = subscriptionId

	retransmitSequenceNumber, err := ReadSimpleField(ctx, "retransmitSequenceNumber", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'retransmitSequenceNumber' field"))
	}
	m.RetransmitSequenceNumber = retransmitSequenceNumber

	if closeErr := readBuffer.CloseContext("RepublishRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RepublishRequest")
	}

	return m, nil
}

func (m *_RepublishRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RepublishRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RepublishRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RepublishRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "subscriptionId", m.GetSubscriptionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "retransmitSequenceNumber", m.GetRetransmitSequenceNumber(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'retransmitSequenceNumber' field")
		}

		if popErr := writeBuffer.PopContext("RepublishRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RepublishRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RepublishRequest) IsRepublishRequest() {}

func (m *_RepublishRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RepublishRequest) deepCopy() *_RepublishRequest {
	if m == nil {
		return nil
	}
	_RepublishRequestCopy := &_RepublishRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.SubscriptionId,
		m.RetransmitSequenceNumber,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _RepublishRequestCopy
}

func (m *_RepublishRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

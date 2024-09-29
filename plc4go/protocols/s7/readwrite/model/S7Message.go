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

// Constant values.
const S7Message_PROTOCOLID uint8 = 0x32

// S7Message is the corresponding interface of S7Message
type S7Message interface {
	S7MessageContract
	S7MessageRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsS7Message is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7Message()
	// CreateBuilder creates a S7MessageBuilder
	CreateS7MessageBuilder() S7MessageBuilder
}

// S7MessageContract provides a set of functions which can be overwritten by a sub struct
type S7MessageContract interface {
	// GetTpduReference returns TpduReference (property field)
	GetTpduReference() uint16
	// GetParameter returns Parameter (property field)
	GetParameter() S7Parameter
	// GetPayload returns Payload (property field)
	GetPayload() S7Payload
	// IsS7Message is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7Message()
	// CreateBuilder creates a S7MessageBuilder
	CreateS7MessageBuilder() S7MessageBuilder
}

// S7MessageRequirements provides a set of functions which need to be implemented by a sub struct
type S7MessageRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() uint8
}

// _S7Message is the data-structure of this message
type _S7Message struct {
	_SubType      S7Message
	TpduReference uint16
	Parameter     S7Parameter
	Payload       S7Payload
	// Reserved Fields
	reservedField0 *uint16
}

var _ S7MessageContract = (*_S7Message)(nil)

// NewS7Message factory function for _S7Message
func NewS7Message(tpduReference uint16, parameter S7Parameter, payload S7Payload) *_S7Message {
	return &_S7Message{TpduReference: tpduReference, Parameter: parameter, Payload: payload}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7MessageBuilder is a builder for S7Message
type S7MessageBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(tpduReference uint16) S7MessageBuilder
	// WithTpduReference adds TpduReference (property field)
	WithTpduReference(uint16) S7MessageBuilder
	// WithParameter adds Parameter (property field)
	WithOptionalParameter(S7Parameter) S7MessageBuilder
	// WithOptionalParameterBuilder adds Parameter (property field) which is build by the builder
	WithOptionalParameterBuilder(func(S7ParameterBuilder) S7ParameterBuilder) S7MessageBuilder
	// WithPayload adds Payload (property field)
	WithOptionalPayload(S7Payload) S7MessageBuilder
	// WithOptionalPayloadBuilder adds Payload (property field) which is build by the builder
	WithOptionalPayloadBuilder(func(S7PayloadBuilder) S7PayloadBuilder) S7MessageBuilder
	// AsS7MessageRequest converts this build to a subType of S7Message. It is always possible to return to current builder using Done()
	AsS7MessageRequest() interface {
		S7MessageRequestBuilder
		Done() S7MessageBuilder
	}
	// AsS7MessageResponse converts this build to a subType of S7Message. It is always possible to return to current builder using Done()
	AsS7MessageResponse() interface {
		S7MessageResponseBuilder
		Done() S7MessageBuilder
	}
	// AsS7MessageResponseData converts this build to a subType of S7Message. It is always possible to return to current builder using Done()
	AsS7MessageResponseData() interface {
		S7MessageResponseDataBuilder
		Done() S7MessageBuilder
	}
	// AsS7MessageUserData converts this build to a subType of S7Message. It is always possible to return to current builder using Done()
	AsS7MessageUserData() interface {
		S7MessageUserDataBuilder
		Done() S7MessageBuilder
	}
	// Build builds the S7Message or returns an error if something is wrong
	PartialBuild() (S7MessageContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() S7MessageContract
	// Build builds the S7Message or returns an error if something is wrong
	Build() (S7Message, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7Message
}

// NewS7MessageBuilder() creates a S7MessageBuilder
func NewS7MessageBuilder() S7MessageBuilder {
	return &_S7MessageBuilder{_S7Message: new(_S7Message)}
}

type _S7MessageChildBuilder interface {
	utils.Copyable
	setParent(S7MessageContract)
	buildForS7Message() (S7Message, error)
}

type _S7MessageBuilder struct {
	*_S7Message

	childBuilder _S7MessageChildBuilder

	err *utils.MultiError
}

var _ (S7MessageBuilder) = (*_S7MessageBuilder)(nil)

func (b *_S7MessageBuilder) WithMandatoryFields(tpduReference uint16) S7MessageBuilder {
	return b.WithTpduReference(tpduReference)
}

func (b *_S7MessageBuilder) WithTpduReference(tpduReference uint16) S7MessageBuilder {
	b.TpduReference = tpduReference
	return b
}

func (b *_S7MessageBuilder) WithOptionalParameter(parameter S7Parameter) S7MessageBuilder {
	b.Parameter = parameter
	return b
}

func (b *_S7MessageBuilder) WithOptionalParameterBuilder(builderSupplier func(S7ParameterBuilder) S7ParameterBuilder) S7MessageBuilder {
	builder := builderSupplier(b.Parameter.CreateS7ParameterBuilder())
	var err error
	b.Parameter, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "S7ParameterBuilder failed"))
	}
	return b
}

func (b *_S7MessageBuilder) WithOptionalPayload(payload S7Payload) S7MessageBuilder {
	b.Payload = payload
	return b
}

func (b *_S7MessageBuilder) WithOptionalPayloadBuilder(builderSupplier func(S7PayloadBuilder) S7PayloadBuilder) S7MessageBuilder {
	builder := builderSupplier(b.Payload.CreateS7PayloadBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "S7PayloadBuilder failed"))
	}
	return b
}

func (b *_S7MessageBuilder) PartialBuild() (S7MessageContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7Message.deepCopy(), nil
}

func (b *_S7MessageBuilder) PartialMustBuild() S7MessageContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7MessageBuilder) AsS7MessageRequest() interface {
	S7MessageRequestBuilder
	Done() S7MessageBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		S7MessageRequestBuilder
		Done() S7MessageBuilder
	}); ok {
		return cb
	}
	cb := NewS7MessageRequestBuilder().(*_S7MessageRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_S7MessageBuilder) AsS7MessageResponse() interface {
	S7MessageResponseBuilder
	Done() S7MessageBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		S7MessageResponseBuilder
		Done() S7MessageBuilder
	}); ok {
		return cb
	}
	cb := NewS7MessageResponseBuilder().(*_S7MessageResponseBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_S7MessageBuilder) AsS7MessageResponseData() interface {
	S7MessageResponseDataBuilder
	Done() S7MessageBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		S7MessageResponseDataBuilder
		Done() S7MessageBuilder
	}); ok {
		return cb
	}
	cb := NewS7MessageResponseDataBuilder().(*_S7MessageResponseDataBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_S7MessageBuilder) AsS7MessageUserData() interface {
	S7MessageUserDataBuilder
	Done() S7MessageBuilder
} {
	if cb, ok := b.childBuilder.(interface {
		S7MessageUserDataBuilder
		Done() S7MessageBuilder
	}); ok {
		return cb
	}
	cb := NewS7MessageUserDataBuilder().(*_S7MessageUserDataBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_S7MessageBuilder) Build() (S7Message, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForS7Message()
}

func (b *_S7MessageBuilder) MustBuild() S7Message {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7MessageBuilder) DeepCopy() any {
	_copy := b.CreateS7MessageBuilder().(*_S7MessageBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_S7MessageChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7MessageBuilder creates a S7MessageBuilder
func (b *_S7Message) CreateS7MessageBuilder() S7MessageBuilder {
	if b == nil {
		return NewS7MessageBuilder()
	}
	return &_S7MessageBuilder{_S7Message: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7Message) GetTpduReference() uint16 {
	return m.TpduReference
}

func (m *_S7Message) GetParameter() S7Parameter {
	return m.Parameter
}

func (m *_S7Message) GetPayload() S7Payload {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_S7Message) GetProtocolId() uint8 {
	return S7Message_PROTOCOLID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7Message(structType any) S7Message {
	if casted, ok := structType.(S7Message); ok {
		return casted
	}
	if casted, ok := structType.(*S7Message); ok {
		return *casted
	}
	return nil
}

func (m *_S7Message) GetTypeName() string {
	return "S7Message"
}

func (m *_S7Message) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (protocolId)
	lengthInBits += 8
	// Discriminator Field (messageType)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 16

	// Simple field (tpduReference)
	lengthInBits += 16

	// Implicit Field (parameterLength)
	lengthInBits += 16

	// Implicit Field (payloadLength)
	lengthInBits += 16

	// Optional Field (parameter)
	if m.Parameter != nil {
		lengthInBits += m.Parameter.GetLengthInBits(ctx)
	}

	// Optional Field (payload)
	if m.Payload != nil {
		lengthInBits += m.Payload.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_S7Message) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func S7MessageParse[T S7Message](ctx context.Context, theBytes []byte) (T, error) {
	return S7MessageParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func S7MessageParseWithBufferProducer[T S7Message]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := S7MessageParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func S7MessageParseWithBuffer[T S7Message](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_S7Message{}).parse(ctx, readBuffer)
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

func (m *_S7Message) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__s7Message S7Message, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7Message"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7Message")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	protocolId, err := ReadConstField[uint8](ctx, "protocolId", ReadUnsignedByte(readBuffer, uint8(8)), S7Message_PROTOCOLID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolId' field"))
	}
	_ = protocolId

	messageType, err := ReadDiscriminatorField[uint8](ctx, "messageType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageType' field"))
	}

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedShort(readBuffer, uint8(16)), uint16(0x0000))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	tpduReference, err := ReadSimpleField(ctx, "tpduReference", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tpduReference' field"))
	}
	m.TpduReference = tpduReference

	parameterLength, err := ReadImplicitField[uint16](ctx, "parameterLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameterLength' field"))
	}
	_ = parameterLength

	payloadLength, err := ReadImplicitField[uint16](ctx, "payloadLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payloadLength' field"))
	}
	_ = payloadLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child S7Message
	switch {
	case messageType == 0x01: // S7MessageRequest
		if _child, err = new(_S7MessageRequest).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7MessageRequest for type-switch of S7Message")
		}
	case messageType == 0x02: // S7MessageResponse
		if _child, err = new(_S7MessageResponse).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7MessageResponse for type-switch of S7Message")
		}
	case messageType == 0x03: // S7MessageResponseData
		if _child, err = new(_S7MessageResponseData).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7MessageResponseData for type-switch of S7Message")
		}
	case messageType == 0x07: // S7MessageUserData
		if _child, err = new(_S7MessageUserData).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7MessageUserData for type-switch of S7Message")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [messageType=%v]", messageType)
	}

	var parameter S7Parameter
	_parameter, err := ReadOptionalField[S7Parameter](ctx, "parameter", ReadComplex[S7Parameter](S7ParameterParseWithBufferProducer[S7Parameter]((uint8)(messageType)), readBuffer), bool((parameterLength) > (0)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameter' field"))
	}
	if _parameter != nil {
		parameter = *_parameter
		m.Parameter = parameter
	}

	var payload S7Payload
	_payload, err := ReadOptionalField[S7Payload](ctx, "payload", ReadComplex[S7Payload](S7PayloadParseWithBufferProducer[S7Payload]((uint8)(messageType), (S7Parameter)((parameter))), readBuffer), bool((payloadLength) > (0)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	if _payload != nil {
		payload = *_payload
		m.Payload = payload
	}

	if closeErr := readBuffer.CloseContext("S7Message"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7Message")
	}

	return _child, nil
}

func (pm *_S7Message) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7Message, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("S7Message"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7Message")
	}

	if err := WriteConstField(ctx, "protocolId", S7Message_PROTOCOLID, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'protocolId' field")
	}

	if err := WriteDiscriminatorField(ctx, "messageType", m.GetMessageType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'messageType' field")
	}

	if err := WriteReservedField[uint16](ctx, "reserved", uint16(0x0000), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteSimpleField[uint16](ctx, "tpduReference", m.GetTpduReference(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'tpduReference' field")
	}
	parameterLength := uint16(utils.InlineIf(bool((m.GetParameter()) != (nil)), func() any { return uint16((m.GetParameter()).GetLengthInBytes(ctx)) }, func() any { return uint16(uint16(0)) }).(uint16))
	if err := WriteImplicitField(ctx, "parameterLength", parameterLength, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'parameterLength' field")
	}
	payloadLength := uint16(utils.InlineIf(bool((m.GetPayload()) != (nil)), func() any { return uint16((m.GetPayload()).GetLengthInBytes(ctx)) }, func() any { return uint16(uint16(0)) }).(uint16))
	if err := WriteImplicitField(ctx, "payloadLength", payloadLength, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'payloadLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteOptionalField[S7Parameter](ctx, "parameter", GetRef(m.GetParameter()), WriteComplex[S7Parameter](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'parameter' field")
	}

	if err := WriteOptionalField[S7Payload](ctx, "payload", GetRef(m.GetPayload()), WriteComplex[S7Payload](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'payload' field")
	}

	if popErr := writeBuffer.PopContext("S7Message"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7Message")
	}
	return nil
}

func (m *_S7Message) IsS7Message() {}

func (m *_S7Message) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7Message) deepCopy() *_S7Message {
	if m == nil {
		return nil
	}
	_S7MessageCopy := &_S7Message{
		nil, // will be set by child
		m.TpduReference,
		m.Parameter.DeepCopy().(S7Parameter),
		m.Payload.DeepCopy().(S7Payload),
		m.reservedField0,
	}
	return _S7MessageCopy
}

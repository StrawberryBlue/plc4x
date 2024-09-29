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

// EndpointConfiguration is the corresponding interface of EndpointConfiguration
type EndpointConfiguration interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetOperationTimeout returns OperationTimeout (property field)
	GetOperationTimeout() int32
	// GetUseBinaryEncoding returns UseBinaryEncoding (property field)
	GetUseBinaryEncoding() bool
	// GetMaxStringLength returns MaxStringLength (property field)
	GetMaxStringLength() int32
	// GetMaxByteStringLength returns MaxByteStringLength (property field)
	GetMaxByteStringLength() int32
	// GetMaxArrayLength returns MaxArrayLength (property field)
	GetMaxArrayLength() int32
	// GetMaxMessageSize returns MaxMessageSize (property field)
	GetMaxMessageSize() int32
	// GetMaxBufferSize returns MaxBufferSize (property field)
	GetMaxBufferSize() int32
	// GetChannelLifetime returns ChannelLifetime (property field)
	GetChannelLifetime() int32
	// GetSecurityTokenLifetime returns SecurityTokenLifetime (property field)
	GetSecurityTokenLifetime() int32
	// IsEndpointConfiguration is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEndpointConfiguration()
	// CreateBuilder creates a EndpointConfigurationBuilder
	CreateEndpointConfigurationBuilder() EndpointConfigurationBuilder
}

// _EndpointConfiguration is the data-structure of this message
type _EndpointConfiguration struct {
	ExtensionObjectDefinitionContract
	OperationTimeout      int32
	UseBinaryEncoding     bool
	MaxStringLength       int32
	MaxByteStringLength   int32
	MaxArrayLength        int32
	MaxMessageSize        int32
	MaxBufferSize         int32
	ChannelLifetime       int32
	SecurityTokenLifetime int32
	// Reserved Fields
	reservedField0 *uint8
}

var _ EndpointConfiguration = (*_EndpointConfiguration)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_EndpointConfiguration)(nil)

// NewEndpointConfiguration factory function for _EndpointConfiguration
func NewEndpointConfiguration(operationTimeout int32, useBinaryEncoding bool, maxStringLength int32, maxByteStringLength int32, maxArrayLength int32, maxMessageSize int32, maxBufferSize int32, channelLifetime int32, securityTokenLifetime int32) *_EndpointConfiguration {
	_result := &_EndpointConfiguration{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		OperationTimeout:                  operationTimeout,
		UseBinaryEncoding:                 useBinaryEncoding,
		MaxStringLength:                   maxStringLength,
		MaxByteStringLength:               maxByteStringLength,
		MaxArrayLength:                    maxArrayLength,
		MaxMessageSize:                    maxMessageSize,
		MaxBufferSize:                     maxBufferSize,
		ChannelLifetime:                   channelLifetime,
		SecurityTokenLifetime:             securityTokenLifetime,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// EndpointConfigurationBuilder is a builder for EndpointConfiguration
type EndpointConfigurationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(operationTimeout int32, useBinaryEncoding bool, maxStringLength int32, maxByteStringLength int32, maxArrayLength int32, maxMessageSize int32, maxBufferSize int32, channelLifetime int32, securityTokenLifetime int32) EndpointConfigurationBuilder
	// WithOperationTimeout adds OperationTimeout (property field)
	WithOperationTimeout(int32) EndpointConfigurationBuilder
	// WithUseBinaryEncoding adds UseBinaryEncoding (property field)
	WithUseBinaryEncoding(bool) EndpointConfigurationBuilder
	// WithMaxStringLength adds MaxStringLength (property field)
	WithMaxStringLength(int32) EndpointConfigurationBuilder
	// WithMaxByteStringLength adds MaxByteStringLength (property field)
	WithMaxByteStringLength(int32) EndpointConfigurationBuilder
	// WithMaxArrayLength adds MaxArrayLength (property field)
	WithMaxArrayLength(int32) EndpointConfigurationBuilder
	// WithMaxMessageSize adds MaxMessageSize (property field)
	WithMaxMessageSize(int32) EndpointConfigurationBuilder
	// WithMaxBufferSize adds MaxBufferSize (property field)
	WithMaxBufferSize(int32) EndpointConfigurationBuilder
	// WithChannelLifetime adds ChannelLifetime (property field)
	WithChannelLifetime(int32) EndpointConfigurationBuilder
	// WithSecurityTokenLifetime adds SecurityTokenLifetime (property field)
	WithSecurityTokenLifetime(int32) EndpointConfigurationBuilder
	// Build builds the EndpointConfiguration or returns an error if something is wrong
	Build() (EndpointConfiguration, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() EndpointConfiguration
}

// NewEndpointConfigurationBuilder() creates a EndpointConfigurationBuilder
func NewEndpointConfigurationBuilder() EndpointConfigurationBuilder {
	return &_EndpointConfigurationBuilder{_EndpointConfiguration: new(_EndpointConfiguration)}
}

type _EndpointConfigurationBuilder struct {
	*_EndpointConfiguration

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (EndpointConfigurationBuilder) = (*_EndpointConfigurationBuilder)(nil)

func (b *_EndpointConfigurationBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_EndpointConfigurationBuilder) WithMandatoryFields(operationTimeout int32, useBinaryEncoding bool, maxStringLength int32, maxByteStringLength int32, maxArrayLength int32, maxMessageSize int32, maxBufferSize int32, channelLifetime int32, securityTokenLifetime int32) EndpointConfigurationBuilder {
	return b.WithOperationTimeout(operationTimeout).WithUseBinaryEncoding(useBinaryEncoding).WithMaxStringLength(maxStringLength).WithMaxByteStringLength(maxByteStringLength).WithMaxArrayLength(maxArrayLength).WithMaxMessageSize(maxMessageSize).WithMaxBufferSize(maxBufferSize).WithChannelLifetime(channelLifetime).WithSecurityTokenLifetime(securityTokenLifetime)
}

func (b *_EndpointConfigurationBuilder) WithOperationTimeout(operationTimeout int32) EndpointConfigurationBuilder {
	b.OperationTimeout = operationTimeout
	return b
}

func (b *_EndpointConfigurationBuilder) WithUseBinaryEncoding(useBinaryEncoding bool) EndpointConfigurationBuilder {
	b.UseBinaryEncoding = useBinaryEncoding
	return b
}

func (b *_EndpointConfigurationBuilder) WithMaxStringLength(maxStringLength int32) EndpointConfigurationBuilder {
	b.MaxStringLength = maxStringLength
	return b
}

func (b *_EndpointConfigurationBuilder) WithMaxByteStringLength(maxByteStringLength int32) EndpointConfigurationBuilder {
	b.MaxByteStringLength = maxByteStringLength
	return b
}

func (b *_EndpointConfigurationBuilder) WithMaxArrayLength(maxArrayLength int32) EndpointConfigurationBuilder {
	b.MaxArrayLength = maxArrayLength
	return b
}

func (b *_EndpointConfigurationBuilder) WithMaxMessageSize(maxMessageSize int32) EndpointConfigurationBuilder {
	b.MaxMessageSize = maxMessageSize
	return b
}

func (b *_EndpointConfigurationBuilder) WithMaxBufferSize(maxBufferSize int32) EndpointConfigurationBuilder {
	b.MaxBufferSize = maxBufferSize
	return b
}

func (b *_EndpointConfigurationBuilder) WithChannelLifetime(channelLifetime int32) EndpointConfigurationBuilder {
	b.ChannelLifetime = channelLifetime
	return b
}

func (b *_EndpointConfigurationBuilder) WithSecurityTokenLifetime(securityTokenLifetime int32) EndpointConfigurationBuilder {
	b.SecurityTokenLifetime = securityTokenLifetime
	return b
}

func (b *_EndpointConfigurationBuilder) Build() (EndpointConfiguration, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._EndpointConfiguration.deepCopy(), nil
}

func (b *_EndpointConfigurationBuilder) MustBuild() EndpointConfiguration {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_EndpointConfigurationBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_EndpointConfigurationBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_EndpointConfigurationBuilder) DeepCopy() any {
	_copy := b.CreateEndpointConfigurationBuilder().(*_EndpointConfigurationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateEndpointConfigurationBuilder creates a EndpointConfigurationBuilder
func (b *_EndpointConfiguration) CreateEndpointConfigurationBuilder() EndpointConfigurationBuilder {
	if b == nil {
		return NewEndpointConfigurationBuilder()
	}
	return &_EndpointConfigurationBuilder{_EndpointConfiguration: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EndpointConfiguration) GetIdentifier() string {
	return "333"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EndpointConfiguration) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EndpointConfiguration) GetOperationTimeout() int32 {
	return m.OperationTimeout
}

func (m *_EndpointConfiguration) GetUseBinaryEncoding() bool {
	return m.UseBinaryEncoding
}

func (m *_EndpointConfiguration) GetMaxStringLength() int32 {
	return m.MaxStringLength
}

func (m *_EndpointConfiguration) GetMaxByteStringLength() int32 {
	return m.MaxByteStringLength
}

func (m *_EndpointConfiguration) GetMaxArrayLength() int32 {
	return m.MaxArrayLength
}

func (m *_EndpointConfiguration) GetMaxMessageSize() int32 {
	return m.MaxMessageSize
}

func (m *_EndpointConfiguration) GetMaxBufferSize() int32 {
	return m.MaxBufferSize
}

func (m *_EndpointConfiguration) GetChannelLifetime() int32 {
	return m.ChannelLifetime
}

func (m *_EndpointConfiguration) GetSecurityTokenLifetime() int32 {
	return m.SecurityTokenLifetime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastEndpointConfiguration(structType any) EndpointConfiguration {
	if casted, ok := structType.(EndpointConfiguration); ok {
		return casted
	}
	if casted, ok := structType.(*EndpointConfiguration); ok {
		return *casted
	}
	return nil
}

func (m *_EndpointConfiguration) GetTypeName() string {
	return "EndpointConfiguration"
}

func (m *_EndpointConfiguration) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (operationTimeout)
	lengthInBits += 32

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (useBinaryEncoding)
	lengthInBits += 1

	// Simple field (maxStringLength)
	lengthInBits += 32

	// Simple field (maxByteStringLength)
	lengthInBits += 32

	// Simple field (maxArrayLength)
	lengthInBits += 32

	// Simple field (maxMessageSize)
	lengthInBits += 32

	// Simple field (maxBufferSize)
	lengthInBits += 32

	// Simple field (channelLifetime)
	lengthInBits += 32

	// Simple field (securityTokenLifetime)
	lengthInBits += 32

	return lengthInBits
}

func (m *_EndpointConfiguration) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EndpointConfiguration) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__endpointConfiguration EndpointConfiguration, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EndpointConfiguration"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EndpointConfiguration")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	operationTimeout, err := ReadSimpleField(ctx, "operationTimeout", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operationTimeout' field"))
	}
	m.OperationTimeout = operationTimeout

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	useBinaryEncoding, err := ReadSimpleField(ctx, "useBinaryEncoding", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'useBinaryEncoding' field"))
	}
	m.UseBinaryEncoding = useBinaryEncoding

	maxStringLength, err := ReadSimpleField(ctx, "maxStringLength", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxStringLength' field"))
	}
	m.MaxStringLength = maxStringLength

	maxByteStringLength, err := ReadSimpleField(ctx, "maxByteStringLength", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxByteStringLength' field"))
	}
	m.MaxByteStringLength = maxByteStringLength

	maxArrayLength, err := ReadSimpleField(ctx, "maxArrayLength", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxArrayLength' field"))
	}
	m.MaxArrayLength = maxArrayLength

	maxMessageSize, err := ReadSimpleField(ctx, "maxMessageSize", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxMessageSize' field"))
	}
	m.MaxMessageSize = maxMessageSize

	maxBufferSize, err := ReadSimpleField(ctx, "maxBufferSize", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxBufferSize' field"))
	}
	m.MaxBufferSize = maxBufferSize

	channelLifetime, err := ReadSimpleField(ctx, "channelLifetime", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelLifetime' field"))
	}
	m.ChannelLifetime = channelLifetime

	securityTokenLifetime, err := ReadSimpleField(ctx, "securityTokenLifetime", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityTokenLifetime' field"))
	}
	m.SecurityTokenLifetime = securityTokenLifetime

	if closeErr := readBuffer.CloseContext("EndpointConfiguration"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EndpointConfiguration")
	}

	return m, nil
}

func (m *_EndpointConfiguration) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EndpointConfiguration) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EndpointConfiguration"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EndpointConfiguration")
		}

		if err := WriteSimpleField[int32](ctx, "operationTimeout", m.GetOperationTimeout(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'operationTimeout' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "useBinaryEncoding", m.GetUseBinaryEncoding(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'useBinaryEncoding' field")
		}

		if err := WriteSimpleField[int32](ctx, "maxStringLength", m.GetMaxStringLength(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxStringLength' field")
		}

		if err := WriteSimpleField[int32](ctx, "maxByteStringLength", m.GetMaxByteStringLength(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxByteStringLength' field")
		}

		if err := WriteSimpleField[int32](ctx, "maxArrayLength", m.GetMaxArrayLength(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxArrayLength' field")
		}

		if err := WriteSimpleField[int32](ctx, "maxMessageSize", m.GetMaxMessageSize(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxMessageSize' field")
		}

		if err := WriteSimpleField[int32](ctx, "maxBufferSize", m.GetMaxBufferSize(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxBufferSize' field")
		}

		if err := WriteSimpleField[int32](ctx, "channelLifetime", m.GetChannelLifetime(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'channelLifetime' field")
		}

		if err := WriteSimpleField[int32](ctx, "securityTokenLifetime", m.GetSecurityTokenLifetime(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityTokenLifetime' field")
		}

		if popErr := writeBuffer.PopContext("EndpointConfiguration"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EndpointConfiguration")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EndpointConfiguration) IsEndpointConfiguration() {}

func (m *_EndpointConfiguration) DeepCopy() any {
	return m.deepCopy()
}

func (m *_EndpointConfiguration) deepCopy() *_EndpointConfiguration {
	if m == nil {
		return nil
	}
	_EndpointConfigurationCopy := &_EndpointConfiguration{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.OperationTimeout,
		m.UseBinaryEncoding,
		m.MaxStringLength,
		m.MaxByteStringLength,
		m.MaxArrayLength,
		m.MaxMessageSize,
		m.MaxBufferSize,
		m.ChannelLifetime,
		m.SecurityTokenLifetime,
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _EndpointConfigurationCopy
}

func (m *_EndpointConfiguration) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

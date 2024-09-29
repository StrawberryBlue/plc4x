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

// S7ParameterSetupCommunication is the corresponding interface of S7ParameterSetupCommunication
type S7ParameterSetupCommunication interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7Parameter
	// GetMaxAmqCaller returns MaxAmqCaller (property field)
	GetMaxAmqCaller() uint16
	// GetMaxAmqCallee returns MaxAmqCallee (property field)
	GetMaxAmqCallee() uint16
	// GetPduLength returns PduLength (property field)
	GetPduLength() uint16
	// IsS7ParameterSetupCommunication is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7ParameterSetupCommunication()
	// CreateBuilder creates a S7ParameterSetupCommunicationBuilder
	CreateS7ParameterSetupCommunicationBuilder() S7ParameterSetupCommunicationBuilder
}

// _S7ParameterSetupCommunication is the data-structure of this message
type _S7ParameterSetupCommunication struct {
	S7ParameterContract
	MaxAmqCaller uint16
	MaxAmqCallee uint16
	PduLength    uint16
	// Reserved Fields
	reservedField0 *uint8
}

var _ S7ParameterSetupCommunication = (*_S7ParameterSetupCommunication)(nil)
var _ S7ParameterRequirements = (*_S7ParameterSetupCommunication)(nil)

// NewS7ParameterSetupCommunication factory function for _S7ParameterSetupCommunication
func NewS7ParameterSetupCommunication(maxAmqCaller uint16, maxAmqCallee uint16, pduLength uint16) *_S7ParameterSetupCommunication {
	_result := &_S7ParameterSetupCommunication{
		S7ParameterContract: NewS7Parameter(),
		MaxAmqCaller:        maxAmqCaller,
		MaxAmqCallee:        maxAmqCallee,
		PduLength:           pduLength,
	}
	_result.S7ParameterContract.(*_S7Parameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7ParameterSetupCommunicationBuilder is a builder for S7ParameterSetupCommunication
type S7ParameterSetupCommunicationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maxAmqCaller uint16, maxAmqCallee uint16, pduLength uint16) S7ParameterSetupCommunicationBuilder
	// WithMaxAmqCaller adds MaxAmqCaller (property field)
	WithMaxAmqCaller(uint16) S7ParameterSetupCommunicationBuilder
	// WithMaxAmqCallee adds MaxAmqCallee (property field)
	WithMaxAmqCallee(uint16) S7ParameterSetupCommunicationBuilder
	// WithPduLength adds PduLength (property field)
	WithPduLength(uint16) S7ParameterSetupCommunicationBuilder
	// Build builds the S7ParameterSetupCommunication or returns an error if something is wrong
	Build() (S7ParameterSetupCommunication, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7ParameterSetupCommunication
}

// NewS7ParameterSetupCommunicationBuilder() creates a S7ParameterSetupCommunicationBuilder
func NewS7ParameterSetupCommunicationBuilder() S7ParameterSetupCommunicationBuilder {
	return &_S7ParameterSetupCommunicationBuilder{_S7ParameterSetupCommunication: new(_S7ParameterSetupCommunication)}
}

type _S7ParameterSetupCommunicationBuilder struct {
	*_S7ParameterSetupCommunication

	parentBuilder *_S7ParameterBuilder

	err *utils.MultiError
}

var _ (S7ParameterSetupCommunicationBuilder) = (*_S7ParameterSetupCommunicationBuilder)(nil)

func (b *_S7ParameterSetupCommunicationBuilder) setParent(contract S7ParameterContract) {
	b.S7ParameterContract = contract
}

func (b *_S7ParameterSetupCommunicationBuilder) WithMandatoryFields(maxAmqCaller uint16, maxAmqCallee uint16, pduLength uint16) S7ParameterSetupCommunicationBuilder {
	return b.WithMaxAmqCaller(maxAmqCaller).WithMaxAmqCallee(maxAmqCallee).WithPduLength(pduLength)
}

func (b *_S7ParameterSetupCommunicationBuilder) WithMaxAmqCaller(maxAmqCaller uint16) S7ParameterSetupCommunicationBuilder {
	b.MaxAmqCaller = maxAmqCaller
	return b
}

func (b *_S7ParameterSetupCommunicationBuilder) WithMaxAmqCallee(maxAmqCallee uint16) S7ParameterSetupCommunicationBuilder {
	b.MaxAmqCallee = maxAmqCallee
	return b
}

func (b *_S7ParameterSetupCommunicationBuilder) WithPduLength(pduLength uint16) S7ParameterSetupCommunicationBuilder {
	b.PduLength = pduLength
	return b
}

func (b *_S7ParameterSetupCommunicationBuilder) Build() (S7ParameterSetupCommunication, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7ParameterSetupCommunication.deepCopy(), nil
}

func (b *_S7ParameterSetupCommunicationBuilder) MustBuild() S7ParameterSetupCommunication {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_S7ParameterSetupCommunicationBuilder) Done() S7ParameterBuilder {
	return b.parentBuilder
}

func (b *_S7ParameterSetupCommunicationBuilder) buildForS7Parameter() (S7Parameter, error) {
	return b.Build()
}

func (b *_S7ParameterSetupCommunicationBuilder) DeepCopy() any {
	_copy := b.CreateS7ParameterSetupCommunicationBuilder().(*_S7ParameterSetupCommunicationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7ParameterSetupCommunicationBuilder creates a S7ParameterSetupCommunicationBuilder
func (b *_S7ParameterSetupCommunication) CreateS7ParameterSetupCommunicationBuilder() S7ParameterSetupCommunicationBuilder {
	if b == nil {
		return NewS7ParameterSetupCommunicationBuilder()
	}
	return &_S7ParameterSetupCommunicationBuilder{_S7ParameterSetupCommunication: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterSetupCommunication) GetParameterType() uint8 {
	return 0xF0
}

func (m *_S7ParameterSetupCommunication) GetMessageType() uint8 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterSetupCommunication) GetParent() S7ParameterContract {
	return m.S7ParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterSetupCommunication) GetMaxAmqCaller() uint16 {
	return m.MaxAmqCaller
}

func (m *_S7ParameterSetupCommunication) GetMaxAmqCallee() uint16 {
	return m.MaxAmqCallee
}

func (m *_S7ParameterSetupCommunication) GetPduLength() uint16 {
	return m.PduLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7ParameterSetupCommunication(structType any) S7ParameterSetupCommunication {
	if casted, ok := structType.(S7ParameterSetupCommunication); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterSetupCommunication); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterSetupCommunication) GetTypeName() string {
	return "S7ParameterSetupCommunication"
}

func (m *_S7ParameterSetupCommunication) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7ParameterContract.(*_S7Parameter).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (maxAmqCaller)
	lengthInBits += 16

	// Simple field (maxAmqCallee)
	lengthInBits += 16

	// Simple field (pduLength)
	lengthInBits += 16

	return lengthInBits
}

func (m *_S7ParameterSetupCommunication) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7ParameterSetupCommunication) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7Parameter, messageType uint8) (__s7ParameterSetupCommunication S7ParameterSetupCommunication, err error) {
	m.S7ParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterSetupCommunication"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterSetupCommunication")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	maxAmqCaller, err := ReadSimpleField(ctx, "maxAmqCaller", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxAmqCaller' field"))
	}
	m.MaxAmqCaller = maxAmqCaller

	maxAmqCallee, err := ReadSimpleField(ctx, "maxAmqCallee", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxAmqCallee' field"))
	}
	m.MaxAmqCallee = maxAmqCallee

	pduLength, err := ReadSimpleField(ctx, "pduLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pduLength' field"))
	}
	m.PduLength = pduLength

	if closeErr := readBuffer.CloseContext("S7ParameterSetupCommunication"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterSetupCommunication")
	}

	return m, nil
}

func (m *_S7ParameterSetupCommunication) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterSetupCommunication) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterSetupCommunication"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterSetupCommunication")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint16](ctx, "maxAmqCaller", m.GetMaxAmqCaller(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxAmqCaller' field")
		}

		if err := WriteSimpleField[uint16](ctx, "maxAmqCallee", m.GetMaxAmqCallee(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxAmqCallee' field")
		}

		if err := WriteSimpleField[uint16](ctx, "pduLength", m.GetPduLength(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'pduLength' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterSetupCommunication"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterSetupCommunication")
		}
		return nil
	}
	return m.S7ParameterContract.(*_S7Parameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7ParameterSetupCommunication) IsS7ParameterSetupCommunication() {}

func (m *_S7ParameterSetupCommunication) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7ParameterSetupCommunication) deepCopy() *_S7ParameterSetupCommunication {
	if m == nil {
		return nil
	}
	_S7ParameterSetupCommunicationCopy := &_S7ParameterSetupCommunication{
		m.S7ParameterContract.(*_S7Parameter).deepCopy(),
		m.MaxAmqCaller,
		m.MaxAmqCallee,
		m.PduLength,
		m.reservedField0,
	}
	m.S7ParameterContract.(*_S7Parameter)._SubType = m
	return _S7ParameterSetupCommunicationCopy
}

func (m *_S7ParameterSetupCommunication) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

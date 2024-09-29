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

// ModbusPDUWriteSingleRegisterResponse is the corresponding interface of ModbusPDUWriteSingleRegisterResponse
type ModbusPDUWriteSingleRegisterResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// GetValue returns Value (property field)
	GetValue() uint16
	// IsModbusPDUWriteSingleRegisterResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUWriteSingleRegisterResponse()
	// CreateBuilder creates a ModbusPDUWriteSingleRegisterResponseBuilder
	CreateModbusPDUWriteSingleRegisterResponseBuilder() ModbusPDUWriteSingleRegisterResponseBuilder
}

// _ModbusPDUWriteSingleRegisterResponse is the data-structure of this message
type _ModbusPDUWriteSingleRegisterResponse struct {
	ModbusPDUContract
	Address uint16
	Value   uint16
}

var _ ModbusPDUWriteSingleRegisterResponse = (*_ModbusPDUWriteSingleRegisterResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUWriteSingleRegisterResponse)(nil)

// NewModbusPDUWriteSingleRegisterResponse factory function for _ModbusPDUWriteSingleRegisterResponse
func NewModbusPDUWriteSingleRegisterResponse(address uint16, value uint16) *_ModbusPDUWriteSingleRegisterResponse {
	_result := &_ModbusPDUWriteSingleRegisterResponse{
		ModbusPDUContract: NewModbusPDU(),
		Address:           address,
		Value:             value,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUWriteSingleRegisterResponseBuilder is a builder for ModbusPDUWriteSingleRegisterResponse
type ModbusPDUWriteSingleRegisterResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(address uint16, value uint16) ModbusPDUWriteSingleRegisterResponseBuilder
	// WithAddress adds Address (property field)
	WithAddress(uint16) ModbusPDUWriteSingleRegisterResponseBuilder
	// WithValue adds Value (property field)
	WithValue(uint16) ModbusPDUWriteSingleRegisterResponseBuilder
	// Build builds the ModbusPDUWriteSingleRegisterResponse or returns an error if something is wrong
	Build() (ModbusPDUWriteSingleRegisterResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUWriteSingleRegisterResponse
}

// NewModbusPDUWriteSingleRegisterResponseBuilder() creates a ModbusPDUWriteSingleRegisterResponseBuilder
func NewModbusPDUWriteSingleRegisterResponseBuilder() ModbusPDUWriteSingleRegisterResponseBuilder {
	return &_ModbusPDUWriteSingleRegisterResponseBuilder{_ModbusPDUWriteSingleRegisterResponse: new(_ModbusPDUWriteSingleRegisterResponse)}
}

type _ModbusPDUWriteSingleRegisterResponseBuilder struct {
	*_ModbusPDUWriteSingleRegisterResponse

	parentBuilder *_ModbusPDUBuilder

	err *utils.MultiError
}

var _ (ModbusPDUWriteSingleRegisterResponseBuilder) = (*_ModbusPDUWriteSingleRegisterResponseBuilder)(nil)

func (b *_ModbusPDUWriteSingleRegisterResponseBuilder) setParent(contract ModbusPDUContract) {
	b.ModbusPDUContract = contract
}

func (b *_ModbusPDUWriteSingleRegisterResponseBuilder) WithMandatoryFields(address uint16, value uint16) ModbusPDUWriteSingleRegisterResponseBuilder {
	return b.WithAddress(address).WithValue(value)
}

func (b *_ModbusPDUWriteSingleRegisterResponseBuilder) WithAddress(address uint16) ModbusPDUWriteSingleRegisterResponseBuilder {
	b.Address = address
	return b
}

func (b *_ModbusPDUWriteSingleRegisterResponseBuilder) WithValue(value uint16) ModbusPDUWriteSingleRegisterResponseBuilder {
	b.Value = value
	return b
}

func (b *_ModbusPDUWriteSingleRegisterResponseBuilder) Build() (ModbusPDUWriteSingleRegisterResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ModbusPDUWriteSingleRegisterResponse.deepCopy(), nil
}

func (b *_ModbusPDUWriteSingleRegisterResponseBuilder) MustBuild() ModbusPDUWriteSingleRegisterResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ModbusPDUWriteSingleRegisterResponseBuilder) Done() ModbusPDUBuilder {
	return b.parentBuilder
}

func (b *_ModbusPDUWriteSingleRegisterResponseBuilder) buildForModbusPDU() (ModbusPDU, error) {
	return b.Build()
}

func (b *_ModbusPDUWriteSingleRegisterResponseBuilder) DeepCopy() any {
	_copy := b.CreateModbusPDUWriteSingleRegisterResponseBuilder().(*_ModbusPDUWriteSingleRegisterResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateModbusPDUWriteSingleRegisterResponseBuilder creates a ModbusPDUWriteSingleRegisterResponseBuilder
func (b *_ModbusPDUWriteSingleRegisterResponse) CreateModbusPDUWriteSingleRegisterResponseBuilder() ModbusPDUWriteSingleRegisterResponseBuilder {
	if b == nil {
		return NewModbusPDUWriteSingleRegisterResponseBuilder()
	}
	return &_ModbusPDUWriteSingleRegisterResponseBuilder{_ModbusPDUWriteSingleRegisterResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUWriteSingleRegisterResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUWriteSingleRegisterResponse) GetFunctionFlag() uint8 {
	return 0x06
}

func (m *_ModbusPDUWriteSingleRegisterResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUWriteSingleRegisterResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteSingleRegisterResponse) GetAddress() uint16 {
	return m.Address
}

func (m *_ModbusPDUWriteSingleRegisterResponse) GetValue() uint16 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteSingleRegisterResponse(structType any) ModbusPDUWriteSingleRegisterResponse {
	if casted, ok := structType.(ModbusPDUWriteSingleRegisterResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteSingleRegisterResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteSingleRegisterResponse) GetTypeName() string {
	return "ModbusPDUWriteSingleRegisterResponse"
}

func (m *_ModbusPDUWriteSingleRegisterResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (address)
	lengthInBits += 16

	// Simple field (value)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUWriteSingleRegisterResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUWriteSingleRegisterResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUWriteSingleRegisterResponse ModbusPDUWriteSingleRegisterResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteSingleRegisterResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteSingleRegisterResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	address, err := ReadSimpleField(ctx, "address", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	value, err := ReadSimpleField(ctx, "value", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteSingleRegisterResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteSingleRegisterResponse")
	}

	return m, nil
}

func (m *_ModbusPDUWriteSingleRegisterResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteSingleRegisterResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteSingleRegisterResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteSingleRegisterResponse")
		}

		if err := WriteSimpleField[uint16](ctx, "address", m.GetAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if err := WriteSimpleField[uint16](ctx, "value", m.GetValue(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteSingleRegisterResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteSingleRegisterResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUWriteSingleRegisterResponse) IsModbusPDUWriteSingleRegisterResponse() {}

func (m *_ModbusPDUWriteSingleRegisterResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUWriteSingleRegisterResponse) deepCopy() *_ModbusPDUWriteSingleRegisterResponse {
	if m == nil {
		return nil
	}
	_ModbusPDUWriteSingleRegisterResponseCopy := &_ModbusPDUWriteSingleRegisterResponse{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		m.Address,
		m.Value,
	}
	m.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUWriteSingleRegisterResponseCopy
}

func (m *_ModbusPDUWriteSingleRegisterResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

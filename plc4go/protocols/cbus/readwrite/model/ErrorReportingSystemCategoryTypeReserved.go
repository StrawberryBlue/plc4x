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

// ErrorReportingSystemCategoryTypeReserved is the corresponding interface of ErrorReportingSystemCategoryTypeReserved
type ErrorReportingSystemCategoryTypeReserved interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ErrorReportingSystemCategoryType
	// GetReservedValue returns ReservedValue (property field)
	GetReservedValue() uint8
	// IsErrorReportingSystemCategoryTypeReserved is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorReportingSystemCategoryTypeReserved()
	// CreateBuilder creates a ErrorReportingSystemCategoryTypeReservedBuilder
	CreateErrorReportingSystemCategoryTypeReservedBuilder() ErrorReportingSystemCategoryTypeReservedBuilder
}

// _ErrorReportingSystemCategoryTypeReserved is the data-structure of this message
type _ErrorReportingSystemCategoryTypeReserved struct {
	ErrorReportingSystemCategoryTypeContract
	ReservedValue uint8
}

var _ ErrorReportingSystemCategoryTypeReserved = (*_ErrorReportingSystemCategoryTypeReserved)(nil)
var _ ErrorReportingSystemCategoryTypeRequirements = (*_ErrorReportingSystemCategoryTypeReserved)(nil)

// NewErrorReportingSystemCategoryTypeReserved factory function for _ErrorReportingSystemCategoryTypeReserved
func NewErrorReportingSystemCategoryTypeReserved(reservedValue uint8) *_ErrorReportingSystemCategoryTypeReserved {
	_result := &_ErrorReportingSystemCategoryTypeReserved{
		ErrorReportingSystemCategoryTypeContract: NewErrorReportingSystemCategoryType(),
		ReservedValue:                            reservedValue,
	}
	_result.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ErrorReportingSystemCategoryTypeReservedBuilder is a builder for ErrorReportingSystemCategoryTypeReserved
type ErrorReportingSystemCategoryTypeReservedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(reservedValue uint8) ErrorReportingSystemCategoryTypeReservedBuilder
	// WithReservedValue adds ReservedValue (property field)
	WithReservedValue(uint8) ErrorReportingSystemCategoryTypeReservedBuilder
	// Build builds the ErrorReportingSystemCategoryTypeReserved or returns an error if something is wrong
	Build() (ErrorReportingSystemCategoryTypeReserved, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ErrorReportingSystemCategoryTypeReserved
}

// NewErrorReportingSystemCategoryTypeReservedBuilder() creates a ErrorReportingSystemCategoryTypeReservedBuilder
func NewErrorReportingSystemCategoryTypeReservedBuilder() ErrorReportingSystemCategoryTypeReservedBuilder {
	return &_ErrorReportingSystemCategoryTypeReservedBuilder{_ErrorReportingSystemCategoryTypeReserved: new(_ErrorReportingSystemCategoryTypeReserved)}
}

type _ErrorReportingSystemCategoryTypeReservedBuilder struct {
	*_ErrorReportingSystemCategoryTypeReserved

	parentBuilder *_ErrorReportingSystemCategoryTypeBuilder

	err *utils.MultiError
}

var _ (ErrorReportingSystemCategoryTypeReservedBuilder) = (*_ErrorReportingSystemCategoryTypeReservedBuilder)(nil)

func (b *_ErrorReportingSystemCategoryTypeReservedBuilder) setParent(contract ErrorReportingSystemCategoryTypeContract) {
	b.ErrorReportingSystemCategoryTypeContract = contract
}

func (b *_ErrorReportingSystemCategoryTypeReservedBuilder) WithMandatoryFields(reservedValue uint8) ErrorReportingSystemCategoryTypeReservedBuilder {
	return b.WithReservedValue(reservedValue)
}

func (b *_ErrorReportingSystemCategoryTypeReservedBuilder) WithReservedValue(reservedValue uint8) ErrorReportingSystemCategoryTypeReservedBuilder {
	b.ReservedValue = reservedValue
	return b
}

func (b *_ErrorReportingSystemCategoryTypeReservedBuilder) Build() (ErrorReportingSystemCategoryTypeReserved, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ErrorReportingSystemCategoryTypeReserved.deepCopy(), nil
}

func (b *_ErrorReportingSystemCategoryTypeReservedBuilder) MustBuild() ErrorReportingSystemCategoryTypeReserved {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ErrorReportingSystemCategoryTypeReservedBuilder) Done() ErrorReportingSystemCategoryTypeBuilder {
	return b.parentBuilder
}

func (b *_ErrorReportingSystemCategoryTypeReservedBuilder) buildForErrorReportingSystemCategoryType() (ErrorReportingSystemCategoryType, error) {
	return b.Build()
}

func (b *_ErrorReportingSystemCategoryTypeReservedBuilder) DeepCopy() any {
	_copy := b.CreateErrorReportingSystemCategoryTypeReservedBuilder().(*_ErrorReportingSystemCategoryTypeReservedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateErrorReportingSystemCategoryTypeReservedBuilder creates a ErrorReportingSystemCategoryTypeReservedBuilder
func (b *_ErrorReportingSystemCategoryTypeReserved) CreateErrorReportingSystemCategoryTypeReservedBuilder() ErrorReportingSystemCategoryTypeReservedBuilder {
	if b == nil {
		return NewErrorReportingSystemCategoryTypeReservedBuilder()
	}
	return &_ErrorReportingSystemCategoryTypeReservedBuilder{_ErrorReportingSystemCategoryTypeReserved: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeReserved) GetErrorReportingSystemCategoryClass() ErrorReportingSystemCategoryClass {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ErrorReportingSystemCategoryTypeReserved) GetParent() ErrorReportingSystemCategoryTypeContract {
	return m.ErrorReportingSystemCategoryTypeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeReserved) GetReservedValue() uint8 {
	return m.ReservedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastErrorReportingSystemCategoryTypeReserved(structType any) ErrorReportingSystemCategoryTypeReserved {
	if casted, ok := structType.(ErrorReportingSystemCategoryTypeReserved); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingSystemCategoryTypeReserved); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingSystemCategoryTypeReserved) GetTypeName() string {
	return "ErrorReportingSystemCategoryTypeReserved"
}

func (m *_ErrorReportingSystemCategoryTypeReserved) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).getLengthInBits(ctx))

	// Simple field (reservedValue)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ErrorReportingSystemCategoryTypeReserved) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ErrorReportingSystemCategoryTypeReserved) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ErrorReportingSystemCategoryType, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (__errorReportingSystemCategoryTypeReserved ErrorReportingSystemCategoryTypeReserved, err error) {
	m.ErrorReportingSystemCategoryTypeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorReportingSystemCategoryTypeReserved"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingSystemCategoryTypeReserved")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedValue, err := ReadSimpleField(ctx, "reservedValue", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reservedValue' field"))
	}
	m.ReservedValue = reservedValue

	if closeErr := readBuffer.CloseContext("ErrorReportingSystemCategoryTypeReserved"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingSystemCategoryTypeReserved")
	}

	return m, nil
}

func (m *_ErrorReportingSystemCategoryTypeReserved) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorReportingSystemCategoryTypeReserved) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ErrorReportingSystemCategoryTypeReserved"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ErrorReportingSystemCategoryTypeReserved")
		}

		if err := WriteSimpleField[uint8](ctx, "reservedValue", m.GetReservedValue(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'reservedValue' field")
		}

		if popErr := writeBuffer.PopContext("ErrorReportingSystemCategoryTypeReserved"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ErrorReportingSystemCategoryTypeReserved")
		}
		return nil
	}
	return m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ErrorReportingSystemCategoryTypeReserved) IsErrorReportingSystemCategoryTypeReserved() {}

func (m *_ErrorReportingSystemCategoryTypeReserved) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ErrorReportingSystemCategoryTypeReserved) deepCopy() *_ErrorReportingSystemCategoryTypeReserved {
	if m == nil {
		return nil
	}
	_ErrorReportingSystemCategoryTypeReservedCopy := &_ErrorReportingSystemCategoryTypeReserved{
		m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).deepCopy(),
		m.ReservedValue,
	}
	m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType)._SubType = m
	return _ErrorReportingSystemCategoryTypeReservedCopy
}

func (m *_ErrorReportingSystemCategoryTypeReserved) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}

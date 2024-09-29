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

// HVACRawLevels is the corresponding interface of HVACRawLevels
type HVACRawLevels interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetRawValue returns RawValue (property field)
	GetRawValue() int16
	// GetValueInPercent returns ValueInPercent (virtual field)
	GetValueInPercent() float32
	// IsHVACRawLevels is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHVACRawLevels()
	// CreateBuilder creates a HVACRawLevelsBuilder
	CreateHVACRawLevelsBuilder() HVACRawLevelsBuilder
}

// _HVACRawLevels is the data-structure of this message
type _HVACRawLevels struct {
	RawValue int16
}

var _ HVACRawLevels = (*_HVACRawLevels)(nil)

// NewHVACRawLevels factory function for _HVACRawLevels
func NewHVACRawLevels(rawValue int16) *_HVACRawLevels {
	return &_HVACRawLevels{RawValue: rawValue}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HVACRawLevelsBuilder is a builder for HVACRawLevels
type HVACRawLevelsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(rawValue int16) HVACRawLevelsBuilder
	// WithRawValue adds RawValue (property field)
	WithRawValue(int16) HVACRawLevelsBuilder
	// Build builds the HVACRawLevels or returns an error if something is wrong
	Build() (HVACRawLevels, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HVACRawLevels
}

// NewHVACRawLevelsBuilder() creates a HVACRawLevelsBuilder
func NewHVACRawLevelsBuilder() HVACRawLevelsBuilder {
	return &_HVACRawLevelsBuilder{_HVACRawLevels: new(_HVACRawLevels)}
}

type _HVACRawLevelsBuilder struct {
	*_HVACRawLevels

	err *utils.MultiError
}

var _ (HVACRawLevelsBuilder) = (*_HVACRawLevelsBuilder)(nil)

func (b *_HVACRawLevelsBuilder) WithMandatoryFields(rawValue int16) HVACRawLevelsBuilder {
	return b.WithRawValue(rawValue)
}

func (b *_HVACRawLevelsBuilder) WithRawValue(rawValue int16) HVACRawLevelsBuilder {
	b.RawValue = rawValue
	return b
}

func (b *_HVACRawLevelsBuilder) Build() (HVACRawLevels, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HVACRawLevels.deepCopy(), nil
}

func (b *_HVACRawLevelsBuilder) MustBuild() HVACRawLevels {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_HVACRawLevelsBuilder) DeepCopy() any {
	_copy := b.CreateHVACRawLevelsBuilder().(*_HVACRawLevelsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHVACRawLevelsBuilder creates a HVACRawLevelsBuilder
func (b *_HVACRawLevels) CreateHVACRawLevelsBuilder() HVACRawLevelsBuilder {
	if b == nil {
		return NewHVACRawLevelsBuilder()
	}
	return &_HVACRawLevelsBuilder{_HVACRawLevels: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACRawLevels) GetRawValue() int16 {
	return m.RawValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACRawLevels) GetValueInPercent() float32 {
	ctx := context.Background()
	_ = ctx
	return float32(float32(m.GetRawValue()) / float32(float32(32767)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHVACRawLevels(structType any) HVACRawLevels {
	if casted, ok := structType.(HVACRawLevels); ok {
		return casted
	}
	if casted, ok := structType.(*HVACRawLevels); ok {
		return *casted
	}
	return nil
}

func (m *_HVACRawLevels) GetTypeName() string {
	return "HVACRawLevels"
}

func (m *_HVACRawLevels) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (rawValue)
	lengthInBits += 16

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_HVACRawLevels) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACRawLevelsParse(ctx context.Context, theBytes []byte) (HVACRawLevels, error) {
	return HVACRawLevelsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACRawLevelsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACRawLevels, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACRawLevels, error) {
		return HVACRawLevelsParseWithBuffer(ctx, readBuffer)
	}
}

func HVACRawLevelsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACRawLevels, error) {
	v, err := (&_HVACRawLevels{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_HVACRawLevels) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__hVACRawLevels HVACRawLevels, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HVACRawLevels"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACRawLevels")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	rawValue, err := ReadSimpleField(ctx, "rawValue", ReadSignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rawValue' field"))
	}
	m.RawValue = rawValue

	valueInPercent, err := ReadVirtualField[float32](ctx, "valueInPercent", (*float32)(nil), float32(rawValue)/float32(float32(32767)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueInPercent' field"))
	}
	_ = valueInPercent

	if closeErr := readBuffer.CloseContext("HVACRawLevels"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACRawLevels")
	}

	return m, nil
}

func (m *_HVACRawLevels) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACRawLevels) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("HVACRawLevels"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACRawLevels")
	}

	if err := WriteSimpleField[int16](ctx, "rawValue", m.GetRawValue(), WriteSignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'rawValue' field")
	}
	// Virtual field
	valueInPercent := m.GetValueInPercent()
	_ = valueInPercent
	if _valueInPercentErr := writeBuffer.WriteVirtual(ctx, "valueInPercent", m.GetValueInPercent()); _valueInPercentErr != nil {
		return errors.Wrap(_valueInPercentErr, "Error serializing 'valueInPercent' field")
	}

	if popErr := writeBuffer.PopContext("HVACRawLevels"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACRawLevels")
	}
	return nil
}

func (m *_HVACRawLevels) IsHVACRawLevels() {}

func (m *_HVACRawLevels) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HVACRawLevels) deepCopy() *_HVACRawLevels {
	if m == nil {
		return nil
	}
	_HVACRawLevelsCopy := &_HVACRawLevels{
		m.RawValue,
	}
	return _HVACRawLevelsCopy
}

func (m *_HVACRawLevels) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

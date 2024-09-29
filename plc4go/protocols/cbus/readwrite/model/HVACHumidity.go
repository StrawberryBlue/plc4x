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

// HVACHumidity is the corresponding interface of HVACHumidity
type HVACHumidity interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHumidityValue returns HumidityValue (property field)
	GetHumidityValue() uint16
	// GetHumidityInPercent returns HumidityInPercent (virtual field)
	GetHumidityInPercent() float32
	// IsHVACHumidity is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHVACHumidity()
	// CreateBuilder creates a HVACHumidityBuilder
	CreateHVACHumidityBuilder() HVACHumidityBuilder
}

// _HVACHumidity is the data-structure of this message
type _HVACHumidity struct {
	HumidityValue uint16
}

var _ HVACHumidity = (*_HVACHumidity)(nil)

// NewHVACHumidity factory function for _HVACHumidity
func NewHVACHumidity(humidityValue uint16) *_HVACHumidity {
	return &_HVACHumidity{HumidityValue: humidityValue}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HVACHumidityBuilder is a builder for HVACHumidity
type HVACHumidityBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(humidityValue uint16) HVACHumidityBuilder
	// WithHumidityValue adds HumidityValue (property field)
	WithHumidityValue(uint16) HVACHumidityBuilder
	// Build builds the HVACHumidity or returns an error if something is wrong
	Build() (HVACHumidity, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HVACHumidity
}

// NewHVACHumidityBuilder() creates a HVACHumidityBuilder
func NewHVACHumidityBuilder() HVACHumidityBuilder {
	return &_HVACHumidityBuilder{_HVACHumidity: new(_HVACHumidity)}
}

type _HVACHumidityBuilder struct {
	*_HVACHumidity

	err *utils.MultiError
}

var _ (HVACHumidityBuilder) = (*_HVACHumidityBuilder)(nil)

func (b *_HVACHumidityBuilder) WithMandatoryFields(humidityValue uint16) HVACHumidityBuilder {
	return b.WithHumidityValue(humidityValue)
}

func (b *_HVACHumidityBuilder) WithHumidityValue(humidityValue uint16) HVACHumidityBuilder {
	b.HumidityValue = humidityValue
	return b
}

func (b *_HVACHumidityBuilder) Build() (HVACHumidity, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HVACHumidity.deepCopy(), nil
}

func (b *_HVACHumidityBuilder) MustBuild() HVACHumidity {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_HVACHumidityBuilder) DeepCopy() any {
	_copy := b.CreateHVACHumidityBuilder().(*_HVACHumidityBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHVACHumidityBuilder creates a HVACHumidityBuilder
func (b *_HVACHumidity) CreateHVACHumidityBuilder() HVACHumidityBuilder {
	if b == nil {
		return NewHVACHumidityBuilder()
	}
	return &_HVACHumidityBuilder{_HVACHumidity: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HVACHumidity) GetHumidityValue() uint16 {
	return m.HumidityValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_HVACHumidity) GetHumidityInPercent() float32 {
	ctx := context.Background()
	_ = ctx
	return float32(float32(m.GetHumidityValue()) / float32(float32(65535)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHVACHumidity(structType any) HVACHumidity {
	if casted, ok := structType.(HVACHumidity); ok {
		return casted
	}
	if casted, ok := structType.(*HVACHumidity); ok {
		return *casted
	}
	return nil
}

func (m *_HVACHumidity) GetTypeName() string {
	return "HVACHumidity"
}

func (m *_HVACHumidity) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (humidityValue)
	lengthInBits += 16

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_HVACHumidity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HVACHumidityParse(ctx context.Context, theBytes []byte) (HVACHumidity, error) {
	return HVACHumidityParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HVACHumidityParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACHumidity, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (HVACHumidity, error) {
		return HVACHumidityParseWithBuffer(ctx, readBuffer)
	}
}

func HVACHumidityParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HVACHumidity, error) {
	v, err := (&_HVACHumidity{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_HVACHumidity) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__hVACHumidity HVACHumidity, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HVACHumidity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HVACHumidity")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	humidityValue, err := ReadSimpleField(ctx, "humidityValue", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidityValue' field"))
	}
	m.HumidityValue = humidityValue

	humidityInPercent, err := ReadVirtualField[float32](ctx, "humidityInPercent", (*float32)(nil), float32(humidityValue)/float32(float32(65535)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidityInPercent' field"))
	}
	_ = humidityInPercent

	if closeErr := readBuffer.CloseContext("HVACHumidity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HVACHumidity")
	}

	return m, nil
}

func (m *_HVACHumidity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HVACHumidity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("HVACHumidity"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HVACHumidity")
	}

	if err := WriteSimpleField[uint16](ctx, "humidityValue", m.GetHumidityValue(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'humidityValue' field")
	}
	// Virtual field
	humidityInPercent := m.GetHumidityInPercent()
	_ = humidityInPercent
	if _humidityInPercentErr := writeBuffer.WriteVirtual(ctx, "humidityInPercent", m.GetHumidityInPercent()); _humidityInPercentErr != nil {
		return errors.Wrap(_humidityInPercentErr, "Error serializing 'humidityInPercent' field")
	}

	if popErr := writeBuffer.PopContext("HVACHumidity"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HVACHumidity")
	}
	return nil
}

func (m *_HVACHumidity) IsHVACHumidity() {}

func (m *_HVACHumidity) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HVACHumidity) deepCopy() *_HVACHumidity {
	if m == nil {
		return nil
	}
	_HVACHumidityCopy := &_HVACHumidity{
		m.HumidityValue,
	}
	return _HVACHumidityCopy
}

func (m *_HVACHumidity) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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

// BACnetConstructedDataStartTime is the corresponding interface of BACnetConstructedDataStartTime
type BACnetConstructedDataStartTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetStartTime returns StartTime (property field)
	GetStartTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
	// IsBACnetConstructedDataStartTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataStartTime()
	// CreateBuilder creates a BACnetConstructedDataStartTimeBuilder
	CreateBACnetConstructedDataStartTimeBuilder() BACnetConstructedDataStartTimeBuilder
}

// _BACnetConstructedDataStartTime is the data-structure of this message
type _BACnetConstructedDataStartTime struct {
	BACnetConstructedDataContract
	StartTime BACnetDateTime
}

var _ BACnetConstructedDataStartTime = (*_BACnetConstructedDataStartTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataStartTime)(nil)

// NewBACnetConstructedDataStartTime factory function for _BACnetConstructedDataStartTime
func NewBACnetConstructedDataStartTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, startTime BACnetDateTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStartTime {
	if startTime == nil {
		panic("startTime of type BACnetDateTime for BACnetConstructedDataStartTime must not be nil")
	}
	_result := &_BACnetConstructedDataStartTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		StartTime:                     startTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataStartTimeBuilder is a builder for BACnetConstructedDataStartTime
type BACnetConstructedDataStartTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(startTime BACnetDateTime) BACnetConstructedDataStartTimeBuilder
	// WithStartTime adds StartTime (property field)
	WithStartTime(BACnetDateTime) BACnetConstructedDataStartTimeBuilder
	// WithStartTimeBuilder adds StartTime (property field) which is build by the builder
	WithStartTimeBuilder(func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataStartTimeBuilder
	// Build builds the BACnetConstructedDataStartTime or returns an error if something is wrong
	Build() (BACnetConstructedDataStartTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataStartTime
}

// NewBACnetConstructedDataStartTimeBuilder() creates a BACnetConstructedDataStartTimeBuilder
func NewBACnetConstructedDataStartTimeBuilder() BACnetConstructedDataStartTimeBuilder {
	return &_BACnetConstructedDataStartTimeBuilder{_BACnetConstructedDataStartTime: new(_BACnetConstructedDataStartTime)}
}

type _BACnetConstructedDataStartTimeBuilder struct {
	*_BACnetConstructedDataStartTime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataStartTimeBuilder) = (*_BACnetConstructedDataStartTimeBuilder)(nil)

func (b *_BACnetConstructedDataStartTimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataStartTimeBuilder) WithMandatoryFields(startTime BACnetDateTime) BACnetConstructedDataStartTimeBuilder {
	return b.WithStartTime(startTime)
}

func (b *_BACnetConstructedDataStartTimeBuilder) WithStartTime(startTime BACnetDateTime) BACnetConstructedDataStartTimeBuilder {
	b.StartTime = startTime
	return b
}

func (b *_BACnetConstructedDataStartTimeBuilder) WithStartTimeBuilder(builderSupplier func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataStartTimeBuilder {
	builder := builderSupplier(b.StartTime.CreateBACnetDateTimeBuilder())
	var err error
	b.StartTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDateTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataStartTimeBuilder) Build() (BACnetConstructedDataStartTime, error) {
	if b.StartTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'startTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataStartTime.deepCopy(), nil
}

func (b *_BACnetConstructedDataStartTimeBuilder) MustBuild() BACnetConstructedDataStartTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataStartTimeBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataStartTimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataStartTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataStartTimeBuilder().(*_BACnetConstructedDataStartTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataStartTimeBuilder creates a BACnetConstructedDataStartTimeBuilder
func (b *_BACnetConstructedDataStartTime) CreateBACnetConstructedDataStartTimeBuilder() BACnetConstructedDataStartTimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataStartTimeBuilder()
	}
	return &_BACnetConstructedDataStartTimeBuilder{_BACnetConstructedDataStartTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStartTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataStartTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_START_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStartTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataStartTime) GetStartTime() BACnetDateTime {
	return m.StartTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataStartTime) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetStartTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStartTime(structType any) BACnetConstructedDataStartTime {
	if casted, ok := structType.(BACnetConstructedDataStartTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStartTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStartTime) GetTypeName() string {
	return "BACnetConstructedDataStartTime"
}

func (m *_BACnetConstructedDataStartTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (startTime)
	lengthInBits += m.StartTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataStartTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataStartTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataStartTime BACnetConstructedDataStartTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStartTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStartTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	startTime, err := ReadSimpleField[BACnetDateTime](ctx, "startTime", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startTime' field"))
	}
	m.StartTime = startTime

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), startTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStartTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStartTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataStartTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataStartTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStartTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStartTime")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "startTime", m.GetStartTime(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'startTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStartTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStartTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStartTime) IsBACnetConstructedDataStartTime() {}

func (m *_BACnetConstructedDataStartTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataStartTime) deepCopy() *_BACnetConstructedDataStartTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataStartTimeCopy := &_BACnetConstructedDataStartTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.StartTime.DeepCopy().(BACnetDateTime),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataStartTimeCopy
}

func (m *_BACnetConstructedDataStartTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

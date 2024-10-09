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

// BACnetConstructedDataTimeOfStateCountReset is the corresponding interface of BACnetConstructedDataTimeOfStateCountReset
type BACnetConstructedDataTimeOfStateCountReset interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetTimeOfStateCountReset returns TimeOfStateCountReset (property field)
	GetTimeOfStateCountReset() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
	// IsBACnetConstructedDataTimeOfStateCountReset is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimeOfStateCountReset()
	// CreateBuilder creates a BACnetConstructedDataTimeOfStateCountResetBuilder
	CreateBACnetConstructedDataTimeOfStateCountResetBuilder() BACnetConstructedDataTimeOfStateCountResetBuilder
}

// _BACnetConstructedDataTimeOfStateCountReset is the data-structure of this message
type _BACnetConstructedDataTimeOfStateCountReset struct {
	BACnetConstructedDataContract
	TimeOfStateCountReset BACnetDateTime
}

var _ BACnetConstructedDataTimeOfStateCountReset = (*_BACnetConstructedDataTimeOfStateCountReset)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimeOfStateCountReset)(nil)

// NewBACnetConstructedDataTimeOfStateCountReset factory function for _BACnetConstructedDataTimeOfStateCountReset
func NewBACnetConstructedDataTimeOfStateCountReset(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, timeOfStateCountReset BACnetDateTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeOfStateCountReset {
	if timeOfStateCountReset == nil {
		panic("timeOfStateCountReset of type BACnetDateTime for BACnetConstructedDataTimeOfStateCountReset must not be nil")
	}
	_result := &_BACnetConstructedDataTimeOfStateCountReset{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		TimeOfStateCountReset:         timeOfStateCountReset,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTimeOfStateCountResetBuilder is a builder for BACnetConstructedDataTimeOfStateCountReset
type BACnetConstructedDataTimeOfStateCountResetBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timeOfStateCountReset BACnetDateTime) BACnetConstructedDataTimeOfStateCountResetBuilder
	// WithTimeOfStateCountReset adds TimeOfStateCountReset (property field)
	WithTimeOfStateCountReset(BACnetDateTime) BACnetConstructedDataTimeOfStateCountResetBuilder
	// WithTimeOfStateCountResetBuilder adds TimeOfStateCountReset (property field) which is build by the builder
	WithTimeOfStateCountResetBuilder(func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataTimeOfStateCountResetBuilder
	// Build builds the BACnetConstructedDataTimeOfStateCountReset or returns an error if something is wrong
	Build() (BACnetConstructedDataTimeOfStateCountReset, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTimeOfStateCountReset
}

// NewBACnetConstructedDataTimeOfStateCountResetBuilder() creates a BACnetConstructedDataTimeOfStateCountResetBuilder
func NewBACnetConstructedDataTimeOfStateCountResetBuilder() BACnetConstructedDataTimeOfStateCountResetBuilder {
	return &_BACnetConstructedDataTimeOfStateCountResetBuilder{_BACnetConstructedDataTimeOfStateCountReset: new(_BACnetConstructedDataTimeOfStateCountReset)}
}

type _BACnetConstructedDataTimeOfStateCountResetBuilder struct {
	*_BACnetConstructedDataTimeOfStateCountReset

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataTimeOfStateCountResetBuilder) = (*_BACnetConstructedDataTimeOfStateCountResetBuilder)(nil)

func (b *_BACnetConstructedDataTimeOfStateCountResetBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataTimeOfStateCountResetBuilder) WithMandatoryFields(timeOfStateCountReset BACnetDateTime) BACnetConstructedDataTimeOfStateCountResetBuilder {
	return b.WithTimeOfStateCountReset(timeOfStateCountReset)
}

func (b *_BACnetConstructedDataTimeOfStateCountResetBuilder) WithTimeOfStateCountReset(timeOfStateCountReset BACnetDateTime) BACnetConstructedDataTimeOfStateCountResetBuilder {
	b.TimeOfStateCountReset = timeOfStateCountReset
	return b
}

func (b *_BACnetConstructedDataTimeOfStateCountResetBuilder) WithTimeOfStateCountResetBuilder(builderSupplier func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataTimeOfStateCountResetBuilder {
	builder := builderSupplier(b.TimeOfStateCountReset.CreateBACnetDateTimeBuilder())
	var err error
	b.TimeOfStateCountReset, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDateTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataTimeOfStateCountResetBuilder) Build() (BACnetConstructedDataTimeOfStateCountReset, error) {
	if b.TimeOfStateCountReset == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeOfStateCountReset' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataTimeOfStateCountReset.deepCopy(), nil
}

func (b *_BACnetConstructedDataTimeOfStateCountResetBuilder) MustBuild() BACnetConstructedDataTimeOfStateCountReset {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataTimeOfStateCountResetBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataTimeOfStateCountResetBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataTimeOfStateCountResetBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataTimeOfStateCountResetBuilder().(*_BACnetConstructedDataTimeOfStateCountResetBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataTimeOfStateCountResetBuilder creates a BACnetConstructedDataTimeOfStateCountResetBuilder
func (b *_BACnetConstructedDataTimeOfStateCountReset) CreateBACnetConstructedDataTimeOfStateCountResetBuilder() BACnetConstructedDataTimeOfStateCountResetBuilder {
	if b == nil {
		return NewBACnetConstructedDataTimeOfStateCountResetBuilder()
	}
	return &_BACnetConstructedDataTimeOfStateCountResetBuilder{_BACnetConstructedDataTimeOfStateCountReset: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_OF_STATE_COUNT_RESET
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetTimeOfStateCountReset() BACnetDateTime {
	return m.TimeOfStateCountReset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetTimeOfStateCountReset())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeOfStateCountReset(structType any) BACnetConstructedDataTimeOfStateCountReset {
	if casted, ok := structType.(BACnetConstructedDataTimeOfStateCountReset); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeOfStateCountReset); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetTypeName() string {
	return "BACnetConstructedDataTimeOfStateCountReset"
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (timeOfStateCountReset)
	lengthInBits += m.TimeOfStateCountReset.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimeOfStateCountReset BACnetConstructedDataTimeOfStateCountReset, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeOfStateCountReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeOfStateCountReset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeOfStateCountReset, err := ReadSimpleField[BACnetDateTime](ctx, "timeOfStateCountReset", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeOfStateCountReset' field"))
	}
	m.TimeOfStateCountReset = timeOfStateCountReset

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), timeOfStateCountReset)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeOfStateCountReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeOfStateCountReset")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeOfStateCountReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeOfStateCountReset")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "timeOfStateCountReset", m.GetTimeOfStateCountReset(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeOfStateCountReset' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeOfStateCountReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeOfStateCountReset")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) IsBACnetConstructedDataTimeOfStateCountReset() {
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) deepCopy() *_BACnetConstructedDataTimeOfStateCountReset {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTimeOfStateCountResetCopy := &_BACnetConstructedDataTimeOfStateCountReset{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.TimeOfStateCountReset.DeepCopy().(BACnetDateTime),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTimeOfStateCountResetCopy
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) String() string {
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

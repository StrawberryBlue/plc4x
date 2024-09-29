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

// BACnetConstructedDataChannelPresentValue is the corresponding interface of BACnetConstructedDataChannelPresentValue
type BACnetConstructedDataChannelPresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetChannelValue
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetChannelValue
	// IsBACnetConstructedDataChannelPresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataChannelPresentValue()
	// CreateBuilder creates a BACnetConstructedDataChannelPresentValueBuilder
	CreateBACnetConstructedDataChannelPresentValueBuilder() BACnetConstructedDataChannelPresentValueBuilder
}

// _BACnetConstructedDataChannelPresentValue is the data-structure of this message
type _BACnetConstructedDataChannelPresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetChannelValue
}

var _ BACnetConstructedDataChannelPresentValue = (*_BACnetConstructedDataChannelPresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataChannelPresentValue)(nil)

// NewBACnetConstructedDataChannelPresentValue factory function for _BACnetConstructedDataChannelPresentValue
func NewBACnetConstructedDataChannelPresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetChannelValue, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataChannelPresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetChannelValue for BACnetConstructedDataChannelPresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataChannelPresentValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PresentValue:                  presentValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataChannelPresentValueBuilder is a builder for BACnetConstructedDataChannelPresentValue
type BACnetConstructedDataChannelPresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetChannelValue) BACnetConstructedDataChannelPresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetChannelValue) BACnetConstructedDataChannelPresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetChannelValueBuilder) BACnetChannelValueBuilder) BACnetConstructedDataChannelPresentValueBuilder
	// Build builds the BACnetConstructedDataChannelPresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataChannelPresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataChannelPresentValue
}

// NewBACnetConstructedDataChannelPresentValueBuilder() creates a BACnetConstructedDataChannelPresentValueBuilder
func NewBACnetConstructedDataChannelPresentValueBuilder() BACnetConstructedDataChannelPresentValueBuilder {
	return &_BACnetConstructedDataChannelPresentValueBuilder{_BACnetConstructedDataChannelPresentValue: new(_BACnetConstructedDataChannelPresentValue)}
}

type _BACnetConstructedDataChannelPresentValueBuilder struct {
	*_BACnetConstructedDataChannelPresentValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataChannelPresentValueBuilder) = (*_BACnetConstructedDataChannelPresentValueBuilder)(nil)

func (b *_BACnetConstructedDataChannelPresentValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataChannelPresentValueBuilder) WithMandatoryFields(presentValue BACnetChannelValue) BACnetConstructedDataChannelPresentValueBuilder {
	return b.WithPresentValue(presentValue)
}

func (b *_BACnetConstructedDataChannelPresentValueBuilder) WithPresentValue(presentValue BACnetChannelValue) BACnetConstructedDataChannelPresentValueBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetConstructedDataChannelPresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetChannelValueBuilder) BACnetChannelValueBuilder) BACnetConstructedDataChannelPresentValueBuilder {
	builder := builderSupplier(b.PresentValue.CreateBACnetChannelValueBuilder())
	var err error
	b.PresentValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetChannelValueBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataChannelPresentValueBuilder) Build() (BACnetConstructedDataChannelPresentValue, error) {
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataChannelPresentValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataChannelPresentValueBuilder) MustBuild() BACnetConstructedDataChannelPresentValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataChannelPresentValueBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataChannelPresentValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataChannelPresentValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataChannelPresentValueBuilder().(*_BACnetConstructedDataChannelPresentValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataChannelPresentValueBuilder creates a BACnetConstructedDataChannelPresentValueBuilder
func (b *_BACnetConstructedDataChannelPresentValue) CreateBACnetConstructedDataChannelPresentValueBuilder() BACnetConstructedDataChannelPresentValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataChannelPresentValueBuilder()
	}
	return &_BACnetConstructedDataChannelPresentValueBuilder{_BACnetConstructedDataChannelPresentValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataChannelPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_CHANNEL
}

func (m *_BACnetConstructedDataChannelPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataChannelPresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataChannelPresentValue) GetPresentValue() BACnetChannelValue {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataChannelPresentValue) GetActualValue() BACnetChannelValue {
	ctx := context.Background()
	_ = ctx
	return CastBACnetChannelValue(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataChannelPresentValue(structType any) BACnetConstructedDataChannelPresentValue {
	if casted, ok := structType.(BACnetConstructedDataChannelPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataChannelPresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataChannelPresentValue) GetTypeName() string {
	return "BACnetConstructedDataChannelPresentValue"
}

func (m *_BACnetConstructedDataChannelPresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataChannelPresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataChannelPresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataChannelPresentValue BACnetConstructedDataChannelPresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataChannelPresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataChannelPresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetChannelValue](ctx, "presentValue", ReadComplex[BACnetChannelValue](BACnetChannelValueParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetChannelValue](ctx, "actualValue", (*BACnetChannelValue)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataChannelPresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataChannelPresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataChannelPresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataChannelPresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataChannelPresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataChannelPresentValue")
		}

		if err := WriteSimpleField[BACnetChannelValue](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetChannelValue](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataChannelPresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataChannelPresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataChannelPresentValue) IsBACnetConstructedDataChannelPresentValue() {}

func (m *_BACnetConstructedDataChannelPresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataChannelPresentValue) deepCopy() *_BACnetConstructedDataChannelPresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataChannelPresentValueCopy := &_BACnetConstructedDataChannelPresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.PresentValue.DeepCopy().(BACnetChannelValue),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataChannelPresentValueCopy
}

func (m *_BACnetConstructedDataChannelPresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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

// BACnetConstructedDataBinaryLightingOutputPresentValue is the corresponding interface of BACnetConstructedDataBinaryLightingOutputPresentValue
type BACnetConstructedDataBinaryLightingOutputPresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetBinaryLightingPVTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetBinaryLightingPVTagged
	// IsBACnetConstructedDataBinaryLightingOutputPresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBinaryLightingOutputPresentValue()
	// CreateBuilder creates a BACnetConstructedDataBinaryLightingOutputPresentValueBuilder
	CreateBACnetConstructedDataBinaryLightingOutputPresentValueBuilder() BACnetConstructedDataBinaryLightingOutputPresentValueBuilder
}

// _BACnetConstructedDataBinaryLightingOutputPresentValue is the data-structure of this message
type _BACnetConstructedDataBinaryLightingOutputPresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetBinaryLightingPVTagged
}

var _ BACnetConstructedDataBinaryLightingOutputPresentValue = (*_BACnetConstructedDataBinaryLightingOutputPresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBinaryLightingOutputPresentValue)(nil)

// NewBACnetConstructedDataBinaryLightingOutputPresentValue factory function for _BACnetConstructedDataBinaryLightingOutputPresentValue
func NewBACnetConstructedDataBinaryLightingOutputPresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetBinaryLightingPVTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBinaryLightingOutputPresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetBinaryLightingPVTagged for BACnetConstructedDataBinaryLightingOutputPresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataBinaryLightingOutputPresentValue{
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

// BACnetConstructedDataBinaryLightingOutputPresentValueBuilder is a builder for BACnetConstructedDataBinaryLightingOutputPresentValue
type BACnetConstructedDataBinaryLightingOutputPresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetBinaryLightingPVTagged) BACnetConstructedDataBinaryLightingOutputPresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetBinaryLightingPVTagged) BACnetConstructedDataBinaryLightingOutputPresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetBinaryLightingPVTaggedBuilder) BACnetBinaryLightingPVTaggedBuilder) BACnetConstructedDataBinaryLightingOutputPresentValueBuilder
	// Build builds the BACnetConstructedDataBinaryLightingOutputPresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataBinaryLightingOutputPresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBinaryLightingOutputPresentValue
}

// NewBACnetConstructedDataBinaryLightingOutputPresentValueBuilder() creates a BACnetConstructedDataBinaryLightingOutputPresentValueBuilder
func NewBACnetConstructedDataBinaryLightingOutputPresentValueBuilder() BACnetConstructedDataBinaryLightingOutputPresentValueBuilder {
	return &_BACnetConstructedDataBinaryLightingOutputPresentValueBuilder{_BACnetConstructedDataBinaryLightingOutputPresentValue: new(_BACnetConstructedDataBinaryLightingOutputPresentValue)}
}

type _BACnetConstructedDataBinaryLightingOutputPresentValueBuilder struct {
	*_BACnetConstructedDataBinaryLightingOutputPresentValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBinaryLightingOutputPresentValueBuilder) = (*_BACnetConstructedDataBinaryLightingOutputPresentValueBuilder)(nil)

func (b *_BACnetConstructedDataBinaryLightingOutputPresentValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataBinaryLightingOutputPresentValueBuilder) WithMandatoryFields(presentValue BACnetBinaryLightingPVTagged) BACnetConstructedDataBinaryLightingOutputPresentValueBuilder {
	return b.WithPresentValue(presentValue)
}

func (b *_BACnetConstructedDataBinaryLightingOutputPresentValueBuilder) WithPresentValue(presentValue BACnetBinaryLightingPVTagged) BACnetConstructedDataBinaryLightingOutputPresentValueBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetConstructedDataBinaryLightingOutputPresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetBinaryLightingPVTaggedBuilder) BACnetBinaryLightingPVTaggedBuilder) BACnetConstructedDataBinaryLightingOutputPresentValueBuilder {
	builder := builderSupplier(b.PresentValue.CreateBACnetBinaryLightingPVTaggedBuilder())
	var err error
	b.PresentValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetBinaryLightingPVTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBinaryLightingOutputPresentValueBuilder) Build() (BACnetConstructedDataBinaryLightingOutputPresentValue, error) {
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBinaryLightingOutputPresentValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataBinaryLightingOutputPresentValueBuilder) MustBuild() BACnetConstructedDataBinaryLightingOutputPresentValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataBinaryLightingOutputPresentValueBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBinaryLightingOutputPresentValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBinaryLightingOutputPresentValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBinaryLightingOutputPresentValueBuilder().(*_BACnetConstructedDataBinaryLightingOutputPresentValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBinaryLightingOutputPresentValueBuilder creates a BACnetConstructedDataBinaryLightingOutputPresentValueBuilder
func (b *_BACnetConstructedDataBinaryLightingOutputPresentValue) CreateBACnetConstructedDataBinaryLightingOutputPresentValueBuilder() BACnetConstructedDataBinaryLightingOutputPresentValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataBinaryLightingOutputPresentValueBuilder()
	}
	return &_BACnetConstructedDataBinaryLightingOutputPresentValueBuilder{_BACnetConstructedDataBinaryLightingOutputPresentValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_LIGHTING_OUTPUT
}

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) GetPresentValue() BACnetBinaryLightingPVTagged {
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

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) GetActualValue() BACnetBinaryLightingPVTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetBinaryLightingPVTagged(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryLightingOutputPresentValue(structType any) BACnetConstructedDataBinaryLightingOutputPresentValue {
	if casted, ok := structType.(BACnetConstructedDataBinaryLightingOutputPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryLightingOutputPresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) GetTypeName() string {
	return "BACnetConstructedDataBinaryLightingOutputPresentValue"
}

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBinaryLightingOutputPresentValue BACnetConstructedDataBinaryLightingOutputPresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryLightingOutputPresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryLightingOutputPresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetBinaryLightingPVTagged](ctx, "presentValue", ReadComplex[BACnetBinaryLightingPVTagged](BACnetBinaryLightingPVTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetBinaryLightingPVTagged](ctx, "actualValue", (*BACnetBinaryLightingPVTagged)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryLightingOutputPresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryLightingOutputPresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryLightingOutputPresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryLightingOutputPresentValue")
		}

		if err := WriteSimpleField[BACnetBinaryLightingPVTagged](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetBinaryLightingPVTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryLightingOutputPresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryLightingOutputPresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) IsBACnetConstructedDataBinaryLightingOutputPresentValue() {
}

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) deepCopy() *_BACnetConstructedDataBinaryLightingOutputPresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBinaryLightingOutputPresentValueCopy := &_BACnetConstructedDataBinaryLightingOutputPresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.PresentValue.DeepCopy().(BACnetBinaryLightingPVTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBinaryLightingOutputPresentValueCopy
}

func (m *_BACnetConstructedDataBinaryLightingOutputPresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

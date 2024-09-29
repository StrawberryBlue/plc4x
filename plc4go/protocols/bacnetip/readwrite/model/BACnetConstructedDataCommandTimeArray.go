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

// BACnetConstructedDataCommandTimeArray is the corresponding interface of BACnetConstructedDataCommandTimeArray
type BACnetConstructedDataCommandTimeArray interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetCommandTimeArray returns CommandTimeArray (property field)
	GetCommandTimeArray() []BACnetTimeStamp
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataCommandTimeArray is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCommandTimeArray()
	// CreateBuilder creates a BACnetConstructedDataCommandTimeArrayBuilder
	CreateBACnetConstructedDataCommandTimeArrayBuilder() BACnetConstructedDataCommandTimeArrayBuilder
}

// _BACnetConstructedDataCommandTimeArray is the data-structure of this message
type _BACnetConstructedDataCommandTimeArray struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	CommandTimeArray     []BACnetTimeStamp
}

var _ BACnetConstructedDataCommandTimeArray = (*_BACnetConstructedDataCommandTimeArray)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCommandTimeArray)(nil)

// NewBACnetConstructedDataCommandTimeArray factory function for _BACnetConstructedDataCommandTimeArray
func NewBACnetConstructedDataCommandTimeArray(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, commandTimeArray []BACnetTimeStamp, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCommandTimeArray {
	_result := &_BACnetConstructedDataCommandTimeArray{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		CommandTimeArray:              commandTimeArray,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCommandTimeArrayBuilder is a builder for BACnetConstructedDataCommandTimeArray
type BACnetConstructedDataCommandTimeArrayBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(commandTimeArray []BACnetTimeStamp) BACnetConstructedDataCommandTimeArrayBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataCommandTimeArrayBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataCommandTimeArrayBuilder
	// WithCommandTimeArray adds CommandTimeArray (property field)
	WithCommandTimeArray(...BACnetTimeStamp) BACnetConstructedDataCommandTimeArrayBuilder
	// Build builds the BACnetConstructedDataCommandTimeArray or returns an error if something is wrong
	Build() (BACnetConstructedDataCommandTimeArray, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCommandTimeArray
}

// NewBACnetConstructedDataCommandTimeArrayBuilder() creates a BACnetConstructedDataCommandTimeArrayBuilder
func NewBACnetConstructedDataCommandTimeArrayBuilder() BACnetConstructedDataCommandTimeArrayBuilder {
	return &_BACnetConstructedDataCommandTimeArrayBuilder{_BACnetConstructedDataCommandTimeArray: new(_BACnetConstructedDataCommandTimeArray)}
}

type _BACnetConstructedDataCommandTimeArrayBuilder struct {
	*_BACnetConstructedDataCommandTimeArray

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCommandTimeArrayBuilder) = (*_BACnetConstructedDataCommandTimeArrayBuilder)(nil)

func (b *_BACnetConstructedDataCommandTimeArrayBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataCommandTimeArrayBuilder) WithMandatoryFields(commandTimeArray []BACnetTimeStamp) BACnetConstructedDataCommandTimeArrayBuilder {
	return b.WithCommandTimeArray(commandTimeArray...)
}

func (b *_BACnetConstructedDataCommandTimeArrayBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataCommandTimeArrayBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataCommandTimeArrayBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataCommandTimeArrayBuilder {
	builder := builderSupplier(b.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataCommandTimeArrayBuilder) WithCommandTimeArray(commandTimeArray ...BACnetTimeStamp) BACnetConstructedDataCommandTimeArrayBuilder {
	b.CommandTimeArray = commandTimeArray
	return b
}

func (b *_BACnetConstructedDataCommandTimeArrayBuilder) Build() (BACnetConstructedDataCommandTimeArray, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCommandTimeArray.deepCopy(), nil
}

func (b *_BACnetConstructedDataCommandTimeArrayBuilder) MustBuild() BACnetConstructedDataCommandTimeArray {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataCommandTimeArrayBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCommandTimeArrayBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCommandTimeArrayBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCommandTimeArrayBuilder().(*_BACnetConstructedDataCommandTimeArrayBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCommandTimeArrayBuilder creates a BACnetConstructedDataCommandTimeArrayBuilder
func (b *_BACnetConstructedDataCommandTimeArray) CreateBACnetConstructedDataCommandTimeArrayBuilder() BACnetConstructedDataCommandTimeArrayBuilder {
	if b == nil {
		return NewBACnetConstructedDataCommandTimeArrayBuilder()
	}
	return &_BACnetConstructedDataCommandTimeArrayBuilder{_BACnetConstructedDataCommandTimeArray: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCommandTimeArray) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCommandTimeArray) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COMMAND_TIME_ARRAY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCommandTimeArray) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCommandTimeArray) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataCommandTimeArray) GetCommandTimeArray() []BACnetTimeStamp {
	return m.CommandTimeArray
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCommandTimeArray) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCommandTimeArray(structType any) BACnetConstructedDataCommandTimeArray {
	if casted, ok := structType.(BACnetConstructedDataCommandTimeArray); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCommandTimeArray); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCommandTimeArray) GetTypeName() string {
	return "BACnetConstructedDataCommandTimeArray"
}

func (m *_BACnetConstructedDataCommandTimeArray) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.CommandTimeArray) > 0 {
		for _, element := range m.CommandTimeArray {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataCommandTimeArray) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCommandTimeArray) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCommandTimeArray BACnetConstructedDataCommandTimeArray, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCommandTimeArray"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCommandTimeArray")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	commandTimeArray, err := ReadTerminatedArrayField[BACnetTimeStamp](ctx, "commandTimeArray", ReadComplex[BACnetTimeStamp](BACnetTimeStampParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTimeArray' field"))
	}
	m.CommandTimeArray = commandTimeArray

	// Validation
	if !(bool(bool((arrayIndexArgument) != (nil))) || bool(bool((len(commandTimeArray)) == (16)))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "commandTimeArray should have exactly 16 values"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCommandTimeArray"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCommandTimeArray")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCommandTimeArray) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCommandTimeArray) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCommandTimeArray"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCommandTimeArray")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "commandTimeArray", m.GetCommandTimeArray(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'commandTimeArray' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCommandTimeArray"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCommandTimeArray")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCommandTimeArray) IsBACnetConstructedDataCommandTimeArray() {}

func (m *_BACnetConstructedDataCommandTimeArray) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCommandTimeArray) deepCopy() *_BACnetConstructedDataCommandTimeArray {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCommandTimeArrayCopy := &_BACnetConstructedDataCommandTimeArray{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetTimeStamp, BACnetTimeStamp](m.CommandTimeArray),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCommandTimeArrayCopy
}

func (m *_BACnetConstructedDataCommandTimeArray) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

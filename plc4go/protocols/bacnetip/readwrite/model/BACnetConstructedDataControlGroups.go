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

// BACnetConstructedDataControlGroups is the corresponding interface of BACnetConstructedDataControlGroups
type BACnetConstructedDataControlGroups interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetControlGroups returns ControlGroups (property field)
	GetControlGroups() []BACnetApplicationTagUnsignedInteger
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataControlGroups is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataControlGroups()
	// CreateBuilder creates a BACnetConstructedDataControlGroupsBuilder
	CreateBACnetConstructedDataControlGroupsBuilder() BACnetConstructedDataControlGroupsBuilder
}

// _BACnetConstructedDataControlGroups is the data-structure of this message
type _BACnetConstructedDataControlGroups struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	ControlGroups        []BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataControlGroups = (*_BACnetConstructedDataControlGroups)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataControlGroups)(nil)

// NewBACnetConstructedDataControlGroups factory function for _BACnetConstructedDataControlGroups
func NewBACnetConstructedDataControlGroups(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, controlGroups []BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataControlGroups {
	_result := &_BACnetConstructedDataControlGroups{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		ControlGroups:                 controlGroups,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataControlGroupsBuilder is a builder for BACnetConstructedDataControlGroups
type BACnetConstructedDataControlGroupsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(controlGroups []BACnetApplicationTagUnsignedInteger) BACnetConstructedDataControlGroupsBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataControlGroupsBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataControlGroupsBuilder
	// WithControlGroups adds ControlGroups (property field)
	WithControlGroups(...BACnetApplicationTagUnsignedInteger) BACnetConstructedDataControlGroupsBuilder
	// Build builds the BACnetConstructedDataControlGroups or returns an error if something is wrong
	Build() (BACnetConstructedDataControlGroups, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataControlGroups
}

// NewBACnetConstructedDataControlGroupsBuilder() creates a BACnetConstructedDataControlGroupsBuilder
func NewBACnetConstructedDataControlGroupsBuilder() BACnetConstructedDataControlGroupsBuilder {
	return &_BACnetConstructedDataControlGroupsBuilder{_BACnetConstructedDataControlGroups: new(_BACnetConstructedDataControlGroups)}
}

type _BACnetConstructedDataControlGroupsBuilder struct {
	*_BACnetConstructedDataControlGroups

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataControlGroupsBuilder) = (*_BACnetConstructedDataControlGroupsBuilder)(nil)

func (b *_BACnetConstructedDataControlGroupsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataControlGroupsBuilder) WithMandatoryFields(controlGroups []BACnetApplicationTagUnsignedInteger) BACnetConstructedDataControlGroupsBuilder {
	return b.WithControlGroups(controlGroups...)
}

func (b *_BACnetConstructedDataControlGroupsBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataControlGroupsBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataControlGroupsBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataControlGroupsBuilder {
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

func (b *_BACnetConstructedDataControlGroupsBuilder) WithControlGroups(controlGroups ...BACnetApplicationTagUnsignedInteger) BACnetConstructedDataControlGroupsBuilder {
	b.ControlGroups = controlGroups
	return b
}

func (b *_BACnetConstructedDataControlGroupsBuilder) Build() (BACnetConstructedDataControlGroups, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataControlGroups.deepCopy(), nil
}

func (b *_BACnetConstructedDataControlGroupsBuilder) MustBuild() BACnetConstructedDataControlGroups {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataControlGroupsBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataControlGroupsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataControlGroupsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataControlGroupsBuilder().(*_BACnetConstructedDataControlGroupsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataControlGroupsBuilder creates a BACnetConstructedDataControlGroupsBuilder
func (b *_BACnetConstructedDataControlGroups) CreateBACnetConstructedDataControlGroupsBuilder() BACnetConstructedDataControlGroupsBuilder {
	if b == nil {
		return NewBACnetConstructedDataControlGroupsBuilder()
	}
	return &_BACnetConstructedDataControlGroupsBuilder{_BACnetConstructedDataControlGroups: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataControlGroups) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataControlGroups) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CONTROL_GROUPS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataControlGroups) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataControlGroups) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataControlGroups) GetControlGroups() []BACnetApplicationTagUnsignedInteger {
	return m.ControlGroups
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataControlGroups) GetZero() uint64 {
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
func CastBACnetConstructedDataControlGroups(structType any) BACnetConstructedDataControlGroups {
	if casted, ok := structType.(BACnetConstructedDataControlGroups); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataControlGroups); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataControlGroups) GetTypeName() string {
	return "BACnetConstructedDataControlGroups"
}

func (m *_BACnetConstructedDataControlGroups) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.ControlGroups) > 0 {
		for _, element := range m.ControlGroups {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataControlGroups) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataControlGroups) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataControlGroups BACnetConstructedDataControlGroups, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataControlGroups"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataControlGroups")
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

	controlGroups, err := ReadTerminatedArrayField[BACnetApplicationTagUnsignedInteger](ctx, "controlGroups", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'controlGroups' field"))
	}
	m.ControlGroups = controlGroups

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataControlGroups"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataControlGroups")
	}

	return m, nil
}

func (m *_BACnetConstructedDataControlGroups) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataControlGroups) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataControlGroups"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataControlGroups")
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

		if err := WriteComplexTypeArrayField(ctx, "controlGroups", m.GetControlGroups(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'controlGroups' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataControlGroups"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataControlGroups")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataControlGroups) IsBACnetConstructedDataControlGroups() {}

func (m *_BACnetConstructedDataControlGroups) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataControlGroups) deepCopy() *_BACnetConstructedDataControlGroups {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataControlGroupsCopy := &_BACnetConstructedDataControlGroups{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetApplicationTagUnsignedInteger, BACnetApplicationTagUnsignedInteger](m.ControlGroups),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataControlGroupsCopy
}

func (m *_BACnetConstructedDataControlGroups) String() string {
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

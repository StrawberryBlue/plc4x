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

// BACnetConstructedDataCommandAction is the corresponding interface of BACnetConstructedDataCommandAction
type BACnetConstructedDataCommandAction interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetActionLists returns ActionLists (property field)
	GetActionLists() []BACnetActionList
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataCommandAction is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCommandAction()
	// CreateBuilder creates a BACnetConstructedDataCommandActionBuilder
	CreateBACnetConstructedDataCommandActionBuilder() BACnetConstructedDataCommandActionBuilder
}

// _BACnetConstructedDataCommandAction is the data-structure of this message
type _BACnetConstructedDataCommandAction struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	ActionLists          []BACnetActionList
}

var _ BACnetConstructedDataCommandAction = (*_BACnetConstructedDataCommandAction)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCommandAction)(nil)

// NewBACnetConstructedDataCommandAction factory function for _BACnetConstructedDataCommandAction
func NewBACnetConstructedDataCommandAction(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, actionLists []BACnetActionList, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCommandAction {
	_result := &_BACnetConstructedDataCommandAction{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		ActionLists:                   actionLists,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCommandActionBuilder is a builder for BACnetConstructedDataCommandAction
type BACnetConstructedDataCommandActionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(actionLists []BACnetActionList) BACnetConstructedDataCommandActionBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataCommandActionBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataCommandActionBuilder
	// WithActionLists adds ActionLists (property field)
	WithActionLists(...BACnetActionList) BACnetConstructedDataCommandActionBuilder
	// Build builds the BACnetConstructedDataCommandAction or returns an error if something is wrong
	Build() (BACnetConstructedDataCommandAction, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCommandAction
}

// NewBACnetConstructedDataCommandActionBuilder() creates a BACnetConstructedDataCommandActionBuilder
func NewBACnetConstructedDataCommandActionBuilder() BACnetConstructedDataCommandActionBuilder {
	return &_BACnetConstructedDataCommandActionBuilder{_BACnetConstructedDataCommandAction: new(_BACnetConstructedDataCommandAction)}
}

type _BACnetConstructedDataCommandActionBuilder struct {
	*_BACnetConstructedDataCommandAction

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCommandActionBuilder) = (*_BACnetConstructedDataCommandActionBuilder)(nil)

func (b *_BACnetConstructedDataCommandActionBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataCommandActionBuilder) WithMandatoryFields(actionLists []BACnetActionList) BACnetConstructedDataCommandActionBuilder {
	return b.WithActionLists(actionLists...)
}

func (b *_BACnetConstructedDataCommandActionBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataCommandActionBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataCommandActionBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataCommandActionBuilder {
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

func (b *_BACnetConstructedDataCommandActionBuilder) WithActionLists(actionLists ...BACnetActionList) BACnetConstructedDataCommandActionBuilder {
	b.ActionLists = actionLists
	return b
}

func (b *_BACnetConstructedDataCommandActionBuilder) Build() (BACnetConstructedDataCommandAction, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCommandAction.deepCopy(), nil
}

func (b *_BACnetConstructedDataCommandActionBuilder) MustBuild() BACnetConstructedDataCommandAction {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataCommandActionBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCommandActionBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCommandActionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCommandActionBuilder().(*_BACnetConstructedDataCommandActionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCommandActionBuilder creates a BACnetConstructedDataCommandActionBuilder
func (b *_BACnetConstructedDataCommandAction) CreateBACnetConstructedDataCommandActionBuilder() BACnetConstructedDataCommandActionBuilder {
	if b == nil {
		return NewBACnetConstructedDataCommandActionBuilder()
	}
	return &_BACnetConstructedDataCommandActionBuilder{_BACnetConstructedDataCommandAction: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCommandAction) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_COMMAND
}

func (m *_BACnetConstructedDataCommandAction) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCommandAction) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCommandAction) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataCommandAction) GetActionLists() []BACnetActionList {
	return m.ActionLists
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCommandAction) GetZero() uint64 {
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
func CastBACnetConstructedDataCommandAction(structType any) BACnetConstructedDataCommandAction {
	if casted, ok := structType.(BACnetConstructedDataCommandAction); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCommandAction); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCommandAction) GetTypeName() string {
	return "BACnetConstructedDataCommandAction"
}

func (m *_BACnetConstructedDataCommandAction) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.ActionLists) > 0 {
		for _, element := range m.ActionLists {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataCommandAction) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCommandAction) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCommandAction BACnetConstructedDataCommandAction, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCommandAction"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCommandAction")
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

	actionLists, err := ReadTerminatedArrayField[BACnetActionList](ctx, "actionLists", ReadComplex[BACnetActionList](BACnetActionListParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actionLists' field"))
	}
	m.ActionLists = actionLists

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCommandAction"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCommandAction")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCommandAction) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCommandAction) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCommandAction"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCommandAction")
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

		if err := WriteComplexTypeArrayField(ctx, "actionLists", m.GetActionLists(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'actionLists' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCommandAction"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCommandAction")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCommandAction) IsBACnetConstructedDataCommandAction() {}

func (m *_BACnetConstructedDataCommandAction) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCommandAction) deepCopy() *_BACnetConstructedDataCommandAction {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCommandActionCopy := &_BACnetConstructedDataCommandAction{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetActionList, BACnetActionList](m.ActionLists),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCommandActionCopy
}

func (m *_BACnetConstructedDataCommandAction) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

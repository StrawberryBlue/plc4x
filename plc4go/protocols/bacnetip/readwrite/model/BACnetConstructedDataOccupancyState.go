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

// BACnetConstructedDataOccupancyState is the corresponding interface of BACnetConstructedDataOccupancyState
type BACnetConstructedDataOccupancyState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetOccupancyState returns OccupancyState (property field)
	GetOccupancyState() BACnetAccessZoneOccupancyStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessZoneOccupancyStateTagged
	// IsBACnetConstructedDataOccupancyState is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataOccupancyState()
	// CreateBuilder creates a BACnetConstructedDataOccupancyStateBuilder
	CreateBACnetConstructedDataOccupancyStateBuilder() BACnetConstructedDataOccupancyStateBuilder
}

// _BACnetConstructedDataOccupancyState is the data-structure of this message
type _BACnetConstructedDataOccupancyState struct {
	BACnetConstructedDataContract
	OccupancyState BACnetAccessZoneOccupancyStateTagged
}

var _ BACnetConstructedDataOccupancyState = (*_BACnetConstructedDataOccupancyState)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataOccupancyState)(nil)

// NewBACnetConstructedDataOccupancyState factory function for _BACnetConstructedDataOccupancyState
func NewBACnetConstructedDataOccupancyState(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, occupancyState BACnetAccessZoneOccupancyStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOccupancyState {
	if occupancyState == nil {
		panic("occupancyState of type BACnetAccessZoneOccupancyStateTagged for BACnetConstructedDataOccupancyState must not be nil")
	}
	_result := &_BACnetConstructedDataOccupancyState{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		OccupancyState:                occupancyState,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataOccupancyStateBuilder is a builder for BACnetConstructedDataOccupancyState
type BACnetConstructedDataOccupancyStateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(occupancyState BACnetAccessZoneOccupancyStateTagged) BACnetConstructedDataOccupancyStateBuilder
	// WithOccupancyState adds OccupancyState (property field)
	WithOccupancyState(BACnetAccessZoneOccupancyStateTagged) BACnetConstructedDataOccupancyStateBuilder
	// WithOccupancyStateBuilder adds OccupancyState (property field) which is build by the builder
	WithOccupancyStateBuilder(func(BACnetAccessZoneOccupancyStateTaggedBuilder) BACnetAccessZoneOccupancyStateTaggedBuilder) BACnetConstructedDataOccupancyStateBuilder
	// Build builds the BACnetConstructedDataOccupancyState or returns an error if something is wrong
	Build() (BACnetConstructedDataOccupancyState, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataOccupancyState
}

// NewBACnetConstructedDataOccupancyStateBuilder() creates a BACnetConstructedDataOccupancyStateBuilder
func NewBACnetConstructedDataOccupancyStateBuilder() BACnetConstructedDataOccupancyStateBuilder {
	return &_BACnetConstructedDataOccupancyStateBuilder{_BACnetConstructedDataOccupancyState: new(_BACnetConstructedDataOccupancyState)}
}

type _BACnetConstructedDataOccupancyStateBuilder struct {
	*_BACnetConstructedDataOccupancyState

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataOccupancyStateBuilder) = (*_BACnetConstructedDataOccupancyStateBuilder)(nil)

func (b *_BACnetConstructedDataOccupancyStateBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataOccupancyStateBuilder) WithMandatoryFields(occupancyState BACnetAccessZoneOccupancyStateTagged) BACnetConstructedDataOccupancyStateBuilder {
	return b.WithOccupancyState(occupancyState)
}

func (b *_BACnetConstructedDataOccupancyStateBuilder) WithOccupancyState(occupancyState BACnetAccessZoneOccupancyStateTagged) BACnetConstructedDataOccupancyStateBuilder {
	b.OccupancyState = occupancyState
	return b
}

func (b *_BACnetConstructedDataOccupancyStateBuilder) WithOccupancyStateBuilder(builderSupplier func(BACnetAccessZoneOccupancyStateTaggedBuilder) BACnetAccessZoneOccupancyStateTaggedBuilder) BACnetConstructedDataOccupancyStateBuilder {
	builder := builderSupplier(b.OccupancyState.CreateBACnetAccessZoneOccupancyStateTaggedBuilder())
	var err error
	b.OccupancyState, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetAccessZoneOccupancyStateTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataOccupancyStateBuilder) Build() (BACnetConstructedDataOccupancyState, error) {
	if b.OccupancyState == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'occupancyState' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataOccupancyState.deepCopy(), nil
}

func (b *_BACnetConstructedDataOccupancyStateBuilder) MustBuild() BACnetConstructedDataOccupancyState {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataOccupancyStateBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataOccupancyStateBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataOccupancyStateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataOccupancyStateBuilder().(*_BACnetConstructedDataOccupancyStateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataOccupancyStateBuilder creates a BACnetConstructedDataOccupancyStateBuilder
func (b *_BACnetConstructedDataOccupancyState) CreateBACnetConstructedDataOccupancyStateBuilder() BACnetConstructedDataOccupancyStateBuilder {
	if b == nil {
		return NewBACnetConstructedDataOccupancyStateBuilder()
	}
	return &_BACnetConstructedDataOccupancyStateBuilder{_BACnetConstructedDataOccupancyState: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOccupancyState) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOccupancyState) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_STATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOccupancyState) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyState) GetOccupancyState() BACnetAccessZoneOccupancyStateTagged {
	return m.OccupancyState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyState) GetActualValue() BACnetAccessZoneOccupancyStateTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAccessZoneOccupancyStateTagged(m.GetOccupancyState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOccupancyState(structType any) BACnetConstructedDataOccupancyState {
	if casted, ok := structType.(BACnetConstructedDataOccupancyState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOccupancyState) GetTypeName() string {
	return "BACnetConstructedDataOccupancyState"
}

func (m *_BACnetConstructedDataOccupancyState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (occupancyState)
	lengthInBits += m.OccupancyState.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOccupancyState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataOccupancyState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataOccupancyState BACnetConstructedDataOccupancyState, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOccupancyState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	occupancyState, err := ReadSimpleField[BACnetAccessZoneOccupancyStateTagged](ctx, "occupancyState", ReadComplex[BACnetAccessZoneOccupancyStateTagged](BACnetAccessZoneOccupancyStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'occupancyState' field"))
	}
	m.OccupancyState = occupancyState

	actualValue, err := ReadVirtualField[BACnetAccessZoneOccupancyStateTagged](ctx, "actualValue", (*BACnetAccessZoneOccupancyStateTagged)(nil), occupancyState)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOccupancyState")
	}

	return m, nil
}

func (m *_BACnetConstructedDataOccupancyState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOccupancyState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOccupancyState")
		}

		if err := WriteSimpleField[BACnetAccessZoneOccupancyStateTagged](ctx, "occupancyState", m.GetOccupancyState(), WriteComplex[BACnetAccessZoneOccupancyStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'occupancyState' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOccupancyState")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOccupancyState) IsBACnetConstructedDataOccupancyState() {}

func (m *_BACnetConstructedDataOccupancyState) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataOccupancyState) deepCopy() *_BACnetConstructedDataOccupancyState {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataOccupancyStateCopy := &_BACnetConstructedDataOccupancyState{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.OccupancyState.DeepCopy().(BACnetAccessZoneOccupancyStateTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataOccupancyStateCopy
}

func (m *_BACnetConstructedDataOccupancyState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

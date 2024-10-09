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

// BACnetPropertyStatesDoorAlarmState is the corresponding interface of BACnetPropertyStatesDoorAlarmState
type BACnetPropertyStatesDoorAlarmState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetDoorAlarmState returns DoorAlarmState (property field)
	GetDoorAlarmState() BACnetDoorAlarmStateTagged
	// IsBACnetPropertyStatesDoorAlarmState is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesDoorAlarmState()
	// CreateBuilder creates a BACnetPropertyStatesDoorAlarmStateBuilder
	CreateBACnetPropertyStatesDoorAlarmStateBuilder() BACnetPropertyStatesDoorAlarmStateBuilder
}

// _BACnetPropertyStatesDoorAlarmState is the data-structure of this message
type _BACnetPropertyStatesDoorAlarmState struct {
	BACnetPropertyStatesContract
	DoorAlarmState BACnetDoorAlarmStateTagged
}

var _ BACnetPropertyStatesDoorAlarmState = (*_BACnetPropertyStatesDoorAlarmState)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesDoorAlarmState)(nil)

// NewBACnetPropertyStatesDoorAlarmState factory function for _BACnetPropertyStatesDoorAlarmState
func NewBACnetPropertyStatesDoorAlarmState(peekedTagHeader BACnetTagHeader, doorAlarmState BACnetDoorAlarmStateTagged) *_BACnetPropertyStatesDoorAlarmState {
	if doorAlarmState == nil {
		panic("doorAlarmState of type BACnetDoorAlarmStateTagged for BACnetPropertyStatesDoorAlarmState must not be nil")
	}
	_result := &_BACnetPropertyStatesDoorAlarmState{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		DoorAlarmState:               doorAlarmState,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesDoorAlarmStateBuilder is a builder for BACnetPropertyStatesDoorAlarmState
type BACnetPropertyStatesDoorAlarmStateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(doorAlarmState BACnetDoorAlarmStateTagged) BACnetPropertyStatesDoorAlarmStateBuilder
	// WithDoorAlarmState adds DoorAlarmState (property field)
	WithDoorAlarmState(BACnetDoorAlarmStateTagged) BACnetPropertyStatesDoorAlarmStateBuilder
	// WithDoorAlarmStateBuilder adds DoorAlarmState (property field) which is build by the builder
	WithDoorAlarmStateBuilder(func(BACnetDoorAlarmStateTaggedBuilder) BACnetDoorAlarmStateTaggedBuilder) BACnetPropertyStatesDoorAlarmStateBuilder
	// Build builds the BACnetPropertyStatesDoorAlarmState or returns an error if something is wrong
	Build() (BACnetPropertyStatesDoorAlarmState, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesDoorAlarmState
}

// NewBACnetPropertyStatesDoorAlarmStateBuilder() creates a BACnetPropertyStatesDoorAlarmStateBuilder
func NewBACnetPropertyStatesDoorAlarmStateBuilder() BACnetPropertyStatesDoorAlarmStateBuilder {
	return &_BACnetPropertyStatesDoorAlarmStateBuilder{_BACnetPropertyStatesDoorAlarmState: new(_BACnetPropertyStatesDoorAlarmState)}
}

type _BACnetPropertyStatesDoorAlarmStateBuilder struct {
	*_BACnetPropertyStatesDoorAlarmState

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesDoorAlarmStateBuilder) = (*_BACnetPropertyStatesDoorAlarmStateBuilder)(nil)

func (b *_BACnetPropertyStatesDoorAlarmStateBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesDoorAlarmStateBuilder) WithMandatoryFields(doorAlarmState BACnetDoorAlarmStateTagged) BACnetPropertyStatesDoorAlarmStateBuilder {
	return b.WithDoorAlarmState(doorAlarmState)
}

func (b *_BACnetPropertyStatesDoorAlarmStateBuilder) WithDoorAlarmState(doorAlarmState BACnetDoorAlarmStateTagged) BACnetPropertyStatesDoorAlarmStateBuilder {
	b.DoorAlarmState = doorAlarmState
	return b
}

func (b *_BACnetPropertyStatesDoorAlarmStateBuilder) WithDoorAlarmStateBuilder(builderSupplier func(BACnetDoorAlarmStateTaggedBuilder) BACnetDoorAlarmStateTaggedBuilder) BACnetPropertyStatesDoorAlarmStateBuilder {
	builder := builderSupplier(b.DoorAlarmState.CreateBACnetDoorAlarmStateTaggedBuilder())
	var err error
	b.DoorAlarmState, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDoorAlarmStateTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesDoorAlarmStateBuilder) Build() (BACnetPropertyStatesDoorAlarmState, error) {
	if b.DoorAlarmState == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'doorAlarmState' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesDoorAlarmState.deepCopy(), nil
}

func (b *_BACnetPropertyStatesDoorAlarmStateBuilder) MustBuild() BACnetPropertyStatesDoorAlarmState {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPropertyStatesDoorAlarmStateBuilder) Done() BACnetPropertyStatesBuilder {
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesDoorAlarmStateBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesDoorAlarmStateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesDoorAlarmStateBuilder().(*_BACnetPropertyStatesDoorAlarmStateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesDoorAlarmStateBuilder creates a BACnetPropertyStatesDoorAlarmStateBuilder
func (b *_BACnetPropertyStatesDoorAlarmState) CreateBACnetPropertyStatesDoorAlarmStateBuilder() BACnetPropertyStatesDoorAlarmStateBuilder {
	if b == nil {
		return NewBACnetPropertyStatesDoorAlarmStateBuilder()
	}
	return &_BACnetPropertyStatesDoorAlarmStateBuilder{_BACnetPropertyStatesDoorAlarmState: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesDoorAlarmState) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesDoorAlarmState) GetDoorAlarmState() BACnetDoorAlarmStateTagged {
	return m.DoorAlarmState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesDoorAlarmState(structType any) BACnetPropertyStatesDoorAlarmState {
	if casted, ok := structType.(BACnetPropertyStatesDoorAlarmState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesDoorAlarmState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesDoorAlarmState) GetTypeName() string {
	return "BACnetPropertyStatesDoorAlarmState"
}

func (m *_BACnetPropertyStatesDoorAlarmState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (doorAlarmState)
	lengthInBits += m.DoorAlarmState.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesDoorAlarmState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesDoorAlarmState) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesDoorAlarmState BACnetPropertyStatesDoorAlarmState, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesDoorAlarmState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesDoorAlarmState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doorAlarmState, err := ReadSimpleField[BACnetDoorAlarmStateTagged](ctx, "doorAlarmState", ReadComplex[BACnetDoorAlarmStateTagged](BACnetDoorAlarmStateTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doorAlarmState' field"))
	}
	m.DoorAlarmState = doorAlarmState

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesDoorAlarmState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesDoorAlarmState")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesDoorAlarmState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesDoorAlarmState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesDoorAlarmState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesDoorAlarmState")
		}

		if err := WriteSimpleField[BACnetDoorAlarmStateTagged](ctx, "doorAlarmState", m.GetDoorAlarmState(), WriteComplex[BACnetDoorAlarmStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doorAlarmState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesDoorAlarmState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesDoorAlarmState")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesDoorAlarmState) IsBACnetPropertyStatesDoorAlarmState() {}

func (m *_BACnetPropertyStatesDoorAlarmState) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesDoorAlarmState) deepCopy() *_BACnetPropertyStatesDoorAlarmState {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesDoorAlarmStateCopy := &_BACnetPropertyStatesDoorAlarmState{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.DoorAlarmState.DeepCopy().(BACnetDoorAlarmStateTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesDoorAlarmStateCopy
}

func (m *_BACnetPropertyStatesDoorAlarmState) String() string {
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

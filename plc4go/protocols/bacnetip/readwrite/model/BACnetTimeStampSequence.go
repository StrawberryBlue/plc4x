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

// BACnetTimeStampSequence is the corresponding interface of BACnetTimeStampSequence
type BACnetTimeStampSequence interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetTimeStamp
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() BACnetContextTagUnsignedInteger
	// IsBACnetTimeStampSequence is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimeStampSequence()
	// CreateBuilder creates a BACnetTimeStampSequenceBuilder
	CreateBACnetTimeStampSequenceBuilder() BACnetTimeStampSequenceBuilder
}

// _BACnetTimeStampSequence is the data-structure of this message
type _BACnetTimeStampSequence struct {
	BACnetTimeStampContract
	SequenceNumber BACnetContextTagUnsignedInteger
}

var _ BACnetTimeStampSequence = (*_BACnetTimeStampSequence)(nil)
var _ BACnetTimeStampRequirements = (*_BACnetTimeStampSequence)(nil)

// NewBACnetTimeStampSequence factory function for _BACnetTimeStampSequence
func NewBACnetTimeStampSequence(peekedTagHeader BACnetTagHeader, sequenceNumber BACnetContextTagUnsignedInteger) *_BACnetTimeStampSequence {
	if sequenceNumber == nil {
		panic("sequenceNumber of type BACnetContextTagUnsignedInteger for BACnetTimeStampSequence must not be nil")
	}
	_result := &_BACnetTimeStampSequence{
		BACnetTimeStampContract: NewBACnetTimeStamp(peekedTagHeader),
		SequenceNumber:          sequenceNumber,
	}
	_result.BACnetTimeStampContract.(*_BACnetTimeStamp)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimeStampSequenceBuilder is a builder for BACnetTimeStampSequence
type BACnetTimeStampSequenceBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(sequenceNumber BACnetContextTagUnsignedInteger) BACnetTimeStampSequenceBuilder
	// WithSequenceNumber adds SequenceNumber (property field)
	WithSequenceNumber(BACnetContextTagUnsignedInteger) BACnetTimeStampSequenceBuilder
	// WithSequenceNumberBuilder adds SequenceNumber (property field) which is build by the builder
	WithSequenceNumberBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetTimeStampSequenceBuilder
	// Build builds the BACnetTimeStampSequence or returns an error if something is wrong
	Build() (BACnetTimeStampSequence, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimeStampSequence
}

// NewBACnetTimeStampSequenceBuilder() creates a BACnetTimeStampSequenceBuilder
func NewBACnetTimeStampSequenceBuilder() BACnetTimeStampSequenceBuilder {
	return &_BACnetTimeStampSequenceBuilder{_BACnetTimeStampSequence: new(_BACnetTimeStampSequence)}
}

type _BACnetTimeStampSequenceBuilder struct {
	*_BACnetTimeStampSequence

	parentBuilder *_BACnetTimeStampBuilder

	err *utils.MultiError
}

var _ (BACnetTimeStampSequenceBuilder) = (*_BACnetTimeStampSequenceBuilder)(nil)

func (b *_BACnetTimeStampSequenceBuilder) setParent(contract BACnetTimeStampContract) {
	b.BACnetTimeStampContract = contract
}

func (b *_BACnetTimeStampSequenceBuilder) WithMandatoryFields(sequenceNumber BACnetContextTagUnsignedInteger) BACnetTimeStampSequenceBuilder {
	return b.WithSequenceNumber(sequenceNumber)
}

func (b *_BACnetTimeStampSequenceBuilder) WithSequenceNumber(sequenceNumber BACnetContextTagUnsignedInteger) BACnetTimeStampSequenceBuilder {
	b.SequenceNumber = sequenceNumber
	return b
}

func (b *_BACnetTimeStampSequenceBuilder) WithSequenceNumberBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetTimeStampSequenceBuilder {
	builder := builderSupplier(b.SequenceNumber.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.SequenceNumber, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetTimeStampSequenceBuilder) Build() (BACnetTimeStampSequence, error) {
	if b.SequenceNumber == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'sequenceNumber' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTimeStampSequence.deepCopy(), nil
}

func (b *_BACnetTimeStampSequenceBuilder) MustBuild() BACnetTimeStampSequence {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetTimeStampSequenceBuilder) Done() BACnetTimeStampBuilder {
	return b.parentBuilder
}

func (b *_BACnetTimeStampSequenceBuilder) buildForBACnetTimeStamp() (BACnetTimeStamp, error) {
	return b.Build()
}

func (b *_BACnetTimeStampSequenceBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTimeStampSequenceBuilder().(*_BACnetTimeStampSequenceBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTimeStampSequenceBuilder creates a BACnetTimeStampSequenceBuilder
func (b *_BACnetTimeStampSequence) CreateBACnetTimeStampSequenceBuilder() BACnetTimeStampSequenceBuilder {
	if b == nil {
		return NewBACnetTimeStampSequenceBuilder()
	}
	return &_BACnetTimeStampSequenceBuilder{_BACnetTimeStampSequence: b.deepCopy()}
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

func (m *_BACnetTimeStampSequence) GetParent() BACnetTimeStampContract {
	return m.BACnetTimeStampContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimeStampSequence) GetSequenceNumber() BACnetContextTagUnsignedInteger {
	return m.SequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimeStampSequence(structType any) BACnetTimeStampSequence {
	if casted, ok := structType.(BACnetTimeStampSequence); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimeStampSequence); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimeStampSequence) GetTypeName() string {
	return "BACnetTimeStampSequence"
}

func (m *_BACnetTimeStampSequence) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimeStampContract.(*_BACnetTimeStamp).getLengthInBits(ctx))

	// Simple field (sequenceNumber)
	lengthInBits += m.SequenceNumber.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimeStampSequence) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimeStampSequence) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimeStamp) (__bACnetTimeStampSequence BACnetTimeStampSequence, err error) {
	m.BACnetTimeStampContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimeStampSequence"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimeStampSequence")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sequenceNumber, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "sequenceNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}
	m.SequenceNumber = sequenceNumber

	if closeErr := readBuffer.CloseContext("BACnetTimeStampSequence"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimeStampSequence")
	}

	return m, nil
}

func (m *_BACnetTimeStampSequence) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimeStampSequence) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimeStampSequence"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimeStampSequence")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "sequenceNumber", m.GetSequenceNumber(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'sequenceNumber' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimeStampSequence"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimeStampSequence")
		}
		return nil
	}
	return m.BACnetTimeStampContract.(*_BACnetTimeStamp).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimeStampSequence) IsBACnetTimeStampSequence() {}

func (m *_BACnetTimeStampSequence) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimeStampSequence) deepCopy() *_BACnetTimeStampSequence {
	if m == nil {
		return nil
	}
	_BACnetTimeStampSequenceCopy := &_BACnetTimeStampSequence{
		m.BACnetTimeStampContract.(*_BACnetTimeStamp).deepCopy(),
		m.SequenceNumber.DeepCopy().(BACnetContextTagUnsignedInteger),
	}
	m.BACnetTimeStampContract.(*_BACnetTimeStamp)._SubType = m
	return _BACnetTimeStampSequenceCopy
}

func (m *_BACnetTimeStampSequence) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

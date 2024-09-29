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

// BACnetEventLogRecordLogDatumTimeChange is the corresponding interface of BACnetEventLogRecordLogDatumTimeChange
type BACnetEventLogRecordLogDatumTimeChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventLogRecordLogDatum
	// GetTimeChange returns TimeChange (property field)
	GetTimeChange() BACnetContextTagReal
	// IsBACnetEventLogRecordLogDatumTimeChange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventLogRecordLogDatumTimeChange()
	// CreateBuilder creates a BACnetEventLogRecordLogDatumTimeChangeBuilder
	CreateBACnetEventLogRecordLogDatumTimeChangeBuilder() BACnetEventLogRecordLogDatumTimeChangeBuilder
}

// _BACnetEventLogRecordLogDatumTimeChange is the data-structure of this message
type _BACnetEventLogRecordLogDatumTimeChange struct {
	BACnetEventLogRecordLogDatumContract
	TimeChange BACnetContextTagReal
}

var _ BACnetEventLogRecordLogDatumTimeChange = (*_BACnetEventLogRecordLogDatumTimeChange)(nil)
var _ BACnetEventLogRecordLogDatumRequirements = (*_BACnetEventLogRecordLogDatumTimeChange)(nil)

// NewBACnetEventLogRecordLogDatumTimeChange factory function for _BACnetEventLogRecordLogDatumTimeChange
func NewBACnetEventLogRecordLogDatumTimeChange(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, timeChange BACnetContextTagReal, tagNumber uint8) *_BACnetEventLogRecordLogDatumTimeChange {
	if timeChange == nil {
		panic("timeChange of type BACnetContextTagReal for BACnetEventLogRecordLogDatumTimeChange must not be nil")
	}
	_result := &_BACnetEventLogRecordLogDatumTimeChange{
		BACnetEventLogRecordLogDatumContract: NewBACnetEventLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
		TimeChange:                           timeChange,
	}
	_result.BACnetEventLogRecordLogDatumContract.(*_BACnetEventLogRecordLogDatum)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventLogRecordLogDatumTimeChangeBuilder is a builder for BACnetEventLogRecordLogDatumTimeChange
type BACnetEventLogRecordLogDatumTimeChangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timeChange BACnetContextTagReal) BACnetEventLogRecordLogDatumTimeChangeBuilder
	// WithTimeChange adds TimeChange (property field)
	WithTimeChange(BACnetContextTagReal) BACnetEventLogRecordLogDatumTimeChangeBuilder
	// WithTimeChangeBuilder adds TimeChange (property field) which is build by the builder
	WithTimeChangeBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetEventLogRecordLogDatumTimeChangeBuilder
	// Build builds the BACnetEventLogRecordLogDatumTimeChange or returns an error if something is wrong
	Build() (BACnetEventLogRecordLogDatumTimeChange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventLogRecordLogDatumTimeChange
}

// NewBACnetEventLogRecordLogDatumTimeChangeBuilder() creates a BACnetEventLogRecordLogDatumTimeChangeBuilder
func NewBACnetEventLogRecordLogDatumTimeChangeBuilder() BACnetEventLogRecordLogDatumTimeChangeBuilder {
	return &_BACnetEventLogRecordLogDatumTimeChangeBuilder{_BACnetEventLogRecordLogDatumTimeChange: new(_BACnetEventLogRecordLogDatumTimeChange)}
}

type _BACnetEventLogRecordLogDatumTimeChangeBuilder struct {
	*_BACnetEventLogRecordLogDatumTimeChange

	parentBuilder *_BACnetEventLogRecordLogDatumBuilder

	err *utils.MultiError
}

var _ (BACnetEventLogRecordLogDatumTimeChangeBuilder) = (*_BACnetEventLogRecordLogDatumTimeChangeBuilder)(nil)

func (b *_BACnetEventLogRecordLogDatumTimeChangeBuilder) setParent(contract BACnetEventLogRecordLogDatumContract) {
	b.BACnetEventLogRecordLogDatumContract = contract
}

func (b *_BACnetEventLogRecordLogDatumTimeChangeBuilder) WithMandatoryFields(timeChange BACnetContextTagReal) BACnetEventLogRecordLogDatumTimeChangeBuilder {
	return b.WithTimeChange(timeChange)
}

func (b *_BACnetEventLogRecordLogDatumTimeChangeBuilder) WithTimeChange(timeChange BACnetContextTagReal) BACnetEventLogRecordLogDatumTimeChangeBuilder {
	b.TimeChange = timeChange
	return b
}

func (b *_BACnetEventLogRecordLogDatumTimeChangeBuilder) WithTimeChangeBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetEventLogRecordLogDatumTimeChangeBuilder {
	builder := builderSupplier(b.TimeChange.CreateBACnetContextTagRealBuilder())
	var err error
	b.TimeChange, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetEventLogRecordLogDatumTimeChangeBuilder) Build() (BACnetEventLogRecordLogDatumTimeChange, error) {
	if b.TimeChange == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeChange' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetEventLogRecordLogDatumTimeChange.deepCopy(), nil
}

func (b *_BACnetEventLogRecordLogDatumTimeChangeBuilder) MustBuild() BACnetEventLogRecordLogDatumTimeChange {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetEventLogRecordLogDatumTimeChangeBuilder) Done() BACnetEventLogRecordLogDatumBuilder {
	return b.parentBuilder
}

func (b *_BACnetEventLogRecordLogDatumTimeChangeBuilder) buildForBACnetEventLogRecordLogDatum() (BACnetEventLogRecordLogDatum, error) {
	return b.Build()
}

func (b *_BACnetEventLogRecordLogDatumTimeChangeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventLogRecordLogDatumTimeChangeBuilder().(*_BACnetEventLogRecordLogDatumTimeChangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventLogRecordLogDatumTimeChangeBuilder creates a BACnetEventLogRecordLogDatumTimeChangeBuilder
func (b *_BACnetEventLogRecordLogDatumTimeChange) CreateBACnetEventLogRecordLogDatumTimeChangeBuilder() BACnetEventLogRecordLogDatumTimeChangeBuilder {
	if b == nil {
		return NewBACnetEventLogRecordLogDatumTimeChangeBuilder()
	}
	return &_BACnetEventLogRecordLogDatumTimeChangeBuilder{_BACnetEventLogRecordLogDatumTimeChange: b.deepCopy()}
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

func (m *_BACnetEventLogRecordLogDatumTimeChange) GetParent() BACnetEventLogRecordLogDatumContract {
	return m.BACnetEventLogRecordLogDatumContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventLogRecordLogDatumTimeChange) GetTimeChange() BACnetContextTagReal {
	return m.TimeChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventLogRecordLogDatumTimeChange(structType any) BACnetEventLogRecordLogDatumTimeChange {
	if casted, ok := structType.(BACnetEventLogRecordLogDatumTimeChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventLogRecordLogDatumTimeChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) GetTypeName() string {
	return "BACnetEventLogRecordLogDatumTimeChange"
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventLogRecordLogDatumContract.(*_BACnetEventLogRecordLogDatum).getLengthInBits(ctx))

	// Simple field (timeChange)
	lengthInBits += m.TimeChange.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventLogRecordLogDatum, tagNumber uint8) (__bACnetEventLogRecordLogDatumTimeChange BACnetEventLogRecordLogDatumTimeChange, err error) {
	m.BACnetEventLogRecordLogDatumContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventLogRecordLogDatumTimeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventLogRecordLogDatumTimeChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeChange, err := ReadSimpleField[BACnetContextTagReal](ctx, "timeChange", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeChange' field"))
	}
	m.TimeChange = timeChange

	if closeErr := readBuffer.CloseContext("BACnetEventLogRecordLogDatumTimeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventLogRecordLogDatumTimeChange")
	}

	return m, nil
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventLogRecordLogDatumTimeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventLogRecordLogDatumTimeChange")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "timeChange", m.GetTimeChange(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeChange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventLogRecordLogDatumTimeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventLogRecordLogDatumTimeChange")
		}
		return nil
	}
	return m.BACnetEventLogRecordLogDatumContract.(*_BACnetEventLogRecordLogDatum).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) IsBACnetEventLogRecordLogDatumTimeChange() {}

func (m *_BACnetEventLogRecordLogDatumTimeChange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) deepCopy() *_BACnetEventLogRecordLogDatumTimeChange {
	if m == nil {
		return nil
	}
	_BACnetEventLogRecordLogDatumTimeChangeCopy := &_BACnetEventLogRecordLogDatumTimeChange{
		m.BACnetEventLogRecordLogDatumContract.(*_BACnetEventLogRecordLogDatum).deepCopy(),
		m.TimeChange.DeepCopy().(BACnetContextTagReal),
	}
	m.BACnetEventLogRecordLogDatumContract.(*_BACnetEventLogRecordLogDatum)._SubType = m
	return _BACnetEventLogRecordLogDatumTimeChangeCopy
}

func (m *_BACnetEventLogRecordLogDatumTimeChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

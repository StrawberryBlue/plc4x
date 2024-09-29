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

// BACnetLogDataLogDataEntryUnsignedValue is the corresponding interface of BACnetLogDataLogDataEntryUnsignedValue
type BACnetLogDataLogDataEntryUnsignedValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetLogDataLogDataEntry
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetContextTagUnsignedInteger
	// IsBACnetLogDataLogDataEntryUnsignedValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogDataLogDataEntryUnsignedValue()
	// CreateBuilder creates a BACnetLogDataLogDataEntryUnsignedValueBuilder
	CreateBACnetLogDataLogDataEntryUnsignedValueBuilder() BACnetLogDataLogDataEntryUnsignedValueBuilder
}

// _BACnetLogDataLogDataEntryUnsignedValue is the data-structure of this message
type _BACnetLogDataLogDataEntryUnsignedValue struct {
	BACnetLogDataLogDataEntryContract
	UnsignedValue BACnetContextTagUnsignedInteger
}

var _ BACnetLogDataLogDataEntryUnsignedValue = (*_BACnetLogDataLogDataEntryUnsignedValue)(nil)
var _ BACnetLogDataLogDataEntryRequirements = (*_BACnetLogDataLogDataEntryUnsignedValue)(nil)

// NewBACnetLogDataLogDataEntryUnsignedValue factory function for _BACnetLogDataLogDataEntryUnsignedValue
func NewBACnetLogDataLogDataEntryUnsignedValue(peekedTagHeader BACnetTagHeader, unsignedValue BACnetContextTagUnsignedInteger) *_BACnetLogDataLogDataEntryUnsignedValue {
	if unsignedValue == nil {
		panic("unsignedValue of type BACnetContextTagUnsignedInteger for BACnetLogDataLogDataEntryUnsignedValue must not be nil")
	}
	_result := &_BACnetLogDataLogDataEntryUnsignedValue{
		BACnetLogDataLogDataEntryContract: NewBACnetLogDataLogDataEntry(peekedTagHeader),
		UnsignedValue:                     unsignedValue,
	}
	_result.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLogDataLogDataEntryUnsignedValueBuilder is a builder for BACnetLogDataLogDataEntryUnsignedValue
type BACnetLogDataLogDataEntryUnsignedValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(unsignedValue BACnetContextTagUnsignedInteger) BACnetLogDataLogDataEntryUnsignedValueBuilder
	// WithUnsignedValue adds UnsignedValue (property field)
	WithUnsignedValue(BACnetContextTagUnsignedInteger) BACnetLogDataLogDataEntryUnsignedValueBuilder
	// WithUnsignedValueBuilder adds UnsignedValue (property field) which is build by the builder
	WithUnsignedValueBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetLogDataLogDataEntryUnsignedValueBuilder
	// Build builds the BACnetLogDataLogDataEntryUnsignedValue or returns an error if something is wrong
	Build() (BACnetLogDataLogDataEntryUnsignedValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLogDataLogDataEntryUnsignedValue
}

// NewBACnetLogDataLogDataEntryUnsignedValueBuilder() creates a BACnetLogDataLogDataEntryUnsignedValueBuilder
func NewBACnetLogDataLogDataEntryUnsignedValueBuilder() BACnetLogDataLogDataEntryUnsignedValueBuilder {
	return &_BACnetLogDataLogDataEntryUnsignedValueBuilder{_BACnetLogDataLogDataEntryUnsignedValue: new(_BACnetLogDataLogDataEntryUnsignedValue)}
}

type _BACnetLogDataLogDataEntryUnsignedValueBuilder struct {
	*_BACnetLogDataLogDataEntryUnsignedValue

	parentBuilder *_BACnetLogDataLogDataEntryBuilder

	err *utils.MultiError
}

var _ (BACnetLogDataLogDataEntryUnsignedValueBuilder) = (*_BACnetLogDataLogDataEntryUnsignedValueBuilder)(nil)

func (b *_BACnetLogDataLogDataEntryUnsignedValueBuilder) setParent(contract BACnetLogDataLogDataEntryContract) {
	b.BACnetLogDataLogDataEntryContract = contract
}

func (b *_BACnetLogDataLogDataEntryUnsignedValueBuilder) WithMandatoryFields(unsignedValue BACnetContextTagUnsignedInteger) BACnetLogDataLogDataEntryUnsignedValueBuilder {
	return b.WithUnsignedValue(unsignedValue)
}

func (b *_BACnetLogDataLogDataEntryUnsignedValueBuilder) WithUnsignedValue(unsignedValue BACnetContextTagUnsignedInteger) BACnetLogDataLogDataEntryUnsignedValueBuilder {
	b.UnsignedValue = unsignedValue
	return b
}

func (b *_BACnetLogDataLogDataEntryUnsignedValueBuilder) WithUnsignedValueBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetLogDataLogDataEntryUnsignedValueBuilder {
	builder := builderSupplier(b.UnsignedValue.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.UnsignedValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetLogDataLogDataEntryUnsignedValueBuilder) Build() (BACnetLogDataLogDataEntryUnsignedValue, error) {
	if b.UnsignedValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'unsignedValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLogDataLogDataEntryUnsignedValue.deepCopy(), nil
}

func (b *_BACnetLogDataLogDataEntryUnsignedValueBuilder) MustBuild() BACnetLogDataLogDataEntryUnsignedValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetLogDataLogDataEntryUnsignedValueBuilder) Done() BACnetLogDataLogDataEntryBuilder {
	return b.parentBuilder
}

func (b *_BACnetLogDataLogDataEntryUnsignedValueBuilder) buildForBACnetLogDataLogDataEntry() (BACnetLogDataLogDataEntry, error) {
	return b.Build()
}

func (b *_BACnetLogDataLogDataEntryUnsignedValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLogDataLogDataEntryUnsignedValueBuilder().(*_BACnetLogDataLogDataEntryUnsignedValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLogDataLogDataEntryUnsignedValueBuilder creates a BACnetLogDataLogDataEntryUnsignedValueBuilder
func (b *_BACnetLogDataLogDataEntryUnsignedValue) CreateBACnetLogDataLogDataEntryUnsignedValueBuilder() BACnetLogDataLogDataEntryUnsignedValueBuilder {
	if b == nil {
		return NewBACnetLogDataLogDataEntryUnsignedValueBuilder()
	}
	return &_BACnetLogDataLogDataEntryUnsignedValueBuilder{_BACnetLogDataLogDataEntryUnsignedValue: b.deepCopy()}
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

func (m *_BACnetLogDataLogDataEntryUnsignedValue) GetParent() BACnetLogDataLogDataEntryContract {
	return m.BACnetLogDataLogDataEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryUnsignedValue) GetUnsignedValue() BACnetContextTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryUnsignedValue(structType any) BACnetLogDataLogDataEntryUnsignedValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryUnsignedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryUnsignedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryUnsignedValue"
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).getLengthInBits(ctx))

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetLogDataLogDataEntry) (__bACnetLogDataLogDataEntryUnsignedValue BACnetLogDataLogDataEntryUnsignedValue, err error) {
	m.BACnetLogDataLogDataEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryUnsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryUnsignedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unsignedValue, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "unsignedValue", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unsignedValue' field"))
	}
	m.UnsignedValue = unsignedValue

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryUnsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryUnsignedValue")
	}

	return m, nil
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryUnsignedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryUnsignedValue")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "unsignedValue", m.GetUnsignedValue(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryUnsignedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryUnsignedValue")
		}
		return nil
	}
	return m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) IsBACnetLogDataLogDataEntryUnsignedValue() {}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) deepCopy() *_BACnetLogDataLogDataEntryUnsignedValue {
	if m == nil {
		return nil
	}
	_BACnetLogDataLogDataEntryUnsignedValueCopy := &_BACnetLogDataLogDataEntryUnsignedValue{
		m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry).deepCopy(),
		m.UnsignedValue.DeepCopy().(BACnetContextTagUnsignedInteger),
	}
	m.BACnetLogDataLogDataEntryContract.(*_BACnetLogDataLogDataEntry)._SubType = m
	return _BACnetLogDataLogDataEntryUnsignedValueCopy
}

func (m *_BACnetLogDataLogDataEntryUnsignedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetTagPayloadBoolean is the corresponding interface of BACnetTagPayloadBoolean
type BACnetTagPayloadBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetValue returns Value (virtual field)
	GetValue() bool
	// GetIsTrue returns IsTrue (virtual field)
	GetIsTrue() bool
	// GetIsFalse returns IsFalse (virtual field)
	GetIsFalse() bool
	// IsBACnetTagPayloadBoolean is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTagPayloadBoolean()
	// CreateBuilder creates a BACnetTagPayloadBooleanBuilder
	CreateBACnetTagPayloadBooleanBuilder() BACnetTagPayloadBooleanBuilder
}

// _BACnetTagPayloadBoolean is the data-structure of this message
type _BACnetTagPayloadBoolean struct {

	// Arguments.
	ActualLength uint32
}

var _ BACnetTagPayloadBoolean = (*_BACnetTagPayloadBoolean)(nil)

// NewBACnetTagPayloadBoolean factory function for _BACnetTagPayloadBoolean
func NewBACnetTagPayloadBoolean(actualLength uint32) *_BACnetTagPayloadBoolean {
	return &_BACnetTagPayloadBoolean{ActualLength: actualLength}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTagPayloadBooleanBuilder is a builder for BACnetTagPayloadBoolean
type BACnetTagPayloadBooleanBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetTagPayloadBooleanBuilder
	// Build builds the BACnetTagPayloadBoolean or returns an error if something is wrong
	Build() (BACnetTagPayloadBoolean, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTagPayloadBoolean
}

// NewBACnetTagPayloadBooleanBuilder() creates a BACnetTagPayloadBooleanBuilder
func NewBACnetTagPayloadBooleanBuilder() BACnetTagPayloadBooleanBuilder {
	return &_BACnetTagPayloadBooleanBuilder{_BACnetTagPayloadBoolean: new(_BACnetTagPayloadBoolean)}
}

type _BACnetTagPayloadBooleanBuilder struct {
	*_BACnetTagPayloadBoolean

	err *utils.MultiError
}

var _ (BACnetTagPayloadBooleanBuilder) = (*_BACnetTagPayloadBooleanBuilder)(nil)

func (b *_BACnetTagPayloadBooleanBuilder) WithMandatoryFields() BACnetTagPayloadBooleanBuilder {
	return b
}

func (b *_BACnetTagPayloadBooleanBuilder) Build() (BACnetTagPayloadBoolean, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTagPayloadBoolean.deepCopy(), nil
}

func (b *_BACnetTagPayloadBooleanBuilder) MustBuild() BACnetTagPayloadBoolean {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetTagPayloadBooleanBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTagPayloadBooleanBuilder().(*_BACnetTagPayloadBooleanBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTagPayloadBooleanBuilder creates a BACnetTagPayloadBooleanBuilder
func (b *_BACnetTagPayloadBoolean) CreateBACnetTagPayloadBooleanBuilder() BACnetTagPayloadBooleanBuilder {
	if b == nil {
		return NewBACnetTagPayloadBooleanBuilder()
	}
	return &_BACnetTagPayloadBooleanBuilder{_BACnetTagPayloadBoolean: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetTagPayloadBoolean) GetValue() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetActualLength()) == (1)))
}

func (m *_BACnetTagPayloadBoolean) GetIsTrue() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetValue())
}

func (m *_BACnetTagPayloadBoolean) GetIsFalse() bool {
	ctx := context.Background()
	_ = ctx
	return bool(!(m.GetValue()))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadBoolean(structType any) BACnetTagPayloadBoolean {
	if casted, ok := structType.(BACnetTagPayloadBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadBoolean) GetTypeName() string {
	return "BACnetTagPayloadBoolean"
}

func (m *_BACnetTagPayloadBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetTagPayloadBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagPayloadBooleanParse(ctx context.Context, theBytes []byte, actualLength uint32) (BACnetTagPayloadBoolean, error) {
	return BACnetTagPayloadBooleanParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), actualLength)
}

func BACnetTagPayloadBooleanParseWithBufferProducer(actualLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadBoolean, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadBoolean, error) {
		return BACnetTagPayloadBooleanParseWithBuffer(ctx, readBuffer, actualLength)
	}
}

func BACnetTagPayloadBooleanParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (BACnetTagPayloadBoolean, error) {
	v, err := (&_BACnetTagPayloadBoolean{ActualLength: actualLength}).parse(ctx, readBuffer, actualLength)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTagPayloadBoolean) parse(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (__bACnetTagPayloadBoolean BACnetTagPayloadBoolean, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	value, err := ReadVirtualField[bool](ctx, "value", (*bool)(nil), bool((actualLength) == (1)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	_ = value

	isTrue, err := ReadVirtualField[bool](ctx, "isTrue", (*bool)(nil), value)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isTrue' field"))
	}
	_ = isTrue

	isFalse, err := ReadVirtualField[bool](ctx, "isFalse", (*bool)(nil), !(value))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isFalse' field"))
	}
	_ = isFalse

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadBoolean")
	}

	return m, nil
}

func (m *_BACnetTagPayloadBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadBoolean"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadBoolean")
	}
	// Virtual field
	value := m.GetValue()
	_ = value
	if _valueErr := writeBuffer.WriteVirtual(ctx, "value", m.GetValue()); _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}
	// Virtual field
	isTrue := m.GetIsTrue()
	_ = isTrue
	if _isTrueErr := writeBuffer.WriteVirtual(ctx, "isTrue", m.GetIsTrue()); _isTrueErr != nil {
		return errors.Wrap(_isTrueErr, "Error serializing 'isTrue' field")
	}
	// Virtual field
	isFalse := m.GetIsFalse()
	_ = isFalse
	if _isFalseErr := writeBuffer.WriteVirtual(ctx, "isFalse", m.GetIsFalse()); _isFalseErr != nil {
		return errors.Wrap(_isFalseErr, "Error serializing 'isFalse' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadBoolean"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadBoolean")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetTagPayloadBoolean) GetActualLength() uint32 {
	return m.ActualLength
}

//
////

func (m *_BACnetTagPayloadBoolean) IsBACnetTagPayloadBoolean() {}

func (m *_BACnetTagPayloadBoolean) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTagPayloadBoolean) deepCopy() *_BACnetTagPayloadBoolean {
	if m == nil {
		return nil
	}
	_BACnetTagPayloadBooleanCopy := &_BACnetTagPayloadBoolean{
		m.ActualLength,
	}
	return _BACnetTagPayloadBooleanCopy
}

func (m *_BACnetTagPayloadBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

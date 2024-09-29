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

// PascalByteString is the corresponding interface of PascalByteString
type PascalByteString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetStringLength returns StringLength (property field)
	GetStringLength() int32
	// GetStringValue returns StringValue (property field)
	GetStringValue() []byte
	// IsPascalByteString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPascalByteString()
	// CreateBuilder creates a PascalByteStringBuilder
	CreatePascalByteStringBuilder() PascalByteStringBuilder
}

// _PascalByteString is the data-structure of this message
type _PascalByteString struct {
	StringLength int32
	StringValue  []byte
}

var _ PascalByteString = (*_PascalByteString)(nil)

// NewPascalByteString factory function for _PascalByteString
func NewPascalByteString(stringLength int32, stringValue []byte) *_PascalByteString {
	return &_PascalByteString{StringLength: stringLength, StringValue: stringValue}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// PascalByteStringBuilder is a builder for PascalByteString
type PascalByteStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(stringLength int32, stringValue []byte) PascalByteStringBuilder
	// WithStringLength adds StringLength (property field)
	WithStringLength(int32) PascalByteStringBuilder
	// WithStringValue adds StringValue (property field)
	WithStringValue(...byte) PascalByteStringBuilder
	// Build builds the PascalByteString or returns an error if something is wrong
	Build() (PascalByteString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() PascalByteString
}

// NewPascalByteStringBuilder() creates a PascalByteStringBuilder
func NewPascalByteStringBuilder() PascalByteStringBuilder {
	return &_PascalByteStringBuilder{_PascalByteString: new(_PascalByteString)}
}

type _PascalByteStringBuilder struct {
	*_PascalByteString

	err *utils.MultiError
}

var _ (PascalByteStringBuilder) = (*_PascalByteStringBuilder)(nil)

func (b *_PascalByteStringBuilder) WithMandatoryFields(stringLength int32, stringValue []byte) PascalByteStringBuilder {
	return b.WithStringLength(stringLength).WithStringValue(stringValue...)
}

func (b *_PascalByteStringBuilder) WithStringLength(stringLength int32) PascalByteStringBuilder {
	b.StringLength = stringLength
	return b
}

func (b *_PascalByteStringBuilder) WithStringValue(stringValue ...byte) PascalByteStringBuilder {
	b.StringValue = stringValue
	return b
}

func (b *_PascalByteStringBuilder) Build() (PascalByteString, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._PascalByteString.deepCopy(), nil
}

func (b *_PascalByteStringBuilder) MustBuild() PascalByteString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_PascalByteStringBuilder) DeepCopy() any {
	_copy := b.CreatePascalByteStringBuilder().(*_PascalByteStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreatePascalByteStringBuilder creates a PascalByteStringBuilder
func (b *_PascalByteString) CreatePascalByteStringBuilder() PascalByteStringBuilder {
	if b == nil {
		return NewPascalByteStringBuilder()
	}
	return &_PascalByteStringBuilder{_PascalByteString: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PascalByteString) GetStringLength() int32 {
	return m.StringLength
}

func (m *_PascalByteString) GetStringValue() []byte {
	return m.StringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPascalByteString(structType any) PascalByteString {
	if casted, ok := structType.(PascalByteString); ok {
		return casted
	}
	if casted, ok := structType.(*PascalByteString); ok {
		return *casted
	}
	return nil
}

func (m *_PascalByteString) GetTypeName() string {
	return "PascalByteString"
}

func (m *_PascalByteString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (stringLength)
	lengthInBits += 32

	// Array field
	if len(m.StringValue) > 0 {
		lengthInBits += 8 * uint16(len(m.StringValue))
	}

	return lengthInBits
}

func (m *_PascalByteString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PascalByteStringParse(ctx context.Context, theBytes []byte) (PascalByteString, error) {
	return PascalByteStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PascalByteStringParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (PascalByteString, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (PascalByteString, error) {
		return PascalByteStringParseWithBuffer(ctx, readBuffer)
	}
}

func PascalByteStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PascalByteString, error) {
	v, err := (&_PascalByteString{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_PascalByteString) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__pascalByteString PascalByteString, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PascalByteString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PascalByteString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	stringLength, err := ReadSimpleField(ctx, "stringLength", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'stringLength' field"))
	}
	m.StringLength = stringLength

	stringValue, err := readBuffer.ReadByteArray("stringValue", int(utils.InlineIf(bool((stringLength) == (-(1))), func() any { return int32(int32(0)) }, func() any { return int32(stringLength) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'stringValue' field"))
	}
	m.StringValue = stringValue

	if closeErr := readBuffer.CloseContext("PascalByteString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PascalByteString")
	}

	return m, nil
}

func (m *_PascalByteString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PascalByteString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("PascalByteString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for PascalByteString")
	}

	if err := WriteSimpleField[int32](ctx, "stringLength", m.GetStringLength(), WriteSignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'stringLength' field")
	}

	if err := WriteByteArrayField(ctx, "stringValue", m.GetStringValue(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'stringValue' field")
	}

	if popErr := writeBuffer.PopContext("PascalByteString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for PascalByteString")
	}
	return nil
}

func (m *_PascalByteString) IsPascalByteString() {}

func (m *_PascalByteString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PascalByteString) deepCopy() *_PascalByteString {
	if m == nil {
		return nil
	}
	_PascalByteStringCopy := &_PascalByteString{
		m.StringLength,
		utils.DeepCopySlice[byte, byte](m.StringValue),
	}
	return _PascalByteStringCopy
}

func (m *_PascalByteString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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

// ErrorClassTagged is the corresponding interface of ErrorClassTagged
type ErrorClassTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() ErrorClass
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsErrorClassTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorClassTagged()
	// CreateBuilder creates a ErrorClassTaggedBuilder
	CreateErrorClassTaggedBuilder() ErrorClassTaggedBuilder
}

// _ErrorClassTagged is the data-structure of this message
type _ErrorClassTagged struct {
	Header           BACnetTagHeader
	Value            ErrorClass
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ ErrorClassTagged = (*_ErrorClassTagged)(nil)

// NewErrorClassTagged factory function for _ErrorClassTagged
func NewErrorClassTagged(header BACnetTagHeader, value ErrorClass, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_ErrorClassTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for ErrorClassTagged must not be nil")
	}
	return &_ErrorClassTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ErrorClassTaggedBuilder is a builder for ErrorClassTagged
type ErrorClassTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value ErrorClass, proprietaryValue uint32) ErrorClassTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) ErrorClassTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) ErrorClassTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(ErrorClass) ErrorClassTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) ErrorClassTaggedBuilder
	// Build builds the ErrorClassTagged or returns an error if something is wrong
	Build() (ErrorClassTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ErrorClassTagged
}

// NewErrorClassTaggedBuilder() creates a ErrorClassTaggedBuilder
func NewErrorClassTaggedBuilder() ErrorClassTaggedBuilder {
	return &_ErrorClassTaggedBuilder{_ErrorClassTagged: new(_ErrorClassTagged)}
}

type _ErrorClassTaggedBuilder struct {
	*_ErrorClassTagged

	err *utils.MultiError
}

var _ (ErrorClassTaggedBuilder) = (*_ErrorClassTaggedBuilder)(nil)

func (b *_ErrorClassTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value ErrorClass, proprietaryValue uint32) ErrorClassTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_ErrorClassTaggedBuilder) WithHeader(header BACnetTagHeader) ErrorClassTaggedBuilder {
	b.Header = header
	return b
}

func (b *_ErrorClassTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) ErrorClassTaggedBuilder {
	builder := builderSupplier(b.Header.CreateBACnetTagHeaderBuilder())
	var err error
	b.Header, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_ErrorClassTaggedBuilder) WithValue(value ErrorClass) ErrorClassTaggedBuilder {
	b.Value = value
	return b
}

func (b *_ErrorClassTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) ErrorClassTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_ErrorClassTaggedBuilder) Build() (ErrorClassTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ErrorClassTagged.deepCopy(), nil
}

func (b *_ErrorClassTaggedBuilder) MustBuild() ErrorClassTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ErrorClassTaggedBuilder) DeepCopy() any {
	_copy := b.CreateErrorClassTaggedBuilder().(*_ErrorClassTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateErrorClassTaggedBuilder creates a ErrorClassTaggedBuilder
func (b *_ErrorClassTagged) CreateErrorClassTaggedBuilder() ErrorClassTaggedBuilder {
	if b == nil {
		return NewErrorClassTaggedBuilder()
	}
	return &_ErrorClassTaggedBuilder{_ErrorClassTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorClassTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_ErrorClassTagged) GetValue() ErrorClass {
	return m.Value
}

func (m *_ErrorClassTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_ErrorClassTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (ErrorClass_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastErrorClassTagged(structType any) ErrorClassTagged {
	if casted, ok := structType.(ErrorClassTagged); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorClassTagged); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorClassTagged) GetTypeName() string {
	return "ErrorClassTagged"
}

func (m *_ErrorClassTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_ErrorClassTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorClassTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (ErrorClassTagged, error) {
	return ErrorClassTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func ErrorClassTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorClassTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorClassTagged, error) {
		return ErrorClassTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func ErrorClassTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (ErrorClassTagged, error) {
	v, err := (&_ErrorClassTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ErrorClassTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__errorClassTagged ErrorClassTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorClassTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorClassTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[ErrorClass](ctx, "value", readBuffer, EnsureType[ErrorClass](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), ErrorClass_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (ErrorClass_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("ErrorClassTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorClassTagged")
	}

	return m, nil
}

func (m *_ErrorClassTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorClassTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ErrorClassTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ErrorClassTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[ErrorClass](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	if err := WriteManualField[uint32](ctx, "proprietaryValue", func(ctx context.Context) error {
		return WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	}, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("ErrorClassTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ErrorClassTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_ErrorClassTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_ErrorClassTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_ErrorClassTagged) IsErrorClassTagged() {}

func (m *_ErrorClassTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ErrorClassTagged) deepCopy() *_ErrorClassTagged {
	if m == nil {
		return nil
	}
	_ErrorClassTaggedCopy := &_ErrorClassTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _ErrorClassTaggedCopy
}

func (m *_ErrorClassTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

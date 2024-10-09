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

// BACnetObjectTypeTagged is the corresponding interface of BACnetObjectTypeTagged
type BACnetObjectTypeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetObjectType
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetObjectTypeTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetObjectTypeTagged()
	// CreateBuilder creates a BACnetObjectTypeTaggedBuilder
	CreateBACnetObjectTypeTaggedBuilder() BACnetObjectTypeTaggedBuilder
}

// _BACnetObjectTypeTagged is the data-structure of this message
type _BACnetObjectTypeTagged struct {
	Header           BACnetTagHeader
	Value            BACnetObjectType
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetObjectTypeTagged = (*_BACnetObjectTypeTagged)(nil)

// NewBACnetObjectTypeTagged factory function for _BACnetObjectTypeTagged
func NewBACnetObjectTypeTagged(header BACnetTagHeader, value BACnetObjectType, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetObjectTypeTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetObjectTypeTagged must not be nil")
	}
	return &_BACnetObjectTypeTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetObjectTypeTaggedBuilder is a builder for BACnetObjectTypeTagged
type BACnetObjectTypeTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetObjectType, proprietaryValue uint32) BACnetObjectTypeTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetObjectTypeTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetObjectTypeTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetObjectType) BACnetObjectTypeTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetObjectTypeTaggedBuilder
	// Build builds the BACnetObjectTypeTagged or returns an error if something is wrong
	Build() (BACnetObjectTypeTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetObjectTypeTagged
}

// NewBACnetObjectTypeTaggedBuilder() creates a BACnetObjectTypeTaggedBuilder
func NewBACnetObjectTypeTaggedBuilder() BACnetObjectTypeTaggedBuilder {
	return &_BACnetObjectTypeTaggedBuilder{_BACnetObjectTypeTagged: new(_BACnetObjectTypeTagged)}
}

type _BACnetObjectTypeTaggedBuilder struct {
	*_BACnetObjectTypeTagged

	err *utils.MultiError
}

var _ (BACnetObjectTypeTaggedBuilder) = (*_BACnetObjectTypeTaggedBuilder)(nil)

func (b *_BACnetObjectTypeTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetObjectType, proprietaryValue uint32) BACnetObjectTypeTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetObjectTypeTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetObjectTypeTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetObjectTypeTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetObjectTypeTaggedBuilder {
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

func (b *_BACnetObjectTypeTaggedBuilder) WithValue(value BACnetObjectType) BACnetObjectTypeTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetObjectTypeTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetObjectTypeTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetObjectTypeTaggedBuilder) Build() (BACnetObjectTypeTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetObjectTypeTagged.deepCopy(), nil
}

func (b *_BACnetObjectTypeTaggedBuilder) MustBuild() BACnetObjectTypeTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetObjectTypeTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetObjectTypeTaggedBuilder().(*_BACnetObjectTypeTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetObjectTypeTaggedBuilder creates a BACnetObjectTypeTaggedBuilder
func (b *_BACnetObjectTypeTagged) CreateBACnetObjectTypeTaggedBuilder() BACnetObjectTypeTaggedBuilder {
	if b == nil {
		return NewBACnetObjectTypeTaggedBuilder()
	}
	return &_BACnetObjectTypeTaggedBuilder{_BACnetObjectTypeTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetObjectTypeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetObjectTypeTagged) GetValue() BACnetObjectType {
	return m.Value
}

func (m *_BACnetObjectTypeTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetObjectTypeTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetObjectType_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetObjectTypeTagged(structType any) BACnetObjectTypeTagged {
	if casted, ok := structType.(BACnetObjectTypeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetObjectTypeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetObjectTypeTagged) GetTypeName() string {
	return "BACnetObjectTypeTagged"
}

func (m *_BACnetObjectTypeTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetObjectTypeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetObjectTypeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetObjectTypeTagged, error) {
	return BACnetObjectTypeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetObjectTypeTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetObjectTypeTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetObjectTypeTagged, error) {
		return BACnetObjectTypeTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetObjectTypeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetObjectTypeTagged, error) {
	v, err := (&_BACnetObjectTypeTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetObjectTypeTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetObjectTypeTagged BACnetObjectTypeTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetObjectTypeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetObjectTypeTagged")
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

	value, err := ReadManualField[BACnetObjectType](ctx, "value", readBuffer, EnsureType[BACnetObjectType](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetObjectType_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetObjectType_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetObjectTypeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetObjectTypeTagged")
	}

	return m, nil
}

func (m *_BACnetObjectTypeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetObjectTypeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetObjectTypeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetObjectTypeTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetObjectType](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetObjectTypeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetObjectTypeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetObjectTypeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetObjectTypeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetObjectTypeTagged) IsBACnetObjectTypeTagged() {}

func (m *_BACnetObjectTypeTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetObjectTypeTagged) deepCopy() *_BACnetObjectTypeTagged {
	if m == nil {
		return nil
	}
	_BACnetObjectTypeTaggedCopy := &_BACnetObjectTypeTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetObjectTypeTaggedCopy
}

func (m *_BACnetObjectTypeTagged) String() string {
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

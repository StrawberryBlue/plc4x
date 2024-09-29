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

// BACnetAuthorizationExemptionTagged is the corresponding interface of BACnetAuthorizationExemptionTagged
type BACnetAuthorizationExemptionTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAuthorizationExemption
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetAuthorizationExemptionTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAuthorizationExemptionTagged()
	// CreateBuilder creates a BACnetAuthorizationExemptionTaggedBuilder
	CreateBACnetAuthorizationExemptionTaggedBuilder() BACnetAuthorizationExemptionTaggedBuilder
}

// _BACnetAuthorizationExemptionTagged is the data-structure of this message
type _BACnetAuthorizationExemptionTagged struct {
	Header           BACnetTagHeader
	Value            BACnetAuthorizationExemption
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetAuthorizationExemptionTagged = (*_BACnetAuthorizationExemptionTagged)(nil)

// NewBACnetAuthorizationExemptionTagged factory function for _BACnetAuthorizationExemptionTagged
func NewBACnetAuthorizationExemptionTagged(header BACnetTagHeader, value BACnetAuthorizationExemption, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetAuthorizationExemptionTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetAuthorizationExemptionTagged must not be nil")
	}
	return &_BACnetAuthorizationExemptionTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAuthorizationExemptionTaggedBuilder is a builder for BACnetAuthorizationExemptionTagged
type BACnetAuthorizationExemptionTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetAuthorizationExemption, proprietaryValue uint32) BACnetAuthorizationExemptionTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetAuthorizationExemptionTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAuthorizationExemptionTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetAuthorizationExemption) BACnetAuthorizationExemptionTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetAuthorizationExemptionTaggedBuilder
	// Build builds the BACnetAuthorizationExemptionTagged or returns an error if something is wrong
	Build() (BACnetAuthorizationExemptionTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAuthorizationExemptionTagged
}

// NewBACnetAuthorizationExemptionTaggedBuilder() creates a BACnetAuthorizationExemptionTaggedBuilder
func NewBACnetAuthorizationExemptionTaggedBuilder() BACnetAuthorizationExemptionTaggedBuilder {
	return &_BACnetAuthorizationExemptionTaggedBuilder{_BACnetAuthorizationExemptionTagged: new(_BACnetAuthorizationExemptionTagged)}
}

type _BACnetAuthorizationExemptionTaggedBuilder struct {
	*_BACnetAuthorizationExemptionTagged

	err *utils.MultiError
}

var _ (BACnetAuthorizationExemptionTaggedBuilder) = (*_BACnetAuthorizationExemptionTaggedBuilder)(nil)

func (b *_BACnetAuthorizationExemptionTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetAuthorizationExemption, proprietaryValue uint32) BACnetAuthorizationExemptionTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetAuthorizationExemptionTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetAuthorizationExemptionTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetAuthorizationExemptionTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAuthorizationExemptionTaggedBuilder {
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

func (b *_BACnetAuthorizationExemptionTaggedBuilder) WithValue(value BACnetAuthorizationExemption) BACnetAuthorizationExemptionTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetAuthorizationExemptionTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetAuthorizationExemptionTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetAuthorizationExemptionTaggedBuilder) Build() (BACnetAuthorizationExemptionTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAuthorizationExemptionTagged.deepCopy(), nil
}

func (b *_BACnetAuthorizationExemptionTaggedBuilder) MustBuild() BACnetAuthorizationExemptionTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAuthorizationExemptionTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAuthorizationExemptionTaggedBuilder().(*_BACnetAuthorizationExemptionTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAuthorizationExemptionTaggedBuilder creates a BACnetAuthorizationExemptionTaggedBuilder
func (b *_BACnetAuthorizationExemptionTagged) CreateBACnetAuthorizationExemptionTaggedBuilder() BACnetAuthorizationExemptionTaggedBuilder {
	if b == nil {
		return NewBACnetAuthorizationExemptionTaggedBuilder()
	}
	return &_BACnetAuthorizationExemptionTaggedBuilder{_BACnetAuthorizationExemptionTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAuthorizationExemptionTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAuthorizationExemptionTagged) GetValue() BACnetAuthorizationExemption {
	return m.Value
}

func (m *_BACnetAuthorizationExemptionTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetAuthorizationExemptionTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetAuthorizationExemption_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAuthorizationExemptionTagged(structType any) BACnetAuthorizationExemptionTagged {
	if casted, ok := structType.(BACnetAuthorizationExemptionTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAuthorizationExemptionTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAuthorizationExemptionTagged) GetTypeName() string {
	return "BACnetAuthorizationExemptionTagged"
}

func (m *_BACnetAuthorizationExemptionTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetAuthorizationExemptionTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthorizationExemptionTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAuthorizationExemptionTagged, error) {
	return BACnetAuthorizationExemptionTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAuthorizationExemptionTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthorizationExemptionTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthorizationExemptionTagged, error) {
		return BACnetAuthorizationExemptionTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetAuthorizationExemptionTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAuthorizationExemptionTagged, error) {
	v, err := (&_BACnetAuthorizationExemptionTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAuthorizationExemptionTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetAuthorizationExemptionTagged BACnetAuthorizationExemptionTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthorizationExemptionTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthorizationExemptionTagged")
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

	value, err := ReadManualField[BACnetAuthorizationExemption](ctx, "value", readBuffer, EnsureType[BACnetAuthorizationExemption](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetAuthorizationExemption_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetAuthorizationExemption_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetAuthorizationExemptionTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthorizationExemptionTagged")
	}

	return m, nil
}

func (m *_BACnetAuthorizationExemptionTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAuthorizationExemptionTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAuthorizationExemptionTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthorizationExemptionTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetAuthorizationExemption](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetAuthorizationExemptionTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthorizationExemptionTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAuthorizationExemptionTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAuthorizationExemptionTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAuthorizationExemptionTagged) IsBACnetAuthorizationExemptionTagged() {}

func (m *_BACnetAuthorizationExemptionTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAuthorizationExemptionTagged) deepCopy() *_BACnetAuthorizationExemptionTagged {
	if m == nil {
		return nil
	}
	_BACnetAuthorizationExemptionTaggedCopy := &_BACnetAuthorizationExemptionTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetAuthorizationExemptionTaggedCopy
}

func (m *_BACnetAuthorizationExemptionTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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

// BACnetAccessAuthenticationFactorDisableTagged is the corresponding interface of BACnetAccessAuthenticationFactorDisableTagged
type BACnetAccessAuthenticationFactorDisableTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAccessAuthenticationFactorDisable
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetAccessAuthenticationFactorDisableTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAccessAuthenticationFactorDisableTagged()
	// CreateBuilder creates a BACnetAccessAuthenticationFactorDisableTaggedBuilder
	CreateBACnetAccessAuthenticationFactorDisableTaggedBuilder() BACnetAccessAuthenticationFactorDisableTaggedBuilder
}

// _BACnetAccessAuthenticationFactorDisableTagged is the data-structure of this message
type _BACnetAccessAuthenticationFactorDisableTagged struct {
	Header           BACnetTagHeader
	Value            BACnetAccessAuthenticationFactorDisable
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetAccessAuthenticationFactorDisableTagged = (*_BACnetAccessAuthenticationFactorDisableTagged)(nil)

// NewBACnetAccessAuthenticationFactorDisableTagged factory function for _BACnetAccessAuthenticationFactorDisableTagged
func NewBACnetAccessAuthenticationFactorDisableTagged(header BACnetTagHeader, value BACnetAccessAuthenticationFactorDisable, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetAccessAuthenticationFactorDisableTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetAccessAuthenticationFactorDisableTagged must not be nil")
	}
	return &_BACnetAccessAuthenticationFactorDisableTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAccessAuthenticationFactorDisableTaggedBuilder is a builder for BACnetAccessAuthenticationFactorDisableTagged
type BACnetAccessAuthenticationFactorDisableTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetAccessAuthenticationFactorDisable, proprietaryValue uint32) BACnetAccessAuthenticationFactorDisableTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetAccessAuthenticationFactorDisableTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessAuthenticationFactorDisableTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetAccessAuthenticationFactorDisable) BACnetAccessAuthenticationFactorDisableTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetAccessAuthenticationFactorDisableTaggedBuilder
	// Build builds the BACnetAccessAuthenticationFactorDisableTagged or returns an error if something is wrong
	Build() (BACnetAccessAuthenticationFactorDisableTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAccessAuthenticationFactorDisableTagged
}

// NewBACnetAccessAuthenticationFactorDisableTaggedBuilder() creates a BACnetAccessAuthenticationFactorDisableTaggedBuilder
func NewBACnetAccessAuthenticationFactorDisableTaggedBuilder() BACnetAccessAuthenticationFactorDisableTaggedBuilder {
	return &_BACnetAccessAuthenticationFactorDisableTaggedBuilder{_BACnetAccessAuthenticationFactorDisableTagged: new(_BACnetAccessAuthenticationFactorDisableTagged)}
}

type _BACnetAccessAuthenticationFactorDisableTaggedBuilder struct {
	*_BACnetAccessAuthenticationFactorDisableTagged

	err *utils.MultiError
}

var _ (BACnetAccessAuthenticationFactorDisableTaggedBuilder) = (*_BACnetAccessAuthenticationFactorDisableTaggedBuilder)(nil)

func (b *_BACnetAccessAuthenticationFactorDisableTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetAccessAuthenticationFactorDisable, proprietaryValue uint32) BACnetAccessAuthenticationFactorDisableTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetAccessAuthenticationFactorDisableTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetAccessAuthenticationFactorDisableTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetAccessAuthenticationFactorDisableTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessAuthenticationFactorDisableTaggedBuilder {
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

func (b *_BACnetAccessAuthenticationFactorDisableTaggedBuilder) WithValue(value BACnetAccessAuthenticationFactorDisable) BACnetAccessAuthenticationFactorDisableTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetAccessAuthenticationFactorDisableTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetAccessAuthenticationFactorDisableTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetAccessAuthenticationFactorDisableTaggedBuilder) Build() (BACnetAccessAuthenticationFactorDisableTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAccessAuthenticationFactorDisableTagged.deepCopy(), nil
}

func (b *_BACnetAccessAuthenticationFactorDisableTaggedBuilder) MustBuild() BACnetAccessAuthenticationFactorDisableTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAccessAuthenticationFactorDisableTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAccessAuthenticationFactorDisableTaggedBuilder().(*_BACnetAccessAuthenticationFactorDisableTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAccessAuthenticationFactorDisableTaggedBuilder creates a BACnetAccessAuthenticationFactorDisableTaggedBuilder
func (b *_BACnetAccessAuthenticationFactorDisableTagged) CreateBACnetAccessAuthenticationFactorDisableTaggedBuilder() BACnetAccessAuthenticationFactorDisableTaggedBuilder {
	if b == nil {
		return NewBACnetAccessAuthenticationFactorDisableTaggedBuilder()
	}
	return &_BACnetAccessAuthenticationFactorDisableTaggedBuilder{_BACnetAccessAuthenticationFactorDisableTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetValue() BACnetAccessAuthenticationFactorDisable {
	return m.Value
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetAccessAuthenticationFactorDisable_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAccessAuthenticationFactorDisableTagged(structType any) BACnetAccessAuthenticationFactorDisableTagged {
	if casted, ok := structType.(BACnetAccessAuthenticationFactorDisableTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessAuthenticationFactorDisableTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetTypeName() string {
	return "BACnetAccessAuthenticationFactorDisableTagged"
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessAuthenticationFactorDisableTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAccessAuthenticationFactorDisableTagged, error) {
	return BACnetAccessAuthenticationFactorDisableTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAccessAuthenticationFactorDisableTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessAuthenticationFactorDisableTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessAuthenticationFactorDisableTagged, error) {
		return BACnetAccessAuthenticationFactorDisableTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetAccessAuthenticationFactorDisableTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAccessAuthenticationFactorDisableTagged, error) {
	v, err := (&_BACnetAccessAuthenticationFactorDisableTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetAccessAuthenticationFactorDisableTagged BACnetAccessAuthenticationFactorDisableTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccessAuthenticationFactorDisableTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessAuthenticationFactorDisableTagged")
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

	value, err := ReadManualField[BACnetAccessAuthenticationFactorDisable](ctx, "value", readBuffer, EnsureType[BACnetAccessAuthenticationFactorDisable](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetAccessAuthenticationFactorDisable_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetAccessAuthenticationFactorDisable_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetAccessAuthenticationFactorDisableTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessAuthenticationFactorDisableTagged")
	}

	return m, nil
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccessAuthenticationFactorDisableTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessAuthenticationFactorDisableTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetAccessAuthenticationFactorDisable](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
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

	if popErr := writeBuffer.PopContext("BACnetAccessAuthenticationFactorDisableTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessAuthenticationFactorDisableTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAccessAuthenticationFactorDisableTagged) IsBACnetAccessAuthenticationFactorDisableTagged() {
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) deepCopy() *_BACnetAccessAuthenticationFactorDisableTagged {
	if m == nil {
		return nil
	}
	_BACnetAccessAuthenticationFactorDisableTaggedCopy := &_BACnetAccessAuthenticationFactorDisableTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetAccessAuthenticationFactorDisableTaggedCopy
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

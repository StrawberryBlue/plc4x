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

// BACnetProgramRequestTagged is the corresponding interface of BACnetProgramRequestTagged
type BACnetProgramRequestTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetProgramRequest
	// IsBACnetProgramRequestTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetProgramRequestTagged()
	// CreateBuilder creates a BACnetProgramRequestTaggedBuilder
	CreateBACnetProgramRequestTaggedBuilder() BACnetProgramRequestTaggedBuilder
}

// _BACnetProgramRequestTagged is the data-structure of this message
type _BACnetProgramRequestTagged struct {
	Header BACnetTagHeader
	Value  BACnetProgramRequest

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetProgramRequestTagged = (*_BACnetProgramRequestTagged)(nil)

// NewBACnetProgramRequestTagged factory function for _BACnetProgramRequestTagged
func NewBACnetProgramRequestTagged(header BACnetTagHeader, value BACnetProgramRequest, tagNumber uint8, tagClass TagClass) *_BACnetProgramRequestTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetProgramRequestTagged must not be nil")
	}
	return &_BACnetProgramRequestTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetProgramRequestTaggedBuilder is a builder for BACnetProgramRequestTagged
type BACnetProgramRequestTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetProgramRequest) BACnetProgramRequestTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetProgramRequestTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetProgramRequestTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetProgramRequest) BACnetProgramRequestTaggedBuilder
	// Build builds the BACnetProgramRequestTagged or returns an error if something is wrong
	Build() (BACnetProgramRequestTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetProgramRequestTagged
}

// NewBACnetProgramRequestTaggedBuilder() creates a BACnetProgramRequestTaggedBuilder
func NewBACnetProgramRequestTaggedBuilder() BACnetProgramRequestTaggedBuilder {
	return &_BACnetProgramRequestTaggedBuilder{_BACnetProgramRequestTagged: new(_BACnetProgramRequestTagged)}
}

type _BACnetProgramRequestTaggedBuilder struct {
	*_BACnetProgramRequestTagged

	err *utils.MultiError
}

var _ (BACnetProgramRequestTaggedBuilder) = (*_BACnetProgramRequestTaggedBuilder)(nil)

func (b *_BACnetProgramRequestTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetProgramRequest) BACnetProgramRequestTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_BACnetProgramRequestTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetProgramRequestTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetProgramRequestTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetProgramRequestTaggedBuilder {
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

func (b *_BACnetProgramRequestTaggedBuilder) WithValue(value BACnetProgramRequest) BACnetProgramRequestTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetProgramRequestTaggedBuilder) Build() (BACnetProgramRequestTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetProgramRequestTagged.deepCopy(), nil
}

func (b *_BACnetProgramRequestTaggedBuilder) MustBuild() BACnetProgramRequestTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetProgramRequestTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetProgramRequestTaggedBuilder().(*_BACnetProgramRequestTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetProgramRequestTaggedBuilder creates a BACnetProgramRequestTaggedBuilder
func (b *_BACnetProgramRequestTagged) CreateBACnetProgramRequestTaggedBuilder() BACnetProgramRequestTaggedBuilder {
	if b == nil {
		return NewBACnetProgramRequestTaggedBuilder()
	}
	return &_BACnetProgramRequestTaggedBuilder{_BACnetProgramRequestTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetProgramRequestTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetProgramRequestTagged) GetValue() BACnetProgramRequest {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetProgramRequestTagged(structType any) BACnetProgramRequestTagged {
	if casted, ok := structType.(BACnetProgramRequestTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetProgramRequestTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetProgramRequestTagged) GetTypeName() string {
	return "BACnetProgramRequestTagged"
}

func (m *_BACnetProgramRequestTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetProgramRequestTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetProgramRequestTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetProgramRequestTagged, error) {
	return BACnetProgramRequestTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetProgramRequestTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetProgramRequestTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetProgramRequestTagged, error) {
		return BACnetProgramRequestTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetProgramRequestTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetProgramRequestTagged, error) {
	v, err := (&_BACnetProgramRequestTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetProgramRequestTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetProgramRequestTagged BACnetProgramRequestTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProgramRequestTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProgramRequestTagged")
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

	value, err := ReadManualField[BACnetProgramRequest](ctx, "value", readBuffer, EnsureType[BACnetProgramRequest](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetProgramRequest_READY)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetProgramRequestTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProgramRequestTagged")
	}

	return m, nil
}

func (m *_BACnetProgramRequestTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetProgramRequestTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetProgramRequestTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetProgramRequestTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetProgramRequest](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetProgramRequestTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetProgramRequestTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetProgramRequestTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetProgramRequestTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetProgramRequestTagged) IsBACnetProgramRequestTagged() {}

func (m *_BACnetProgramRequestTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetProgramRequestTagged) deepCopy() *_BACnetProgramRequestTagged {
	if m == nil {
		return nil
	}
	_BACnetProgramRequestTaggedCopy := &_BACnetProgramRequestTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetProgramRequestTaggedCopy
}

func (m *_BACnetProgramRequestTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

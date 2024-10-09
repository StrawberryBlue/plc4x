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

// BACnetTimerTransitionTagged is the corresponding interface of BACnetTimerTransitionTagged
type BACnetTimerTransitionTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetTimerTransition
	// IsBACnetTimerTransitionTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerTransitionTagged()
	// CreateBuilder creates a BACnetTimerTransitionTaggedBuilder
	CreateBACnetTimerTransitionTaggedBuilder() BACnetTimerTransitionTaggedBuilder
}

// _BACnetTimerTransitionTagged is the data-structure of this message
type _BACnetTimerTransitionTagged struct {
	Header BACnetTagHeader
	Value  BACnetTimerTransition

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetTimerTransitionTagged = (*_BACnetTimerTransitionTagged)(nil)

// NewBACnetTimerTransitionTagged factory function for _BACnetTimerTransitionTagged
func NewBACnetTimerTransitionTagged(header BACnetTagHeader, value BACnetTimerTransition, tagNumber uint8, tagClass TagClass) *_BACnetTimerTransitionTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetTimerTransitionTagged must not be nil")
	}
	return &_BACnetTimerTransitionTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimerTransitionTaggedBuilder is a builder for BACnetTimerTransitionTagged
type BACnetTimerTransitionTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetTimerTransition) BACnetTimerTransitionTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetTimerTransitionTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetTimerTransitionTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetTimerTransition) BACnetTimerTransitionTaggedBuilder
	// Build builds the BACnetTimerTransitionTagged or returns an error if something is wrong
	Build() (BACnetTimerTransitionTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimerTransitionTagged
}

// NewBACnetTimerTransitionTaggedBuilder() creates a BACnetTimerTransitionTaggedBuilder
func NewBACnetTimerTransitionTaggedBuilder() BACnetTimerTransitionTaggedBuilder {
	return &_BACnetTimerTransitionTaggedBuilder{_BACnetTimerTransitionTagged: new(_BACnetTimerTransitionTagged)}
}

type _BACnetTimerTransitionTaggedBuilder struct {
	*_BACnetTimerTransitionTagged

	err *utils.MultiError
}

var _ (BACnetTimerTransitionTaggedBuilder) = (*_BACnetTimerTransitionTaggedBuilder)(nil)

func (b *_BACnetTimerTransitionTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetTimerTransition) BACnetTimerTransitionTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_BACnetTimerTransitionTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetTimerTransitionTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetTimerTransitionTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetTimerTransitionTaggedBuilder {
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

func (b *_BACnetTimerTransitionTaggedBuilder) WithValue(value BACnetTimerTransition) BACnetTimerTransitionTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetTimerTransitionTaggedBuilder) Build() (BACnetTimerTransitionTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTimerTransitionTagged.deepCopy(), nil
}

func (b *_BACnetTimerTransitionTaggedBuilder) MustBuild() BACnetTimerTransitionTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetTimerTransitionTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTimerTransitionTaggedBuilder().(*_BACnetTimerTransitionTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTimerTransitionTaggedBuilder creates a BACnetTimerTransitionTaggedBuilder
func (b *_BACnetTimerTransitionTagged) CreateBACnetTimerTransitionTaggedBuilder() BACnetTimerTransitionTaggedBuilder {
	if b == nil {
		return NewBACnetTimerTransitionTaggedBuilder()
	}
	return &_BACnetTimerTransitionTaggedBuilder{_BACnetTimerTransitionTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerTransitionTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetTimerTransitionTagged) GetValue() BACnetTimerTransition {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimerTransitionTagged(structType any) BACnetTimerTransitionTagged {
	if casted, ok := structType.(BACnetTimerTransitionTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerTransitionTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerTransitionTagged) GetTypeName() string {
	return "BACnetTimerTransitionTagged"
}

func (m *_BACnetTimerTransitionTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetTimerTransitionTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimerTransitionTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetTimerTransitionTagged, error) {
	return BACnetTimerTransitionTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetTimerTransitionTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimerTransitionTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimerTransitionTagged, error) {
		return BACnetTimerTransitionTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetTimerTransitionTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetTimerTransitionTagged, error) {
	v, err := (&_BACnetTimerTransitionTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTimerTransitionTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetTimerTransitionTagged BACnetTimerTransitionTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerTransitionTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerTransitionTagged")
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

	value, err := ReadManualField[BACnetTimerTransition](ctx, "value", readBuffer, EnsureType[BACnetTimerTransition](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetTimerTransition_NONE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetTimerTransitionTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerTransitionTagged")
	}

	return m, nil
}

func (m *_BACnetTimerTransitionTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerTransitionTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTimerTransitionTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTimerTransitionTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetTimerTransition](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTimerTransitionTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTimerTransitionTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetTimerTransitionTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetTimerTransitionTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetTimerTransitionTagged) IsBACnetTimerTransitionTagged() {}

func (m *_BACnetTimerTransitionTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimerTransitionTagged) deepCopy() *_BACnetTimerTransitionTagged {
	if m == nil {
		return nil
	}
	_BACnetTimerTransitionTaggedCopy := &_BACnetTimerTransitionTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetTimerTransitionTaggedCopy
}

func (m *_BACnetTimerTransitionTagged) String() string {
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

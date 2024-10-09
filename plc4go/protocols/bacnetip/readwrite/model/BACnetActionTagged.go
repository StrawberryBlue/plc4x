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

// BACnetActionTagged is the corresponding interface of BACnetActionTagged
type BACnetActionTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAction
	// IsBACnetActionTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetActionTagged()
	// CreateBuilder creates a BACnetActionTaggedBuilder
	CreateBACnetActionTaggedBuilder() BACnetActionTaggedBuilder
}

// _BACnetActionTagged is the data-structure of this message
type _BACnetActionTagged struct {
	Header BACnetTagHeader
	Value  BACnetAction

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetActionTagged = (*_BACnetActionTagged)(nil)

// NewBACnetActionTagged factory function for _BACnetActionTagged
func NewBACnetActionTagged(header BACnetTagHeader, value BACnetAction, tagNumber uint8, tagClass TagClass) *_BACnetActionTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetActionTagged must not be nil")
	}
	return &_BACnetActionTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetActionTaggedBuilder is a builder for BACnetActionTagged
type BACnetActionTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetAction) BACnetActionTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetActionTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetActionTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetAction) BACnetActionTaggedBuilder
	// Build builds the BACnetActionTagged or returns an error if something is wrong
	Build() (BACnetActionTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetActionTagged
}

// NewBACnetActionTaggedBuilder() creates a BACnetActionTaggedBuilder
func NewBACnetActionTaggedBuilder() BACnetActionTaggedBuilder {
	return &_BACnetActionTaggedBuilder{_BACnetActionTagged: new(_BACnetActionTagged)}
}

type _BACnetActionTaggedBuilder struct {
	*_BACnetActionTagged

	err *utils.MultiError
}

var _ (BACnetActionTaggedBuilder) = (*_BACnetActionTaggedBuilder)(nil)

func (b *_BACnetActionTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetAction) BACnetActionTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_BACnetActionTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetActionTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetActionTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetActionTaggedBuilder {
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

func (b *_BACnetActionTaggedBuilder) WithValue(value BACnetAction) BACnetActionTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetActionTaggedBuilder) Build() (BACnetActionTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetActionTagged.deepCopy(), nil
}

func (b *_BACnetActionTaggedBuilder) MustBuild() BACnetActionTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetActionTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetActionTaggedBuilder().(*_BACnetActionTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetActionTaggedBuilder creates a BACnetActionTaggedBuilder
func (b *_BACnetActionTagged) CreateBACnetActionTaggedBuilder() BACnetActionTaggedBuilder {
	if b == nil {
		return NewBACnetActionTaggedBuilder()
	}
	return &_BACnetActionTaggedBuilder{_BACnetActionTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetActionTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetActionTagged) GetValue() BACnetAction {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetActionTagged(structType any) BACnetActionTagged {
	if casted, ok := structType.(BACnetActionTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetActionTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetActionTagged) GetTypeName() string {
	return "BACnetActionTagged"
}

func (m *_BACnetActionTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetActionTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetActionTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetActionTagged, error) {
	return BACnetActionTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetActionTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetActionTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetActionTagged, error) {
		return BACnetActionTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetActionTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetActionTagged, error) {
	v, err := (&_BACnetActionTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetActionTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetActionTagged BACnetActionTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetActionTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetActionTagged")
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

	value, err := ReadManualField[BACnetAction](ctx, "value", readBuffer, EnsureType[BACnetAction](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetAction_DIRECT)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetActionTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetActionTagged")
	}

	return m, nil
}

func (m *_BACnetActionTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetActionTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetActionTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetActionTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetAction](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetActionTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetActionTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetActionTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetActionTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetActionTagged) IsBACnetActionTagged() {}

func (m *_BACnetActionTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetActionTagged) deepCopy() *_BACnetActionTagged {
	if m == nil {
		return nil
	}
	_BACnetActionTaggedCopy := &_BACnetActionTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetActionTaggedCopy
}

func (m *_BACnetActionTagged) String() string {
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

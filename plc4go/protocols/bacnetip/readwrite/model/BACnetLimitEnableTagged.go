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

// BACnetLimitEnableTagged is the corresponding interface of BACnetLimitEnableTagged
type BACnetLimitEnableTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
	// GetLowLimitEnable returns LowLimitEnable (virtual field)
	GetLowLimitEnable() bool
	// GetHighLimitEnable returns HighLimitEnable (virtual field)
	GetHighLimitEnable() bool
	// IsBACnetLimitEnableTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLimitEnableTagged()
	// CreateBuilder creates a BACnetLimitEnableTaggedBuilder
	CreateBACnetLimitEnableTaggedBuilder() BACnetLimitEnableTaggedBuilder
}

// _BACnetLimitEnableTagged is the data-structure of this message
type _BACnetLimitEnableTagged struct {
	Header  BACnetTagHeader
	Payload BACnetTagPayloadBitString

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetLimitEnableTagged = (*_BACnetLimitEnableTagged)(nil)

// NewBACnetLimitEnableTagged factory function for _BACnetLimitEnableTagged
func NewBACnetLimitEnableTagged(header BACnetTagHeader, payload BACnetTagPayloadBitString, tagNumber uint8, tagClass TagClass) *_BACnetLimitEnableTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetLimitEnableTagged must not be nil")
	}
	if payload == nil {
		panic("payload of type BACnetTagPayloadBitString for BACnetLimitEnableTagged must not be nil")
	}
	return &_BACnetLimitEnableTagged{Header: header, Payload: payload, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLimitEnableTaggedBuilder is a builder for BACnetLimitEnableTagged
type BACnetLimitEnableTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, payload BACnetTagPayloadBitString) BACnetLimitEnableTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetLimitEnableTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetLimitEnableTaggedBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadBitString) BACnetLimitEnableTaggedBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadBitStringBuilder) BACnetTagPayloadBitStringBuilder) BACnetLimitEnableTaggedBuilder
	// Build builds the BACnetLimitEnableTagged or returns an error if something is wrong
	Build() (BACnetLimitEnableTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLimitEnableTagged
}

// NewBACnetLimitEnableTaggedBuilder() creates a BACnetLimitEnableTaggedBuilder
func NewBACnetLimitEnableTaggedBuilder() BACnetLimitEnableTaggedBuilder {
	return &_BACnetLimitEnableTaggedBuilder{_BACnetLimitEnableTagged: new(_BACnetLimitEnableTagged)}
}

type _BACnetLimitEnableTaggedBuilder struct {
	*_BACnetLimitEnableTagged

	err *utils.MultiError
}

var _ (BACnetLimitEnableTaggedBuilder) = (*_BACnetLimitEnableTaggedBuilder)(nil)

func (b *_BACnetLimitEnableTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, payload BACnetTagPayloadBitString) BACnetLimitEnableTaggedBuilder {
	return b.WithHeader(header).WithPayload(payload)
}

func (b *_BACnetLimitEnableTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetLimitEnableTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetLimitEnableTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetLimitEnableTaggedBuilder {
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

func (b *_BACnetLimitEnableTaggedBuilder) WithPayload(payload BACnetTagPayloadBitString) BACnetLimitEnableTaggedBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetLimitEnableTaggedBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadBitStringBuilder) BACnetTagPayloadBitStringBuilder) BACnetLimitEnableTaggedBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadBitStringBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadBitStringBuilder failed"))
	}
	return b
}

func (b *_BACnetLimitEnableTaggedBuilder) Build() (BACnetLimitEnableTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLimitEnableTagged.deepCopy(), nil
}

func (b *_BACnetLimitEnableTaggedBuilder) MustBuild() BACnetLimitEnableTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLimitEnableTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLimitEnableTaggedBuilder().(*_BACnetLimitEnableTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLimitEnableTaggedBuilder creates a BACnetLimitEnableTaggedBuilder
func (b *_BACnetLimitEnableTagged) CreateBACnetLimitEnableTaggedBuilder() BACnetLimitEnableTaggedBuilder {
	if b == nil {
		return NewBACnetLimitEnableTaggedBuilder()
	}
	return &_BACnetLimitEnableTaggedBuilder{_BACnetLimitEnableTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLimitEnableTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLimitEnableTagged) GetPayload() BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetLimitEnableTagged) GetLowLimitEnable() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (0))), func() any { return bool(m.GetPayload().GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetLimitEnableTagged) GetHighLimitEnable() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (1))), func() any { return bool(m.GetPayload().GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLimitEnableTagged(structType any) BACnetLimitEnableTagged {
	if casted, ok := structType.(BACnetLimitEnableTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLimitEnableTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLimitEnableTagged) GetTypeName() string {
	return "BACnetLimitEnableTagged"
}

func (m *_BACnetLimitEnableTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetLimitEnableTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLimitEnableTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetLimitEnableTagged, error) {
	return BACnetLimitEnableTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetLimitEnableTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLimitEnableTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLimitEnableTagged, error) {
		return BACnetLimitEnableTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetLimitEnableTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLimitEnableTagged, error) {
	v, err := (&_BACnetLimitEnableTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetLimitEnableTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetLimitEnableTagged BACnetLimitEnableTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLimitEnableTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLimitEnableTagged")
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

	payload, err := ReadSimpleField[BACnetTagPayloadBitString](ctx, "payload", ReadComplex[BACnetTagPayloadBitString](BACnetTagPayloadBitStringParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	lowLimitEnable, err := ReadVirtualField[bool](ctx, "lowLimitEnable", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (0))), func() any { return bool(payload.GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowLimitEnable' field"))
	}
	_ = lowLimitEnable

	highLimitEnable, err := ReadVirtualField[bool](ctx, "highLimitEnable", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (1))), func() any { return bool(payload.GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'highLimitEnable' field"))
	}
	_ = highLimitEnable

	if closeErr := readBuffer.CloseContext("BACnetLimitEnableTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLimitEnableTagged")
	}

	return m, nil
}

func (m *_BACnetLimitEnableTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLimitEnableTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLimitEnableTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLimitEnableTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteSimpleField[BACnetTagPayloadBitString](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadBitString](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'payload' field")
	}
	// Virtual field
	lowLimitEnable := m.GetLowLimitEnable()
	_ = lowLimitEnable
	if _lowLimitEnableErr := writeBuffer.WriteVirtual(ctx, "lowLimitEnable", m.GetLowLimitEnable()); _lowLimitEnableErr != nil {
		return errors.Wrap(_lowLimitEnableErr, "Error serializing 'lowLimitEnable' field")
	}
	// Virtual field
	highLimitEnable := m.GetHighLimitEnable()
	_ = highLimitEnable
	if _highLimitEnableErr := writeBuffer.WriteVirtual(ctx, "highLimitEnable", m.GetHighLimitEnable()); _highLimitEnableErr != nil {
		return errors.Wrap(_highLimitEnableErr, "Error serializing 'highLimitEnable' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLimitEnableTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLimitEnableTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetLimitEnableTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLimitEnableTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetLimitEnableTagged) IsBACnetLimitEnableTagged() {}

func (m *_BACnetLimitEnableTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLimitEnableTagged) deepCopy() *_BACnetLimitEnableTagged {
	if m == nil {
		return nil
	}
	_BACnetLimitEnableTaggedCopy := &_BACnetLimitEnableTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Payload.DeepCopy().(BACnetTagPayloadBitString),
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetLimitEnableTaggedCopy
}

func (m *_BACnetLimitEnableTagged) String() string {
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

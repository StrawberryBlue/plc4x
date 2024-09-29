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

// BACnetAddressEnclosed is the corresponding interface of BACnetAddressEnclosed
type BACnetAddressEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetAddress returns Address (property field)
	GetAddress() BACnetAddress
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetAddressEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAddressEnclosed()
	// CreateBuilder creates a BACnetAddressEnclosedBuilder
	CreateBACnetAddressEnclosedBuilder() BACnetAddressEnclosedBuilder
}

// _BACnetAddressEnclosed is the data-structure of this message
type _BACnetAddressEnclosed struct {
	OpeningTag BACnetOpeningTag
	Address    BACnetAddress
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetAddressEnclosed = (*_BACnetAddressEnclosed)(nil)

// NewBACnetAddressEnclosed factory function for _BACnetAddressEnclosed
func NewBACnetAddressEnclosed(openingTag BACnetOpeningTag, address BACnetAddress, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetAddressEnclosed {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetAddressEnclosed must not be nil")
	}
	if address == nil {
		panic("address of type BACnetAddress for BACnetAddressEnclosed must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetAddressEnclosed must not be nil")
	}
	return &_BACnetAddressEnclosed{OpeningTag: openingTag, Address: address, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAddressEnclosedBuilder is a builder for BACnetAddressEnclosed
type BACnetAddressEnclosedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, address BACnetAddress, closingTag BACnetClosingTag) BACnetAddressEnclosedBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetAddressEnclosedBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetAddressEnclosedBuilder
	// WithAddress adds Address (property field)
	WithAddress(BACnetAddress) BACnetAddressEnclosedBuilder
	// WithAddressBuilder adds Address (property field) which is build by the builder
	WithAddressBuilder(func(BACnetAddressBuilder) BACnetAddressBuilder) BACnetAddressEnclosedBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetAddressEnclosedBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetAddressEnclosedBuilder
	// Build builds the BACnetAddressEnclosed or returns an error if something is wrong
	Build() (BACnetAddressEnclosed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAddressEnclosed
}

// NewBACnetAddressEnclosedBuilder() creates a BACnetAddressEnclosedBuilder
func NewBACnetAddressEnclosedBuilder() BACnetAddressEnclosedBuilder {
	return &_BACnetAddressEnclosedBuilder{_BACnetAddressEnclosed: new(_BACnetAddressEnclosed)}
}

type _BACnetAddressEnclosedBuilder struct {
	*_BACnetAddressEnclosed

	err *utils.MultiError
}

var _ (BACnetAddressEnclosedBuilder) = (*_BACnetAddressEnclosedBuilder)(nil)

func (b *_BACnetAddressEnclosedBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, address BACnetAddress, closingTag BACnetClosingTag) BACnetAddressEnclosedBuilder {
	return b.WithOpeningTag(openingTag).WithAddress(address).WithClosingTag(closingTag)
}

func (b *_BACnetAddressEnclosedBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetAddressEnclosedBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetAddressEnclosedBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetAddressEnclosedBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetAddressEnclosedBuilder) WithAddress(address BACnetAddress) BACnetAddressEnclosedBuilder {
	b.Address = address
	return b
}

func (b *_BACnetAddressEnclosedBuilder) WithAddressBuilder(builderSupplier func(BACnetAddressBuilder) BACnetAddressBuilder) BACnetAddressEnclosedBuilder {
	builder := builderSupplier(b.Address.CreateBACnetAddressBuilder())
	var err error
	b.Address, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetAddressBuilder failed"))
	}
	return b
}

func (b *_BACnetAddressEnclosedBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetAddressEnclosedBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetAddressEnclosedBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetAddressEnclosedBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetAddressEnclosedBuilder) Build() (BACnetAddressEnclosed, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.Address == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'address' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAddressEnclosed.deepCopy(), nil
}

func (b *_BACnetAddressEnclosedBuilder) MustBuild() BACnetAddressEnclosed {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAddressEnclosedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAddressEnclosedBuilder().(*_BACnetAddressEnclosedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAddressEnclosedBuilder creates a BACnetAddressEnclosedBuilder
func (b *_BACnetAddressEnclosed) CreateBACnetAddressEnclosedBuilder() BACnetAddressEnclosedBuilder {
	if b == nil {
		return NewBACnetAddressEnclosedBuilder()
	}
	return &_BACnetAddressEnclosedBuilder{_BACnetAddressEnclosed: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAddressEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetAddressEnclosed) GetAddress() BACnetAddress {
	return m.Address
}

func (m *_BACnetAddressEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAddressEnclosed(structType any) BACnetAddressEnclosed {
	if casted, ok := structType.(BACnetAddressEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAddressEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAddressEnclosed) GetTypeName() string {
	return "BACnetAddressEnclosed"
}

func (m *_BACnetAddressEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (address)
	lengthInBits += m.Address.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAddressEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAddressEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetAddressEnclosed, error) {
	return BACnetAddressEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetAddressEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAddressEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAddressEnclosed, error) {
		return BACnetAddressEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetAddressEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetAddressEnclosed, error) {
	v, err := (&_BACnetAddressEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAddressEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetAddressEnclosed BACnetAddressEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAddressEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAddressEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	address, err := ReadSimpleField[BACnetAddress](ctx, "address", ReadComplex[BACnetAddress](BACnetAddressParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetAddressEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAddressEnclosed")
	}

	return m, nil
}

func (m *_BACnetAddressEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAddressEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAddressEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAddressEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetAddress](ctx, "address", m.GetAddress(), WriteComplex[BACnetAddress](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'address' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAddressEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAddressEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAddressEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetAddressEnclosed) IsBACnetAddressEnclosed() {}

func (m *_BACnetAddressEnclosed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAddressEnclosed) deepCopy() *_BACnetAddressEnclosed {
	if m == nil {
		return nil
	}
	_BACnetAddressEnclosedCopy := &_BACnetAddressEnclosed{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.Address.DeepCopy().(BACnetAddress),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetAddressEnclosedCopy
}

func (m *_BACnetAddressEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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

// BACnetPropertyReferenceEnclosed is the corresponding interface of BACnetPropertyReferenceEnclosed
type BACnetPropertyReferenceEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetReference returns Reference (property field)
	GetReference() BACnetPropertyReference
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetPropertyReferenceEnclosed is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyReferenceEnclosed()
	// CreateBuilder creates a BACnetPropertyReferenceEnclosedBuilder
	CreateBACnetPropertyReferenceEnclosedBuilder() BACnetPropertyReferenceEnclosedBuilder
}

// _BACnetPropertyReferenceEnclosed is the data-structure of this message
type _BACnetPropertyReferenceEnclosed struct {
	OpeningTag BACnetOpeningTag
	Reference  BACnetPropertyReference
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetPropertyReferenceEnclosed = (*_BACnetPropertyReferenceEnclosed)(nil)

// NewBACnetPropertyReferenceEnclosed factory function for _BACnetPropertyReferenceEnclosed
func NewBACnetPropertyReferenceEnclosed(openingTag BACnetOpeningTag, reference BACnetPropertyReference, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetPropertyReferenceEnclosed {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetPropertyReferenceEnclosed must not be nil")
	}
	if reference == nil {
		panic("reference of type BACnetPropertyReference for BACnetPropertyReferenceEnclosed must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetPropertyReferenceEnclosed must not be nil")
	}
	return &_BACnetPropertyReferenceEnclosed{OpeningTag: openingTag, Reference: reference, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyReferenceEnclosedBuilder is a builder for BACnetPropertyReferenceEnclosed
type BACnetPropertyReferenceEnclosedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, reference BACnetPropertyReference, closingTag BACnetClosingTag) BACnetPropertyReferenceEnclosedBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetPropertyReferenceEnclosedBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetPropertyReferenceEnclosedBuilder
	// WithReference adds Reference (property field)
	WithReference(BACnetPropertyReference) BACnetPropertyReferenceEnclosedBuilder
	// WithReferenceBuilder adds Reference (property field) which is build by the builder
	WithReferenceBuilder(func(BACnetPropertyReferenceBuilder) BACnetPropertyReferenceBuilder) BACnetPropertyReferenceEnclosedBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetPropertyReferenceEnclosedBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetPropertyReferenceEnclosedBuilder
	// Build builds the BACnetPropertyReferenceEnclosed or returns an error if something is wrong
	Build() (BACnetPropertyReferenceEnclosed, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyReferenceEnclosed
}

// NewBACnetPropertyReferenceEnclosedBuilder() creates a BACnetPropertyReferenceEnclosedBuilder
func NewBACnetPropertyReferenceEnclosedBuilder() BACnetPropertyReferenceEnclosedBuilder {
	return &_BACnetPropertyReferenceEnclosedBuilder{_BACnetPropertyReferenceEnclosed: new(_BACnetPropertyReferenceEnclosed)}
}

type _BACnetPropertyReferenceEnclosedBuilder struct {
	*_BACnetPropertyReferenceEnclosed

	err *utils.MultiError
}

var _ (BACnetPropertyReferenceEnclosedBuilder) = (*_BACnetPropertyReferenceEnclosedBuilder)(nil)

func (b *_BACnetPropertyReferenceEnclosedBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, reference BACnetPropertyReference, closingTag BACnetClosingTag) BACnetPropertyReferenceEnclosedBuilder {
	return b.WithOpeningTag(openingTag).WithReference(reference).WithClosingTag(closingTag)
}

func (b *_BACnetPropertyReferenceEnclosedBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetPropertyReferenceEnclosedBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetPropertyReferenceEnclosedBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetPropertyReferenceEnclosedBuilder {
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

func (b *_BACnetPropertyReferenceEnclosedBuilder) WithReference(reference BACnetPropertyReference) BACnetPropertyReferenceEnclosedBuilder {
	b.Reference = reference
	return b
}

func (b *_BACnetPropertyReferenceEnclosedBuilder) WithReferenceBuilder(builderSupplier func(BACnetPropertyReferenceBuilder) BACnetPropertyReferenceBuilder) BACnetPropertyReferenceEnclosedBuilder {
	builder := builderSupplier(b.Reference.CreateBACnetPropertyReferenceBuilder())
	var err error
	b.Reference, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetPropertyReferenceBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyReferenceEnclosedBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetPropertyReferenceEnclosedBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetPropertyReferenceEnclosedBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetPropertyReferenceEnclosedBuilder {
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

func (b *_BACnetPropertyReferenceEnclosedBuilder) Build() (BACnetPropertyReferenceEnclosed, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.Reference == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'reference' not set"))
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
	return b._BACnetPropertyReferenceEnclosed.deepCopy(), nil
}

func (b *_BACnetPropertyReferenceEnclosedBuilder) MustBuild() BACnetPropertyReferenceEnclosed {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyReferenceEnclosedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyReferenceEnclosedBuilder().(*_BACnetPropertyReferenceEnclosedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyReferenceEnclosedBuilder creates a BACnetPropertyReferenceEnclosedBuilder
func (b *_BACnetPropertyReferenceEnclosed) CreateBACnetPropertyReferenceEnclosedBuilder() BACnetPropertyReferenceEnclosedBuilder {
	if b == nil {
		return NewBACnetPropertyReferenceEnclosedBuilder()
	}
	return &_BACnetPropertyReferenceEnclosedBuilder{_BACnetPropertyReferenceEnclosed: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyReferenceEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetPropertyReferenceEnclosed) GetReference() BACnetPropertyReference {
	return m.Reference
}

func (m *_BACnetPropertyReferenceEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyReferenceEnclosed(structType any) BACnetPropertyReferenceEnclosed {
	if casted, ok := structType.(BACnetPropertyReferenceEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyReferenceEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyReferenceEnclosed) GetTypeName() string {
	return "BACnetPropertyReferenceEnclosed"
}

func (m *_BACnetPropertyReferenceEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (reference)
	lengthInBits += m.Reference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyReferenceEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyReferenceEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetPropertyReferenceEnclosed, error) {
	return BACnetPropertyReferenceEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetPropertyReferenceEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyReferenceEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyReferenceEnclosed, error) {
		return BACnetPropertyReferenceEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetPropertyReferenceEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetPropertyReferenceEnclosed, error) {
	v, err := (&_BACnetPropertyReferenceEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetPropertyReferenceEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetPropertyReferenceEnclosed BACnetPropertyReferenceEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyReferenceEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyReferenceEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	reference, err := ReadSimpleField[BACnetPropertyReference](ctx, "reference", ReadComplex[BACnetPropertyReference](BACnetPropertyReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reference' field"))
	}
	m.Reference = reference

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetPropertyReferenceEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyReferenceEnclosed")
	}

	return m, nil
}

func (m *_BACnetPropertyReferenceEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyReferenceEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPropertyReferenceEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyReferenceEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetPropertyReference](ctx, "reference", m.GetReference(), WriteComplex[BACnetPropertyReference](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reference' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyReferenceEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyReferenceEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetPropertyReferenceEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetPropertyReferenceEnclosed) IsBACnetPropertyReferenceEnclosed() {}

func (m *_BACnetPropertyReferenceEnclosed) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyReferenceEnclosed) deepCopy() *_BACnetPropertyReferenceEnclosed {
	if m == nil {
		return nil
	}
	_BACnetPropertyReferenceEnclosedCopy := &_BACnetPropertyReferenceEnclosed{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.Reference.DeepCopy().(BACnetPropertyReference),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetPropertyReferenceEnclosedCopy
}

func (m *_BACnetPropertyReferenceEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

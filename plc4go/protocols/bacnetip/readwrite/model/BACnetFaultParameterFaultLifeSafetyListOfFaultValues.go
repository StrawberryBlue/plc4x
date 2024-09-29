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

// BACnetFaultParameterFaultLifeSafetyListOfFaultValues is the corresponding interface of BACnetFaultParameterFaultLifeSafetyListOfFaultValues
type BACnetFaultParameterFaultLifeSafetyListOfFaultValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListIfFaultValues returns ListIfFaultValues (property field)
	GetListIfFaultValues() []BACnetLifeSafetyStateTagged
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetFaultParameterFaultLifeSafetyListOfFaultValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultLifeSafetyListOfFaultValues()
	// CreateBuilder creates a BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder
	CreateBACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder() BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder
}

// _BACnetFaultParameterFaultLifeSafetyListOfFaultValues is the data-structure of this message
type _BACnetFaultParameterFaultLifeSafetyListOfFaultValues struct {
	OpeningTag        BACnetOpeningTag
	ListIfFaultValues []BACnetLifeSafetyStateTagged
	ClosingTag        BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetFaultParameterFaultLifeSafetyListOfFaultValues = (*_BACnetFaultParameterFaultLifeSafetyListOfFaultValues)(nil)

// NewBACnetFaultParameterFaultLifeSafetyListOfFaultValues factory function for _BACnetFaultParameterFaultLifeSafetyListOfFaultValues
func NewBACnetFaultParameterFaultLifeSafetyListOfFaultValues(openingTag BACnetOpeningTag, listIfFaultValues []BACnetLifeSafetyStateTagged, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetFaultParameterFaultLifeSafetyListOfFaultValues must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetFaultParameterFaultLifeSafetyListOfFaultValues must not be nil")
	}
	return &_BACnetFaultParameterFaultLifeSafetyListOfFaultValues{OpeningTag: openingTag, ListIfFaultValues: listIfFaultValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder is a builder for BACnetFaultParameterFaultLifeSafetyListOfFaultValues
type BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, listIfFaultValues []BACnetLifeSafetyStateTagged, closingTag BACnetClosingTag) BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder
	// WithListIfFaultValues adds ListIfFaultValues (property field)
	WithListIfFaultValues(...BACnetLifeSafetyStateTagged) BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder
	// Build builds the BACnetFaultParameterFaultLifeSafetyListOfFaultValues or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultLifeSafetyListOfFaultValues, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultLifeSafetyListOfFaultValues
}

// NewBACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder() creates a BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder
func NewBACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder() BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder {
	return &_BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder{_BACnetFaultParameterFaultLifeSafetyListOfFaultValues: new(_BACnetFaultParameterFaultLifeSafetyListOfFaultValues)}
}

type _BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder struct {
	*_BACnetFaultParameterFaultLifeSafetyListOfFaultValues

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder) = (*_BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder)(nil)

func (b *_BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, listIfFaultValues []BACnetLifeSafetyStateTagged, closingTag BACnetClosingTag) BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder {
	return b.WithOpeningTag(openingTag).WithListIfFaultValues(listIfFaultValues...).WithClosingTag(closingTag)
}

func (b *_BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder {
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

func (b *_BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder) WithListIfFaultValues(listIfFaultValues ...BACnetLifeSafetyStateTagged) BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder {
	b.ListIfFaultValues = listIfFaultValues
	return b
}

func (b *_BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder {
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

func (b *_BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder) Build() (BACnetFaultParameterFaultLifeSafetyListOfFaultValues, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
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
	return b._BACnetFaultParameterFaultLifeSafetyListOfFaultValues.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder) MustBuild() BACnetFaultParameterFaultLifeSafetyListOfFaultValues {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder().(*_BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder creates a BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder
func (b *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) CreateBACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder() BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder()
	}
	return &_BACnetFaultParameterFaultLifeSafetyListOfFaultValuesBuilder{_BACnetFaultParameterFaultLifeSafetyListOfFaultValues: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) GetListIfFaultValues() []BACnetLifeSafetyStateTagged {
	return m.ListIfFaultValues
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultLifeSafetyListOfFaultValues(structType any) BACnetFaultParameterFaultLifeSafetyListOfFaultValues {
	if casted, ok := structType.(BACnetFaultParameterFaultLifeSafetyListOfFaultValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultLifeSafetyListOfFaultValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) GetTypeName() string {
	return "BACnetFaultParameterFaultLifeSafetyListOfFaultValues"
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListIfFaultValues) > 0 {
		for _, element := range m.ListIfFaultValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultLifeSafetyListOfFaultValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetFaultParameterFaultLifeSafetyListOfFaultValues, error) {
	return BACnetFaultParameterFaultLifeSafetyListOfFaultValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetFaultParameterFaultLifeSafetyListOfFaultValuesParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultLifeSafetyListOfFaultValues, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultLifeSafetyListOfFaultValues, error) {
		return BACnetFaultParameterFaultLifeSafetyListOfFaultValuesParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetFaultParameterFaultLifeSafetyListOfFaultValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultLifeSafetyListOfFaultValues, error) {
	v, err := (&_BACnetFaultParameterFaultLifeSafetyListOfFaultValues{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetFaultParameterFaultLifeSafetyListOfFaultValues BACnetFaultParameterFaultLifeSafetyListOfFaultValues, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultLifeSafetyListOfFaultValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultLifeSafetyListOfFaultValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listIfFaultValues, err := ReadTerminatedArrayField[BACnetLifeSafetyStateTagged](ctx, "listIfFaultValues", ReadComplex[BACnetLifeSafetyStateTagged](BACnetLifeSafetyStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listIfFaultValues' field"))
	}
	m.ListIfFaultValues = listIfFaultValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultLifeSafetyListOfFaultValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultLifeSafetyListOfFaultValues")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultLifeSafetyListOfFaultValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultLifeSafetyListOfFaultValues")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listIfFaultValues", m.GetListIfFaultValues(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listIfFaultValues' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultLifeSafetyListOfFaultValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultLifeSafetyListOfFaultValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) IsBACnetFaultParameterFaultLifeSafetyListOfFaultValues() {
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) deepCopy() *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultLifeSafetyListOfFaultValuesCopy := &_BACnetFaultParameterFaultLifeSafetyListOfFaultValues{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		utils.DeepCopySlice[BACnetLifeSafetyStateTagged, BACnetLifeSafetyStateTagged](m.ListIfFaultValues),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetFaultParameterFaultLifeSafetyListOfFaultValuesCopy
}

func (m *_BACnetFaultParameterFaultLifeSafetyListOfFaultValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

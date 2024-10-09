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

// BACnetEventParameterChangeOfCharacterStringListOfAlarmValues is the corresponding interface of BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
type BACnetEventParameterChangeOfCharacterStringListOfAlarmValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfAlarmValues returns ListOfAlarmValues (property field)
	GetListOfAlarmValues() []BACnetApplicationTagCharacterString
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterChangeOfCharacterStringListOfAlarmValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterChangeOfCharacterStringListOfAlarmValues()
	// CreateBuilder creates a BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder
	CreateBACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder() BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder
}

// _BACnetEventParameterChangeOfCharacterStringListOfAlarmValues is the data-structure of this message
type _BACnetEventParameterChangeOfCharacterStringListOfAlarmValues struct {
	OpeningTag        BACnetOpeningTag
	ListOfAlarmValues []BACnetApplicationTagCharacterString
	ClosingTag        BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetEventParameterChangeOfCharacterStringListOfAlarmValues = (*_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues)(nil)

// NewBACnetEventParameterChangeOfCharacterStringListOfAlarmValues factory function for _BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
func NewBACnetEventParameterChangeOfCharacterStringListOfAlarmValues(openingTag BACnetOpeningTag, listOfAlarmValues []BACnetApplicationTagCharacterString, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues must not be nil")
	}
	return &_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues{OpeningTag: openingTag, ListOfAlarmValues: listOfAlarmValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder is a builder for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
type BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, listOfAlarmValues []BACnetApplicationTagCharacterString, closingTag BACnetClosingTag) BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder
	// WithListOfAlarmValues adds ListOfAlarmValues (property field)
	WithListOfAlarmValues(...BACnetApplicationTagCharacterString) BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder
	// Build builds the BACnetEventParameterChangeOfCharacterStringListOfAlarmValues or returns an error if something is wrong
	Build() (BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterChangeOfCharacterStringListOfAlarmValues
}

// NewBACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder() creates a BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder
func NewBACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder() BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder {
	return &_BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder{_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues: new(_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues)}
}

type _BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder struct {
	*_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues

	err *utils.MultiError
}

var _ (BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder) = (*_BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder)(nil)

func (b *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, listOfAlarmValues []BACnetApplicationTagCharacterString, closingTag BACnetClosingTag) BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder {
	return b.WithOpeningTag(openingTag).WithListOfAlarmValues(listOfAlarmValues...).WithClosingTag(closingTag)
}

func (b *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder {
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

func (b *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder) WithListOfAlarmValues(listOfAlarmValues ...BACnetApplicationTagCharacterString) BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder {
	b.ListOfAlarmValues = listOfAlarmValues
	return b
}

func (b *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder {
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

func (b *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder) Build() (BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, error) {
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
	return b._BACnetEventParameterChangeOfCharacterStringListOfAlarmValues.deepCopy(), nil
}

func (b *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder) MustBuild() BACnetEventParameterChangeOfCharacterStringListOfAlarmValues {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder().(*_BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder creates a BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder
func (b *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) CreateBACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder() BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder {
	if b == nil {
		return NewBACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder()
	}
	return &_BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesBuilder{_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetListOfAlarmValues() []BACnetApplicationTagCharacterString {
	return m.ListOfAlarmValues
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfCharacterStringListOfAlarmValues(structType any) BACnetEventParameterChangeOfCharacterStringListOfAlarmValues {
	if casted, ok := structType.(BACnetEventParameterChangeOfCharacterStringListOfAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfCharacterStringListOfAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetTypeName() string {
	return "BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfAlarmValues) > 0 {
		for _, element := range m.ListOfAlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, error) {
	return BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, error) {
		return BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, error) {
	v, err := (&_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetEventParameterChangeOfCharacterStringListOfAlarmValues BACnetEventParameterChangeOfCharacterStringListOfAlarmValues, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfAlarmValues, err := ReadTerminatedArrayField[BACnetApplicationTagCharacterString](ctx, "listOfAlarmValues", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfAlarmValues' field"))
	}
	m.ListOfAlarmValues = listOfAlarmValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
	}

	return m, nil
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfAlarmValues", m.GetListOfAlarmValues(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfAlarmValues' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfCharacterStringListOfAlarmValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfCharacterStringListOfAlarmValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) IsBACnetEventParameterChangeOfCharacterStringListOfAlarmValues() {
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) deepCopy() *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues {
	if m == nil {
		return nil
	}
	_BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesCopy := &_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		utils.DeepCopySlice[BACnetApplicationTagCharacterString, BACnetApplicationTagCharacterString](m.ListOfAlarmValues),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetEventParameterChangeOfCharacterStringListOfAlarmValuesCopy
}

func (m *_BACnetEventParameterChangeOfCharacterStringListOfAlarmValues) String() string {
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

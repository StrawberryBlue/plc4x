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

// BACnetFaultParameterFaultExtendedParametersEntryCharacterString is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryCharacterString
type BACnetFaultParameterFaultExtendedParametersEntryCharacterString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetCharacterStringValue returns CharacterStringValue (property field)
	GetCharacterStringValue() BACnetApplicationTagCharacterString
	// IsBACnetFaultParameterFaultExtendedParametersEntryCharacterString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultExtendedParametersEntryCharacterString()
	// CreateBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder
	CreateBACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder() BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder
}

// _BACnetFaultParameterFaultExtendedParametersEntryCharacterString is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryCharacterString struct {
	BACnetFaultParameterFaultExtendedParametersEntryContract
	CharacterStringValue BACnetApplicationTagCharacterString
}

var _ BACnetFaultParameterFaultExtendedParametersEntryCharacterString = (*_BACnetFaultParameterFaultExtendedParametersEntryCharacterString)(nil)
var _ BACnetFaultParameterFaultExtendedParametersEntryRequirements = (*_BACnetFaultParameterFaultExtendedParametersEntryCharacterString)(nil)

// NewBACnetFaultParameterFaultExtendedParametersEntryCharacterString factory function for _BACnetFaultParameterFaultExtendedParametersEntryCharacterString
func NewBACnetFaultParameterFaultExtendedParametersEntryCharacterString(peekedTagHeader BACnetTagHeader, characterStringValue BACnetApplicationTagCharacterString) *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString {
	if characterStringValue == nil {
		panic("characterStringValue of type BACnetApplicationTagCharacterString for BACnetFaultParameterFaultExtendedParametersEntryCharacterString must not be nil")
	}
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryCharacterString{
		BACnetFaultParameterFaultExtendedParametersEntryContract: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
		CharacterStringValue: characterStringValue,
	}
	_result.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder is a builder for BACnetFaultParameterFaultExtendedParametersEntryCharacterString
type BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(characterStringValue BACnetApplicationTagCharacterString) BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder
	// WithCharacterStringValue adds CharacterStringValue (property field)
	WithCharacterStringValue(BACnetApplicationTagCharacterString) BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder
	// WithCharacterStringValueBuilder adds CharacterStringValue (property field) which is build by the builder
	WithCharacterStringValueBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder
	// Build builds the BACnetFaultParameterFaultExtendedParametersEntryCharacterString or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultExtendedParametersEntryCharacterString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultExtendedParametersEntryCharacterString
}

// NewBACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder() creates a BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder
func NewBACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder() BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder {
	return &_BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder{_BACnetFaultParameterFaultExtendedParametersEntryCharacterString: new(_BACnetFaultParameterFaultExtendedParametersEntryCharacterString)}
}

type _BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder struct {
	*_BACnetFaultParameterFaultExtendedParametersEntryCharacterString

	parentBuilder *_BACnetFaultParameterFaultExtendedParametersEntryBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder) = (*_BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder)(nil)

func (b *_BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder) setParent(contract BACnetFaultParameterFaultExtendedParametersEntryContract) {
	b.BACnetFaultParameterFaultExtendedParametersEntryContract = contract
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder) WithMandatoryFields(characterStringValue BACnetApplicationTagCharacterString) BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder {
	return b.WithCharacterStringValue(characterStringValue)
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder) WithCharacterStringValue(characterStringValue BACnetApplicationTagCharacterString) BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder {
	b.CharacterStringValue = characterStringValue
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder) WithCharacterStringValueBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder {
	builder := builderSupplier(b.CharacterStringValue.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.CharacterStringValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder) Build() (BACnetFaultParameterFaultExtendedParametersEntryCharacterString, error) {
	if b.CharacterStringValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'characterStringValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultExtendedParametersEntryCharacterString.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder) MustBuild() BACnetFaultParameterFaultExtendedParametersEntryCharacterString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder) Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder {
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder) buildForBACnetFaultParameterFaultExtendedParametersEntry() (BACnetFaultParameterFaultExtendedParametersEntry, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder
func (b *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString) CreateBACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder() BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder()
	}
	return &_BACnetFaultParameterFaultExtendedParametersEntryCharacterStringBuilder{_BACnetFaultParameterFaultExtendedParametersEntryCharacterString: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString) GetParent() BACnetFaultParameterFaultExtendedParametersEntryContract {
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString) GetCharacterStringValue() BACnetApplicationTagCharacterString {
	return m.CharacterStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryCharacterString(structType any) BACnetFaultParameterFaultExtendedParametersEntryCharacterString {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryCharacterString"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).getLengthInBits(ctx))

	// Simple field (characterStringValue)
	lengthInBits += m.CharacterStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultExtendedParametersEntry) (__bACnetFaultParameterFaultExtendedParametersEntryCharacterString BACnetFaultParameterFaultExtendedParametersEntryCharacterString, err error) {
	m.BACnetFaultParameterFaultExtendedParametersEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	characterStringValue, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "characterStringValue", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'characterStringValue' field"))
	}
	m.CharacterStringValue = characterStringValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryCharacterString")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryCharacterString")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "characterStringValue", m.GetCharacterStringValue(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'characterStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryCharacterString")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString) IsBACnetFaultParameterFaultExtendedParametersEntryCharacterString() {
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString) deepCopy() *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultExtendedParametersEntryCharacterStringCopy := &_BACnetFaultParameterFaultExtendedParametersEntryCharacterString{
		m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).deepCopy(),
		m.CharacterStringValue.DeepCopy().(BACnetApplicationTagCharacterString),
	}
	m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = m
	return _BACnetFaultParameterFaultExtendedParametersEntryCharacterStringCopy
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

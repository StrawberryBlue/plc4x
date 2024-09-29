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

// BACnetFaultParameterFaultExtendedParametersEntryUnsigned is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryUnsigned
type BACnetFaultParameterFaultExtendedParametersEntryUnsigned interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetFaultParameterFaultExtendedParametersEntryUnsigned is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultExtendedParametersEntryUnsigned()
	// CreateBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder
	CreateBACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder() BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder
}

// _BACnetFaultParameterFaultExtendedParametersEntryUnsigned is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryUnsigned struct {
	BACnetFaultParameterFaultExtendedParametersEntryContract
	UnsignedValue BACnetApplicationTagUnsignedInteger
}

var _ BACnetFaultParameterFaultExtendedParametersEntryUnsigned = (*_BACnetFaultParameterFaultExtendedParametersEntryUnsigned)(nil)
var _ BACnetFaultParameterFaultExtendedParametersEntryRequirements = (*_BACnetFaultParameterFaultExtendedParametersEntryUnsigned)(nil)

// NewBACnetFaultParameterFaultExtendedParametersEntryUnsigned factory function for _BACnetFaultParameterFaultExtendedParametersEntryUnsigned
func NewBACnetFaultParameterFaultExtendedParametersEntryUnsigned(peekedTagHeader BACnetTagHeader, unsignedValue BACnetApplicationTagUnsignedInteger) *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned {
	if unsignedValue == nil {
		panic("unsignedValue of type BACnetApplicationTagUnsignedInteger for BACnetFaultParameterFaultExtendedParametersEntryUnsigned must not be nil")
	}
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryUnsigned{
		BACnetFaultParameterFaultExtendedParametersEntryContract: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
		UnsignedValue: unsignedValue,
	}
	_result.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder is a builder for BACnetFaultParameterFaultExtendedParametersEntryUnsigned
type BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(unsignedValue BACnetApplicationTagUnsignedInteger) BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder
	// WithUnsignedValue adds UnsignedValue (property field)
	WithUnsignedValue(BACnetApplicationTagUnsignedInteger) BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder
	// WithUnsignedValueBuilder adds UnsignedValue (property field) which is build by the builder
	WithUnsignedValueBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder
	// Build builds the BACnetFaultParameterFaultExtendedParametersEntryUnsigned or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultExtendedParametersEntryUnsigned, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultExtendedParametersEntryUnsigned
}

// NewBACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder() creates a BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder
func NewBACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder() BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder {
	return &_BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder{_BACnetFaultParameterFaultExtendedParametersEntryUnsigned: new(_BACnetFaultParameterFaultExtendedParametersEntryUnsigned)}
}

type _BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder struct {
	*_BACnetFaultParameterFaultExtendedParametersEntryUnsigned

	parentBuilder *_BACnetFaultParameterFaultExtendedParametersEntryBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder) = (*_BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder)(nil)

func (b *_BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder) setParent(contract BACnetFaultParameterFaultExtendedParametersEntryContract) {
	b.BACnetFaultParameterFaultExtendedParametersEntryContract = contract
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder) WithMandatoryFields(unsignedValue BACnetApplicationTagUnsignedInteger) BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder {
	return b.WithUnsignedValue(unsignedValue)
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder) WithUnsignedValue(unsignedValue BACnetApplicationTagUnsignedInteger) BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder {
	b.UnsignedValue = unsignedValue
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder) WithUnsignedValueBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder {
	builder := builderSupplier(b.UnsignedValue.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.UnsignedValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder) Build() (BACnetFaultParameterFaultExtendedParametersEntryUnsigned, error) {
	if b.UnsignedValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'unsignedValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultExtendedParametersEntryUnsigned.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder) MustBuild() BACnetFaultParameterFaultExtendedParametersEntryUnsigned {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder) Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder {
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder) buildForBACnetFaultParameterFaultExtendedParametersEntry() (BACnetFaultParameterFaultExtendedParametersEntry, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder
func (b *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned) CreateBACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder() BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder()
	}
	return &_BACnetFaultParameterFaultExtendedParametersEntryUnsignedBuilder{_BACnetFaultParameterFaultExtendedParametersEntryUnsigned: b.deepCopy()}
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

func (m *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned) GetParent() BACnetFaultParameterFaultExtendedParametersEntryContract {
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryUnsigned(structType any) BACnetFaultParameterFaultExtendedParametersEntryUnsigned {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryUnsigned); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryUnsigned"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).getLengthInBits(ctx))

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultExtendedParametersEntry) (__bACnetFaultParameterFaultExtendedParametersEntryUnsigned BACnetFaultParameterFaultExtendedParametersEntryUnsigned, err error) {
	m.BACnetFaultParameterFaultExtendedParametersEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryUnsigned"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryUnsigned")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	unsignedValue, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unsignedValue' field"))
	}
	m.UnsignedValue = unsignedValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryUnsigned"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryUnsigned")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryUnsigned"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryUnsigned")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", m.GetUnsignedValue(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryUnsigned"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryUnsigned")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned) IsBACnetFaultParameterFaultExtendedParametersEntryUnsigned() {
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned) deepCopy() *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultExtendedParametersEntryUnsignedCopy := &_BACnetFaultParameterFaultExtendedParametersEntryUnsigned{
		m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).deepCopy(),
		m.UnsignedValue.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = m
	return _BACnetFaultParameterFaultExtendedParametersEntryUnsignedCopy
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryUnsigned) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

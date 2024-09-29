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

// BACnetConstructedDataMaxAPDULengthAccepted is the corresponding interface of BACnetConstructedDataMaxAPDULengthAccepted
type BACnetConstructedDataMaxAPDULengthAccepted interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaxApduLengthAccepted returns MaxApduLengthAccepted (property field)
	GetMaxApduLengthAccepted() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataMaxAPDULengthAccepted is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMaxAPDULengthAccepted()
	// CreateBuilder creates a BACnetConstructedDataMaxAPDULengthAcceptedBuilder
	CreateBACnetConstructedDataMaxAPDULengthAcceptedBuilder() BACnetConstructedDataMaxAPDULengthAcceptedBuilder
}

// _BACnetConstructedDataMaxAPDULengthAccepted is the data-structure of this message
type _BACnetConstructedDataMaxAPDULengthAccepted struct {
	BACnetConstructedDataContract
	MaxApduLengthAccepted BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataMaxAPDULengthAccepted = (*_BACnetConstructedDataMaxAPDULengthAccepted)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMaxAPDULengthAccepted)(nil)

// NewBACnetConstructedDataMaxAPDULengthAccepted factory function for _BACnetConstructedDataMaxAPDULengthAccepted
func NewBACnetConstructedDataMaxAPDULengthAccepted(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maxApduLengthAccepted BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaxAPDULengthAccepted {
	if maxApduLengthAccepted == nil {
		panic("maxApduLengthAccepted of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataMaxAPDULengthAccepted must not be nil")
	}
	_result := &_BACnetConstructedDataMaxAPDULengthAccepted{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaxApduLengthAccepted:         maxApduLengthAccepted,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMaxAPDULengthAcceptedBuilder is a builder for BACnetConstructedDataMaxAPDULengthAccepted
type BACnetConstructedDataMaxAPDULengthAcceptedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maxApduLengthAccepted BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMaxAPDULengthAcceptedBuilder
	// WithMaxApduLengthAccepted adds MaxApduLengthAccepted (property field)
	WithMaxApduLengthAccepted(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMaxAPDULengthAcceptedBuilder
	// WithMaxApduLengthAcceptedBuilder adds MaxApduLengthAccepted (property field) which is build by the builder
	WithMaxApduLengthAcceptedBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataMaxAPDULengthAcceptedBuilder
	// Build builds the BACnetConstructedDataMaxAPDULengthAccepted or returns an error if something is wrong
	Build() (BACnetConstructedDataMaxAPDULengthAccepted, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMaxAPDULengthAccepted
}

// NewBACnetConstructedDataMaxAPDULengthAcceptedBuilder() creates a BACnetConstructedDataMaxAPDULengthAcceptedBuilder
func NewBACnetConstructedDataMaxAPDULengthAcceptedBuilder() BACnetConstructedDataMaxAPDULengthAcceptedBuilder {
	return &_BACnetConstructedDataMaxAPDULengthAcceptedBuilder{_BACnetConstructedDataMaxAPDULengthAccepted: new(_BACnetConstructedDataMaxAPDULengthAccepted)}
}

type _BACnetConstructedDataMaxAPDULengthAcceptedBuilder struct {
	*_BACnetConstructedDataMaxAPDULengthAccepted

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMaxAPDULengthAcceptedBuilder) = (*_BACnetConstructedDataMaxAPDULengthAcceptedBuilder)(nil)

func (b *_BACnetConstructedDataMaxAPDULengthAcceptedBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataMaxAPDULengthAcceptedBuilder) WithMandatoryFields(maxApduLengthAccepted BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMaxAPDULengthAcceptedBuilder {
	return b.WithMaxApduLengthAccepted(maxApduLengthAccepted)
}

func (b *_BACnetConstructedDataMaxAPDULengthAcceptedBuilder) WithMaxApduLengthAccepted(maxApduLengthAccepted BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMaxAPDULengthAcceptedBuilder {
	b.MaxApduLengthAccepted = maxApduLengthAccepted
	return b
}

func (b *_BACnetConstructedDataMaxAPDULengthAcceptedBuilder) WithMaxApduLengthAcceptedBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataMaxAPDULengthAcceptedBuilder {
	builder := builderSupplier(b.MaxApduLengthAccepted.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.MaxApduLengthAccepted, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataMaxAPDULengthAcceptedBuilder) Build() (BACnetConstructedDataMaxAPDULengthAccepted, error) {
	if b.MaxApduLengthAccepted == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'maxApduLengthAccepted' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMaxAPDULengthAccepted.deepCopy(), nil
}

func (b *_BACnetConstructedDataMaxAPDULengthAcceptedBuilder) MustBuild() BACnetConstructedDataMaxAPDULengthAccepted {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataMaxAPDULengthAcceptedBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMaxAPDULengthAcceptedBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMaxAPDULengthAcceptedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMaxAPDULengthAcceptedBuilder().(*_BACnetConstructedDataMaxAPDULengthAcceptedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMaxAPDULengthAcceptedBuilder creates a BACnetConstructedDataMaxAPDULengthAcceptedBuilder
func (b *_BACnetConstructedDataMaxAPDULengthAccepted) CreateBACnetConstructedDataMaxAPDULengthAcceptedBuilder() BACnetConstructedDataMaxAPDULengthAcceptedBuilder {
	if b == nil {
		return NewBACnetConstructedDataMaxAPDULengthAcceptedBuilder()
	}
	return &_BACnetConstructedDataMaxAPDULengthAcceptedBuilder{_BACnetConstructedDataMaxAPDULengthAccepted: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_APDU_LENGTH_ACCEPTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetMaxApduLengthAccepted() BACnetApplicationTagUnsignedInteger {
	return m.MaxApduLengthAccepted
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxApduLengthAccepted())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaxAPDULengthAccepted(structType any) BACnetConstructedDataMaxAPDULengthAccepted {
	if casted, ok := structType.(BACnetConstructedDataMaxAPDULengthAccepted); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaxAPDULengthAccepted); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetTypeName() string {
	return "BACnetConstructedDataMaxAPDULengthAccepted"
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maxApduLengthAccepted)
	lengthInBits += m.MaxApduLengthAccepted.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMaxAPDULengthAccepted BACnetConstructedDataMaxAPDULengthAccepted, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaxAPDULengthAccepted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaxAPDULengthAccepted")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maxApduLengthAccepted, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxApduLengthAccepted", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxApduLengthAccepted' field"))
	}
	m.MaxApduLengthAccepted = maxApduLengthAccepted

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), maxApduLengthAccepted)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaxAPDULengthAccepted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaxAPDULengthAccepted")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaxAPDULengthAccepted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaxAPDULengthAccepted")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxApduLengthAccepted", m.GetMaxApduLengthAccepted(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxApduLengthAccepted' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaxAPDULengthAccepted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaxAPDULengthAccepted")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) IsBACnetConstructedDataMaxAPDULengthAccepted() {
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) deepCopy() *_BACnetConstructedDataMaxAPDULengthAccepted {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMaxAPDULengthAcceptedCopy := &_BACnetConstructedDataMaxAPDULengthAccepted{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.MaxApduLengthAccepted.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMaxAPDULengthAcceptedCopy
}

func (m *_BACnetConstructedDataMaxAPDULengthAccepted) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

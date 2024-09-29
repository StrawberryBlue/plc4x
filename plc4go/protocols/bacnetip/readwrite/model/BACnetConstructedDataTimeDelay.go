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

// BACnetConstructedDataTimeDelay is the corresponding interface of BACnetConstructedDataTimeDelay
type BACnetConstructedDataTimeDelay interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataTimeDelay is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimeDelay()
	// CreateBuilder creates a BACnetConstructedDataTimeDelayBuilder
	CreateBACnetConstructedDataTimeDelayBuilder() BACnetConstructedDataTimeDelayBuilder
}

// _BACnetConstructedDataTimeDelay is the data-structure of this message
type _BACnetConstructedDataTimeDelay struct {
	BACnetConstructedDataContract
	TimeDelay BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataTimeDelay = (*_BACnetConstructedDataTimeDelay)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimeDelay)(nil)

// NewBACnetConstructedDataTimeDelay factory function for _BACnetConstructedDataTimeDelay
func NewBACnetConstructedDataTimeDelay(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, timeDelay BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeDelay {
	if timeDelay == nil {
		panic("timeDelay of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataTimeDelay must not be nil")
	}
	_result := &_BACnetConstructedDataTimeDelay{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		TimeDelay:                     timeDelay,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTimeDelayBuilder is a builder for BACnetConstructedDataTimeDelay
type BACnetConstructedDataTimeDelayBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timeDelay BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTimeDelayBuilder
	// WithTimeDelay adds TimeDelay (property field)
	WithTimeDelay(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTimeDelayBuilder
	// WithTimeDelayBuilder adds TimeDelay (property field) which is build by the builder
	WithTimeDelayBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataTimeDelayBuilder
	// Build builds the BACnetConstructedDataTimeDelay or returns an error if something is wrong
	Build() (BACnetConstructedDataTimeDelay, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTimeDelay
}

// NewBACnetConstructedDataTimeDelayBuilder() creates a BACnetConstructedDataTimeDelayBuilder
func NewBACnetConstructedDataTimeDelayBuilder() BACnetConstructedDataTimeDelayBuilder {
	return &_BACnetConstructedDataTimeDelayBuilder{_BACnetConstructedDataTimeDelay: new(_BACnetConstructedDataTimeDelay)}
}

type _BACnetConstructedDataTimeDelayBuilder struct {
	*_BACnetConstructedDataTimeDelay

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataTimeDelayBuilder) = (*_BACnetConstructedDataTimeDelayBuilder)(nil)

func (b *_BACnetConstructedDataTimeDelayBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataTimeDelayBuilder) WithMandatoryFields(timeDelay BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTimeDelayBuilder {
	return b.WithTimeDelay(timeDelay)
}

func (b *_BACnetConstructedDataTimeDelayBuilder) WithTimeDelay(timeDelay BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTimeDelayBuilder {
	b.TimeDelay = timeDelay
	return b
}

func (b *_BACnetConstructedDataTimeDelayBuilder) WithTimeDelayBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataTimeDelayBuilder {
	builder := builderSupplier(b.TimeDelay.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.TimeDelay, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataTimeDelayBuilder) Build() (BACnetConstructedDataTimeDelay, error) {
	if b.TimeDelay == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeDelay' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataTimeDelay.deepCopy(), nil
}

func (b *_BACnetConstructedDataTimeDelayBuilder) MustBuild() BACnetConstructedDataTimeDelay {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataTimeDelayBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataTimeDelayBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataTimeDelayBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataTimeDelayBuilder().(*_BACnetConstructedDataTimeDelayBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataTimeDelayBuilder creates a BACnetConstructedDataTimeDelayBuilder
func (b *_BACnetConstructedDataTimeDelay) CreateBACnetConstructedDataTimeDelayBuilder() BACnetConstructedDataTimeDelayBuilder {
	if b == nil {
		return NewBACnetConstructedDataTimeDelayBuilder()
	}
	return &_BACnetConstructedDataTimeDelayBuilder{_BACnetConstructedDataTimeDelay: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeDelay) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimeDelay) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_DELAY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeDelay) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeDelay) GetTimeDelay() BACnetApplicationTagUnsignedInteger {
	return m.TimeDelay
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimeDelay) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetTimeDelay())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeDelay(structType any) BACnetConstructedDataTimeDelay {
	if casted, ok := structType.(BACnetConstructedDataTimeDelay); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeDelay); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeDelay) GetTypeName() string {
	return "BACnetConstructedDataTimeDelay"
}

func (m *_BACnetConstructedDataTimeDelay) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeDelay) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimeDelay) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimeDelay BACnetConstructedDataTimeDelay, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeDelay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeDelay")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeDelay, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), timeDelay)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeDelay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeDelay")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimeDelay) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeDelay) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeDelay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeDelay")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeDelay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeDelay")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeDelay) IsBACnetConstructedDataTimeDelay() {}

func (m *_BACnetConstructedDataTimeDelay) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTimeDelay) deepCopy() *_BACnetConstructedDataTimeDelay {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTimeDelayCopy := &_BACnetConstructedDataTimeDelay{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.TimeDelay.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTimeDelayCopy
}

func (m *_BACnetConstructedDataTimeDelay) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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

// BACnetConstructedDataPositiveIntegerValueHighLimit is the corresponding interface of BACnetConstructedDataPositiveIntegerValueHighLimit
type BACnetConstructedDataPositiveIntegerValueHighLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataPositiveIntegerValueHighLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataPositiveIntegerValueHighLimit()
	// CreateBuilder creates a BACnetConstructedDataPositiveIntegerValueHighLimitBuilder
	CreateBACnetConstructedDataPositiveIntegerValueHighLimitBuilder() BACnetConstructedDataPositiveIntegerValueHighLimitBuilder
}

// _BACnetConstructedDataPositiveIntegerValueHighLimit is the data-structure of this message
type _BACnetConstructedDataPositiveIntegerValueHighLimit struct {
	BACnetConstructedDataContract
	HighLimit BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataPositiveIntegerValueHighLimit = (*_BACnetConstructedDataPositiveIntegerValueHighLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataPositiveIntegerValueHighLimit)(nil)

// NewBACnetConstructedDataPositiveIntegerValueHighLimit factory function for _BACnetConstructedDataPositiveIntegerValueHighLimit
func NewBACnetConstructedDataPositiveIntegerValueHighLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, highLimit BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPositiveIntegerValueHighLimit {
	if highLimit == nil {
		panic("highLimit of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataPositiveIntegerValueHighLimit must not be nil")
	}
	_result := &_BACnetConstructedDataPositiveIntegerValueHighLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		HighLimit:                     highLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataPositiveIntegerValueHighLimitBuilder is a builder for BACnetConstructedDataPositiveIntegerValueHighLimit
type BACnetConstructedDataPositiveIntegerValueHighLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(highLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueHighLimitBuilder
	// WithHighLimit adds HighLimit (property field)
	WithHighLimit(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueHighLimitBuilder
	// WithHighLimitBuilder adds HighLimit (property field) which is build by the builder
	WithHighLimitBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPositiveIntegerValueHighLimitBuilder
	// Build builds the BACnetConstructedDataPositiveIntegerValueHighLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataPositiveIntegerValueHighLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataPositiveIntegerValueHighLimit
}

// NewBACnetConstructedDataPositiveIntegerValueHighLimitBuilder() creates a BACnetConstructedDataPositiveIntegerValueHighLimitBuilder
func NewBACnetConstructedDataPositiveIntegerValueHighLimitBuilder() BACnetConstructedDataPositiveIntegerValueHighLimitBuilder {
	return &_BACnetConstructedDataPositiveIntegerValueHighLimitBuilder{_BACnetConstructedDataPositiveIntegerValueHighLimit: new(_BACnetConstructedDataPositiveIntegerValueHighLimit)}
}

type _BACnetConstructedDataPositiveIntegerValueHighLimitBuilder struct {
	*_BACnetConstructedDataPositiveIntegerValueHighLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataPositiveIntegerValueHighLimitBuilder) = (*_BACnetConstructedDataPositiveIntegerValueHighLimitBuilder)(nil)

func (b *_BACnetConstructedDataPositiveIntegerValueHighLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataPositiveIntegerValueHighLimitBuilder) WithMandatoryFields(highLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueHighLimitBuilder {
	return b.WithHighLimit(highLimit)
}

func (b *_BACnetConstructedDataPositiveIntegerValueHighLimitBuilder) WithHighLimit(highLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueHighLimitBuilder {
	b.HighLimit = highLimit
	return b
}

func (b *_BACnetConstructedDataPositiveIntegerValueHighLimitBuilder) WithHighLimitBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPositiveIntegerValueHighLimitBuilder {
	builder := builderSupplier(b.HighLimit.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.HighLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataPositiveIntegerValueHighLimitBuilder) Build() (BACnetConstructedDataPositiveIntegerValueHighLimit, error) {
	if b.HighLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'highLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataPositiveIntegerValueHighLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataPositiveIntegerValueHighLimitBuilder) MustBuild() BACnetConstructedDataPositiveIntegerValueHighLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataPositiveIntegerValueHighLimitBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataPositiveIntegerValueHighLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataPositiveIntegerValueHighLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataPositiveIntegerValueHighLimitBuilder().(*_BACnetConstructedDataPositiveIntegerValueHighLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataPositiveIntegerValueHighLimitBuilder creates a BACnetConstructedDataPositiveIntegerValueHighLimitBuilder
func (b *_BACnetConstructedDataPositiveIntegerValueHighLimit) CreateBACnetConstructedDataPositiveIntegerValueHighLimitBuilder() BACnetConstructedDataPositiveIntegerValueHighLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataPositiveIntegerValueHighLimitBuilder()
	}
	return &_BACnetConstructedDataPositiveIntegerValueHighLimitBuilder{_BACnetConstructedDataPositiveIntegerValueHighLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_POSITIVE_INTEGER_VALUE
}

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_HIGH_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) GetHighLimit() BACnetApplicationTagUnsignedInteger {
	return m.HighLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetHighLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPositiveIntegerValueHighLimit(structType any) BACnetConstructedDataPositiveIntegerValueHighLimit {
	if casted, ok := structType.(BACnetConstructedDataPositiveIntegerValueHighLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPositiveIntegerValueHighLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) GetTypeName() string {
	return "BACnetConstructedDataPositiveIntegerValueHighLimit"
}

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataPositiveIntegerValueHighLimit BACnetConstructedDataPositiveIntegerValueHighLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPositiveIntegerValueHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPositiveIntegerValueHighLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	highLimit, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "highLimit", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'highLimit' field"))
	}
	m.HighLimit = highLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), highLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPositiveIntegerValueHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPositiveIntegerValueHighLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPositiveIntegerValueHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPositiveIntegerValueHighLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "highLimit", m.GetHighLimit(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'highLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPositiveIntegerValueHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPositiveIntegerValueHighLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) IsBACnetConstructedDataPositiveIntegerValueHighLimit() {
}

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) deepCopy() *_BACnetConstructedDataPositiveIntegerValueHighLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataPositiveIntegerValueHighLimitCopy := &_BACnetConstructedDataPositiveIntegerValueHighLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.HighLimit.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataPositiveIntegerValueHighLimitCopy
}

func (m *_BACnetConstructedDataPositiveIntegerValueHighLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

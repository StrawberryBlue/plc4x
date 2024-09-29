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

// BACnetConstructedDataAttemptedSamples is the corresponding interface of BACnetConstructedDataAttemptedSamples
type BACnetConstructedDataAttemptedSamples interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAttemptedSamples returns AttemptedSamples (property field)
	GetAttemptedSamples() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataAttemptedSamples is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAttemptedSamples()
	// CreateBuilder creates a BACnetConstructedDataAttemptedSamplesBuilder
	CreateBACnetConstructedDataAttemptedSamplesBuilder() BACnetConstructedDataAttemptedSamplesBuilder
}

// _BACnetConstructedDataAttemptedSamples is the data-structure of this message
type _BACnetConstructedDataAttemptedSamples struct {
	BACnetConstructedDataContract
	AttemptedSamples BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataAttemptedSamples = (*_BACnetConstructedDataAttemptedSamples)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAttemptedSamples)(nil)

// NewBACnetConstructedDataAttemptedSamples factory function for _BACnetConstructedDataAttemptedSamples
func NewBACnetConstructedDataAttemptedSamples(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, attemptedSamples BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAttemptedSamples {
	if attemptedSamples == nil {
		panic("attemptedSamples of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataAttemptedSamples must not be nil")
	}
	_result := &_BACnetConstructedDataAttemptedSamples{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AttemptedSamples:              attemptedSamples,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAttemptedSamplesBuilder is a builder for BACnetConstructedDataAttemptedSamples
type BACnetConstructedDataAttemptedSamplesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(attemptedSamples BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAttemptedSamplesBuilder
	// WithAttemptedSamples adds AttemptedSamples (property field)
	WithAttemptedSamples(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAttemptedSamplesBuilder
	// WithAttemptedSamplesBuilder adds AttemptedSamples (property field) which is build by the builder
	WithAttemptedSamplesBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAttemptedSamplesBuilder
	// Build builds the BACnetConstructedDataAttemptedSamples or returns an error if something is wrong
	Build() (BACnetConstructedDataAttemptedSamples, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAttemptedSamples
}

// NewBACnetConstructedDataAttemptedSamplesBuilder() creates a BACnetConstructedDataAttemptedSamplesBuilder
func NewBACnetConstructedDataAttemptedSamplesBuilder() BACnetConstructedDataAttemptedSamplesBuilder {
	return &_BACnetConstructedDataAttemptedSamplesBuilder{_BACnetConstructedDataAttemptedSamples: new(_BACnetConstructedDataAttemptedSamples)}
}

type _BACnetConstructedDataAttemptedSamplesBuilder struct {
	*_BACnetConstructedDataAttemptedSamples

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAttemptedSamplesBuilder) = (*_BACnetConstructedDataAttemptedSamplesBuilder)(nil)

func (b *_BACnetConstructedDataAttemptedSamplesBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataAttemptedSamplesBuilder) WithMandatoryFields(attemptedSamples BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAttemptedSamplesBuilder {
	return b.WithAttemptedSamples(attemptedSamples)
}

func (b *_BACnetConstructedDataAttemptedSamplesBuilder) WithAttemptedSamples(attemptedSamples BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAttemptedSamplesBuilder {
	b.AttemptedSamples = attemptedSamples
	return b
}

func (b *_BACnetConstructedDataAttemptedSamplesBuilder) WithAttemptedSamplesBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAttemptedSamplesBuilder {
	builder := builderSupplier(b.AttemptedSamples.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.AttemptedSamples, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAttemptedSamplesBuilder) Build() (BACnetConstructedDataAttemptedSamples, error) {
	if b.AttemptedSamples == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'attemptedSamples' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAttemptedSamples.deepCopy(), nil
}

func (b *_BACnetConstructedDataAttemptedSamplesBuilder) MustBuild() BACnetConstructedDataAttemptedSamples {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataAttemptedSamplesBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAttemptedSamplesBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAttemptedSamplesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAttemptedSamplesBuilder().(*_BACnetConstructedDataAttemptedSamplesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAttemptedSamplesBuilder creates a BACnetConstructedDataAttemptedSamplesBuilder
func (b *_BACnetConstructedDataAttemptedSamples) CreateBACnetConstructedDataAttemptedSamplesBuilder() BACnetConstructedDataAttemptedSamplesBuilder {
	if b == nil {
		return NewBACnetConstructedDataAttemptedSamplesBuilder()
	}
	return &_BACnetConstructedDataAttemptedSamplesBuilder{_BACnetConstructedDataAttemptedSamples: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAttemptedSamples) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAttemptedSamples) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ATTEMPTED_SAMPLES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAttemptedSamples) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAttemptedSamples) GetAttemptedSamples() BACnetApplicationTagUnsignedInteger {
	return m.AttemptedSamples
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAttemptedSamples) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetAttemptedSamples())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAttemptedSamples(structType any) BACnetConstructedDataAttemptedSamples {
	if casted, ok := structType.(BACnetConstructedDataAttemptedSamples); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAttemptedSamples); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAttemptedSamples) GetTypeName() string {
	return "BACnetConstructedDataAttemptedSamples"
}

func (m *_BACnetConstructedDataAttemptedSamples) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (attemptedSamples)
	lengthInBits += m.AttemptedSamples.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAttemptedSamples) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAttemptedSamples) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAttemptedSamples BACnetConstructedDataAttemptedSamples, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAttemptedSamples"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAttemptedSamples")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	attemptedSamples, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "attemptedSamples", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'attemptedSamples' field"))
	}
	m.AttemptedSamples = attemptedSamples

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), attemptedSamples)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAttemptedSamples"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAttemptedSamples")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAttemptedSamples) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAttemptedSamples) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAttemptedSamples"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAttemptedSamples")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "attemptedSamples", m.GetAttemptedSamples(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'attemptedSamples' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAttemptedSamples"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAttemptedSamples")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAttemptedSamples) IsBACnetConstructedDataAttemptedSamples() {}

func (m *_BACnetConstructedDataAttemptedSamples) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAttemptedSamples) deepCopy() *_BACnetConstructedDataAttemptedSamples {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAttemptedSamplesCopy := &_BACnetConstructedDataAttemptedSamples{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.AttemptedSamples.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAttemptedSamplesCopy
}

func (m *_BACnetConstructedDataAttemptedSamples) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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

// BACnetConstructedDataPositiveIntegerValueResolution is the corresponding interface of BACnetConstructedDataPositiveIntegerValueResolution
type BACnetConstructedDataPositiveIntegerValueResolution interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetResolution returns Resolution (property field)
	GetResolution() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataPositiveIntegerValueResolution is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataPositiveIntegerValueResolution()
	// CreateBuilder creates a BACnetConstructedDataPositiveIntegerValueResolutionBuilder
	CreateBACnetConstructedDataPositiveIntegerValueResolutionBuilder() BACnetConstructedDataPositiveIntegerValueResolutionBuilder
}

// _BACnetConstructedDataPositiveIntegerValueResolution is the data-structure of this message
type _BACnetConstructedDataPositiveIntegerValueResolution struct {
	BACnetConstructedDataContract
	Resolution BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataPositiveIntegerValueResolution = (*_BACnetConstructedDataPositiveIntegerValueResolution)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataPositiveIntegerValueResolution)(nil)

// NewBACnetConstructedDataPositiveIntegerValueResolution factory function for _BACnetConstructedDataPositiveIntegerValueResolution
func NewBACnetConstructedDataPositiveIntegerValueResolution(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, resolution BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPositiveIntegerValueResolution {
	if resolution == nil {
		panic("resolution of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataPositiveIntegerValueResolution must not be nil")
	}
	_result := &_BACnetConstructedDataPositiveIntegerValueResolution{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Resolution:                    resolution,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataPositiveIntegerValueResolutionBuilder is a builder for BACnetConstructedDataPositiveIntegerValueResolution
type BACnetConstructedDataPositiveIntegerValueResolutionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(resolution BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueResolutionBuilder
	// WithResolution adds Resolution (property field)
	WithResolution(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueResolutionBuilder
	// WithResolutionBuilder adds Resolution (property field) which is build by the builder
	WithResolutionBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPositiveIntegerValueResolutionBuilder
	// Build builds the BACnetConstructedDataPositiveIntegerValueResolution or returns an error if something is wrong
	Build() (BACnetConstructedDataPositiveIntegerValueResolution, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataPositiveIntegerValueResolution
}

// NewBACnetConstructedDataPositiveIntegerValueResolutionBuilder() creates a BACnetConstructedDataPositiveIntegerValueResolutionBuilder
func NewBACnetConstructedDataPositiveIntegerValueResolutionBuilder() BACnetConstructedDataPositiveIntegerValueResolutionBuilder {
	return &_BACnetConstructedDataPositiveIntegerValueResolutionBuilder{_BACnetConstructedDataPositiveIntegerValueResolution: new(_BACnetConstructedDataPositiveIntegerValueResolution)}
}

type _BACnetConstructedDataPositiveIntegerValueResolutionBuilder struct {
	*_BACnetConstructedDataPositiveIntegerValueResolution

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataPositiveIntegerValueResolutionBuilder) = (*_BACnetConstructedDataPositiveIntegerValueResolutionBuilder)(nil)

func (b *_BACnetConstructedDataPositiveIntegerValueResolutionBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataPositiveIntegerValueResolutionBuilder) WithMandatoryFields(resolution BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueResolutionBuilder {
	return b.WithResolution(resolution)
}

func (b *_BACnetConstructedDataPositiveIntegerValueResolutionBuilder) WithResolution(resolution BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueResolutionBuilder {
	b.Resolution = resolution
	return b
}

func (b *_BACnetConstructedDataPositiveIntegerValueResolutionBuilder) WithResolutionBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPositiveIntegerValueResolutionBuilder {
	builder := builderSupplier(b.Resolution.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.Resolution, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataPositiveIntegerValueResolutionBuilder) Build() (BACnetConstructedDataPositiveIntegerValueResolution, error) {
	if b.Resolution == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'resolution' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataPositiveIntegerValueResolution.deepCopy(), nil
}

func (b *_BACnetConstructedDataPositiveIntegerValueResolutionBuilder) MustBuild() BACnetConstructedDataPositiveIntegerValueResolution {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataPositiveIntegerValueResolutionBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataPositiveIntegerValueResolutionBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataPositiveIntegerValueResolutionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataPositiveIntegerValueResolutionBuilder().(*_BACnetConstructedDataPositiveIntegerValueResolutionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataPositiveIntegerValueResolutionBuilder creates a BACnetConstructedDataPositiveIntegerValueResolutionBuilder
func (b *_BACnetConstructedDataPositiveIntegerValueResolution) CreateBACnetConstructedDataPositiveIntegerValueResolutionBuilder() BACnetConstructedDataPositiveIntegerValueResolutionBuilder {
	if b == nil {
		return NewBACnetConstructedDataPositiveIntegerValueResolutionBuilder()
	}
	return &_BACnetConstructedDataPositiveIntegerValueResolutionBuilder{_BACnetConstructedDataPositiveIntegerValueResolution: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_POSITIVE_INTEGER_VALUE
}

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RESOLUTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) GetResolution() BACnetApplicationTagUnsignedInteger {
	return m.Resolution
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetResolution())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPositiveIntegerValueResolution(structType any) BACnetConstructedDataPositiveIntegerValueResolution {
	if casted, ok := structType.(BACnetConstructedDataPositiveIntegerValueResolution); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPositiveIntegerValueResolution); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) GetTypeName() string {
	return "BACnetConstructedDataPositiveIntegerValueResolution"
}

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (resolution)
	lengthInBits += m.Resolution.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataPositiveIntegerValueResolution BACnetConstructedDataPositiveIntegerValueResolution, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPositiveIntegerValueResolution"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPositiveIntegerValueResolution")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	resolution, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "resolution", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'resolution' field"))
	}
	m.Resolution = resolution

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), resolution)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPositiveIntegerValueResolution"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPositiveIntegerValueResolution")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPositiveIntegerValueResolution"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPositiveIntegerValueResolution")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "resolution", m.GetResolution(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'resolution' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPositiveIntegerValueResolution"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPositiveIntegerValueResolution")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) IsBACnetConstructedDataPositiveIntegerValueResolution() {
}

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) deepCopy() *_BACnetConstructedDataPositiveIntegerValueResolution {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataPositiveIntegerValueResolutionCopy := &_BACnetConstructedDataPositiveIntegerValueResolution{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.Resolution.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataPositiveIntegerValueResolutionCopy
}

func (m *_BACnetConstructedDataPositiveIntegerValueResolution) String() string {
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

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

// BACnetConstructedDataLargeAnalogValueFaultLowLimit is the corresponding interface of BACnetConstructedDataLargeAnalogValueFaultLowLimit
type BACnetConstructedDataLargeAnalogValueFaultLowLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFaultLowLimit returns FaultLowLimit (property field)
	GetFaultLowLimit() BACnetApplicationTagDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDouble
	// IsBACnetConstructedDataLargeAnalogValueFaultLowLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLargeAnalogValueFaultLowLimit()
	// CreateBuilder creates a BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder
	CreateBACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder() BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder
}

// _BACnetConstructedDataLargeAnalogValueFaultLowLimit is the data-structure of this message
type _BACnetConstructedDataLargeAnalogValueFaultLowLimit struct {
	BACnetConstructedDataContract
	FaultLowLimit BACnetApplicationTagDouble
}

var _ BACnetConstructedDataLargeAnalogValueFaultLowLimit = (*_BACnetConstructedDataLargeAnalogValueFaultLowLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLargeAnalogValueFaultLowLimit)(nil)

// NewBACnetConstructedDataLargeAnalogValueFaultLowLimit factory function for _BACnetConstructedDataLargeAnalogValueFaultLowLimit
func NewBACnetConstructedDataLargeAnalogValueFaultLowLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, faultLowLimit BACnetApplicationTagDouble, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLargeAnalogValueFaultLowLimit {
	if faultLowLimit == nil {
		panic("faultLowLimit of type BACnetApplicationTagDouble for BACnetConstructedDataLargeAnalogValueFaultLowLimit must not be nil")
	}
	_result := &_BACnetConstructedDataLargeAnalogValueFaultLowLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FaultLowLimit:                 faultLowLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder is a builder for BACnetConstructedDataLargeAnalogValueFaultLowLimit
type BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(faultLowLimit BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder
	// WithFaultLowLimit adds FaultLowLimit (property field)
	WithFaultLowLimit(BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder
	// WithFaultLowLimitBuilder adds FaultLowLimit (property field) which is build by the builder
	WithFaultLowLimitBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder
	// Build builds the BACnetConstructedDataLargeAnalogValueFaultLowLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataLargeAnalogValueFaultLowLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLargeAnalogValueFaultLowLimit
}

// NewBACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder() creates a BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder
func NewBACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder() BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder {
	return &_BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder{_BACnetConstructedDataLargeAnalogValueFaultLowLimit: new(_BACnetConstructedDataLargeAnalogValueFaultLowLimit)}
}

type _BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder struct {
	*_BACnetConstructedDataLargeAnalogValueFaultLowLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder) = (*_BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder)(nil)

func (b *_BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder) WithMandatoryFields(faultLowLimit BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder {
	return b.WithFaultLowLimit(faultLowLimit)
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder) WithFaultLowLimit(faultLowLimit BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder {
	b.FaultLowLimit = faultLowLimit
	return b
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder) WithFaultLowLimitBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder {
	builder := builderSupplier(b.FaultLowLimit.CreateBACnetApplicationTagDoubleBuilder())
	var err error
	b.FaultLowLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDoubleBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder) Build() (BACnetConstructedDataLargeAnalogValueFaultLowLimit, error) {
	if b.FaultLowLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'faultLowLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLargeAnalogValueFaultLowLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder) MustBuild() BACnetConstructedDataLargeAnalogValueFaultLowLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder().(*_BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder creates a BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder
func (b *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) CreateBACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder() BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder()
	}
	return &_BACnetConstructedDataLargeAnalogValueFaultLowLimitBuilder{_BACnetConstructedDataLargeAnalogValueFaultLowLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) GetFaultLowLimit() BACnetApplicationTagDouble {
	return m.FaultLowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) GetActualValue() BACnetApplicationTagDouble {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDouble(m.GetFaultLowLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLargeAnalogValueFaultLowLimit(structType any) BACnetConstructedDataLargeAnalogValueFaultLowLimit {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueFaultLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueFaultLowLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueFaultLowLimit"
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (faultLowLimit)
	lengthInBits += m.FaultLowLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLargeAnalogValueFaultLowLimit BACnetConstructedDataLargeAnalogValueFaultLowLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueFaultLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueFaultLowLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultLowLimit, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "faultLowLimit", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultLowLimit' field"))
	}
	m.FaultLowLimit = faultLowLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagDouble](ctx, "actualValue", (*BACnetApplicationTagDouble)(nil), faultLowLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueFaultLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueFaultLowLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueFaultLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueFaultLowLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "faultLowLimit", m.GetFaultLowLimit(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'faultLowLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueFaultLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueFaultLowLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) IsBACnetConstructedDataLargeAnalogValueFaultLowLimit() {
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) deepCopy() *_BACnetConstructedDataLargeAnalogValueFaultLowLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLargeAnalogValueFaultLowLimitCopy := &_BACnetConstructedDataLargeAnalogValueFaultLowLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.FaultLowLimit.DeepCopy().(BACnetApplicationTagDouble),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLargeAnalogValueFaultLowLimitCopy
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) String() string {
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

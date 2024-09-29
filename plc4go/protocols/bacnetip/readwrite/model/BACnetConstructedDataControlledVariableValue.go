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

// BACnetConstructedDataControlledVariableValue is the corresponding interface of BACnetConstructedDataControlledVariableValue
type BACnetConstructedDataControlledVariableValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetControlledVariableValue returns ControlledVariableValue (property field)
	GetControlledVariableValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataControlledVariableValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataControlledVariableValue()
	// CreateBuilder creates a BACnetConstructedDataControlledVariableValueBuilder
	CreateBACnetConstructedDataControlledVariableValueBuilder() BACnetConstructedDataControlledVariableValueBuilder
}

// _BACnetConstructedDataControlledVariableValue is the data-structure of this message
type _BACnetConstructedDataControlledVariableValue struct {
	BACnetConstructedDataContract
	ControlledVariableValue BACnetApplicationTagReal
}

var _ BACnetConstructedDataControlledVariableValue = (*_BACnetConstructedDataControlledVariableValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataControlledVariableValue)(nil)

// NewBACnetConstructedDataControlledVariableValue factory function for _BACnetConstructedDataControlledVariableValue
func NewBACnetConstructedDataControlledVariableValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, controlledVariableValue BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataControlledVariableValue {
	if controlledVariableValue == nil {
		panic("controlledVariableValue of type BACnetApplicationTagReal for BACnetConstructedDataControlledVariableValue must not be nil")
	}
	_result := &_BACnetConstructedDataControlledVariableValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ControlledVariableValue:       controlledVariableValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataControlledVariableValueBuilder is a builder for BACnetConstructedDataControlledVariableValue
type BACnetConstructedDataControlledVariableValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(controlledVariableValue BACnetApplicationTagReal) BACnetConstructedDataControlledVariableValueBuilder
	// WithControlledVariableValue adds ControlledVariableValue (property field)
	WithControlledVariableValue(BACnetApplicationTagReal) BACnetConstructedDataControlledVariableValueBuilder
	// WithControlledVariableValueBuilder adds ControlledVariableValue (property field) which is build by the builder
	WithControlledVariableValueBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataControlledVariableValueBuilder
	// Build builds the BACnetConstructedDataControlledVariableValue or returns an error if something is wrong
	Build() (BACnetConstructedDataControlledVariableValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataControlledVariableValue
}

// NewBACnetConstructedDataControlledVariableValueBuilder() creates a BACnetConstructedDataControlledVariableValueBuilder
func NewBACnetConstructedDataControlledVariableValueBuilder() BACnetConstructedDataControlledVariableValueBuilder {
	return &_BACnetConstructedDataControlledVariableValueBuilder{_BACnetConstructedDataControlledVariableValue: new(_BACnetConstructedDataControlledVariableValue)}
}

type _BACnetConstructedDataControlledVariableValueBuilder struct {
	*_BACnetConstructedDataControlledVariableValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataControlledVariableValueBuilder) = (*_BACnetConstructedDataControlledVariableValueBuilder)(nil)

func (b *_BACnetConstructedDataControlledVariableValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataControlledVariableValueBuilder) WithMandatoryFields(controlledVariableValue BACnetApplicationTagReal) BACnetConstructedDataControlledVariableValueBuilder {
	return b.WithControlledVariableValue(controlledVariableValue)
}

func (b *_BACnetConstructedDataControlledVariableValueBuilder) WithControlledVariableValue(controlledVariableValue BACnetApplicationTagReal) BACnetConstructedDataControlledVariableValueBuilder {
	b.ControlledVariableValue = controlledVariableValue
	return b
}

func (b *_BACnetConstructedDataControlledVariableValueBuilder) WithControlledVariableValueBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataControlledVariableValueBuilder {
	builder := builderSupplier(b.ControlledVariableValue.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.ControlledVariableValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataControlledVariableValueBuilder) Build() (BACnetConstructedDataControlledVariableValue, error) {
	if b.ControlledVariableValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'controlledVariableValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataControlledVariableValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataControlledVariableValueBuilder) MustBuild() BACnetConstructedDataControlledVariableValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataControlledVariableValueBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataControlledVariableValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataControlledVariableValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataControlledVariableValueBuilder().(*_BACnetConstructedDataControlledVariableValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataControlledVariableValueBuilder creates a BACnetConstructedDataControlledVariableValueBuilder
func (b *_BACnetConstructedDataControlledVariableValue) CreateBACnetConstructedDataControlledVariableValueBuilder() BACnetConstructedDataControlledVariableValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataControlledVariableValueBuilder()
	}
	return &_BACnetConstructedDataControlledVariableValueBuilder{_BACnetConstructedDataControlledVariableValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataControlledVariableValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CONTROLLED_VARIABLE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataControlledVariableValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableValue) GetControlledVariableValue() BACnetApplicationTagReal {
	return m.ControlledVariableValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetControlledVariableValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataControlledVariableValue(structType any) BACnetConstructedDataControlledVariableValue {
	if casted, ok := structType.(BACnetConstructedDataControlledVariableValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataControlledVariableValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataControlledVariableValue) GetTypeName() string {
	return "BACnetConstructedDataControlledVariableValue"
}

func (m *_BACnetConstructedDataControlledVariableValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (controlledVariableValue)
	lengthInBits += m.ControlledVariableValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataControlledVariableValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataControlledVariableValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataControlledVariableValue BACnetConstructedDataControlledVariableValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataControlledVariableValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataControlledVariableValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	controlledVariableValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "controlledVariableValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'controlledVariableValue' field"))
	}
	m.ControlledVariableValue = controlledVariableValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), controlledVariableValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataControlledVariableValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataControlledVariableValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataControlledVariableValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataControlledVariableValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataControlledVariableValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataControlledVariableValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "controlledVariableValue", m.GetControlledVariableValue(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'controlledVariableValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataControlledVariableValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataControlledVariableValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataControlledVariableValue) IsBACnetConstructedDataControlledVariableValue() {
}

func (m *_BACnetConstructedDataControlledVariableValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataControlledVariableValue) deepCopy() *_BACnetConstructedDataControlledVariableValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataControlledVariableValueCopy := &_BACnetConstructedDataControlledVariableValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.ControlledVariableValue.DeepCopy().(BACnetApplicationTagReal),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataControlledVariableValueCopy
}

func (m *_BACnetConstructedDataControlledVariableValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

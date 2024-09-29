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

// BACnetConstructedDataOperationExpected is the corresponding interface of BACnetConstructedDataOperationExpected
type BACnetConstructedDataOperationExpected interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLifeSafetyOperations returns LifeSafetyOperations (property field)
	GetLifeSafetyOperations() BACnetLifeSafetyOperationTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLifeSafetyOperationTagged
	// IsBACnetConstructedDataOperationExpected is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataOperationExpected()
	// CreateBuilder creates a BACnetConstructedDataOperationExpectedBuilder
	CreateBACnetConstructedDataOperationExpectedBuilder() BACnetConstructedDataOperationExpectedBuilder
}

// _BACnetConstructedDataOperationExpected is the data-structure of this message
type _BACnetConstructedDataOperationExpected struct {
	BACnetConstructedDataContract
	LifeSafetyOperations BACnetLifeSafetyOperationTagged
}

var _ BACnetConstructedDataOperationExpected = (*_BACnetConstructedDataOperationExpected)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataOperationExpected)(nil)

// NewBACnetConstructedDataOperationExpected factory function for _BACnetConstructedDataOperationExpected
func NewBACnetConstructedDataOperationExpected(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lifeSafetyOperations BACnetLifeSafetyOperationTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOperationExpected {
	if lifeSafetyOperations == nil {
		panic("lifeSafetyOperations of type BACnetLifeSafetyOperationTagged for BACnetConstructedDataOperationExpected must not be nil")
	}
	_result := &_BACnetConstructedDataOperationExpected{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LifeSafetyOperations:          lifeSafetyOperations,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataOperationExpectedBuilder is a builder for BACnetConstructedDataOperationExpected
type BACnetConstructedDataOperationExpectedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lifeSafetyOperations BACnetLifeSafetyOperationTagged) BACnetConstructedDataOperationExpectedBuilder
	// WithLifeSafetyOperations adds LifeSafetyOperations (property field)
	WithLifeSafetyOperations(BACnetLifeSafetyOperationTagged) BACnetConstructedDataOperationExpectedBuilder
	// WithLifeSafetyOperationsBuilder adds LifeSafetyOperations (property field) which is build by the builder
	WithLifeSafetyOperationsBuilder(func(BACnetLifeSafetyOperationTaggedBuilder) BACnetLifeSafetyOperationTaggedBuilder) BACnetConstructedDataOperationExpectedBuilder
	// Build builds the BACnetConstructedDataOperationExpected or returns an error if something is wrong
	Build() (BACnetConstructedDataOperationExpected, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataOperationExpected
}

// NewBACnetConstructedDataOperationExpectedBuilder() creates a BACnetConstructedDataOperationExpectedBuilder
func NewBACnetConstructedDataOperationExpectedBuilder() BACnetConstructedDataOperationExpectedBuilder {
	return &_BACnetConstructedDataOperationExpectedBuilder{_BACnetConstructedDataOperationExpected: new(_BACnetConstructedDataOperationExpected)}
}

type _BACnetConstructedDataOperationExpectedBuilder struct {
	*_BACnetConstructedDataOperationExpected

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataOperationExpectedBuilder) = (*_BACnetConstructedDataOperationExpectedBuilder)(nil)

func (b *_BACnetConstructedDataOperationExpectedBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataOperationExpectedBuilder) WithMandatoryFields(lifeSafetyOperations BACnetLifeSafetyOperationTagged) BACnetConstructedDataOperationExpectedBuilder {
	return b.WithLifeSafetyOperations(lifeSafetyOperations)
}

func (b *_BACnetConstructedDataOperationExpectedBuilder) WithLifeSafetyOperations(lifeSafetyOperations BACnetLifeSafetyOperationTagged) BACnetConstructedDataOperationExpectedBuilder {
	b.LifeSafetyOperations = lifeSafetyOperations
	return b
}

func (b *_BACnetConstructedDataOperationExpectedBuilder) WithLifeSafetyOperationsBuilder(builderSupplier func(BACnetLifeSafetyOperationTaggedBuilder) BACnetLifeSafetyOperationTaggedBuilder) BACnetConstructedDataOperationExpectedBuilder {
	builder := builderSupplier(b.LifeSafetyOperations.CreateBACnetLifeSafetyOperationTaggedBuilder())
	var err error
	b.LifeSafetyOperations, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLifeSafetyOperationTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataOperationExpectedBuilder) Build() (BACnetConstructedDataOperationExpected, error) {
	if b.LifeSafetyOperations == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lifeSafetyOperations' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataOperationExpected.deepCopy(), nil
}

func (b *_BACnetConstructedDataOperationExpectedBuilder) MustBuild() BACnetConstructedDataOperationExpected {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataOperationExpectedBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataOperationExpectedBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataOperationExpectedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataOperationExpectedBuilder().(*_BACnetConstructedDataOperationExpectedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataOperationExpectedBuilder creates a BACnetConstructedDataOperationExpectedBuilder
func (b *_BACnetConstructedDataOperationExpected) CreateBACnetConstructedDataOperationExpectedBuilder() BACnetConstructedDataOperationExpectedBuilder {
	if b == nil {
		return NewBACnetConstructedDataOperationExpectedBuilder()
	}
	return &_BACnetConstructedDataOperationExpectedBuilder{_BACnetConstructedDataOperationExpected: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOperationExpected) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOperationExpected) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OPERATION_EXPECTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOperationExpected) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOperationExpected) GetLifeSafetyOperations() BACnetLifeSafetyOperationTagged {
	return m.LifeSafetyOperations
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOperationExpected) GetActualValue() BACnetLifeSafetyOperationTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLifeSafetyOperationTagged(m.GetLifeSafetyOperations())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOperationExpected(structType any) BACnetConstructedDataOperationExpected {
	if casted, ok := structType.(BACnetConstructedDataOperationExpected); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOperationExpected); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOperationExpected) GetTypeName() string {
	return "BACnetConstructedDataOperationExpected"
}

func (m *_BACnetConstructedDataOperationExpected) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lifeSafetyOperations)
	lengthInBits += m.LifeSafetyOperations.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOperationExpected) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataOperationExpected) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataOperationExpected BACnetConstructedDataOperationExpected, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOperationExpected"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOperationExpected")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lifeSafetyOperations, err := ReadSimpleField[BACnetLifeSafetyOperationTagged](ctx, "lifeSafetyOperations", ReadComplex[BACnetLifeSafetyOperationTagged](BACnetLifeSafetyOperationTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lifeSafetyOperations' field"))
	}
	m.LifeSafetyOperations = lifeSafetyOperations

	actualValue, err := ReadVirtualField[BACnetLifeSafetyOperationTagged](ctx, "actualValue", (*BACnetLifeSafetyOperationTagged)(nil), lifeSafetyOperations)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOperationExpected"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOperationExpected")
	}

	return m, nil
}

func (m *_BACnetConstructedDataOperationExpected) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOperationExpected) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOperationExpected"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOperationExpected")
		}

		if err := WriteSimpleField[BACnetLifeSafetyOperationTagged](ctx, "lifeSafetyOperations", m.GetLifeSafetyOperations(), WriteComplex[BACnetLifeSafetyOperationTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lifeSafetyOperations' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOperationExpected"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOperationExpected")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOperationExpected) IsBACnetConstructedDataOperationExpected() {}

func (m *_BACnetConstructedDataOperationExpected) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataOperationExpected) deepCopy() *_BACnetConstructedDataOperationExpected {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataOperationExpectedCopy := &_BACnetConstructedDataOperationExpected{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.LifeSafetyOperations.DeepCopy().(BACnetLifeSafetyOperationTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataOperationExpectedCopy
}

func (m *_BACnetConstructedDataOperationExpected) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

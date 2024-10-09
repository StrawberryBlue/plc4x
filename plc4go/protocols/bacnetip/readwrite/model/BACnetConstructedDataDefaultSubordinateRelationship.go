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

// BACnetConstructedDataDefaultSubordinateRelationship is the corresponding interface of BACnetConstructedDataDefaultSubordinateRelationship
type BACnetConstructedDataDefaultSubordinateRelationship interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDefaultSubordinateRelationship returns DefaultSubordinateRelationship (property field)
	GetDefaultSubordinateRelationship() BACnetRelationshipTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetRelationshipTagged
	// IsBACnetConstructedDataDefaultSubordinateRelationship is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDefaultSubordinateRelationship()
	// CreateBuilder creates a BACnetConstructedDataDefaultSubordinateRelationshipBuilder
	CreateBACnetConstructedDataDefaultSubordinateRelationshipBuilder() BACnetConstructedDataDefaultSubordinateRelationshipBuilder
}

// _BACnetConstructedDataDefaultSubordinateRelationship is the data-structure of this message
type _BACnetConstructedDataDefaultSubordinateRelationship struct {
	BACnetConstructedDataContract
	DefaultSubordinateRelationship BACnetRelationshipTagged
}

var _ BACnetConstructedDataDefaultSubordinateRelationship = (*_BACnetConstructedDataDefaultSubordinateRelationship)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDefaultSubordinateRelationship)(nil)

// NewBACnetConstructedDataDefaultSubordinateRelationship factory function for _BACnetConstructedDataDefaultSubordinateRelationship
func NewBACnetConstructedDataDefaultSubordinateRelationship(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, defaultSubordinateRelationship BACnetRelationshipTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDefaultSubordinateRelationship {
	if defaultSubordinateRelationship == nil {
		panic("defaultSubordinateRelationship of type BACnetRelationshipTagged for BACnetConstructedDataDefaultSubordinateRelationship must not be nil")
	}
	_result := &_BACnetConstructedDataDefaultSubordinateRelationship{
		BACnetConstructedDataContract:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DefaultSubordinateRelationship: defaultSubordinateRelationship,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDefaultSubordinateRelationshipBuilder is a builder for BACnetConstructedDataDefaultSubordinateRelationship
type BACnetConstructedDataDefaultSubordinateRelationshipBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(defaultSubordinateRelationship BACnetRelationshipTagged) BACnetConstructedDataDefaultSubordinateRelationshipBuilder
	// WithDefaultSubordinateRelationship adds DefaultSubordinateRelationship (property field)
	WithDefaultSubordinateRelationship(BACnetRelationshipTagged) BACnetConstructedDataDefaultSubordinateRelationshipBuilder
	// WithDefaultSubordinateRelationshipBuilder adds DefaultSubordinateRelationship (property field) which is build by the builder
	WithDefaultSubordinateRelationshipBuilder(func(BACnetRelationshipTaggedBuilder) BACnetRelationshipTaggedBuilder) BACnetConstructedDataDefaultSubordinateRelationshipBuilder
	// Build builds the BACnetConstructedDataDefaultSubordinateRelationship or returns an error if something is wrong
	Build() (BACnetConstructedDataDefaultSubordinateRelationship, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDefaultSubordinateRelationship
}

// NewBACnetConstructedDataDefaultSubordinateRelationshipBuilder() creates a BACnetConstructedDataDefaultSubordinateRelationshipBuilder
func NewBACnetConstructedDataDefaultSubordinateRelationshipBuilder() BACnetConstructedDataDefaultSubordinateRelationshipBuilder {
	return &_BACnetConstructedDataDefaultSubordinateRelationshipBuilder{_BACnetConstructedDataDefaultSubordinateRelationship: new(_BACnetConstructedDataDefaultSubordinateRelationship)}
}

type _BACnetConstructedDataDefaultSubordinateRelationshipBuilder struct {
	*_BACnetConstructedDataDefaultSubordinateRelationship

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDefaultSubordinateRelationshipBuilder) = (*_BACnetConstructedDataDefaultSubordinateRelationshipBuilder)(nil)

func (b *_BACnetConstructedDataDefaultSubordinateRelationshipBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataDefaultSubordinateRelationshipBuilder) WithMandatoryFields(defaultSubordinateRelationship BACnetRelationshipTagged) BACnetConstructedDataDefaultSubordinateRelationshipBuilder {
	return b.WithDefaultSubordinateRelationship(defaultSubordinateRelationship)
}

func (b *_BACnetConstructedDataDefaultSubordinateRelationshipBuilder) WithDefaultSubordinateRelationship(defaultSubordinateRelationship BACnetRelationshipTagged) BACnetConstructedDataDefaultSubordinateRelationshipBuilder {
	b.DefaultSubordinateRelationship = defaultSubordinateRelationship
	return b
}

func (b *_BACnetConstructedDataDefaultSubordinateRelationshipBuilder) WithDefaultSubordinateRelationshipBuilder(builderSupplier func(BACnetRelationshipTaggedBuilder) BACnetRelationshipTaggedBuilder) BACnetConstructedDataDefaultSubordinateRelationshipBuilder {
	builder := builderSupplier(b.DefaultSubordinateRelationship.CreateBACnetRelationshipTaggedBuilder())
	var err error
	b.DefaultSubordinateRelationship, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetRelationshipTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataDefaultSubordinateRelationshipBuilder) Build() (BACnetConstructedDataDefaultSubordinateRelationship, error) {
	if b.DefaultSubordinateRelationship == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'defaultSubordinateRelationship' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDefaultSubordinateRelationship.deepCopy(), nil
}

func (b *_BACnetConstructedDataDefaultSubordinateRelationshipBuilder) MustBuild() BACnetConstructedDataDefaultSubordinateRelationship {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataDefaultSubordinateRelationshipBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDefaultSubordinateRelationshipBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDefaultSubordinateRelationshipBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDefaultSubordinateRelationshipBuilder().(*_BACnetConstructedDataDefaultSubordinateRelationshipBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDefaultSubordinateRelationshipBuilder creates a BACnetConstructedDataDefaultSubordinateRelationshipBuilder
func (b *_BACnetConstructedDataDefaultSubordinateRelationship) CreateBACnetConstructedDataDefaultSubordinateRelationshipBuilder() BACnetConstructedDataDefaultSubordinateRelationshipBuilder {
	if b == nil {
		return NewBACnetConstructedDataDefaultSubordinateRelationshipBuilder()
	}
	return &_BACnetConstructedDataDefaultSubordinateRelationshipBuilder{_BACnetConstructedDataDefaultSubordinateRelationship: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEFAULT_SUBORDINATE_RELATIONSHIP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) GetDefaultSubordinateRelationship() BACnetRelationshipTagged {
	return m.DefaultSubordinateRelationship
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) GetActualValue() BACnetRelationshipTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetRelationshipTagged(m.GetDefaultSubordinateRelationship())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDefaultSubordinateRelationship(structType any) BACnetConstructedDataDefaultSubordinateRelationship {
	if casted, ok := structType.(BACnetConstructedDataDefaultSubordinateRelationship); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDefaultSubordinateRelationship); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) GetTypeName() string {
	return "BACnetConstructedDataDefaultSubordinateRelationship"
}

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (defaultSubordinateRelationship)
	lengthInBits += m.DefaultSubordinateRelationship.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDefaultSubordinateRelationship BACnetConstructedDataDefaultSubordinateRelationship, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDefaultSubordinateRelationship"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDefaultSubordinateRelationship")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	defaultSubordinateRelationship, err := ReadSimpleField[BACnetRelationshipTagged](ctx, "defaultSubordinateRelationship", ReadComplex[BACnetRelationshipTagged](BACnetRelationshipTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'defaultSubordinateRelationship' field"))
	}
	m.DefaultSubordinateRelationship = defaultSubordinateRelationship

	actualValue, err := ReadVirtualField[BACnetRelationshipTagged](ctx, "actualValue", (*BACnetRelationshipTagged)(nil), defaultSubordinateRelationship)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDefaultSubordinateRelationship"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDefaultSubordinateRelationship")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDefaultSubordinateRelationship"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDefaultSubordinateRelationship")
		}

		if err := WriteSimpleField[BACnetRelationshipTagged](ctx, "defaultSubordinateRelationship", m.GetDefaultSubordinateRelationship(), WriteComplex[BACnetRelationshipTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'defaultSubordinateRelationship' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDefaultSubordinateRelationship"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDefaultSubordinateRelationship")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) IsBACnetConstructedDataDefaultSubordinateRelationship() {
}

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) deepCopy() *_BACnetConstructedDataDefaultSubordinateRelationship {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDefaultSubordinateRelationshipCopy := &_BACnetConstructedDataDefaultSubordinateRelationship{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.DefaultSubordinateRelationship.DeepCopy().(BACnetRelationshipTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDefaultSubordinateRelationshipCopy
}

func (m *_BACnetConstructedDataDefaultSubordinateRelationship) String() string {
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

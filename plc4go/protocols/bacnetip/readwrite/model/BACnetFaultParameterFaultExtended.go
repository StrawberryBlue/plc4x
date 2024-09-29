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

// BACnetFaultParameterFaultExtended is the corresponding interface of BACnetFaultParameterFaultExtended
type BACnetFaultParameterFaultExtended interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetExtendedFaultType returns ExtendedFaultType (property field)
	GetExtendedFaultType() BACnetContextTagUnsignedInteger
	// GetParameters returns Parameters (property field)
	GetParameters() BACnetFaultParameterFaultExtendedParameters
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetFaultParameterFaultExtended is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultExtended()
	// CreateBuilder creates a BACnetFaultParameterFaultExtendedBuilder
	CreateBACnetFaultParameterFaultExtendedBuilder() BACnetFaultParameterFaultExtendedBuilder
}

// _BACnetFaultParameterFaultExtended is the data-structure of this message
type _BACnetFaultParameterFaultExtended struct {
	BACnetFaultParameterContract
	OpeningTag        BACnetOpeningTag
	VendorId          BACnetVendorIdTagged
	ExtendedFaultType BACnetContextTagUnsignedInteger
	Parameters        BACnetFaultParameterFaultExtendedParameters
	ClosingTag        BACnetClosingTag
}

var _ BACnetFaultParameterFaultExtended = (*_BACnetFaultParameterFaultExtended)(nil)
var _ BACnetFaultParameterRequirements = (*_BACnetFaultParameterFaultExtended)(nil)

// NewBACnetFaultParameterFaultExtended factory function for _BACnetFaultParameterFaultExtended
func NewBACnetFaultParameterFaultExtended(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, vendorId BACnetVendorIdTagged, extendedFaultType BACnetContextTagUnsignedInteger, parameters BACnetFaultParameterFaultExtendedParameters, closingTag BACnetClosingTag) *_BACnetFaultParameterFaultExtended {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetFaultParameterFaultExtended must not be nil")
	}
	if vendorId == nil {
		panic("vendorId of type BACnetVendorIdTagged for BACnetFaultParameterFaultExtended must not be nil")
	}
	if extendedFaultType == nil {
		panic("extendedFaultType of type BACnetContextTagUnsignedInteger for BACnetFaultParameterFaultExtended must not be nil")
	}
	if parameters == nil {
		panic("parameters of type BACnetFaultParameterFaultExtendedParameters for BACnetFaultParameterFaultExtended must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetFaultParameterFaultExtended must not be nil")
	}
	_result := &_BACnetFaultParameterFaultExtended{
		BACnetFaultParameterContract: NewBACnetFaultParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		VendorId:                     vendorId,
		ExtendedFaultType:            extendedFaultType,
		Parameters:                   parameters,
		ClosingTag:                   closingTag,
	}
	_result.BACnetFaultParameterContract.(*_BACnetFaultParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultExtendedBuilder is a builder for BACnetFaultParameterFaultExtended
type BACnetFaultParameterFaultExtendedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, vendorId BACnetVendorIdTagged, extendedFaultType BACnetContextTagUnsignedInteger, parameters BACnetFaultParameterFaultExtendedParameters, closingTag BACnetClosingTag) BACnetFaultParameterFaultExtendedBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetFaultParameterFaultExtendedBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetFaultParameterFaultExtendedBuilder
	// WithVendorId adds VendorId (property field)
	WithVendorId(BACnetVendorIdTagged) BACnetFaultParameterFaultExtendedBuilder
	// WithVendorIdBuilder adds VendorId (property field) which is build by the builder
	WithVendorIdBuilder(func(BACnetVendorIdTaggedBuilder) BACnetVendorIdTaggedBuilder) BACnetFaultParameterFaultExtendedBuilder
	// WithExtendedFaultType adds ExtendedFaultType (property field)
	WithExtendedFaultType(BACnetContextTagUnsignedInteger) BACnetFaultParameterFaultExtendedBuilder
	// WithExtendedFaultTypeBuilder adds ExtendedFaultType (property field) which is build by the builder
	WithExtendedFaultTypeBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetFaultParameterFaultExtendedBuilder
	// WithParameters adds Parameters (property field)
	WithParameters(BACnetFaultParameterFaultExtendedParameters) BACnetFaultParameterFaultExtendedBuilder
	// WithParametersBuilder adds Parameters (property field) which is build by the builder
	WithParametersBuilder(func(BACnetFaultParameterFaultExtendedParametersBuilder) BACnetFaultParameterFaultExtendedParametersBuilder) BACnetFaultParameterFaultExtendedBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetFaultParameterFaultExtendedBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetFaultParameterFaultExtendedBuilder
	// Build builds the BACnetFaultParameterFaultExtended or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultExtended, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultExtended
}

// NewBACnetFaultParameterFaultExtendedBuilder() creates a BACnetFaultParameterFaultExtendedBuilder
func NewBACnetFaultParameterFaultExtendedBuilder() BACnetFaultParameterFaultExtendedBuilder {
	return &_BACnetFaultParameterFaultExtendedBuilder{_BACnetFaultParameterFaultExtended: new(_BACnetFaultParameterFaultExtended)}
}

type _BACnetFaultParameterFaultExtendedBuilder struct {
	*_BACnetFaultParameterFaultExtended

	parentBuilder *_BACnetFaultParameterBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultExtendedBuilder) = (*_BACnetFaultParameterFaultExtendedBuilder)(nil)

func (b *_BACnetFaultParameterFaultExtendedBuilder) setParent(contract BACnetFaultParameterContract) {
	b.BACnetFaultParameterContract = contract
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, vendorId BACnetVendorIdTagged, extendedFaultType BACnetContextTagUnsignedInteger, parameters BACnetFaultParameterFaultExtendedParameters, closingTag BACnetClosingTag) BACnetFaultParameterFaultExtendedBuilder {
	return b.WithOpeningTag(openingTag).WithVendorId(vendorId).WithExtendedFaultType(extendedFaultType).WithParameters(parameters).WithClosingTag(closingTag)
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetFaultParameterFaultExtendedBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetFaultParameterFaultExtendedBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) WithVendorId(vendorId BACnetVendorIdTagged) BACnetFaultParameterFaultExtendedBuilder {
	b.VendorId = vendorId
	return b
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) WithVendorIdBuilder(builderSupplier func(BACnetVendorIdTaggedBuilder) BACnetVendorIdTaggedBuilder) BACnetFaultParameterFaultExtendedBuilder {
	builder := builderSupplier(b.VendorId.CreateBACnetVendorIdTaggedBuilder())
	var err error
	b.VendorId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetVendorIdTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) WithExtendedFaultType(extendedFaultType BACnetContextTagUnsignedInteger) BACnetFaultParameterFaultExtendedBuilder {
	b.ExtendedFaultType = extendedFaultType
	return b
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) WithExtendedFaultTypeBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetFaultParameterFaultExtendedBuilder {
	builder := builderSupplier(b.ExtendedFaultType.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.ExtendedFaultType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) WithParameters(parameters BACnetFaultParameterFaultExtendedParameters) BACnetFaultParameterFaultExtendedBuilder {
	b.Parameters = parameters
	return b
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) WithParametersBuilder(builderSupplier func(BACnetFaultParameterFaultExtendedParametersBuilder) BACnetFaultParameterFaultExtendedParametersBuilder) BACnetFaultParameterFaultExtendedBuilder {
	builder := builderSupplier(b.Parameters.CreateBACnetFaultParameterFaultExtendedParametersBuilder())
	var err error
	b.Parameters, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetFaultParameterFaultExtendedParametersBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetFaultParameterFaultExtendedBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetFaultParameterFaultExtendedBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) Build() (BACnetFaultParameterFaultExtended, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.VendorId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'vendorId' not set"))
	}
	if b.ExtendedFaultType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'extendedFaultType' not set"))
	}
	if b.Parameters == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'parameters' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultExtended.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) MustBuild() BACnetFaultParameterFaultExtended {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetFaultParameterFaultExtendedBuilder) Done() BACnetFaultParameterBuilder {
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) buildForBACnetFaultParameter() (BACnetFaultParameter, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultExtendedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultExtendedBuilder().(*_BACnetFaultParameterFaultExtendedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultExtendedBuilder creates a BACnetFaultParameterFaultExtendedBuilder
func (b *_BACnetFaultParameterFaultExtended) CreateBACnetFaultParameterFaultExtendedBuilder() BACnetFaultParameterFaultExtendedBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultExtendedBuilder()
	}
	return &_BACnetFaultParameterFaultExtendedBuilder{_BACnetFaultParameterFaultExtended: b.deepCopy()}
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

func (m *_BACnetFaultParameterFaultExtended) GetParent() BACnetFaultParameterContract {
	return m.BACnetFaultParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtended) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultExtended) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetFaultParameterFaultExtended) GetExtendedFaultType() BACnetContextTagUnsignedInteger {
	return m.ExtendedFaultType
}

func (m *_BACnetFaultParameterFaultExtended) GetParameters() BACnetFaultParameterFaultExtendedParameters {
	return m.Parameters
}

func (m *_BACnetFaultParameterFaultExtended) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtended(structType any) BACnetFaultParameterFaultExtended {
	if casted, ok := structType.(BACnetFaultParameterFaultExtended); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtended); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtended) GetTypeName() string {
	return "BACnetFaultParameterFaultExtended"
}

func (m *_BACnetFaultParameterFaultExtended) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterContract.(*_BACnetFaultParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits(ctx)

	// Simple field (extendedFaultType)
	lengthInBits += m.ExtendedFaultType.GetLengthInBits(ctx)

	// Simple field (parameters)
	lengthInBits += m.Parameters.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtended) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultExtended) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameter) (__bACnetFaultParameterFaultExtended BACnetFaultParameterFaultExtended, err error) {
	m.BACnetFaultParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	vendorId, err := ReadSimpleField[BACnetVendorIdTagged](ctx, "vendorId", ReadComplex[BACnetVendorIdTagged](BACnetVendorIdTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorId' field"))
	}
	m.VendorId = vendorId

	extendedFaultType, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "extendedFaultType", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extendedFaultType' field"))
	}
	m.ExtendedFaultType = extendedFaultType

	parameters, err := ReadSimpleField[BACnetFaultParameterFaultExtendedParameters](ctx, "parameters", ReadComplex[BACnetFaultParameterFaultExtendedParameters](BACnetFaultParameterFaultExtendedParametersParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameters' field"))
	}
	m.Parameters = parameters

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtended")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultExtended) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtended) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtended")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetVendorIdTagged](ctx, "vendorId", m.GetVendorId(), WriteComplex[BACnetVendorIdTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vendorId' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "extendedFaultType", m.GetExtendedFaultType(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'extendedFaultType' field")
		}

		if err := WriteSimpleField[BACnetFaultParameterFaultExtendedParameters](ctx, "parameters", m.GetParameters(), WriteComplex[BACnetFaultParameterFaultExtendedParameters](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'parameters' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtended")
		}
		return nil
	}
	return m.BACnetFaultParameterContract.(*_BACnetFaultParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtended) IsBACnetFaultParameterFaultExtended() {}

func (m *_BACnetFaultParameterFaultExtended) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultExtended) deepCopy() *_BACnetFaultParameterFaultExtended {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultExtendedCopy := &_BACnetFaultParameterFaultExtended{
		m.BACnetFaultParameterContract.(*_BACnetFaultParameter).deepCopy(),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.VendorId.DeepCopy().(BACnetVendorIdTagged),
		m.ExtendedFaultType.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.Parameters.DeepCopy().(BACnetFaultParameterFaultExtendedParameters),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetFaultParameterContract.(*_BACnetFaultParameter)._SubType = m
	return _BACnetFaultParameterFaultExtendedCopy
}

func (m *_BACnetFaultParameterFaultExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

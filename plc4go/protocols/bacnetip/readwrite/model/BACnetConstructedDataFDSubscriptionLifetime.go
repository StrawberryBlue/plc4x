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

// BACnetConstructedDataFDSubscriptionLifetime is the corresponding interface of BACnetConstructedDataFDSubscriptionLifetime
type BACnetConstructedDataFDSubscriptionLifetime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFdSubscriptionLifetime returns FdSubscriptionLifetime (property field)
	GetFdSubscriptionLifetime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataFDSubscriptionLifetime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataFDSubscriptionLifetime()
	// CreateBuilder creates a BACnetConstructedDataFDSubscriptionLifetimeBuilder
	CreateBACnetConstructedDataFDSubscriptionLifetimeBuilder() BACnetConstructedDataFDSubscriptionLifetimeBuilder
}

// _BACnetConstructedDataFDSubscriptionLifetime is the data-structure of this message
type _BACnetConstructedDataFDSubscriptionLifetime struct {
	BACnetConstructedDataContract
	FdSubscriptionLifetime BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataFDSubscriptionLifetime = (*_BACnetConstructedDataFDSubscriptionLifetime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataFDSubscriptionLifetime)(nil)

// NewBACnetConstructedDataFDSubscriptionLifetime factory function for _BACnetConstructedDataFDSubscriptionLifetime
func NewBACnetConstructedDataFDSubscriptionLifetime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, fdSubscriptionLifetime BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFDSubscriptionLifetime {
	if fdSubscriptionLifetime == nil {
		panic("fdSubscriptionLifetime of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataFDSubscriptionLifetime must not be nil")
	}
	_result := &_BACnetConstructedDataFDSubscriptionLifetime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FdSubscriptionLifetime:        fdSubscriptionLifetime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataFDSubscriptionLifetimeBuilder is a builder for BACnetConstructedDataFDSubscriptionLifetime
type BACnetConstructedDataFDSubscriptionLifetimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(fdSubscriptionLifetime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataFDSubscriptionLifetimeBuilder
	// WithFdSubscriptionLifetime adds FdSubscriptionLifetime (property field)
	WithFdSubscriptionLifetime(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataFDSubscriptionLifetimeBuilder
	// WithFdSubscriptionLifetimeBuilder adds FdSubscriptionLifetime (property field) which is build by the builder
	WithFdSubscriptionLifetimeBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataFDSubscriptionLifetimeBuilder
	// Build builds the BACnetConstructedDataFDSubscriptionLifetime or returns an error if something is wrong
	Build() (BACnetConstructedDataFDSubscriptionLifetime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataFDSubscriptionLifetime
}

// NewBACnetConstructedDataFDSubscriptionLifetimeBuilder() creates a BACnetConstructedDataFDSubscriptionLifetimeBuilder
func NewBACnetConstructedDataFDSubscriptionLifetimeBuilder() BACnetConstructedDataFDSubscriptionLifetimeBuilder {
	return &_BACnetConstructedDataFDSubscriptionLifetimeBuilder{_BACnetConstructedDataFDSubscriptionLifetime: new(_BACnetConstructedDataFDSubscriptionLifetime)}
}

type _BACnetConstructedDataFDSubscriptionLifetimeBuilder struct {
	*_BACnetConstructedDataFDSubscriptionLifetime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataFDSubscriptionLifetimeBuilder) = (*_BACnetConstructedDataFDSubscriptionLifetimeBuilder)(nil)

func (b *_BACnetConstructedDataFDSubscriptionLifetimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataFDSubscriptionLifetimeBuilder) WithMandatoryFields(fdSubscriptionLifetime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataFDSubscriptionLifetimeBuilder {
	return b.WithFdSubscriptionLifetime(fdSubscriptionLifetime)
}

func (b *_BACnetConstructedDataFDSubscriptionLifetimeBuilder) WithFdSubscriptionLifetime(fdSubscriptionLifetime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataFDSubscriptionLifetimeBuilder {
	b.FdSubscriptionLifetime = fdSubscriptionLifetime
	return b
}

func (b *_BACnetConstructedDataFDSubscriptionLifetimeBuilder) WithFdSubscriptionLifetimeBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataFDSubscriptionLifetimeBuilder {
	builder := builderSupplier(b.FdSubscriptionLifetime.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.FdSubscriptionLifetime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataFDSubscriptionLifetimeBuilder) Build() (BACnetConstructedDataFDSubscriptionLifetime, error) {
	if b.FdSubscriptionLifetime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'fdSubscriptionLifetime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataFDSubscriptionLifetime.deepCopy(), nil
}

func (b *_BACnetConstructedDataFDSubscriptionLifetimeBuilder) MustBuild() BACnetConstructedDataFDSubscriptionLifetime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataFDSubscriptionLifetimeBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataFDSubscriptionLifetimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataFDSubscriptionLifetimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataFDSubscriptionLifetimeBuilder().(*_BACnetConstructedDataFDSubscriptionLifetimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataFDSubscriptionLifetimeBuilder creates a BACnetConstructedDataFDSubscriptionLifetimeBuilder
func (b *_BACnetConstructedDataFDSubscriptionLifetime) CreateBACnetConstructedDataFDSubscriptionLifetimeBuilder() BACnetConstructedDataFDSubscriptionLifetimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataFDSubscriptionLifetimeBuilder()
	}
	return &_BACnetConstructedDataFDSubscriptionLifetimeBuilder{_BACnetConstructedDataFDSubscriptionLifetime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FD_SUBSCRIPTION_LIFETIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetFdSubscriptionLifetime() BACnetApplicationTagUnsignedInteger {
	return m.FdSubscriptionLifetime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetFdSubscriptionLifetime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFDSubscriptionLifetime(structType any) BACnetConstructedDataFDSubscriptionLifetime {
	if casted, ok := structType.(BACnetConstructedDataFDSubscriptionLifetime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFDSubscriptionLifetime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetTypeName() string {
	return "BACnetConstructedDataFDSubscriptionLifetime"
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (fdSubscriptionLifetime)
	lengthInBits += m.FdSubscriptionLifetime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataFDSubscriptionLifetime BACnetConstructedDataFDSubscriptionLifetime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFDSubscriptionLifetime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFDSubscriptionLifetime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fdSubscriptionLifetime, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "fdSubscriptionLifetime", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fdSubscriptionLifetime' field"))
	}
	m.FdSubscriptionLifetime = fdSubscriptionLifetime

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), fdSubscriptionLifetime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFDSubscriptionLifetime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFDSubscriptionLifetime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFDSubscriptionLifetime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFDSubscriptionLifetime")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "fdSubscriptionLifetime", m.GetFdSubscriptionLifetime(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fdSubscriptionLifetime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFDSubscriptionLifetime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFDSubscriptionLifetime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) IsBACnetConstructedDataFDSubscriptionLifetime() {
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) deepCopy() *_BACnetConstructedDataFDSubscriptionLifetime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataFDSubscriptionLifetimeCopy := &_BACnetConstructedDataFDSubscriptionLifetime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.FdSubscriptionLifetime.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataFDSubscriptionLifetimeCopy
}

func (m *_BACnetConstructedDataFDSubscriptionLifetime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

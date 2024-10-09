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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataRequired is the corresponding interface of BACnetConstructedDataRequired
type BACnetConstructedDataRequired interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataRequired is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataRequired()
	// CreateBuilder creates a BACnetConstructedDataRequiredBuilder
	CreateBACnetConstructedDataRequiredBuilder() BACnetConstructedDataRequiredBuilder
}

// _BACnetConstructedDataRequired is the data-structure of this message
type _BACnetConstructedDataRequired struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataRequired = (*_BACnetConstructedDataRequired)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataRequired)(nil)

// NewBACnetConstructedDataRequired factory function for _BACnetConstructedDataRequired
func NewBACnetConstructedDataRequired(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataRequired {
	_result := &_BACnetConstructedDataRequired{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataRequiredBuilder is a builder for BACnetConstructedDataRequired
type BACnetConstructedDataRequiredBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataRequiredBuilder
	// Build builds the BACnetConstructedDataRequired or returns an error if something is wrong
	Build() (BACnetConstructedDataRequired, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataRequired
}

// NewBACnetConstructedDataRequiredBuilder() creates a BACnetConstructedDataRequiredBuilder
func NewBACnetConstructedDataRequiredBuilder() BACnetConstructedDataRequiredBuilder {
	return &_BACnetConstructedDataRequiredBuilder{_BACnetConstructedDataRequired: new(_BACnetConstructedDataRequired)}
}

type _BACnetConstructedDataRequiredBuilder struct {
	*_BACnetConstructedDataRequired

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataRequiredBuilder) = (*_BACnetConstructedDataRequiredBuilder)(nil)

func (b *_BACnetConstructedDataRequiredBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataRequiredBuilder) WithMandatoryFields() BACnetConstructedDataRequiredBuilder {
	return b
}

func (b *_BACnetConstructedDataRequiredBuilder) Build() (BACnetConstructedDataRequired, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataRequired.deepCopy(), nil
}

func (b *_BACnetConstructedDataRequiredBuilder) MustBuild() BACnetConstructedDataRequired {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataRequiredBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataRequiredBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataRequiredBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataRequiredBuilder().(*_BACnetConstructedDataRequiredBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataRequiredBuilder creates a BACnetConstructedDataRequiredBuilder
func (b *_BACnetConstructedDataRequired) CreateBACnetConstructedDataRequiredBuilder() BACnetConstructedDataRequiredBuilder {
	if b == nil {
		return NewBACnetConstructedDataRequiredBuilder()
	}
	return &_BACnetConstructedDataRequiredBuilder{_BACnetConstructedDataRequired: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRequired) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataRequired) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_REQUIRED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRequired) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRequired(structType any) BACnetConstructedDataRequired {
	if casted, ok := structType.(BACnetConstructedDataRequired); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRequired); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRequired) GetTypeName() string {
	return "BACnetConstructedDataRequired"
}

func (m *_BACnetConstructedDataRequired) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataRequired) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataRequired) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataRequired BACnetConstructedDataRequired, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRequired"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRequired")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "An property identified by REQUIRED should never occur in the wild"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRequired"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRequired")
	}

	return m, nil
}

func (m *_BACnetConstructedDataRequired) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRequired) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRequired"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRequired")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRequired"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRequired")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataRequired) IsBACnetConstructedDataRequired() {}

func (m *_BACnetConstructedDataRequired) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataRequired) deepCopy() *_BACnetConstructedDataRequired {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataRequiredCopy := &_BACnetConstructedDataRequired{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataRequiredCopy
}

func (m *_BACnetConstructedDataRequired) String() string {
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

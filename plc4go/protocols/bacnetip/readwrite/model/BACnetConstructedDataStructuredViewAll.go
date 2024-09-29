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

// BACnetConstructedDataStructuredViewAll is the corresponding interface of BACnetConstructedDataStructuredViewAll
type BACnetConstructedDataStructuredViewAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataStructuredViewAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataStructuredViewAll()
	// CreateBuilder creates a BACnetConstructedDataStructuredViewAllBuilder
	CreateBACnetConstructedDataStructuredViewAllBuilder() BACnetConstructedDataStructuredViewAllBuilder
}

// _BACnetConstructedDataStructuredViewAll is the data-structure of this message
type _BACnetConstructedDataStructuredViewAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataStructuredViewAll = (*_BACnetConstructedDataStructuredViewAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataStructuredViewAll)(nil)

// NewBACnetConstructedDataStructuredViewAll factory function for _BACnetConstructedDataStructuredViewAll
func NewBACnetConstructedDataStructuredViewAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStructuredViewAll {
	_result := &_BACnetConstructedDataStructuredViewAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataStructuredViewAllBuilder is a builder for BACnetConstructedDataStructuredViewAll
type BACnetConstructedDataStructuredViewAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataStructuredViewAllBuilder
	// Build builds the BACnetConstructedDataStructuredViewAll or returns an error if something is wrong
	Build() (BACnetConstructedDataStructuredViewAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataStructuredViewAll
}

// NewBACnetConstructedDataStructuredViewAllBuilder() creates a BACnetConstructedDataStructuredViewAllBuilder
func NewBACnetConstructedDataStructuredViewAllBuilder() BACnetConstructedDataStructuredViewAllBuilder {
	return &_BACnetConstructedDataStructuredViewAllBuilder{_BACnetConstructedDataStructuredViewAll: new(_BACnetConstructedDataStructuredViewAll)}
}

type _BACnetConstructedDataStructuredViewAllBuilder struct {
	*_BACnetConstructedDataStructuredViewAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataStructuredViewAllBuilder) = (*_BACnetConstructedDataStructuredViewAllBuilder)(nil)

func (b *_BACnetConstructedDataStructuredViewAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataStructuredViewAllBuilder) WithMandatoryFields() BACnetConstructedDataStructuredViewAllBuilder {
	return b
}

func (b *_BACnetConstructedDataStructuredViewAllBuilder) Build() (BACnetConstructedDataStructuredViewAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataStructuredViewAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataStructuredViewAllBuilder) MustBuild() BACnetConstructedDataStructuredViewAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataStructuredViewAllBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataStructuredViewAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataStructuredViewAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataStructuredViewAllBuilder().(*_BACnetConstructedDataStructuredViewAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataStructuredViewAllBuilder creates a BACnetConstructedDataStructuredViewAllBuilder
func (b *_BACnetConstructedDataStructuredViewAll) CreateBACnetConstructedDataStructuredViewAllBuilder() BACnetConstructedDataStructuredViewAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataStructuredViewAllBuilder()
	}
	return &_BACnetConstructedDataStructuredViewAllBuilder{_BACnetConstructedDataStructuredViewAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStructuredViewAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_STRUCTURED_VIEW
}

func (m *_BACnetConstructedDataStructuredViewAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStructuredViewAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStructuredViewAll(structType any) BACnetConstructedDataStructuredViewAll {
	if casted, ok := structType.(BACnetConstructedDataStructuredViewAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStructuredViewAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStructuredViewAll) GetTypeName() string {
	return "BACnetConstructedDataStructuredViewAll"
}

func (m *_BACnetConstructedDataStructuredViewAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataStructuredViewAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataStructuredViewAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataStructuredViewAll BACnetConstructedDataStructuredViewAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStructuredViewAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStructuredViewAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStructuredViewAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStructuredViewAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataStructuredViewAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataStructuredViewAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStructuredViewAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStructuredViewAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStructuredViewAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStructuredViewAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStructuredViewAll) IsBACnetConstructedDataStructuredViewAll() {}

func (m *_BACnetConstructedDataStructuredViewAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataStructuredViewAll) deepCopy() *_BACnetConstructedDataStructuredViewAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataStructuredViewAllCopy := &_BACnetConstructedDataStructuredViewAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataStructuredViewAllCopy
}

func (m *_BACnetConstructedDataStructuredViewAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

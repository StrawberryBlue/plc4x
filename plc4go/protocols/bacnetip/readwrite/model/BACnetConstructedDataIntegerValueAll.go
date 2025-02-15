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

// BACnetConstructedDataIntegerValueAll is the corresponding interface of BACnetConstructedDataIntegerValueAll
type BACnetConstructedDataIntegerValueAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataIntegerValueAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIntegerValueAll()
	// CreateBuilder creates a BACnetConstructedDataIntegerValueAllBuilder
	CreateBACnetConstructedDataIntegerValueAllBuilder() BACnetConstructedDataIntegerValueAllBuilder
}

// _BACnetConstructedDataIntegerValueAll is the data-structure of this message
type _BACnetConstructedDataIntegerValueAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataIntegerValueAll = (*_BACnetConstructedDataIntegerValueAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIntegerValueAll)(nil)

// NewBACnetConstructedDataIntegerValueAll factory function for _BACnetConstructedDataIntegerValueAll
func NewBACnetConstructedDataIntegerValueAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegerValueAll {
	_result := &_BACnetConstructedDataIntegerValueAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIntegerValueAllBuilder is a builder for BACnetConstructedDataIntegerValueAll
type BACnetConstructedDataIntegerValueAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataIntegerValueAllBuilder
	// Build builds the BACnetConstructedDataIntegerValueAll or returns an error if something is wrong
	Build() (BACnetConstructedDataIntegerValueAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIntegerValueAll
}

// NewBACnetConstructedDataIntegerValueAllBuilder() creates a BACnetConstructedDataIntegerValueAllBuilder
func NewBACnetConstructedDataIntegerValueAllBuilder() BACnetConstructedDataIntegerValueAllBuilder {
	return &_BACnetConstructedDataIntegerValueAllBuilder{_BACnetConstructedDataIntegerValueAll: new(_BACnetConstructedDataIntegerValueAll)}
}

type _BACnetConstructedDataIntegerValueAllBuilder struct {
	*_BACnetConstructedDataIntegerValueAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIntegerValueAllBuilder) = (*_BACnetConstructedDataIntegerValueAllBuilder)(nil)

func (b *_BACnetConstructedDataIntegerValueAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataIntegerValueAllBuilder) WithMandatoryFields() BACnetConstructedDataIntegerValueAllBuilder {
	return b
}

func (b *_BACnetConstructedDataIntegerValueAllBuilder) Build() (BACnetConstructedDataIntegerValueAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIntegerValueAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataIntegerValueAllBuilder) MustBuild() BACnetConstructedDataIntegerValueAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataIntegerValueAllBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIntegerValueAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIntegerValueAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIntegerValueAllBuilder().(*_BACnetConstructedDataIntegerValueAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIntegerValueAllBuilder creates a BACnetConstructedDataIntegerValueAllBuilder
func (b *_BACnetConstructedDataIntegerValueAll) CreateBACnetConstructedDataIntegerValueAllBuilder() BACnetConstructedDataIntegerValueAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataIntegerValueAllBuilder()
	}
	return &_BACnetConstructedDataIntegerValueAllBuilder{_BACnetConstructedDataIntegerValueAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *_BACnetConstructedDataIntegerValueAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegerValueAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegerValueAll(structType any) BACnetConstructedDataIntegerValueAll {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegerValueAll) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueAll"
}

func (m *_BACnetConstructedDataIntegerValueAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegerValueAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIntegerValueAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIntegerValueAll BACnetConstructedDataIntegerValueAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegerValueAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegerValueAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIntegerValueAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIntegerValueAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegerValueAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegerValueAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegerValueAll) IsBACnetConstructedDataIntegerValueAll() {}

func (m *_BACnetConstructedDataIntegerValueAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIntegerValueAll) deepCopy() *_BACnetConstructedDataIntegerValueAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIntegerValueAllCopy := &_BACnetConstructedDataIntegerValueAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIntegerValueAllCopy
}

func (m *_BACnetConstructedDataIntegerValueAll) String() string {
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

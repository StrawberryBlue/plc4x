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

// BACnetConstructedDataCharacterstringValueAll is the corresponding interface of BACnetConstructedDataCharacterstringValueAll
type BACnetConstructedDataCharacterstringValueAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataCharacterstringValueAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCharacterstringValueAll()
	// CreateBuilder creates a BACnetConstructedDataCharacterstringValueAllBuilder
	CreateBACnetConstructedDataCharacterstringValueAllBuilder() BACnetConstructedDataCharacterstringValueAllBuilder
}

// _BACnetConstructedDataCharacterstringValueAll is the data-structure of this message
type _BACnetConstructedDataCharacterstringValueAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataCharacterstringValueAll = (*_BACnetConstructedDataCharacterstringValueAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCharacterstringValueAll)(nil)

// NewBACnetConstructedDataCharacterstringValueAll factory function for _BACnetConstructedDataCharacterstringValueAll
func NewBACnetConstructedDataCharacterstringValueAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCharacterstringValueAll {
	_result := &_BACnetConstructedDataCharacterstringValueAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCharacterstringValueAllBuilder is a builder for BACnetConstructedDataCharacterstringValueAll
type BACnetConstructedDataCharacterstringValueAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataCharacterstringValueAllBuilder
	// Build builds the BACnetConstructedDataCharacterstringValueAll or returns an error if something is wrong
	Build() (BACnetConstructedDataCharacterstringValueAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCharacterstringValueAll
}

// NewBACnetConstructedDataCharacterstringValueAllBuilder() creates a BACnetConstructedDataCharacterstringValueAllBuilder
func NewBACnetConstructedDataCharacterstringValueAllBuilder() BACnetConstructedDataCharacterstringValueAllBuilder {
	return &_BACnetConstructedDataCharacterstringValueAllBuilder{_BACnetConstructedDataCharacterstringValueAll: new(_BACnetConstructedDataCharacterstringValueAll)}
}

type _BACnetConstructedDataCharacterstringValueAllBuilder struct {
	*_BACnetConstructedDataCharacterstringValueAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCharacterstringValueAllBuilder) = (*_BACnetConstructedDataCharacterstringValueAllBuilder)(nil)

func (b *_BACnetConstructedDataCharacterstringValueAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataCharacterstringValueAllBuilder) WithMandatoryFields() BACnetConstructedDataCharacterstringValueAllBuilder {
	return b
}

func (b *_BACnetConstructedDataCharacterstringValueAllBuilder) Build() (BACnetConstructedDataCharacterstringValueAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCharacterstringValueAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataCharacterstringValueAllBuilder) MustBuild() BACnetConstructedDataCharacterstringValueAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataCharacterstringValueAllBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCharacterstringValueAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCharacterstringValueAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCharacterstringValueAllBuilder().(*_BACnetConstructedDataCharacterstringValueAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCharacterstringValueAllBuilder creates a BACnetConstructedDataCharacterstringValueAllBuilder
func (b *_BACnetConstructedDataCharacterstringValueAll) CreateBACnetConstructedDataCharacterstringValueAllBuilder() BACnetConstructedDataCharacterstringValueAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataCharacterstringValueAllBuilder()
	}
	return &_BACnetConstructedDataCharacterstringValueAllBuilder{_BACnetConstructedDataCharacterstringValueAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCharacterstringValueAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_CHARACTERSTRING_VALUE
}

func (m *_BACnetConstructedDataCharacterstringValueAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCharacterstringValueAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCharacterstringValueAll(structType any) BACnetConstructedDataCharacterstringValueAll {
	if casted, ok := structType.(BACnetConstructedDataCharacterstringValueAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCharacterstringValueAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCharacterstringValueAll) GetTypeName() string {
	return "BACnetConstructedDataCharacterstringValueAll"
}

func (m *_BACnetConstructedDataCharacterstringValueAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataCharacterstringValueAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCharacterstringValueAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCharacterstringValueAll BACnetConstructedDataCharacterstringValueAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCharacterstringValueAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCharacterstringValueAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCharacterstringValueAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCharacterstringValueAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCharacterstringValueAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCharacterstringValueAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCharacterstringValueAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCharacterstringValueAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCharacterstringValueAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCharacterstringValueAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCharacterstringValueAll) IsBACnetConstructedDataCharacterstringValueAll() {
}

func (m *_BACnetConstructedDataCharacterstringValueAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCharacterstringValueAll) deepCopy() *_BACnetConstructedDataCharacterstringValueAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCharacterstringValueAllCopy := &_BACnetConstructedDataCharacterstringValueAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCharacterstringValueAllCopy
}

func (m *_BACnetConstructedDataCharacterstringValueAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

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

// BACnetValueSourceNone is the corresponding interface of BACnetValueSourceNone
type BACnetValueSourceNone interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetValueSource
	// GetNone returns None (property field)
	GetNone() BACnetContextTagNull
	// IsBACnetValueSourceNone is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetValueSourceNone()
	// CreateBuilder creates a BACnetValueSourceNoneBuilder
	CreateBACnetValueSourceNoneBuilder() BACnetValueSourceNoneBuilder
}

// _BACnetValueSourceNone is the data-structure of this message
type _BACnetValueSourceNone struct {
	BACnetValueSourceContract
	None BACnetContextTagNull
}

var _ BACnetValueSourceNone = (*_BACnetValueSourceNone)(nil)
var _ BACnetValueSourceRequirements = (*_BACnetValueSourceNone)(nil)

// NewBACnetValueSourceNone factory function for _BACnetValueSourceNone
func NewBACnetValueSourceNone(peekedTagHeader BACnetTagHeader, none BACnetContextTagNull) *_BACnetValueSourceNone {
	if none == nil {
		panic("none of type BACnetContextTagNull for BACnetValueSourceNone must not be nil")
	}
	_result := &_BACnetValueSourceNone{
		BACnetValueSourceContract: NewBACnetValueSource(peekedTagHeader),
		None:                      none,
	}
	_result.BACnetValueSourceContract.(*_BACnetValueSource)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetValueSourceNoneBuilder is a builder for BACnetValueSourceNone
type BACnetValueSourceNoneBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(none BACnetContextTagNull) BACnetValueSourceNoneBuilder
	// WithNone adds None (property field)
	WithNone(BACnetContextTagNull) BACnetValueSourceNoneBuilder
	// WithNoneBuilder adds None (property field) which is build by the builder
	WithNoneBuilder(func(BACnetContextTagNullBuilder) BACnetContextTagNullBuilder) BACnetValueSourceNoneBuilder
	// Build builds the BACnetValueSourceNone or returns an error if something is wrong
	Build() (BACnetValueSourceNone, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetValueSourceNone
}

// NewBACnetValueSourceNoneBuilder() creates a BACnetValueSourceNoneBuilder
func NewBACnetValueSourceNoneBuilder() BACnetValueSourceNoneBuilder {
	return &_BACnetValueSourceNoneBuilder{_BACnetValueSourceNone: new(_BACnetValueSourceNone)}
}

type _BACnetValueSourceNoneBuilder struct {
	*_BACnetValueSourceNone

	parentBuilder *_BACnetValueSourceBuilder

	err *utils.MultiError
}

var _ (BACnetValueSourceNoneBuilder) = (*_BACnetValueSourceNoneBuilder)(nil)

func (b *_BACnetValueSourceNoneBuilder) setParent(contract BACnetValueSourceContract) {
	b.BACnetValueSourceContract = contract
}

func (b *_BACnetValueSourceNoneBuilder) WithMandatoryFields(none BACnetContextTagNull) BACnetValueSourceNoneBuilder {
	return b.WithNone(none)
}

func (b *_BACnetValueSourceNoneBuilder) WithNone(none BACnetContextTagNull) BACnetValueSourceNoneBuilder {
	b.None = none
	return b
}

func (b *_BACnetValueSourceNoneBuilder) WithNoneBuilder(builderSupplier func(BACnetContextTagNullBuilder) BACnetContextTagNullBuilder) BACnetValueSourceNoneBuilder {
	builder := builderSupplier(b.None.CreateBACnetContextTagNullBuilder())
	var err error
	b.None, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagNullBuilder failed"))
	}
	return b
}

func (b *_BACnetValueSourceNoneBuilder) Build() (BACnetValueSourceNone, error) {
	if b.None == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'none' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetValueSourceNone.deepCopy(), nil
}

func (b *_BACnetValueSourceNoneBuilder) MustBuild() BACnetValueSourceNone {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetValueSourceNoneBuilder) Done() BACnetValueSourceBuilder {
	return b.parentBuilder
}

func (b *_BACnetValueSourceNoneBuilder) buildForBACnetValueSource() (BACnetValueSource, error) {
	return b.Build()
}

func (b *_BACnetValueSourceNoneBuilder) DeepCopy() any {
	_copy := b.CreateBACnetValueSourceNoneBuilder().(*_BACnetValueSourceNoneBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetValueSourceNoneBuilder creates a BACnetValueSourceNoneBuilder
func (b *_BACnetValueSourceNone) CreateBACnetValueSourceNoneBuilder() BACnetValueSourceNoneBuilder {
	if b == nil {
		return NewBACnetValueSourceNoneBuilder()
	}
	return &_BACnetValueSourceNoneBuilder{_BACnetValueSourceNone: b.deepCopy()}
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

func (m *_BACnetValueSourceNone) GetParent() BACnetValueSourceContract {
	return m.BACnetValueSourceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetValueSourceNone) GetNone() BACnetContextTagNull {
	return m.None
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetValueSourceNone(structType any) BACnetValueSourceNone {
	if casted, ok := structType.(BACnetValueSourceNone); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetValueSourceNone); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetValueSourceNone) GetTypeName() string {
	return "BACnetValueSourceNone"
}

func (m *_BACnetValueSourceNone) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetValueSourceContract.(*_BACnetValueSource).getLengthInBits(ctx))

	// Simple field (none)
	lengthInBits += m.None.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetValueSourceNone) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetValueSourceNone) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetValueSource) (__bACnetValueSourceNone BACnetValueSourceNone, err error) {
	m.BACnetValueSourceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetValueSourceNone"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetValueSourceNone")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	none, err := ReadSimpleField[BACnetContextTagNull](ctx, "none", ReadComplex[BACnetContextTagNull](BACnetContextTagParseWithBufferProducer[BACnetContextTagNull]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_NULL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'none' field"))
	}
	m.None = none

	if closeErr := readBuffer.CloseContext("BACnetValueSourceNone"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetValueSourceNone")
	}

	return m, nil
}

func (m *_BACnetValueSourceNone) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetValueSourceNone) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetValueSourceNone"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetValueSourceNone")
		}

		if err := WriteSimpleField[BACnetContextTagNull](ctx, "none", m.GetNone(), WriteComplex[BACnetContextTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'none' field")
		}

		if popErr := writeBuffer.PopContext("BACnetValueSourceNone"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetValueSourceNone")
		}
		return nil
	}
	return m.BACnetValueSourceContract.(*_BACnetValueSource).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetValueSourceNone) IsBACnetValueSourceNone() {}

func (m *_BACnetValueSourceNone) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetValueSourceNone) deepCopy() *_BACnetValueSourceNone {
	if m == nil {
		return nil
	}
	_BACnetValueSourceNoneCopy := &_BACnetValueSourceNone{
		m.BACnetValueSourceContract.(*_BACnetValueSource).deepCopy(),
		m.None.DeepCopy().(BACnetContextTagNull),
	}
	m.BACnetValueSourceContract.(*_BACnetValueSource)._SubType = m
	return _BACnetValueSourceNoneCopy
}

func (m *_BACnetValueSourceNone) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

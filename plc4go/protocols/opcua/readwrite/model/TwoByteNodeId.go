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

// TwoByteNodeId is the corresponding interface of TwoByteNodeId
type TwoByteNodeId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetIdentifier returns Identifier (property field)
	GetIdentifier() uint8
	// IsTwoByteNodeId is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTwoByteNodeId()
	// CreateBuilder creates a TwoByteNodeIdBuilder
	CreateTwoByteNodeIdBuilder() TwoByteNodeIdBuilder
}

// _TwoByteNodeId is the data-structure of this message
type _TwoByteNodeId struct {
	Identifier uint8
}

var _ TwoByteNodeId = (*_TwoByteNodeId)(nil)

// NewTwoByteNodeId factory function for _TwoByteNodeId
func NewTwoByteNodeId(identifier uint8) *_TwoByteNodeId {
	return &_TwoByteNodeId{Identifier: identifier}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TwoByteNodeIdBuilder is a builder for TwoByteNodeId
type TwoByteNodeIdBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(identifier uint8) TwoByteNodeIdBuilder
	// WithIdentifier adds Identifier (property field)
	WithIdentifier(uint8) TwoByteNodeIdBuilder
	// Build builds the TwoByteNodeId or returns an error if something is wrong
	Build() (TwoByteNodeId, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TwoByteNodeId
}

// NewTwoByteNodeIdBuilder() creates a TwoByteNodeIdBuilder
func NewTwoByteNodeIdBuilder() TwoByteNodeIdBuilder {
	return &_TwoByteNodeIdBuilder{_TwoByteNodeId: new(_TwoByteNodeId)}
}

type _TwoByteNodeIdBuilder struct {
	*_TwoByteNodeId

	err *utils.MultiError
}

var _ (TwoByteNodeIdBuilder) = (*_TwoByteNodeIdBuilder)(nil)

func (b *_TwoByteNodeIdBuilder) WithMandatoryFields(identifier uint8) TwoByteNodeIdBuilder {
	return b.WithIdentifier(identifier)
}

func (b *_TwoByteNodeIdBuilder) WithIdentifier(identifier uint8) TwoByteNodeIdBuilder {
	b.Identifier = identifier
	return b
}

func (b *_TwoByteNodeIdBuilder) Build() (TwoByteNodeId, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TwoByteNodeId.deepCopy(), nil
}

func (b *_TwoByteNodeIdBuilder) MustBuild() TwoByteNodeId {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_TwoByteNodeIdBuilder) DeepCopy() any {
	_copy := b.CreateTwoByteNodeIdBuilder().(*_TwoByteNodeIdBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTwoByteNodeIdBuilder creates a TwoByteNodeIdBuilder
func (b *_TwoByteNodeId) CreateTwoByteNodeIdBuilder() TwoByteNodeIdBuilder {
	if b == nil {
		return NewTwoByteNodeIdBuilder()
	}
	return &_TwoByteNodeIdBuilder{_TwoByteNodeId: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TwoByteNodeId) GetIdentifier() uint8 {
	return m.Identifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTwoByteNodeId(structType any) TwoByteNodeId {
	if casted, ok := structType.(TwoByteNodeId); ok {
		return casted
	}
	if casted, ok := structType.(*TwoByteNodeId); ok {
		return *casted
	}
	return nil
}

func (m *_TwoByteNodeId) GetTypeName() string {
	return "TwoByteNodeId"
}

func (m *_TwoByteNodeId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (identifier)
	lengthInBits += 8

	return lengthInBits
}

func (m *_TwoByteNodeId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TwoByteNodeIdParse(ctx context.Context, theBytes []byte) (TwoByteNodeId, error) {
	return TwoByteNodeIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TwoByteNodeIdParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (TwoByteNodeId, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (TwoByteNodeId, error) {
		return TwoByteNodeIdParseWithBuffer(ctx, readBuffer)
	}
}

func TwoByteNodeIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TwoByteNodeId, error) {
	v, err := (&_TwoByteNodeId{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_TwoByteNodeId) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__twoByteNodeId TwoByteNodeId, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TwoByteNodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TwoByteNodeId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	identifier, err := ReadSimpleField(ctx, "identifier", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	m.Identifier = identifier

	if closeErr := readBuffer.CloseContext("TwoByteNodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TwoByteNodeId")
	}

	return m, nil
}

func (m *_TwoByteNodeId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TwoByteNodeId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TwoByteNodeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TwoByteNodeId")
	}

	if err := WriteSimpleField[uint8](ctx, "identifier", m.GetIdentifier(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'identifier' field")
	}

	if popErr := writeBuffer.PopContext("TwoByteNodeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TwoByteNodeId")
	}
	return nil
}

func (m *_TwoByteNodeId) IsTwoByteNodeId() {}

func (m *_TwoByteNodeId) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TwoByteNodeId) deepCopy() *_TwoByteNodeId {
	if m == nil {
		return nil
	}
	_TwoByteNodeIdCopy := &_TwoByteNodeId{
		m.Identifier,
	}
	return _TwoByteNodeIdCopy
}

func (m *_TwoByteNodeId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}

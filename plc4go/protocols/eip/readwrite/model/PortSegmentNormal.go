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

// PortSegmentNormal is the corresponding interface of PortSegmentNormal
type PortSegmentNormal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	PortSegmentType
	// GetPort returns Port (property field)
	GetPort() uint8
	// GetLinkAddress returns LinkAddress (property field)
	GetLinkAddress() uint8
	// IsPortSegmentNormal is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPortSegmentNormal()
	// CreateBuilder creates a PortSegmentNormalBuilder
	CreatePortSegmentNormalBuilder() PortSegmentNormalBuilder
}

// _PortSegmentNormal is the data-structure of this message
type _PortSegmentNormal struct {
	PortSegmentTypeContract
	Port        uint8
	LinkAddress uint8
}

var _ PortSegmentNormal = (*_PortSegmentNormal)(nil)
var _ PortSegmentTypeRequirements = (*_PortSegmentNormal)(nil)

// NewPortSegmentNormal factory function for _PortSegmentNormal
func NewPortSegmentNormal(port uint8, linkAddress uint8) *_PortSegmentNormal {
	_result := &_PortSegmentNormal{
		PortSegmentTypeContract: NewPortSegmentType(),
		Port:                    port,
		LinkAddress:             linkAddress,
	}
	_result.PortSegmentTypeContract.(*_PortSegmentType)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// PortSegmentNormalBuilder is a builder for PortSegmentNormal
type PortSegmentNormalBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(port uint8, linkAddress uint8) PortSegmentNormalBuilder
	// WithPort adds Port (property field)
	WithPort(uint8) PortSegmentNormalBuilder
	// WithLinkAddress adds LinkAddress (property field)
	WithLinkAddress(uint8) PortSegmentNormalBuilder
	// Build builds the PortSegmentNormal or returns an error if something is wrong
	Build() (PortSegmentNormal, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() PortSegmentNormal
}

// NewPortSegmentNormalBuilder() creates a PortSegmentNormalBuilder
func NewPortSegmentNormalBuilder() PortSegmentNormalBuilder {
	return &_PortSegmentNormalBuilder{_PortSegmentNormal: new(_PortSegmentNormal)}
}

type _PortSegmentNormalBuilder struct {
	*_PortSegmentNormal

	parentBuilder *_PortSegmentTypeBuilder

	err *utils.MultiError
}

var _ (PortSegmentNormalBuilder) = (*_PortSegmentNormalBuilder)(nil)

func (b *_PortSegmentNormalBuilder) setParent(contract PortSegmentTypeContract) {
	b.PortSegmentTypeContract = contract
}

func (b *_PortSegmentNormalBuilder) WithMandatoryFields(port uint8, linkAddress uint8) PortSegmentNormalBuilder {
	return b.WithPort(port).WithLinkAddress(linkAddress)
}

func (b *_PortSegmentNormalBuilder) WithPort(port uint8) PortSegmentNormalBuilder {
	b.Port = port
	return b
}

func (b *_PortSegmentNormalBuilder) WithLinkAddress(linkAddress uint8) PortSegmentNormalBuilder {
	b.LinkAddress = linkAddress
	return b
}

func (b *_PortSegmentNormalBuilder) Build() (PortSegmentNormal, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._PortSegmentNormal.deepCopy(), nil
}

func (b *_PortSegmentNormalBuilder) MustBuild() PortSegmentNormal {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_PortSegmentNormalBuilder) Done() PortSegmentTypeBuilder {
	return b.parentBuilder
}

func (b *_PortSegmentNormalBuilder) buildForPortSegmentType() (PortSegmentType, error) {
	return b.Build()
}

func (b *_PortSegmentNormalBuilder) DeepCopy() any {
	_copy := b.CreatePortSegmentNormalBuilder().(*_PortSegmentNormalBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreatePortSegmentNormalBuilder creates a PortSegmentNormalBuilder
func (b *_PortSegmentNormal) CreatePortSegmentNormalBuilder() PortSegmentNormalBuilder {
	if b == nil {
		return NewPortSegmentNormalBuilder()
	}
	return &_PortSegmentNormalBuilder{_PortSegmentNormal: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PortSegmentNormal) GetExtendedLinkAddress() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PortSegmentNormal) GetParent() PortSegmentTypeContract {
	return m.PortSegmentTypeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PortSegmentNormal) GetPort() uint8 {
	return m.Port
}

func (m *_PortSegmentNormal) GetLinkAddress() uint8 {
	return m.LinkAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPortSegmentNormal(structType any) PortSegmentNormal {
	if casted, ok := structType.(PortSegmentNormal); ok {
		return casted
	}
	if casted, ok := structType.(*PortSegmentNormal); ok {
		return *casted
	}
	return nil
}

func (m *_PortSegmentNormal) GetTypeName() string {
	return "PortSegmentNormal"
}

func (m *_PortSegmentNormal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.PortSegmentTypeContract.(*_PortSegmentType).getLengthInBits(ctx))

	// Simple field (port)
	lengthInBits += 4

	// Simple field (linkAddress)
	lengthInBits += 8

	return lengthInBits
}

func (m *_PortSegmentNormal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PortSegmentNormal) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_PortSegmentType) (__portSegmentNormal PortSegmentNormal, err error) {
	m.PortSegmentTypeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PortSegmentNormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PortSegmentNormal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	port, err := ReadSimpleField(ctx, "port", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'port' field"))
	}
	m.Port = port

	linkAddress, err := ReadSimpleField(ctx, "linkAddress", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'linkAddress' field"))
	}
	m.LinkAddress = linkAddress

	if closeErr := readBuffer.CloseContext("PortSegmentNormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PortSegmentNormal")
	}

	return m, nil
}

func (m *_PortSegmentNormal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PortSegmentNormal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PortSegmentNormal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PortSegmentNormal")
		}

		if err := WriteSimpleField[uint8](ctx, "port", m.GetPort(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'port' field")
		}

		if err := WriteSimpleField[uint8](ctx, "linkAddress", m.GetLinkAddress(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'linkAddress' field")
		}

		if popErr := writeBuffer.PopContext("PortSegmentNormal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PortSegmentNormal")
		}
		return nil
	}
	return m.PortSegmentTypeContract.(*_PortSegmentType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PortSegmentNormal) IsPortSegmentNormal() {}

func (m *_PortSegmentNormal) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PortSegmentNormal) deepCopy() *_PortSegmentNormal {
	if m == nil {
		return nil
	}
	_PortSegmentNormalCopy := &_PortSegmentNormal{
		m.PortSegmentTypeContract.(*_PortSegmentType).deepCopy(),
		m.Port,
		m.LinkAddress,
	}
	m.PortSegmentTypeContract.(*_PortSegmentType)._SubType = m
	return _PortSegmentNormalCopy
}

func (m *_PortSegmentNormal) String() string {
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

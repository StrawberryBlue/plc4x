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

// KnxNetIpRouting is the corresponding interface of KnxNetIpRouting
type KnxNetIpRouting interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
	// IsKnxNetIpRouting is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsKnxNetIpRouting()
	// CreateBuilder creates a KnxNetIpRoutingBuilder
	CreateKnxNetIpRoutingBuilder() KnxNetIpRoutingBuilder
}

// _KnxNetIpRouting is the data-structure of this message
type _KnxNetIpRouting struct {
	ServiceIdContract
	Version uint8
}

var _ KnxNetIpRouting = (*_KnxNetIpRouting)(nil)
var _ ServiceIdRequirements = (*_KnxNetIpRouting)(nil)

// NewKnxNetIpRouting factory function for _KnxNetIpRouting
func NewKnxNetIpRouting(version uint8) *_KnxNetIpRouting {
	_result := &_KnxNetIpRouting{
		ServiceIdContract: NewServiceId(),
		Version:           version,
	}
	_result.ServiceIdContract.(*_ServiceId)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// KnxNetIpRoutingBuilder is a builder for KnxNetIpRouting
type KnxNetIpRoutingBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(version uint8) KnxNetIpRoutingBuilder
	// WithVersion adds Version (property field)
	WithVersion(uint8) KnxNetIpRoutingBuilder
	// Build builds the KnxNetIpRouting or returns an error if something is wrong
	Build() (KnxNetIpRouting, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() KnxNetIpRouting
}

// NewKnxNetIpRoutingBuilder() creates a KnxNetIpRoutingBuilder
func NewKnxNetIpRoutingBuilder() KnxNetIpRoutingBuilder {
	return &_KnxNetIpRoutingBuilder{_KnxNetIpRouting: new(_KnxNetIpRouting)}
}

type _KnxNetIpRoutingBuilder struct {
	*_KnxNetIpRouting

	parentBuilder *_ServiceIdBuilder

	err *utils.MultiError
}

var _ (KnxNetIpRoutingBuilder) = (*_KnxNetIpRoutingBuilder)(nil)

func (b *_KnxNetIpRoutingBuilder) setParent(contract ServiceIdContract) {
	b.ServiceIdContract = contract
}

func (b *_KnxNetIpRoutingBuilder) WithMandatoryFields(version uint8) KnxNetIpRoutingBuilder {
	return b.WithVersion(version)
}

func (b *_KnxNetIpRoutingBuilder) WithVersion(version uint8) KnxNetIpRoutingBuilder {
	b.Version = version
	return b
}

func (b *_KnxNetIpRoutingBuilder) Build() (KnxNetIpRouting, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._KnxNetIpRouting.deepCopy(), nil
}

func (b *_KnxNetIpRoutingBuilder) MustBuild() KnxNetIpRouting {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_KnxNetIpRoutingBuilder) Done() ServiceIdBuilder {
	return b.parentBuilder
}

func (b *_KnxNetIpRoutingBuilder) buildForServiceId() (ServiceId, error) {
	return b.Build()
}

func (b *_KnxNetIpRoutingBuilder) DeepCopy() any {
	_copy := b.CreateKnxNetIpRoutingBuilder().(*_KnxNetIpRoutingBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateKnxNetIpRoutingBuilder creates a KnxNetIpRoutingBuilder
func (b *_KnxNetIpRouting) CreateKnxNetIpRoutingBuilder() KnxNetIpRoutingBuilder {
	if b == nil {
		return NewKnxNetIpRoutingBuilder()
	}
	return &_KnxNetIpRoutingBuilder{_KnxNetIpRouting: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetIpRouting) GetServiceType() uint8 {
	return 0x05
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetIpRouting) GetParent() ServiceIdContract {
	return m.ServiceIdContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetIpRouting) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastKnxNetIpRouting(structType any) KnxNetIpRouting {
	if casted, ok := structType.(KnxNetIpRouting); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetIpRouting); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetIpRouting) GetTypeName() string {
	return "KnxNetIpRouting"
}

func (m *_KnxNetIpRouting) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ServiceIdContract.(*_ServiceId).getLengthInBits(ctx))

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetIpRouting) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_KnxNetIpRouting) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ServiceId) (__knxNetIpRouting KnxNetIpRouting, err error) {
	m.ServiceIdContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetIpRouting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetIpRouting")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	if closeErr := readBuffer.CloseContext("KnxNetIpRouting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetIpRouting")
	}

	return m, nil
}

func (m *_KnxNetIpRouting) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxNetIpRouting) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpRouting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetIpRouting")
		}

		if err := WriteSimpleField[uint8](ctx, "version", m.GetVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpRouting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetIpRouting")
		}
		return nil
	}
	return m.ServiceIdContract.(*_ServiceId).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_KnxNetIpRouting) IsKnxNetIpRouting() {}

func (m *_KnxNetIpRouting) DeepCopy() any {
	return m.deepCopy()
}

func (m *_KnxNetIpRouting) deepCopy() *_KnxNetIpRouting {
	if m == nil {
		return nil
	}
	_KnxNetIpRoutingCopy := &_KnxNetIpRouting{
		m.ServiceIdContract.(*_ServiceId).deepCopy(),
		m.Version,
	}
	m.ServiceIdContract.(*_ServiceId)._SubType = m
	return _KnxNetIpRoutingCopy
}

func (m *_KnxNetIpRouting) String() string {
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
